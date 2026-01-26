// Package helpers provides mock data for integration tests
package helpers

import "time"

// MockPlan returns a mock plan
func MockPlan() map[string]interface{} {
	return map[string]interface{}{
		"id":   "c3.small.x86",
		"type": "plans",
		"attributes": map[string]interface{}{
			"name": "c3.small.x86",
			"slug": "c3-small-x86",
			"line": "c3",
			"specs": map[string]interface{}{
				"cpu": map[string]interface{}{
					"type":  "Intel Xeon E3",
					"clock": 3.4,
					"cores": 4,
					"count": 1,
				},
				"memory": map[string]interface{}{
					"total": 16,
				},
				"nics": []map[string]interface{}{
					{
						"type":  "1Gbps",
						"count": 2,
					},
				},
				"drives": []map[string]interface{}{
					{
						"type":  "SSD",
						"size":  "480 GB",
						"count": 2,
					},
				},
			},
			"regions": []map[string]interface{}{
				{
					"name": "São Paulo",
					"locations": map[string]interface{}{
						"available": []string{"sao-paulo-1"},
						"in_stock":  []string{"sao-paulo-1"},
					},
					"stock_level":       "high",
					"deploys_instantly": []string{"sao-paulo-1"},
				},
			},
			"features": []string{"ssh", "raid"},
			"available_operating_systems": []map[string]interface{}{
				{
					"id":   "ubuntu_22_04_x64",
					"name": "Ubuntu 22.04 LTS",
					"slug": "ubuntu-22-04-x64",
				},
			},
		},
	}
}

// MockPlanList returns a mock list of plans
func MockPlanList() map[string]interface{} {
	return map[string]interface{}{
		"data": []interface{}{
			MockPlan(),
			map[string]interface{}{
				"id":   "c3.large.x86",
				"type": "plans",
				"attributes": map[string]interface{}{
					"name": "c3.large.x86",
					"slug": "c3-large-x86",
					"line": "c3",
					"specs": map[string]interface{}{
						"cpu": map[string]interface{}{
							"type":  "Intel Xeon E5",
							"clock": 3.6,
							"cores": 8,
							"count": 2,
						},
						"memory": map[string]interface{}{
							"total": 64,
						},
						"nics": []map[string]interface{}{
							{
								"type":  "10Gbps",
								"count": 2,
							},
						},
						"drives": []map[string]interface{}{
							{
								"type":  "NVME",
								"size":  "1920 GB",
								"count": 2,
							},
						},
						"gpu": map[string]interface{}{
							"type":  "NVIDIA RTX 4090",
							"count": 2,
						},
					},
					"regions": []map[string]interface{}{
						{
							"name": "São Paulo",
							"locations": map[string]interface{}{
								"available": []string{"sao-paulo-1"},
								"in_stock":  []string{"sao-paulo-1"},
							},
							"stock_level":       "low",
							"deploys_instantly": []string{"sao-paulo-1"},
						},
						{
							"name": "Santiago",
							"locations": map[string]interface{}{
								"available": []string{"santiago-1"},
								"in_stock":  []string{},
							},
							"stock_level":       "unavailable",
							"deploys_instantly": []string{},
						},
					},
					"features": []string{"ssh", "raid", "sev"},
					"available_operating_systems": []map[string]interface{}{
						{
							"id":   "ubuntu_22_04_x64",
							"name": "Ubuntu 22.04 LTS",
							"slug": "ubuntu-22-04-x64",
						},
					},
				},
			},
		},
		"meta": map[string]interface{}{
			"total":        2,
			"current_page": 1,
			"last_page":    1,
			"per_page":     25,
		},
	}
}

// MockAPIKey returns a mock API key
func MockAPIKey() map[string]interface{} {
	return map[string]interface{}{
		"id":   "apikey_test123",
		"type": "api_keys",
		"attributes": map[string]interface{}{
			"name":             "Test API Key",
			"api_version":      "v1",
			"token_last_slice": "abc12",
			"read_only":        false,
			"allowed_ips":      nil,
			"last_used_at":     "2024-01-20T10:30:00Z",
			"user": map[string]interface{}{
				"id":    "user_test123",
				"email": "test@example.com",
			},
			"created_at": "2024-01-15T00:00:00Z",
			"updated_at": "2024-01-20T10:30:00Z",
		},
	}
}

// MockAPIKeyList returns a mock list of API keys
func MockAPIKeyList() map[string]interface{} {
	return map[string]interface{}{
		"data": []interface{}{
			MockAPIKey(),
			map[string]interface{}{
				"id":   "apikey_test456",
				"type": "api_keys",
				"attributes": map[string]interface{}{
					"name":             "Read Only Key",
					"api_version":      "v1",
					"token_last_slice": "xyz99",
					"read_only":        true,
					"allowed_ips":      []string{"192.168.1.0/24"},
					"last_used_at":     nil,
					"user": map[string]interface{}{
						"id":    "user_test123",
						"email": "test@example.com",
					},
					"created_at": "2024-01-16T00:00:00Z",
					"updated_at": "2024-01-16T00:00:00Z",
				},
			},
		},
	}
}

