# Contributing to Integration Tests

Thank you for contributing to the latitudesh-go-sdk integration tests! This guide will help you add new tests and maintain existing ones.

## 📋 Table of Contents

- [Getting Started](#getting-started)
- [Test Structure](#test-structure)
- [Adding New Tests](#adding-new-tests)
- [Patterns and Best Practices](#patterns-and-best-practices)
- [Running Tests](#running-tests)
- [CI/CD](#cicd)
- [Troubleshooting](#troubleshooting)

## 🚀 Getting Started

### Prerequisites

- Go 1.22 or higher
- Git
- Make (optional, but recommended)

### Initial Setup

```bash
# Clone the repository
git clone https://github.com/latitudesh/latitudesh-go-sdk.git
cd latitudesh-go-sdk

# Install dependencies
go mod download

# Navigate to tests
cd tests/integration

# Run the tests
go test -v ./...
```

## 🏗️ Test Structure

```
tests/integration/
├── helpers/
│   ├── mock_http_client.go  # Mock HTTP client
│   ├── mock_data.go          # Mock data
│   └── test_client.go        # Test setup
├── apikeys_test.go           # API Keys tests
├── plans_test.go             # Plans tests
├── projects_test.go          # Projects tests
├── teams_test.go             # Teams tests
├── userdata_test.go          # User Data tests
├── servers_test.go           # Servers tests
├── README.md                 # Documentation
├── Makefile                  # Useful commands
└── .golangci.yml            # Linter config
```

## ➕ Adding New Tests

### 1. Create Mock Data

Add factory functions in `helpers/mock_data.go`:

```go
// MockNewResource returns a mock of the new resource
func MockNewResource() map[string]interface{} {
    return map[string]interface{}{
        "id":   "resource_test123",
        "type": "new_resources",
        "attributes": map[string]interface{}{
            "name":       "Test Resource",
            "created_at": "2024-01-26T00:00:00Z",
            "updated_at": "2024-01-26T00:00:00Z",
        },
    }
}

// MockNewResourceList returns a mock list
func MockNewResourceList() map[string]interface{} {
    return map[string]interface{}{
        "data": []interface{}{
            MockNewResource(),
        },
        "meta": map[string]interface{}{
            "total":        1,
            "current_page": 1,
            "last_page":    1,
            "per_page":     25,
        },
    }
}
```

### 2. Create Test File

Create `newresource_test.go`:

```go
package integration

import (
    "context"
    "testing"

    "github.com/latitudesh/latitudesh-go-sdk/models/operations"
    "github.com/latitudesh/latitudesh-go-sdk/tests/integration/helpers"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestNewResourceIntegration(t *testing.T) {
    testClient := helpers.SetupTest()

    t.Run("List Operations", func(t *testing.T) {
        t.Run("should list all resources", func(t *testing.T) {
            testClient.Reset()

            // Arrange
            testClient.MockClient.MockResponse("GET", "/new_resources", &helpers.MockResponse{
                Status: 200,
                Body:   helpers.MockNewResourceList(),
            })

            // Act
            result, err := testClient.SDK.NewResources.List(context.Background())

            // Assert
            require.NoError(t, err)
            assert.NotNil(t, result)
            assert.Equal(t, 1, len(result.Data))
        })
    })
}
```

### 3. Add to Makefile

Add a specific target:

```makefile
test-newresource: ## Run NewResource tests only
	@echo "Running NewResource integration tests..."
	@go test -v -run TestNewResourceIntegration
```

## 📝 Patterns and Best Practices

### AAA Pattern (Arrange-Act-Assert)

Always follow this pattern:

```go
t.Run("should do something", func(t *testing.T) {
    testClient.Reset() // Clear state

    // ARRANGE - Setup mock
    testClient.MockClient.MockResponse("GET", "/endpoint", &helpers.MockResponse{
        Status: 200,
        Body:   helpers.MockData(),
    })

    // ACT - Execute action
    result, err := testClient.SDK.Resource.Method(context.Background())

    // ASSERT - Verify result
    require.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, expected, actual)
})
```

### Test Naming

Use descriptive names that explain the behavior:

✅ **GOOD:**
```go
t.Run("should return 404 when resource not found", func(t *testing.T) { })
t.Run("should create resource with valid data", func(t *testing.T) { })
```

❌ **BAD:**
```go
t.Run("test1", func(t *testing.T) { })
t.Run("error_case", func(t *testing.T) { })
```

### Test Organization

Group related tests:

```go
t.Run("List Operations", func(t *testing.T) {
    t.Run("should list all items", func(t *testing.T) { })
    t.Run("should handle empty list", func(t *testing.T) { })
})

t.Run("Create Operations", func(t *testing.T) {
    t.Run("should create with valid data", func(t *testing.T) { })
    t.Run("should fail with invalid data", func(t *testing.T) { })
})
```

### Error Handling

Always test success AND error cases:

```go
// Success test
t.Run("should succeed with valid input", func(t *testing.T) {
    result, err := doSomething()
    require.NoError(t, err)
    assert.NotNil(t, result)
})

// Error test
t.Run("should fail with invalid input", func(t *testing.T) {
    result, err := doSomething()
    assert.Error(t, err)
    assert.Nil(t, result)
})
```

### Data Types

JSON deserializes numbers as `float64`, so:

```go
// ❌ WRONG
assert.Equal(t, int64(4), *result.Count)

// ✅ CORRECT
assert.Equal(t, float64(4), *result.Count)
// OR let Go infer
assert.NotNil(t, result.Count)
```

### State Cleanup

Always clean state before each test:

```go
t.Run("test name", func(t *testing.T) {
    testClient.Reset() // ALWAYS do this first
    // ... rest of test
})
```

## 🧪 Running Tests

### Basic Commands

```bash
# All tests
go test -v ./...

# Specific test
go test -v -run TestPlansIntegration

# Specific group
go test -v -run TestPlansIntegration/List_Operations

# With coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# With race detector
go test -v -race ./...
```

### Using Makefile

```bash
make test                 # All tests
make test-verbose         # With detailed output
make test-coverage        # With coverage report
make test-plans           # Plans only
make test-race            # With race detector
make clean                # Clean artifacts
```

### Advanced Filters

```bash
# Only List tests
go test -v -run ".*List.*"

# Only error tests
go test -v -run ".*error.*" -i

# With timeout
go test -v -timeout 30s ./...

# Repeat tests (detect flakiness)
go test -v -count=10 ./...
```

## 🔄 CI/CD

### GitHub Actions

Tests run automatically:

- **Push to main/develop**: Run all tests
- **Pull Requests**: Run tests + add comment with results
- **Manual**: Via GitHub interface

### Local Checks Before Commit

```bash
# Run what CI will run
make check      # Lint + format + vet
make test-race  # Tests with race detector
```

### Badges

Workflows generate badges that can be added to README:

```markdown
![Tests](https://github.com/latitudesh/latitudesh-go-sdk/actions/workflows/sdk_integration_tests.yaml/badge.svg)
```

## 🐛 Troubleshooting

### Tests failing locally

1. **Clean cache:**
   ```bash
   go clean -testcache
   make clean
   ```

2. **Check dependencies:**
   ```bash
   go mod tidy
   go mod verify
   ```

3. **Run with verbose:**
   ```bash
   go test -v -run FailingTestName
   ```

### Race conditions

If race detector reports issues:

1. **Analyze output:**
   ```bash
   go test -race -v ./... 2>&1 | tee race-log.txt
   ```

2. **Look for:**
   - Unsynchronized concurrent access
   - Shared variable usage
   - Missing mutexes

### Mock not working

Check:

1. **Correct URL pattern:**
   ```go
   // Use exact path
   mockClient.MockResponse("GET", "/resource", ...)
   // Not: "http://api.../resource"
   ```

2. **Mock registered before call:**
   ```go
   mockClient.MockResponse(...)  // First
   sdk.Method(...)               // After
   ```

3. **State was cleared:**
   ```go
   testClient.Reset()  // At test start
   ```

## 📚 Additional Resources

- [Go Testing Package](https://pkg.go.dev/testing)
- [Testify Documentation](https://pkg.go.dev/github.com/stretchr/testify)
- [Table Driven Tests](https://go.dev/wiki/TableDrivenTests)
- [Go Test Comments](https://go.dev/blog/subtests)

## 🤝 Code Review Checklist

Before submitting a PR, verify:

- [ ] Tests follow AAA pattern
- [ ] Test names are descriptive
- [ ] Success AND error cases covered
- [ ] State cleaned with `Reset()`
- [ ] Mocks are correct
- [ ] Tests pass locally with `-race`
- [ ] Lint passes without errors
- [ ] Documentation updated (if needed)
- [ ] Coverage didn't decrease

## 💬 Getting Help

- Open an issue on GitHub
- Consult documentation in `README.md`
- See examples in existing test files
- Ask in the development channel

---

**Thank you for contributing! 🎉**
