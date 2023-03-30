# github.com/speakeasy-sdks/mergent-go

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

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/mergent-go"
    "github.com/speakeasy-sdks/mergent-go/pkg/models/shared"
    "github.com/speakeasy-sdks/mergent-go/pkg/models/operations"
)

func main() {
    s := mergent.New(
        mergent.WithSecurity(shared.Security{
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    req := shared.TaskNewInput{
        Delay: "60s",
        Name: "mytask",
        Queue: "process1",
        Request: shared.Request{
            Body: "Body String",
            Headers: map[string]interface{}{
                "deserunt": "porro",
                "nulla": "id",
                "vero": "perspiciatis",
            },
            URL: "http://example.com",
        },
        ScheduledFor: "2021-10-01T15:53:05Z",
    }

    ctx := context.Background()
    res, err := s.Tasks.Create(ctx, req)
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
## SDK Available Operations


### Schedules

* `Create` - Create Schedule
* `Delete` - Delete Schedule
* `Get` - Get Schedule
* `GetTasks` - Get Schedule Tasks
* `List` - Get Schedules
* `Update` - Update Schedule

### Tasks

* `BatchCreate` - Batch Create Tasks (Beta)
* `BatchDelete` - Batch Delete Tasks (Beta)
* `Create` - Create Task
* `Delete` - Delete Task
* `Get` - Get Task
* `List` - Get Tasks
* `Run` - Run Task
* `Update` - Update Task
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
