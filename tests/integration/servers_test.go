// Integration tests for Servers resource
package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"github.com/latitudesh/latitudesh-go-sdk/tests/integration/helpers"
)

func TestServersIntegration(t *testing.T) {
	testClient := helpers.SetupTest()

	t.Run("List Operations", func(t *testing.T) {
		t.Run("should list all servers", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockServerList(),
			})

			// Act
			projectID := "project_test123"
			result, err := testClient.SDK.Servers.List(context.Background(), operations.GetServersRequest{
				FilterProject: &projectID,
			})

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.NotNil(t, result.Servers.Data)
			assert.Equal(t, 1, len(result.Servers.Data))
			assert.Equal(t, "server_test123", *result.Servers.Data[0].ID)
			assert.Equal(t, "test-server-01", *result.Servers.Data[0].Attributes.Hostname)
		})

		t.Run("should handle empty servers list", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers", &helpers.MockResponse{
				Status: 200,
				Body: map[string]interface{}{
					"data": []interface{}{},
					"meta": map[string]interface{}{
						"total":        0,
						"current_page": 1,
						"last_page":    1,
						"per_page":     25,
					},
				},
			})

			// Act
			projectID := "project_test123"
			result, err := testClient.SDK.Servers.List(context.Background(), operations.GetServersRequest{
				FilterProject: &projectID,
			})

			// Assert
			require.NoError(t, err)
			assert.Equal(t, 0, len(result.Servers.Data))
		})

		t.Run("should include server specifications", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockServerList(),
			})

			// Act
			projectID := "project_test123"
			result, err := testClient.SDK.Servers.List(context.Background(), operations.GetServersRequest{
				FilterProject: &projectID,
			})

			// Assert
			require.NoError(t, err)
			server := result.Servers.Data[0]
			assert.NotNil(t, server.Attributes.Specs)
			assert.NotNil(t, server.Attributes.Specs.CPU)
			assert.NotNil(t, server.Attributes.Specs.RAM)
		})

		t.Run("should include server plan information", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockServerList(),
			})

			// Act
			projectID := "project_test123"
			result, err := testClient.SDK.Servers.List(context.Background(), operations.GetServersRequest{
				FilterProject: &projectID,
			})

			// Assert
			require.NoError(t, err)
			server := result.Servers.Data[0]
			assert.NotNil(t, server.Attributes.Plan)
			assert.Equal(t, "c3.small.x86", *server.Attributes.Plan.ID)
			assert.Equal(t, "c3-small-x86", *server.Attributes.Plan.Slug)
		})

		t.Run("should include server region information", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockServerList(),
			})

			// Act
			projectID := "project_test123"
			result, err := testClient.SDK.Servers.List(context.Background(), operations.GetServersRequest{
				FilterProject: &projectID,
			})

			// Assert
			require.NoError(t, err)
			server := result.Servers.Data[0]
			assert.NotNil(t, server.Attributes.Region)
			assert.Equal(t, "São Paulo", *server.Attributes.Region.City)
			assert.Equal(t, "BR", *server.Attributes.Region.Country)
		})

		t.Run("should include IP addresses", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockServerList(),
			})

			// Act
			projectID := "project_test123"
			result, err := testClient.SDK.Servers.List(context.Background(), operations.GetServersRequest{
				FilterProject: &projectID,
			})

			// Assert
			require.NoError(t, err)
			server := result.Servers.Data[0]
			assert.NotNil(t, server.Attributes.PrimaryIpv4)
			assert.Equal(t, "203.0.113.10", *server.Attributes.PrimaryIpv4)
			assert.NotNil(t, server.Attributes.PrimaryIpv6)
			assert.Equal(t, "2001:db8::1", *server.Attributes.PrimaryIpv6)
		})
	})

	t.Run("Get Single Server", func(t *testing.T) {
		t.Run("should retrieve a specific server by ID", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers/server_test123", &helpers.MockResponse{
				Status: 200,
				Body: map[string]interface{}{
					"data": helpers.MockServer(),
				},
			})

			// Act
			result, err := testClient.SDK.Servers.Get(context.Background(), "server_test123", nil)

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "server_test123", *result.Server.Data.ID)
			assert.Equal(t, "test-server-01", *result.Server.Data.Attributes.Hostname)
			assert.Equal(t, components.StatusOn, *result.Server.Data.Attributes.Status)
		})

		t.Run("should handle not found error", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers/nonexistent", &helpers.MockResponse{
				Status: 404,
				Body:   helpers.MockError(),
			})

			// Act
			result, err := testClient.SDK.Servers.Get(context.Background(), "nonexistent", nil)

			// Assert
			assert.Error(t, err)
			assert.Nil(t, result)
		})
	})

	t.Run("Filter Operations", func(t *testing.T) {
		t.Run("should filter servers by hostname", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockServerList(),
			})

			// Act
			projectID := "project_test123"
			hostname := "test-server-01"
			result, err := testClient.SDK.Servers.List(context.Background(), operations.GetServersRequest{
				FilterProject:  &projectID,
				FilterHostname: &hostname,
			})

			// Assert
			require.NoError(t, err)
			assert.Equal(t, 1, len(result.Servers.Data))
			assert.Equal(t, "test-server-01", *result.Servers.Data[0].Attributes.Hostname)
		})

		t.Run("should filter servers by status", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/servers", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockServerList(),
			})

			// Act
			projectID := "project_test123"
			status := "on"
			result, err := testClient.SDK.Servers.List(context.Background(), operations.GetServersRequest{
				FilterProject: &projectID,
				FilterStatus:  &status,
			})

			// Assert
			require.NoError(t, err)
			assert.Equal(t, 1, len(result.Servers.Data))
			assert.Equal(t, components.StatusOn, *result.Servers.Data[0].Attributes.Status)
		})
	})
}
