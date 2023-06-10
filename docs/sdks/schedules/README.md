# Schedules

## Overview

Manage Schedule

### Available Operations

* [Create](#create) - Create Schedule
* [Delete](#delete) - Delete Schedule
* [Get](#get) - Get Schedule
* [GetTasks](#gettasks) - Get Schedule Tasks
* [List](#list) - Get Schedules
* [Update](#update) - Update Schedule

## Create

Create Schedule

### Example Usage

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
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Schedules.Create(ctx, shared.ScheduleNewInput{
        Cron: mergent.String("corrupti"),
        Description: mergent.String("provident"),
        Dtstart: types.MustTimeFromString("2021-04-24T16:27:50.833Z"),
        Name: mergent.String("Ismael Little"),
        Paused: mergent.Bool(false),
        Queue: mergent.String("error"),
        Request: shared.Request{
            Body: mergent.String("Body String"),
            Headers: &shared.RequestHeaders{},
            URL: "http://example.com",
        },
        Rrule: mergent.String("deserunt"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.ScheduleNewInput](../../models/shared/schedulenewinput.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.CreateScheduleResponse](../../models/operations/createscheduleresponse.md), error**


## Delete

Delete Schedule

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/mergent-go"
	"github.com/speakeasy-sdks/mergent-go/pkg/models/operations"
)

func main() {
    s := mergent.New(
        mergent.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Schedules.Delete(ctx, operations.DeleteScheduleRequest{
        ScheduleID: "674e0f46-7cc8-4796-ad15-1a05dfc2ddf7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteScheduleRequest](../../models/operations/deleteschedulerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteScheduleResponse](../../models/operations/deletescheduleresponse.md), error**


## Get

Get Schedule

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/mergent-go"
	"github.com/speakeasy-sdks/mergent-go/pkg/models/operations"
)

func main() {
    s := mergent.New(
        mergent.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Schedules.Get(ctx, operations.GetScheduleRequest{
        ScheduleID: "cc78ca1b-a928-4fc8-9674-2cb739205929",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetScheduleRequest](../../models/operations/getschedulerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetScheduleResponse](../../models/operations/getscheduleresponse.md), error**


## GetTasks

Get Schedule Tasks

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/mergent-go"
	"github.com/speakeasy-sdks/mergent-go/pkg/models/operations"
)

func main() {
    s := mergent.New(
        mergent.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Schedules.GetTasks(ctx, operations.GetScheduleTasksRequest{
        ScheduleID: "396fea75-96eb-410f-aaa2-352c5955907a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Tasks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetScheduleTasksRequest](../../models/operations/getscheduletasksrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetScheduleTasksResponse](../../models/operations/getscheduletasksresponse.md), error**


## List

Get Schedules

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/mergent-go"
)

func main() {
    s := mergent.New(
        mergent.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Schedules.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListSchedulesResponse](../../models/operations/listschedulesresponse.md), error**


## Update

Update Schedule

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/mergent-go"
	"github.com/speakeasy-sdks/mergent-go/pkg/models/operations"
	"github.com/speakeasy-sdks/mergent-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mergent-go/pkg/types"
)

func main() {
    s := mergent.New(
        mergent.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Schedules.Update(ctx, operations.UpdateScheduleRequest{
        ScheduleInput: shared.ScheduleInput{
            Cron: mergent.String("doloribus"),
            Description: mergent.String("sapiente"),
            Dtstart: types.MustTimeFromString("2022-05-07T17:33:24.154Z"),
            Name: mergent.String("Cecilia Crooks"),
            Paused: mergent.Bool(false),
            Queue: mergent.String("occaecati"),
            Request: &shared.Request{
                Body: mergent.String("Body String"),
                Headers: &shared.RequestHeaders{},
                URL: "http://example.com",
            },
            Rrule: mergent.String("numquam"),
        },
        ScheduleID: "67739251-aa52-4c3f-9ad0-19da1ffe78f0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateScheduleRequest](../../models/operations/updateschedulerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateScheduleResponse](../../models/operations/updatescheduleresponse.md), error**

