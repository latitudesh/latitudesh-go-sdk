// Integration tests for Projects resource
package integration

import (
	"context"
	"testing"

	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"github.com/latitudesh/latitudesh-go-sdk/tests/integration/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectsIntegration(t *testing.T) {
	testClient := helpers.SetupTest()

	t.Run("List Operations", func(t *testing.T) {
		t.Run("should list all projects", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/projects", &helpers.MockResponse{
				Status: 200,
				Body:   helpers.MockProjectList(),
			})

			// Act
			result, err := testClient.SDK.Projects.List(context.Background(), operations.GetProjectsRequest{})

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.NotNil(t, result.Projects.Data)
			assert.Equal(t, 1, len(result.Projects.Data))
			assert.Equal(t, "project_test123", *result.Projects.Data[0].ID)
			assert.Equal(t, "Test Project", *result.Projects.Data[0].Attributes.Name)
		})

		t.Run("should handle empty projects list", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/projects", &helpers.MockResponse{
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
			result, err := testClient.SDK.Projects.List(context.Background(), operations.GetProjectsRequest{})

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, 0, len(result.Projects.Data))
		})
	})

	// Note: Projects API does not support Get single project endpoint.
	// Use List with filters to retrieve a specific project if needed.

	t.Run("Create Operations", func(t *testing.T) {
		t.Run("should create a new project", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			newProject := map[string]interface{}{
				"id":   "project_new123",
				"type": "projects",
				"attributes": map[string]interface{}{
					"name":        "New Project",
					"slug":        "new-project",
					"description": "A new project",
					"environment": "Production",
					"created_at":  "2024-01-26T00:00:00Z",
					"updated_at":  "2024-01-26T00:00:00Z",
				},
			}

			testClient.MockClient.MockResponse("POST", "/projects", &helpers.MockResponse{
				Status: 201,
				Body: map[string]interface{}{
					"data": newProject,
				},
			})

			// Act
			result, err := testClient.SDK.Projects.Create(context.Background(), operations.CreateProjectProjectsRequestBody{
				Data: &operations.CreateProjectProjectsData{
					Type: operations.CreateProjectProjectsTypeProjects,
					Attributes: &operations.CreateProjectProjectsAttributes{
						Name:             "New Project",
						ProvisioningType: operations.CreateProjectProvisioningTypeOnDemand,
						Description:      latitudeshgosdk.String("A new project"),
						Environment:      operations.CreateProjectEnvironmentProduction.ToPointer(),
					},
				},
			})

			// Assert
			require.NoError(t, err)
			assert.Equal(t, "project_new123", *result.Object.Data.ID)
			assert.Equal(t, "New Project", *result.Object.Data.Attributes.Name)
		})
	})

	t.Run("Update Operations", func(t *testing.T) {
		t.Run("should update project settings", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			updatedProjectData := helpers.MockProject()
			attrs := updatedProjectData["attributes"].(map[string]interface{})
			attrs["name"] = "Updated Project"
			attrs["description"] = "Updated description"

			testClient.MockClient.MockResponse("PATCH", "/projects/project_test123", &helpers.MockResponse{
				Status: 200,
				Body: map[string]interface{}{
					"data": updatedProjectData,
				},
			})

			// Act
			result, err := testClient.SDK.Projects.Update(context.Background(), "project_test123", &operations.UpdateProjectProjectsRequestBody{
				Data: operations.UpdateProjectProjectsData{
					Type: operations.UpdateProjectProjectsTypeProjects,
					Attributes: &operations.UpdateProjectProjectsAttributes{
						Name:        latitudeshgosdk.String("Updated Project"),
						Description: latitudeshgosdk.String("Updated description"),
					},
				},
			})

			// Assert
			require.NoError(t, err)
			assert.Equal(t, "Updated Project", *result.Object.Data.Attributes.Name)
		})
	})

	t.Run("Delete Operations", func(t *testing.T) {
		t.Run("should delete a project", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("DELETE", "/projects/project_test123", &helpers.MockResponse{
				Status: 204,
				Body:   nil,
			})

			// Act
			_, err := testClient.SDK.Projects.Delete(context.Background(), "project_test123")

			// Assert
			require.NoError(t, err)
			assert.True(t, testClient.MockClient.VerifyRequest("DELETE", "/projects/project_test123"))
		})

		t.Run("should handle not found error on delete", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("DELETE", "/projects/nonexistent", &helpers.MockResponse{
				Status: 404,
				Body:   helpers.MockError(),
			})

			// Act
			_, err := testClient.SDK.Projects.Delete(context.Background(), "nonexistent")

			// Assert
			assert.Error(t, err)
		})
	})
}
