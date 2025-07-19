# Status

`on` - The server is powered ON
`off` - The server is powered OFF
`unknown` - The server power status is unknown
`ready` - The server is in reinstalling state `ready` and should start `disk_erasing` shortly
`disk_erasing` - The server is in reinstalling state `disk_erasing`
`failed_disk_erasing` - The server has failed disk erasing in reinstall
`deploying` - The server is in the last reinstalling stage and is `deploying`
`failed_deployment` - The server has failed deployment in reinstall



## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `StatusOn`                | on                        |
| `StatusOff`               | off                       |
| `StatusUnknown`           | unknown                   |
| `StatusReady`             | ready                     |
| `StatusDiskErasing`       | disk_erasing              |
| `StatusFailedDiskErasing` | failed_disk_erasing       |
| `StatusDeploying`         | deploying                 |
| `StatusFailedDeployment`  | failed_deployment         |