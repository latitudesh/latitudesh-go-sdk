# ServerDataPrefix

**Preview** (`prefixes_api` feature flag). The customer prefix this server is bonded onto, or null. Fetch full details from GET /prefixes/{id}.


## Fields

| Field              | Type               | Required           | Description        | Example            |
| ------------------ | ------------------ | ------------------ | ------------------ | ------------------ |
| `ID`               | `*string`          | :heavy_minus_sign: | N/A                | pfx_2aBcDeFgH      |
| `Ipv4`             | `*string`          | :heavy_minus_sign: | N/A                | 10.90.0.0/26       |
| `Ipv6`             | `*string`          | :heavy_minus_sign: | N/A                | 2001:db8::/64      |