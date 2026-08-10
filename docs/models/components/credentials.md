# Credentials

S3 access credentials. Only included when `extra_fields[object_storages]=credentials` is requested and the requesting user is the bucket's creator.


## Fields

| Field                            | Type                             | Required                         | Description                      |
| -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- |
| `AccessKey`                      | `*string`                        | :heavy_minus_sign:               | S3 access key for authentication |