# UpdateAPIKeyAttributes


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Name`                                                                             | **string*                                                                          | :heavy_minus_sign:                                                                 | Name of the API Key                                                                |
| `ReadOnly`                                                                         | **bool*                                                                            | :heavy_minus_sign:                                                                 | Whether the API Key is read-only. Read-only keys can only perform GET requests.    |
| `AllowedIps`                                                                       | []*string*                                                                         | :heavy_minus_sign:                                                                 | List of allowed IP addresses or CIDR ranges (e.g., "192.168.1.100", "10.0.0.0/24") |