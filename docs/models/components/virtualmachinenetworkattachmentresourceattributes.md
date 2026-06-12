# VirtualMachineNetworkAttachmentResourceAttributes


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `VirtualNetworkID`                                          | `*string`                                                   | :heavy_minus_sign:                                          | The encoded VLAN id_hash                                    |
| `Vid`                                                       | `*int64`                                                    | :heavy_minus_sign:                                          | The 802.1Q VLAN ID                                          |
| `PendingRestart`                                            | `*bool`                                                     | :heavy_minus_sign:                                          | True if the attachment requires a VM restart to take effect |