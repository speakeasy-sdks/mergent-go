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
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            Basic: &shared.SchemeBasic{
                Password: "YOUR_PASSWORD_HERE",
                Username: "YOUR_USERNAME_HERE",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tasks.ListTasks(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Tasks != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### Schedules

* `CreateSchedule` - Create Schedule
* `DeleteSchedule` - Delete Schedule
* `GetSchedule` - Get Schedule
* `GetScheduleTasks` - Get Schedule Tasks
* `ListSchedules` - Get Schedules
* `UpdateSchedule` - Update Schedule

### Tasks

* `BatchCreateTasks` - Batch Create Tasks (Beta)
* `BatchDeleteTasks` - Batch Delete Tasks (Beta)
* `CreateTask` - Create Task
* `DeleteTask` - Delete Task
* `GetTask` - Get Task
* `ListTasks` - Get Tasks
* `RunTask` - Run Task
* `UpdateTask` - Update Task
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
