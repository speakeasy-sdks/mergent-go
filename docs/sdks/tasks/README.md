# Tasks

## Overview

Manage Tasks

### Available Operations

* [BatchCreate](#batchcreate) - Batch Create Tasks (Beta)
* [BatchDelete](#batchdelete) - Batch Delete Tasks (Beta)
* [Create](#create) - Create Task
* [Delete](#delete) - Delete Task
* [Get](#get) - Get Task
* [List](#list) - Get Tasks
* [Run](#run) - Run Task
* [Update](#update) - Update Task

## BatchCreate

A maximum of 100 Tasks are accepted per request.

This operation is atomic: it will succeed for all Tasks or fail for all
Tasks; there is no partial success.

This endpoint is in beta and may change at any time without notice.


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
    res, err := s.Tasks.BatchCreate(ctx, []shared.TaskNewInput{
        shared.TaskNewInput{
            Delay: mergent.String("60s"),
            Name: mergent.String("mytask"),
            Queue: mergent.String("process1"),
            Request: shared.Request{
                Body: mergent.String("Body String"),
                Headers: &shared.RequestHeaders{},
                URL: "http://example.com",
            },
            ScheduledFor: types.MustTimeFromString("2021-10-01T15:53:05Z"),
        },
        shared.TaskNewInput{
            Delay: mergent.String("60s"),
            Name: mergent.String("mytask"),
            Queue: mergent.String("process1"),
            Request: shared.Request{
                Body: mergent.String("Body String"),
                Headers: &shared.RequestHeaders{},
                URL: "http://example.com",
            },
            ScheduledFor: types.MustTimeFromString("2021-10-01T15:53:05Z"),
        },
        shared.TaskNewInput{
            Delay: mergent.String("60s"),
            Name: mergent.String("mytask"),
            Queue: mergent.String("process1"),
            Request: shared.Request{
                Body: mergent.String("Body String"),
                Headers: &shared.RequestHeaders{},
                URL: "http://example.com",
            },
            ScheduledFor: types.MustTimeFromString("2021-10-01T15:53:05Z"),
        },
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]shared.TaskNewInput](../../models//.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.BatchCreateTasksResponse](../../models/operations/batchcreatetasksresponse.md), error**


## BatchDelete

A maximum of 100 Task IDs are accepted per request.

This operation is atomic: it will succeed for all Tasks or fail for all
Tasks; there is no partial success.

This endpoint is in beta and may change at any time without notice.


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
    res, err := s.Tasks.BatchDelete(ctx, []string{
        "b0074f15-471b-45e6-a13b-99d488e1e91e",
        "450ad2ab-d442-4698-82d5-02a94bb4f63c",
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]string](../../models//.md)                         | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.BatchDeleteTasksResponse](../../models/operations/batchdeletetasksresponse.md), error**


## Create

Create Task

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
    res, err := s.Tasks.Create(ctx, shared.TaskNewInput{
        Delay: mergent.String("60s"),
        Name: mergent.String("mytask"),
        Queue: mergent.String("process1"),
        Request: shared.Request{
            Body: mergent.String("Body String"),
            Headers: &shared.RequestHeaders{},
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

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.TaskNewInput](../../models/shared/tasknewinput.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.CreateTaskResponse](../../models/operations/createtaskresponse.md), error**


## Delete

Delete Task

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
    res, err := s.Tasks.Delete(ctx, operations.DeleteTaskRequest{
        TaskID: "969e9a3e-fa77-4dfb-94cd-66ae395efb9b",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeleteTaskRequest](../../models/operations/deletetaskrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.DeleteTaskResponse](../../models/operations/deletetaskresponse.md), error**


## Get

Get Task

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
    res, err := s.Tasks.Get(ctx, operations.GetTaskRequest{
        TaskID: "a88f3a66-9970-474b-a446-9b6e21419598",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetTaskRequest](../../models/operations/gettaskrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetTaskResponse](../../models/operations/gettaskresponse.md), error**


## List

Get Tasks

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
    res, err := s.Tasks.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Tasks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListTasksResponse](../../models/operations/listtasksresponse.md), error**


## Run

Reschedules a queued Task to be run immediately.

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
    res, err := s.Tasks.Run(ctx, operations.RunTaskRequest{
        TaskID: "90afa563-e251-46fe-8c8b-711e5b7fd2ed",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.RunTaskRequest](../../models/operations/runtaskrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.RunTaskResponse](../../models/operations/runtaskresponse.md), error**


## Update

Update Task

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
    res, err := s.Tasks.Update(ctx, operations.UpdateTaskRequest{
        TaskInput: shared.TaskInput{
            Delay: mergent.String("60s"),
            Name: mergent.String("mytask"),
            Queue: mergent.String("process1"),
            Request: &shared.Request{
                Body: mergent.String("Body String"),
                Headers: &shared.RequestHeaders{},
                URL: "http://example.com",
            },
            ScheduledFor: types.MustTimeFromString("2021-10-01T15:53:05Z"),
        },
        TaskID: "028921cd-dc69-4260-9fb5-76b0d5f0d30c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpdateTaskRequest](../../models/operations/updatetaskrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UpdateTaskResponse](../../models/operations/updatetaskresponse.md), error**

