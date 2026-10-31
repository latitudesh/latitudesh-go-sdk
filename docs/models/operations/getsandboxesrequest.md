# GetSandboxesRequest


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ProjectID`                                                   | `string`                                                      | :heavy_check_mark:                                            | ID of the project whose sandboxes are listed                  |
| `PageToken`                                                   | `*string`                                                     | :heavy_minus_sign:                                            | Token returned in `meta.next_page_token` by the previous page |