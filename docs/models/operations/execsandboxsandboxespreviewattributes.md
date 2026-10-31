# ExecSandboxSandboxesPreviewAttributes


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Command`                                               | []`string`                                              | :heavy_check_mark:                                      | Command and arguments to run                            |
| `Environment`                                           | map[string]`string`                                     | :heavy_minus_sign:                                      | Environment variables set for the command               |
| `Workdir`                                               | `*string`                                               | :heavy_minus_sign:                                      | Working directory for the command                       |
| `TimeoutSeconds`                                        | `*int64`                                                | :heavy_minus_sign:                                      | Maximum run time of the command. Defaults to 30 seconds |