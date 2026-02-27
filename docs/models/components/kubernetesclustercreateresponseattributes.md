# KubernetesClusterCreateResponseAttributes


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Name`                                                 | **string*                                              | :heavy_minus_sign:                                     | The cluster name                                       |
| `Status`                                               | **string*                                              | :heavy_minus_sign:                                     | The cluster status (always 'provisioning' on creation) |
| `ControlPlaneEndpoint`                                 | **string*                                              | :heavy_minus_sign:                                     | The URL endpoint for the Kubernetes API server         |