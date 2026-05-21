# ElasticIPDataServer

The server this Elastic IP is assigned to. Returns null when the Elastic IP is not assigned to a server or when the assigned server is not active (e.g., decommissioning or deleted).


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `ID`               | `*string`          | :heavy_minus_sign: | N/A                |
| `Hostname`         | `*string`          | :heavy_minus_sign: | N/A                |
| `PrimaryIpv4`      | `*string`          | :heavy_minus_sign: | N/A                |
| `OperatingSystem`  | `*string`          | :heavy_minus_sign: | N/A                |