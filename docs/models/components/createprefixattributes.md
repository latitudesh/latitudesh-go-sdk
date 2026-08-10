# CreatePrefixAttributes


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ProjectID`                                                                | `string`                                                                   | :heavy_check_mark:                                                         | The project to create the network in                                       |
| `Site`                                                                     | `string`                                                                   | :heavy_check_mark:                                                         | The site slug the network is bound to                                      |
| `Size`                                                                     | [components.CreatePrefixSize](../../models/components/createprefixsize.md) | :heavy_check_mark:                                                         | IPv4 prefix length. Determines how many servers the network can host.      |