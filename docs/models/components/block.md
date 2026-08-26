# Block

NVMe-TCP block mapping of a high performance volume. Null for volumes that are not mapped to a server.


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Nqn`                                     | `*string`                                 | :heavy_minus_sign:                        | NVMe Qualified Name of the mapped server. |
| `Nsid`                                    | `*int64`                                  | :heavy_minus_sign:                        | NVMe namespace ID of the mapping.         |
| `ServerID`                                | `*string`                                 | :heavy_minus_sign:                        | ID of the server the volume is mapped to. |