// MockProject returns a mock project
func MockProject() map[string]interface{} {
	return map[string]interface{}{
		"id":   "project_test123",
		"type": "projects",
		"attributes": map[string]interface{}{
			"name":        "Test Project",
			"slug":        "test-project",
			"description": "A test project",
			"bandwidth_alert": map[string]interface{}{
				"enabled":   true,
				"threshold": 80,
			},
			"environment": "Development",
			"created_at":  "2024-01-10T00:00:00Z",
			"updated_at":  "2024-01-15T00:00:00Z",
		},
	}
}

// MockProjectList returns a mock list of projects
func MockProjectList() map[string]interface{} {
	return map[string]interface{}{
		"data": []interface{}{
			MockProject(),
		},
		"meta": map[string]interface{}{
			"total":        1,
			"current_page": 1,
			"last_page":    1,
			"per_page":     25,
		},
	}
}

// MockTeam returns a mock team
func MockTeam() map[string]interface{} {
	return map[string]interface{}{
		"id":   "team_test123",
		"type": "teams",
		"attributes": map[string]interface{}{
			"name":        "Test Team",
			"slug":        "test-team",
			"description": "A test team",
			"address":     "123 Test St, São Paulo, BR, 01310-100",
			"status":      "active",
			"currency":    "USD",
			"created_at":  "2024-01-01T00:00:00Z",
			"updated_at":  "2024-01-10T00:00:00Z",
		},
	}
}

// MockTeamList returns a mock list of teams
func MockTeamList() map[string]interface{} {
	return map[string]interface{}{
		"data": []interface{}{
			MockTeam(),
		},
	}
}

// MockUserData returns a mock user data
func MockUserData() map[string]interface{} {
	return map[string]interface{}{
		"id":   "userdata_test123",
		"type": "user_data",
		"attributes": map[string]interface{}{
			"description": "Test User Data",
			"content":     "#!/bin/bash\necho 'Hello World'",
			"created_at":  "2024-01-18T00:00:00Z",
			"updated_at":  "2024-01-18T00:00:00Z",
		},
	}
}

// MockUserDataList returns a mock list of user data
func MockUserDataList() map[string]interface{} {
	return map[string]interface{}{
		"data": []interface{}{
			MockUserData(),
		},
		"meta": map[string]interface{}{
			"total":        1,
			"current_page": 1,
			"last_page":    1,
			"per_page":     25,
		},
	}
}

// MockServer returns a mock server
func MockServer() map[string]interface{} {
	return map[string]interface{}{
		"id":   "server_test123",
		"type": "servers",
		"attributes": map[string]interface{}{
			"hostname": "test-server-01",
			"label":    "Test Server",
			"status":   "ready",
			"role":     "server",
			"operating_system": map[string]interface{}{
				"id":   "ubuntu_22_04_x64",
				"name": "Ubuntu 22.04 LTS",
				"slug": "ubuntu-22-04-x64",
			},
			"region": map[string]interface{}{
				"city":    "São Paulo",
				"country": "BR",
				"site": map[string]interface{}{
					"id":       "sao-paulo-1",
					"name":     "São Paulo 1",
					"slug":     "sao-paulo-1",
					"facility": "Equinix SP4",
				},
			},
			"plan": map[string]interface{}{
				"id":   "c3.small.x86",
				"name": "c3.small.x86",
				"slug": "c3-small-x86",
			},
			"specs": map[string]interface{}{
				"cpu":  "Intel Xeon E3 (4 cores)",
				"disk": "2 x 500GB",
				"ram":  "16GB",
				"nic":  "1Gbps",
			},
			"primary_ipv4": "203.0.113.10",
			"primary_ipv6": "2001:db8::1",
			"created_at":   time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
			"updated_at":   time.Now().Format(time.RFC3339),
		},
	}
}

// MockServerList returns a mock list of servers
func MockServerList() map[string]interface{} {
	return map[string]interface{}{
		"data": []interface{}{
			MockServer(),
		},
		"meta": map[string]interface{}{
			"total":        1,
			"current_page": 1,
			"last_page":    1,
			"per_page":     25,
		},
	}
}

// MockError returns a mock error response
func MockError() map[string]interface{} {
	return map[string]interface{}{
		"errors": []map[string]interface{}{
			{
				"status": "404",
				"title":  "Not Found",
				"detail": "Resource not found",
			},
		},
	}
}

// MockValidationError returns a mock validation error response
func MockValidationError() map[string]interface{} {
	return map[string]interface{}{
		"errors": []map[string]interface{}{
			{
				"status": "422",
				"title":  "Validation Error",
				"detail": "Invalid input",
				"source": map[string]interface{}{
					"pointer": "/data/attributes/name",
				},
			},
		},
	}
}

// MockRateLimitError returns a mock rate limit error response
func MockRateLimitError() map[string]interface{} {
	return map[string]interface{}{
		"errors": []map[string]interface{}{
			{
				"status": "429",
				"title":  "Rate Limit Exceeded",
				"detail": "Too many requests",
			},
		},
	}
}

// CreatePaginatedResponse creates a paginated response
func CreatePaginatedResponse(data []interface{}, page, perPage, total int) map[string]interface{} {
	lastPage := (total + perPage - 1) / perPage
	if lastPage < 1 {
		lastPage = 1
	}

	return map[string]interface{}{
		"data": data,
		"meta": map[string]interface{}{
			"total":        total,
			"current_page": page,
			"last_page":    lastPage,
			"per_page":     perPage,
		},
	}
}
