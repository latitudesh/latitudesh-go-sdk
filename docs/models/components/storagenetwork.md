# StorageNetwork

Storage network to configure on the server before mounting the filesystem. Null until the status is "attached".


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Vid`                                                      | `*int64`                                                   | :heavy_minus_sign:                                         | VLAN ID of the storage VLAN.                               |
| `HostCidr`                                                 | `*string`                                                  | :heavy_minus_sign:                                         | Storage IP of the server, in CIDR notation.                |
| `Gateway`                                                  | `*string`                                                  | :heavy_minus_sign:                                         | Gateway of the storage network, used for the routes below. |
| `Routes`                                                   | []`string`                                                 | :heavy_minus_sign:                                         | Storage infrastructure prefixes to route via the gateway.  |