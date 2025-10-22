# TeamLimits

Resource limits for the team. Null values indicate limits are not yet implemented/enforced.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `BareMetal`                                                           | **int64*                                                              | :heavy_minus_sign:                                                    | Maximum number of bare metal servers allowed                          |
| `BareMetalGpu`                                                        | **int64*                                                              | :heavy_minus_sign:                                                    | Maximum number of bare metal GPU servers allowed                      |
| `VirtualMachine`                                                      | **int64*                                                              | :heavy_minus_sign:                                                    | Maximum number of virtual machines allowed (not yet implemented)      |
| `VirtualMachineGpu`                                                   | **int64*                                                              | :heavy_minus_sign:                                                    | Maximum number of virtual machine GPUs allowed                        |
| `Database`                                                            | **int64*                                                              | :heavy_minus_sign:                                                    | Maximum number of databases allowed (not yet implemented)             |
| `Filesystem`                                                          | **int64*                                                              | :heavy_minus_sign:                                                    | Maximum number of filesystems allowed (not yet implemented)           |
| `BlockStorage`                                                        | **int64*                                                              | :heavy_minus_sign:                                                    | Maximum number of block storage volumes allowed (not yet implemented) |