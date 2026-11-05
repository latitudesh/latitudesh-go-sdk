# CreateSandboxSandboxesPreviewAttributes


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ProjectID`                                                      | `string`                                                         | :heavy_check_mark:                                               | ID of the project that owns the sandbox                          |
| `Name`                                                           | `*string`                                                        | :heavy_minus_sign:                                               | Name of the sandbox. Generated when omitted                      |
| `Image`                                                          | `*string`                                                        | :heavy_minus_sign:                                               | Container image to run. The runtime default is used when omitted |
| `Environment`                                                    | map[string]`string`                                              | :heavy_minus_sign:                                               | Environment variables set in the sandbox                         |