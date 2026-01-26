// Package helpers provides utilities for integration tests
package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"
)

// MockResponse represents a mock HTTP response
type MockResponse struct {
	Status  int
	Headers map[string]string
	Body    interface{}
}

// MockRequest represents a recorded request
type MockRequest struct {
	Method  string
	URL     string
	Headers http.Header
	Body    []byte
}

// MockHTTPClient is a mock HTTP client for testing
type MockHTTPClient struct {
	responses map[string][]*MockResponse
	requests  []*MockRequest
	mu        sync.RWMutex
}

// NewMockHTTPClient creates a new mock HTTP client
func NewMockHTTPClient() *MockHTTPClient {
	return &MockHTTPClient{
		responses: make(map[string][]*MockResponse),
		requests:  make([]*MockRequest, 0),
	}
}

// MockResponse registers a mock response for a specific endpoint
func (m *MockHTTPClient) MockResponse(method, urlPattern string, response *MockResponse) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := fmt.Sprintf("%s:%s", method, urlPattern)
	if _, exists := m.responses[key]; !exists {
		m.responses[key] = make([]*MockResponse, 0)
	}
	m.responses[key] = append(m.responses[key], response)
}

// RoundTrip implements http.RoundTripper interface
func (m *MockHTTPClient) RoundTrip(req *http.Request) (*http.Response, error) {
	// Record the request
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore body for potential retries
	}

	m.mu.Lock()
	m.requests = append(m.requests, &MockRequest{
		Method:  req.Method,
		URL:     req.URL.String(),
		Headers: req.Header.Clone(),
		Body:    bodyBytes,
	})
	m.mu.Unlock()

	// Find matching mock response
	mockResp := m.findMockResponse(req.Method, req.URL.String())
	if mockResp == nil {
		return nil, fmt.Errorf("no mock response configured for %s %s", req.Method, req.URL.String())
	}

	// Build response
	headers := mockResp.Headers
	if headers == nil {
		headers = map[string]string{
			"Content-Type": "application/vnd.api+json",
		}
	}

	var body []byte
	var err error
	if mockResp.Body != nil {
		body, err = json.Marshal(mockResp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal mock response body: %w", err)
		}
	}

	resp := &http.Response{
		StatusCode: mockResp.Status,
		Status:     http.StatusText(mockResp.Status),
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewReader(body)),
		Request:    req,
	}

	for k, v := range headers {
		resp.Header.Set(k, v)
	}

	return resp, nil
}

// findMockResponse finds a matching mock response for the given method and URL
func (m *MockHTTPClient) findMockResponse(method, url string) *MockResponse {
	m.mu.Lock()
	defer m.mu.Unlock()

	for key, responses := range m.responses {
		mockMethod, mockURL := parseKey(key)

		if mockMethod != method {
			continue
		}

		// Check if URL matches (either exact match or regex)
		if matchesURL(mockURL, url) && len(responses) > 0 {
			// Pop first response
			resp := responses[0]
			m.responses[key] = responses[1:]
			return resp
		}
	}

	return nil
}

// GetRequests returns all recorded requests
func (m *MockHTTPClient) GetRequests() []*MockRequest {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*MockRequest, len(m.requests))
	copy(result, m.requests)
	return result
}

// GetLastRequest returns the last recorded request
func (m *MockHTTPClient) GetLastRequest() *MockRequest {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(m.requests) == 0 {
		return nil
	}
	return m.requests[len(m.requests)-1]
}

// Reset clears all mock responses and recorded requests
func (m *MockHTTPClient) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.responses = make(map[string][]*MockResponse)
	m.requests = make([]*MockRequest, 0)
}

// VerifyRequest verifies that a request was made
func (m *MockHTTPClient) VerifyRequest(method, urlPattern string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, req := range m.requests {
		if req.Method == method && matchesURL(urlPattern, req.URL) {
			return true
		}
	}
	return false
}

// CountRequests counts requests matching the criteria
func (m *MockHTTPClient) CountRequests(method, urlPattern string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, req := range m.requests {
		if (method == "" || req.Method == method) &&
			(urlPattern == "" || matchesURL(urlPattern, req.URL)) {
			count++
		}
	}
	return count
}

// Helper functions

func parseKey(key string) (method, url string) {
	parts := bytes.SplitN([]byte(key), []byte(":"), 2)
	if len(parts) == 2 {
		return string(parts[0]), string(parts[1])
	}
	return "", ""
}

func matchesURL(pattern, url string) bool {
	// Simple string matching
	if bytes.Contains([]byte(url), []byte(pattern)) {
		return true
	}

	// Try regex matching
	re, err := regexp.Compile(pattern)
	if err == nil && re.MatchString(url) {
		return true
	}

	return false
}
