# Attributes


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Name`                                                          | **string*                                                       | :heavy_minus_sign:                                              | Name of the API Key                                             |
| `APIVersion`                                                    | **string*                                                       | :heavy_minus_sign:                                              | The API version associated with this API Key                    |
| `TokenLastSlice`                                                | **string*                                                       | :heavy_minus_sign:                                              | The last 5 characters of the token created for this API Key     |
| `LastUsedAt`                                                    | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | The last time a request was made to the API using this API Key  |
| `User`                                                          | [*components.APIKeyUser](../../models/components/apikeyuser.md) | :heavy_minus_sign:                                              | The owner of the API Key                                        |
| `CreatedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | The time when the API Key was created                           |
| `UpdatedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | The time when the API Key was updated                           |