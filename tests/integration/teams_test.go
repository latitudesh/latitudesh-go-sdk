// Integration tests for Teams resource
package integration

import (
	"context"
	"testing"

	"github.com/latitudesh/latitudesh-go-sdk/tests/integration/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeamsIntegration(t *testing.T) {
	testClient := helpers.SetupTest()

	t.Run("List Operations", func(t *testing.T) {
		t.Run("should list all teams", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/team", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockTeamList(),
			})

			// Act
			result, err := testClient.SDK.Teams.Get(context.Background())

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.NotNil(t, result.Teams.Data)
			assert.Equal(t, 1, len(result.Teams.Data))
			assert.Equal(t, "team_test123", *result.Teams.Data[0].ID)
			assert.Equal(t, "Test Team", *result.Teams.Data[0].Attributes.Name)
		})

		t.Run("should handle empty teams list", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/team", &helpers.MockResponse{
				Status: 200,
				Body: map[string]interface{}{
					"data": []interface{}{},
				},
			})

			// Act
			result, err := testClient.SDK.Teams.Get(context.Background())

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, 0, len(result.Teams.Data))
		})

		t.Run("should include team information with all fields", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/team", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockTeamList(),
			})

			// Act
			result, err := testClient.SDK.Teams.Get(context.Background())

			// Assert
			require.NoError(t, err)
			team := result.Teams.Data[0]
			assert.NotNil(t, team.Attributes.Name)
			assert.NotNil(t, team.Attributes.Slug)
			// Note: Address is a string field, not an object
			assert.NotNil(t, team.Attributes.Address)
		})
	})

	// Note: Teams.Get() returns current team information.
	// There is no separate GetCurrent method - use Get() instead.
}
