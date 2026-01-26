# Integration Tests for Latitude.sh Go SDK

This directory contains integration tests for the Latitude.sh Go SDK. The tests use a mock HTTP client to simulate API responses without making actual network calls.

## Structure

```
tests/integration/
├── helpers/
│   ├── mock_http_client.go  # Mock HTTP client implementation
│   ├── mock_data.go          # Mock response data
│   └── test_client.go        # Test client setup utilities
├── plans_test.go             # Tests for Plans resource
├── apikeys_test.go           # Tests for API Keys resource
├── projects_test.go          # Tests for Projects resource
├── teams_test.go             # Tests for Teams resource
├── userdata_test.go          # Tests for User Data resource
├── servers_test.go           # Tests for Servers resource
└── README.md                 # This file
```

## Running Tests

### Run All Tests

```bash
cd tests/integration
go test -v ./...
```

### Run Specific Test File

```bash
go test -v -run TestPlansIntegration
go test -v -run TestAPIKeysIntegration
go test -v -run TestProjectsIntegration
```

### Run Specific Test Case

```bash
go test -v -run TestPlansIntegration/List_Operations/should_list_all_plans
```

### Run with Coverage

```bash
go test -v -cover ./...
```

### Generate Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Test Structure

Each test file follows this pattern:

1. **Setup**: Initialize test client with mock HTTP client
2. **Arrange**: Configure mock responses for API endpoints
3. **Act**: Call SDK methods
4. **Assert**: Verify results and behavior

### Example Test

```go
func TestPlansIntegration(t *testing.T) {
    testClient := helpers.SetupTest()

    t.Run("should list all plans", func(t *testing.T) {
        // Reset mock client before test
        testClient.Reset()

        // Arrange - Mock API response
        testClient.MockClient.MockResponse("GET", "/plans", &helpers.MockResponse{
            Status: 200,
            Body:   helpers.MockPlanList(),
        })

        // Act - Call SDK method
        result, err := testClient.SDK.Plans.List(context.Background(), operations.GetPlansRequest{})

        // Assert - Verify results
        require.NoError(t, err)
        assert.NotNil(t, result)
        assert.Equal(t, 2, len(result.Data))
    })
}
```

## Mock HTTP Client

The mock HTTP client (`MockHTTPClient`) intercepts HTTP requests and returns predefined responses. This allows testing without making actual API calls.

### Features

- **Request Recording**: All requests are recorded for verification
- **Response Mocking**: Define responses for specific endpoints
- **Pattern Matching**: Support for exact and regex URL matching
- **Request Verification**: Verify that expected requests were made

### Mock Response Configuration

```go
testClient.MockClient.MockResponse("GET", "/plans", &helpers.MockResponse{
    Status: 200,
    Headers: map[string]string{
        "Content-Type": "application/vnd.api+json",
    },
    Body: map[string]interface{}{
        "data": []interface{}{...},
    },
})
```

### Request Verification

```go
// Verify a request was made
assert.True(t, testClient.MockClient.VerifyRequest("DELETE", "/api_keys/apikey_test123"))

// Count requests matching criteria
count := testClient.MockClient.CountRequests("GET", "/plans")
assert.Equal(t, 1, count)
```

## Mock Data

The `helpers/mock_data.go` file contains factory functions for creating mock API responses:

- `MockPlan()` - Single plan
- `MockPlanList()` - List of plans
- `MockAPIKey()` - Single API key
- `MockAPIKeyList()` - List of API keys
- `MockProject()` - Single project
- `MockProjectList()` - List of projects
- `MockTeam()` - Single team
- `MockTeamList()` - List of teams
- `MockUserData()` - Single user data
- `MockUserDataList()` - List of user data
- `MockServer()` - Single server
- `MockServerList()` - List of servers
- `MockError()` - Error response
- `MockValidationError()` - Validation error response

## Test Coverage

The integration tests cover:

- ✅ **Plans**: List, Get, Filter operations
- ✅ **API Keys**: List, Create, Update, Delete, Security features
- ✅ **Projects**: List, Get, Create, Update, Delete
- ✅ **Teams**: List, Get current team
- ✅ **User Data**: List, Get, Create, Update, Delete
- ✅ **Servers**: List, Get, Filter operations

## Writing New Tests

To add tests for a new resource:

1. Create a new test file (e.g., `resource_test.go`)
2. Import required packages:
   ```go
   import (
       "context"
       "testing"
       "github.com/latitudesh/latitudesh-go-sdk/models/operations"
       "github.com/latitudesh/latitudesh-go-sdk/tests/integration/helpers"
       "github.com/stretchr/testify/assert"
       "github.com/stretchr/testify/require"
   )
   ```
3. Follow the test structure pattern (Setup, Arrange, Act, Assert)
4. Add mock data to `helpers/mock_data.go` if needed

## Best Practices

1. **Always reset** the mock client before each test: `testClient.Reset()`
2. **Use subtests** with `t.Run()` for better organization
3. **Test error cases** in addition to success cases
4. **Verify requests** when testing operations with no response body
5. **Use descriptive test names** that explain what is being tested
6. **Follow AAA pattern**: Arrange, Act, Assert

## Dependencies

- **testify**: Assertion and testing utilities
- **stretchr/testify**: Enhanced testing capabilities

## CI/CD Integration

These tests can be integrated into CI/CD pipelines:

```yaml
# GitHub Actions example
- name: Run Integration Tests
  run: |
    cd tests/integration
    go test -v -cover ./...
```

## Troubleshooting

### Tests Failing with "no mock response configured"

Make sure you're calling `testClient.MockClient.MockResponse()` before making the SDK call, and that the URL pattern matches the actual request URL.

### Import Errors

Run `go mod tidy` in the SDK root directory to ensure all dependencies are installed.

### Test Isolation Issues

Always call `testClient.Reset()` at the beginning of each test to clear previous mock responses and recorded requests.
