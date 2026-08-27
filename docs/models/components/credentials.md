# Credentials

Latest provisioning credentials, lazy loaded. Request it with `extra_fields[servers]=credentials`. Empty when the latest provisioning has no valid credentials.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Username`                                                                     | `*string`                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            |
| `Password`                                                                     | `*string`                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            |
| `SSHKeys`                                                                      | [][components.ServerDataSSHKeys](../../models/components/serverdatasshkeys.md) | :heavy_minus_sign:                                                             | N/A                                                                            |
| `ExpiresAt`                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                     | :heavy_minus_sign:                                                             | N/A                                                                            |