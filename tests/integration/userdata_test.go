// Integration tests for UserData resource
package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"github.com/latitudesh/latitudesh-go-sdk/tests/integration/helpers"
)

func TestUserDataIntegration(t *testing.T) {
	testClient := helpers.SetupTest()

	t.Run("List Operations", func(t *testing.T) {
		t.Run("should list all user data", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/projects/project_test123/user_data", &helpers.MockResponse{
				Body:   helpers.MockUserDataList(),
				Status: 200,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			result, err := testClient.SDK.UserData.GetProjectUsersData(context.Background(), "project_test123", nil)

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.NotNil(t, result.UserData)
			// Note: The SDK's UserData structure has an unusual nested format
			// We verify basic structure without drilling into nested fields
			assert.NotNil(t, result.UserData.Data)
		})

		t.Run("should handle empty user data list", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/projects/project_test123/user_data", &helpers.MockResponse{
				Body: map[string]interface{}{
					"data": []interface{}{},
					"meta": map[string]interface{}{
						"total":        0,
						"current_page": 1,
						"last_page":    1,
						"per_page":     25,
					},
				},
				Status: 200,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			result, err := testClient.SDK.UserData.GetProjectUsersData(context.Background(), "project_test123", nil)

			// Assert
			require.NoError(t, err)
			assert.Equal(t, 0, len(result.UserData.Data))
		})
	})

	t.Run("Get Single User Data", func(t *testing.T) {
		t.Run("should retrieve a specific user data by ID", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/projects/project_test123/user_data/userdata_test123", &helpers.MockResponse{
				Body: map[string]interface{}{
					"data": helpers.MockUserData(),
				},
				Status: 200,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			result, err := testClient.SDK.UserData.GetProjectUserData(context.Background(), "project_test123", "userdata_test123", nil)

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, "userdata_test123", *result.UserDataObject.Data.ID)
			assert.Equal(t, "Test User Data", *result.UserDataObject.Data.Attributes.Description)
			assert.Contains(t, *result.UserDataObject.Data.Attributes.Content, "#!/bin/bash")
		})

		t.Run("should handle not found error", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("GET", "/projects/project_test123/user_data/nonexistent", &helpers.MockResponse{
				Body:   helpers.MockError(),
				Status: 404,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			result, err := testClient.SDK.UserData.GetProjectUserData(context.Background(), "project_test123", "nonexistent", nil)

			// Assert
			assert.Error(t, err)
			assert.Nil(t, result)
		})
	})

	t.Run("Create Operations", func(t *testing.T) {
		t.Run("should create new user data", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			newUserData := map[string]interface{}{
				"id":   "userdata_new123",
				"type": "user_data",
				"attributes": map[string]interface{}{
					"description": "New User Data",
					"content":     "#!/bin/bash\necho 'New Script'",
					"created_at":  "2024-01-26T00:00:00Z",
					"updated_at":  "2024-01-26T00:00:00Z",
				},
			}

			testClient.MockClient.MockResponse("POST", "/projects/project_test123/user_data", &helpers.MockResponse{
				Body: map[string]interface{}{
					"data": newUserData,
				},
				Status: 201,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			result, err := testClient.SDK.UserData.Create(context.Background(), "project_test123", operations.PostProjectUserDataUserDataRequestBody{
				Data: operations.PostProjectUserDataUserDataData{
					Type: operations.PostProjectUserDataUserDataTypeUserData,
					Attributes: &operations.PostProjectUserDataUserDataAttributes{
						Description: "New User Data",
						Content:     "#!/bin/bash\necho 'New Script'",
					},
				},
			})

			// Assert
			require.NoError(t, err)
			assert.Equal(t, "userdata_new123", *result.UserDataObject.Data.ID)
			assert.Equal(t, "New User Data", *result.UserDataObject.Data.Attributes.Description)
			assert.Contains(t, *result.UserDataObject.Data.Attributes.Content, "New Script")
		})

		t.Run("should handle validation error on create", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("POST", "/projects/project_test123/user_data", &helpers.MockResponse{
				Body:   helpers.MockValidationError(),
				Status: 422,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			result, err := testClient.SDK.UserData.Create(context.Background(), "project_test123", operations.PostProjectUserDataUserDataRequestBody{
				Data: operations.PostProjectUserDataUserDataData{
					Type: operations.PostProjectUserDataUserDataTypeUserData,
					Attributes: &operations.PostProjectUserDataUserDataAttributes{
						Description: "",
						Content:     "",
					},
				},
			})

			// Assert
			assert.Error(t, err)
			assert.Nil(t, result)
		})
	})

	t.Run("Update Operations", func(t *testing.T) {
		t.Run("should update user data content", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			updatedUserDataMap := helpers.MockUserData()
			attrs := updatedUserDataMap["attributes"].(map[string]interface{})
			attrs["description"] = "Updated User Data"
			attrs["content"] = "#!/bin/bash\necho 'Updated Script'"

			testClient.MockClient.MockResponse("PATCH", "/projects/project_test123/user_data/userdata_test123", &helpers.MockResponse{
				Body: map[string]interface{}{
					"data": updatedUserDataMap,
				},
				Status: 200,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			result, err := testClient.SDK.UserData.UpdateForProject(context.Background(), "project_test123", "userdata_test123", &operations.PutProjectUserDataUserDataRequestBody{
				Data: operations.PutProjectUserDataUserDataData{
					ID:   "userdata_test123",
					Type: operations.PutProjectUserDataUserDataTypeUserData,
					Attributes: &operations.PutProjectUserDataUserDataAttributes{
						Description: latitudeshgosdk.String("Updated User Data"),
						Content:     latitudeshgosdk.String("#!/bin/bash\necho 'Updated Script'"),
					},
				},
			})

			// Assert
			require.NoError(t, err)
			assert.Equal(t, "Updated User Data", *result.UserDataObject.Data.Attributes.Description)
			assert.Contains(t, *result.UserDataObject.Data.Attributes.Content, "Updated Script")
		})
	})

	t.Run("Delete Operations", func(t *testing.T) {
		t.Run("should delete user data", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("DELETE", "/projects/project_test123/user_data/userdata_test123", &helpers.MockResponse{
				Body:   nil,
				Status: 204,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			_, err := testClient.SDK.UserData.DeleteProjectUserData(context.Background(), "project_test123", "userdata_test123")

			// Assert
			require.NoError(t, err)
			assert.True(t, testClient.MockClient.VerifyRequest("DELETE", "/projects/project_test123/user_data/userdata_test123"))
		})

		t.Run("should handle not found error on delete", func(t *testing.T) {
			testClient.Reset()

			// Arrange
			testClient.MockClient.MockResponse("DELETE", "/projects/project_test123/user_data/nonexistent", &helpers.MockResponse{
				Body:   helpers.MockError(),
				Status: 404,
			})

			// Act
			//nolint:staticcheck // Testing deprecated API
			_, err := testClient.SDK.UserData.DeleteProjectUserData(context.Background(), "project_test123", "nonexistent")

			// Assert
			assert.Error(t, err)
		})
	})
}
