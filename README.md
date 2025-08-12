# latitudesh-go-sdk
<!-- Start Summary [summary] -->
## Summary

Latitude.sh API: The Latitude.sh API is a RESTful API to manage your Latitude.sh account. It allows you to perform the same actions as the Latitude.sh dashboard.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [latitudesh-go-sdk](#latitudesh-go-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/latitudesh/latitudesh-go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := latitudeshgosdk.New(
		latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
	)

	res, err := s.APIKeys.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKey != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  | Environment Variable |
| -------- | ------ | ------- | -------------------- |
| `Bearer` | apiKey | API key | `LATITUDESH_BEARER`  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := latitudeshgosdk.New(
		latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
	)

	res, err := s.APIKeys.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKey != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [APIKeys](docs/sdks/apikeys/README.md)

* [List](docs/sdks/apikeys/README.md#list) - List API Keys
* [Create](docs/sdks/apikeys/README.md#create) - Create API Key
* [Update](docs/sdks/apikeys/README.md#update) - Regenerate API Key
* [Delete](docs/sdks/apikeys/README.md#delete) - Delete API Key

### [Billing](docs/sdks/billing/README.md)

* [ListUsage](docs/sdks/billing/README.md#listusage) - List Billing Usage

### [Events](docs/sdks/events/README.md)

* [List](docs/sdks/events/README.md#list) - List all Events

### [Firewalls](docs/sdks/firewalls/README.md)

* [Create](docs/sdks/firewalls/README.md#create) - Create a firewall
* [List](docs/sdks/firewalls/README.md#list) - List firewalls
* [Get](docs/sdks/firewalls/README.md#get) - Retrieve Firewall
* [Update](docs/sdks/firewalls/README.md#update) - Update Firewall
* [Delete](docs/sdks/firewalls/README.md#delete) - Delete Firewall
* [ListAssignments](docs/sdks/firewalls/README.md#listassignments) - Firewall Assignments
* [DeleteAssignment](docs/sdks/firewalls/README.md#deleteassignment) - Delete Firewall Assignment

#### [Firewalls.Assignments](docs/sdks/assignments/README.md)

* [Create](docs/sdks/assignments/README.md#create) - Firewall Assignment

### [IPAddresses](docs/sdks/ipaddresses/README.md)

* [List](docs/sdks/ipaddresses/README.md#list) - List IPs
* [Get](docs/sdks/ipaddresses/README.md#get) - Retrieve an IP


### [OperatingSystems](docs/sdks/operatingsystems/README.md)

* [ListPlans](docs/sdks/operatingsystems/README.md#listplans) - List all operating systems available

### [Plans](docs/sdks/plans/README.md)

* [List](docs/sdks/plans/README.md#list) - List all Plans
* [Get](docs/sdks/plans/README.md#get) - Retrieve a Plan
* [GetBandwidth](docs/sdks/plans/README.md#getbandwidth) - List all bandwidth plans
* [UpdateBandwidth](docs/sdks/plans/README.md#updatebandwidth) - Buy or remove bandwidth packages
* [GetContainersPlans](docs/sdks/plans/README.md#getcontainersplans) - List containers plans
* [GetContainersPlan](docs/sdks/plans/README.md#getcontainersplan) - Retrieve container plan
* [ListStorage](docs/sdks/plans/README.md#liststorage) - List all Storage Plans

#### [Plans.VM](docs/sdks/vm/README.md)

* [List](docs/sdks/vm/README.md#list) - List all Virtual Machines Plans

### [PrivateNetworks](docs/sdks/privatenetworks/README.md)

* [List](docs/sdks/privatenetworks/README.md#list) - List all Virtual Networks
* [Create](docs/sdks/privatenetworks/README.md#create) - Create a Virtual Network
* [Update](docs/sdks/privatenetworks/README.md#update) - Update a Virtual Network
* [Get](docs/sdks/privatenetworks/README.md#get) - Retrieve a Virtual Network
* [ListAssignments](docs/sdks/privatenetworks/README.md#listassignments) - List all servers assigned to virtual networks
* [Assign](docs/sdks/privatenetworks/README.md#assign) - Assign Virtual network
* [DeleteAssignment](docs/sdks/privatenetworks/README.md#deleteassignment) - Delete Virtual Network Assignment

### [Projects](docs/sdks/projects/README.md)

* [List](docs/sdks/projects/README.md#list) - List all Projects
* [Create](docs/sdks/projects/README.md#create) - Create a Project
* [Update](docs/sdks/projects/README.md#update) - Update a Project
* [Delete](docs/sdks/projects/README.md#delete) - Delete a Project

#### [~~Projects.SSHKeys~~](docs/sdks/latitudeshprojectssshkeys/README.md)

* [~~PostProjectSSHKey~~](docs/sdks/latitudeshprojectssshkeys/README.md#postprojectsshkey) - Create a Project SSH Key :warning: **Deprecated**

### [Regions](docs/sdks/regions/README.md)

* [Get](docs/sdks/regions/README.md#get) - List all Regions
* [Fetch](docs/sdks/regions/README.md#fetch) - Retrieve a Region

### [Roles](docs/sdks/roles/README.md)

* [List](docs/sdks/roles/README.md#list) - List all Roles
* [Get](docs/sdks/roles/README.md#get) - Retrieve Role

### [Servers](docs/sdks/servers/README.md)

* [List](docs/sdks/servers/README.md#list) - List all Servers
* [Create](docs/sdks/servers/README.md#create) - Deploy Server
* [Get](docs/sdks/servers/README.md#get) - Retrieve a Server
* [Update](docs/sdks/servers/README.md#update) - Update Server
* [Delete](docs/sdks/servers/README.md#delete) - Remove Server
* [GetDeployConfig](docs/sdks/servers/README.md#getdeployconfig) - Retrieve Deploy Config
* [UpdateDeployConfig](docs/sdks/servers/README.md#updatedeployconfig) - Update Deploy Config
* [Lock](docs/sdks/servers/README.md#lock) - Lock the server
* [Unlock](docs/sdks/servers/README.md#unlock) - Unlock the server
* [StartOutOfBandConnection](docs/sdks/servers/README.md#startoutofbandconnection) - Start Out of Band Connection
* [GetOutOfBand](docs/sdks/servers/README.md#getoutofband) - List Out of Band Connections
* [RunAction](docs/sdks/servers/README.md#runaction) - Run Server Action
* [CreateIpmiSession](docs/sdks/servers/README.md#createipmisession) - Generate IPMI credentials
* [StartRescueMode](docs/sdks/servers/README.md#startrescuemode) - Puts a Server in rescue mode
* [ExitRescueMode](docs/sdks/servers/README.md#exitrescuemode) - Exits rescue mode for a Server
* [ScheduleDeletion](docs/sdks/servers/README.md#scheduledeletion) - Schedule the server deletion
* [UnscheduleDeletion](docs/sdks/servers/README.md#unscheduledeletion) - Unschedule the server deletion
* [Reinstall](docs/sdks/servers/README.md#reinstall) - Run Server Reinstall

### [SSHKeys](docs/sdks/sshkeys/README.md)

* [~~List~~](docs/sdks/sshkeys/README.md#list) - List all Project SSH Keys :warning: **Deprecated**
* [~~Get~~](docs/sdks/sshkeys/README.md#get) - Retrieve a Project SSH Key :warning: **Deprecated**
* [~~ModifyProjectKey~~](docs/sdks/sshkeys/README.md#modifyprojectkey) - Update a Project SSH Key :warning: **Deprecated**
* [~~RemoveFromProject~~](docs/sdks/sshkeys/README.md#removefromproject) - Delete a Project SSH Key :warning: **Deprecated**
* [ListAll](docs/sdks/sshkeys/README.md#listall) - List all SSH Keys
* [Create](docs/sdks/sshkeys/README.md#create) - Create a SSH Key
* [Retrieve](docs/sdks/sshkeys/README.md#retrieve) - Retrieve a SSH Key
* [Update](docs/sdks/sshkeys/README.md#update) - Update a SSH Key
* [Delete](docs/sdks/sshkeys/README.md#delete) - Delete a SSH Key

### [Storage](docs/sdks/storage/README.md)

* [CreateFilesystem](docs/sdks/storage/README.md#createfilesystem) - Create a filesystem for a project
* [ListFilesystems](docs/sdks/storage/README.md#listfilesystems) - List filesystems
* [DeleteFilesystem](docs/sdks/storage/README.md#deletefilesystem) - Delete a filesystem for a project
* [UpdateFilesystem](docs/sdks/storage/README.md#updatefilesystem) - Update a filesystem for a project

### [Tags](docs/sdks/tags/README.md)

* [List](docs/sdks/tags/README.md#list) - List all Tags
* [Create](docs/sdks/tags/README.md#create) - Create a Tag
* [Update](docs/sdks/tags/README.md#update) - Update Tag
* [Delete](docs/sdks/tags/README.md#delete) - Delete Tag

### [TeamMembers](docs/sdks/teammembers/README.md)

* [PostTeamMembers](docs/sdks/teammembers/README.md#postteammembers) - Add a Team Member
* [Delete](docs/sdks/teammembers/README.md#delete) - Remove a Team Member

### [Teams](docs/sdks/teams/README.md)

* [Get](docs/sdks/teams/README.md#get) - Retrieve the team
* [Create](docs/sdks/teams/README.md#create) - Create a team
* [Update](docs/sdks/teams/README.md#update) - Update a team

#### [Teams.Members](docs/sdks/members/README.md)

* [GetTeamMembers](docs/sdks/members/README.md#getteammembers) - List all Team Members

### [Traffic](docs/sdks/traffic/README.md)

* [Get](docs/sdks/traffic/README.md#get) - Retrieve Traffic consumption
* [GetQuota](docs/sdks/traffic/README.md#getquota) - Retrieve Traffic Quota

### [UserData](docs/sdks/userdata/README.md)

* [~~GetProjectUsersData~~](docs/sdks/userdata/README.md#getprojectusersdata) - List all Project User Data :warning: **Deprecated**
* [~~GetProjectUserData~~](docs/sdks/userdata/README.md#getprojectuserdata) - Retrieve a Project User Data :warning: **Deprecated**
* [~~DeleteProjectUserData~~](docs/sdks/userdata/README.md#deleteprojectuserdata) - Delete a Project User Data :warning: **Deprecated**
* [~~Create~~](docs/sdks/userdata/README.md#create) - Create a Project User Data :warning: **Deprecated**
* [~~UpdateForProject~~](docs/sdks/userdata/README.md#updateforproject) - Update a Project User Data :warning: **Deprecated**
* [List](docs/sdks/userdata/README.md#list) - List all User Data
* [CreateNew](docs/sdks/userdata/README.md#createnew) - Create an User Data
* [Retrieve](docs/sdks/userdata/README.md#retrieve) - Retrieve an User Data
* [Update](docs/sdks/userdata/README.md#update) - Update an User Data
* [Delete](docs/sdks/userdata/README.md#delete) - Delete an User Data

### [UserProfile](docs/sdks/userprofile/README.md)

* [Get](docs/sdks/userprofile/README.md#get) - Get user profile
* [Update](docs/sdks/userprofile/README.md#update) - Update User Profile
* [ListTeams](docs/sdks/userprofile/README.md#listteams) - List User Teams

### [VirtualMachines](docs/sdks/virtualmachines/README.md)

* [Create](docs/sdks/virtualmachines/README.md#create) - Create a Virtual Machine
* [List](docs/sdks/virtualmachines/README.md#list) - Get Teams Virtual Machines
* [Get](docs/sdks/virtualmachines/README.md#get) - Get a Virtual Machine
* [Delete](docs/sdks/virtualmachines/README.md#delete) - Destroy a Virtual Machine

### [VirtualNetworks](docs/sdks/virtualnetworks/README.md)

* [Delete](docs/sdks/virtualnetworks/README.md#delete) - Delete a Virtual Network

### [VpnSessions](docs/sdks/vpnsessions/README.md)

* [List](docs/sdks/vpnsessions/README.md#list) - List all Active VPN Sessions
* [Create](docs/sdks/vpnsessions/README.md#create) - Create a VPN Session
* [RefreshPassword](docs/sdks/vpnsessions/README.md#refreshpassword) - Refresh a VPN Session
* [Delete](docs/sdks/vpnsessions/README.md#delete) - Delete a VPN Session

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := latitudeshgosdk.New(
		latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
	)

	res, err := s.Events.List(ctx, operations.GetEventsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := latitudeshgosdk.New(
		latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
	)

	res, err := s.APIKeys.List(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKey != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := latitudeshgosdk.New(
		latitudeshgosdk.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
	)

	res, err := s.APIKeys.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKey != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `components.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `List` function may return the following errors:

| Error Type          | Status Code | Content Type |
| ------------------- | ----------- | ------------ |
| components.APIError | 4XX, 5XX    | \*/\*        |

### Example

```go
package main

import (
	"context"
	"errors"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := latitudeshgosdk.New(
		latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
	)

	res, err := s.APIKeys.List(ctx)
	if err != nil {

		var e *components.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                    | Variables          | Description |
| --- | ------------------------- | ------------------ | ----------- |
| 0   | `https://api.latitude.sh` | `latitude_api_key` |             |
| 1   | `http://api.latitude.sh`  | `latitude_api_key` |             |

If the selected server has variables, you may override its default values using the associated option(s):

| Variable           | Option                                      | Default                        | Description |
| ------------------ | ------------------------------------------- | ------------------------------ | ----------- |
| `latitude_api_key` | `WithLatitudeAPIKey(latitudeAPIKey string)` | `"<insert your api key here>"` |             |

#### Example

```go
package main

import (
	"context"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := latitudeshgosdk.New(
		latitudeshgosdk.WithServerIndex(1),
		latitudeshgosdk.WithLatitudeAPIKey("<value>"),
		latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
	)

	res, err := s.APIKeys.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKey != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := latitudeshgosdk.New(
		latitudeshgosdk.WithServerURL("http://api.latitude.sh"),
		latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
	)

	res, err := s.APIKeys.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKey != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/latitudesh/latitudesh-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = latitudeshgosdk.New(latitudeshgosdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
