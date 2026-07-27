# BucketPermissions


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `BucketID`                                                             | `string`                                                               | :heavy_check_mark:                                                     | Bucket (object storage) ID to grant access to.                         |
| `Permission`                                                           | [operations.Permission](../../models/operations/permission.md)         | :heavy_check_mark:                                                     | `readonly` grants read-only access; `rw` grants read and write access. |