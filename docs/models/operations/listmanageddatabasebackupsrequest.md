# ListManagedDatabaseBackupsRequest


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ManagedDatabaseID`                                                       | `string`                                                                  | :heavy_check_mark:                                                        | Managed database ID                                                       |
| `Phase`                                                                   | `*string`                                                                 | :heavy_minus_sign:                                                        | Filter backups by phase. Use 'completed' to return only finished backups. |