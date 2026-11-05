# SandboxExecAttributes


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Stdout`                                                           | `*string`                                                          | :heavy_minus_sign:                                                 | Standard output, capped at 1 MiB                                   |
| `Stderr`                                                           | `*string`                                                          | :heavy_minus_sign:                                                 | Standard error, capped at 1 MiB                                    |
| `ExitCode`                                                         | `*int64`                                                           | :heavy_minus_sign:                                                 | Exit code of the command. Null when the command did not report one |
| `Truncated`                                                        | `*bool`                                                            | :heavy_minus_sign:                                                 | Whether stdout or stderr exceeded the cap and was truncated        |