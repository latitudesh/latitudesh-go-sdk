# CreateBgpSessionAttributes


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ServerID`                                                  | `string`                                                    | :heavy_check_mark:                                          | The server that will announce the elastic IP over BGP       |
| `RequestorID`                                               | `*string`                                                   | :heavy_minus_sign:                                          | Optional identifier of the requestor, recorded for auditing |