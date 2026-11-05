# VirtualMachineAttributesFeatures

Features supported by this operating system


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Raid`                                                              | `*bool`                                                             | :heavy_minus_sign:                                                  | Whether RAID is supported                                           |
| `SSHKeys`                                                           | `*bool`                                                             | :heavy_minus_sign:                                                  | Whether SSH keys are supported                                      |
| `UserData`                                                          | `*bool`                                                             | :heavy_minus_sign:                                                  | Whether user data is supported                                      |
| `Rescue`                                                            | `*bool`                                                             | :heavy_minus_sign:                                                  | Whether rescue mode is supported                                    |
| `Arm64`                                                             | `*bool`                                                             | :heavy_minus_sign:                                                  | Whether the operating system can be deployed on ARM (arm64) servers |
| `Workflow`                                                          | `*bool`                                                             | :heavy_minus_sign:                                                  | Whether workflow is supported                                       |