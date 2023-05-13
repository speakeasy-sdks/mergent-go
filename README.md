<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/228782760-eae443cd-8197-4667-91f5-656d90633db1.svg" media="(prefers-color-scheme: dark)" width="100">
        <img src="https://user-images.githubusercontent.com/6267663/228782746-c33cbcc8-4767-42c8-9a50-af6552703414.svg" width="100">
    </picture>
    <h1>Mergent Go SDK</h1>
   <p>Developer infrastructure for internal tools</p>
   <a href="https://docs.mergent.co/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/mergent-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/mergent-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/mergent-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/mergent-go?sort=semver&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/mergent-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/mergent-go"
	"github.com/speakeasy-sdks/mergent-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mergent-go/pkg/types"
)

func main() {
    s := mergent.New(
        mergent.WithSecurity(shared.Security{
            APIKey: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Tasks.Create(ctx, shared.TaskNewInput{
        Delay: mergent.String("60s"),
        Name: mergent.String("mytask"),
        Queue: mergent.String("process1"),
        Request: shared.Request{
            Body: mergent.String("Body String"),
            Headers: map[string]interface{}{
                "provident": "distinctio",
                "quibusdam": "unde",
                "nulla": "corrupti",
            },
            URL: "http://example.com",
        },
        ScheduledFor: types.MustTimeFromString("2021-10-01T15:53:05Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Task != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Schedules](docs/schedules/README.md)

* [Create](docs/schedules/README.md#create) - Create Schedule
* [Delete](docs/schedules/README.md#delete) - Delete Schedule
* [Get](docs/schedules/README.md#get) - Get Schedule
* [GetTasks](docs/schedules/README.md#gettasks) - Get Schedule Tasks
* [List](docs/schedules/README.md#list) - Get Schedules
* [Update](docs/schedules/README.md#update) - Update Schedule

### [Tasks](docs/tasks/README.md)

* [BatchCreate](docs/tasks/README.md#batchcreate) - Batch Create Tasks (Beta)
* [BatchDelete](docs/tasks/README.md#batchdelete) - Batch Delete Tasks (Beta)
* [Create](docs/tasks/README.md#create) - Create Task
* [Delete](docs/tasks/README.md#delete) - Delete Task
* [Get](docs/tasks/README.md#get) - Get Task
* [List](docs/tasks/README.md#list) - Get Tasks
* [Run](docs/tasks/README.md#run) - Run Task
* [Update](docs/tasks/README.md#update) - Update Task
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
