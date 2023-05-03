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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
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
                Headers: map[string]interface{}{
                    "modi": "praesentium",
                    "rem": "voluptates",
                    "quasi": "repudiandae",
                    "sint": "veritatis",
                },
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
                Headers: map[string]interface{}{
                    "incidunt": "enim",
                    "consequatur": "est",
                    "quibusdam": "explicabo",
                    "deserunt": "distinctio",
                },
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
                Headers: map[string]interface{}{
                    "labore": "modi",
                    "qui": "aliquid",
                    "cupiditate": "quos",
                    "perferendis": "magni",
                },
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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Tasks.BatchDelete(ctx, []string{
        "502a94bb-4f63-4c96-9e9a-3efa77dfb14c",
        "d66ae395-efb9-4ba8-8f3a-66997074ba44",
        "69b6e214-1959-4890-afa5-63e2516fe4c8",
        "b711e5b7-fd2e-4d02-8921-cddc692601fb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
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
                "voluptate": "autem",
                "nam": "eaque",
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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Tasks.Delete(ctx, operations.DeleteTaskRequest{
        TaskID: "d5f0d30c-5fbb-4258-b053-202c73d5fe9b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Tasks.Get(ctx, operations.GetTaskRequest{
        TaskID: "90c28909-b3fe-449a-8d9c-bf48633323f9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Task != nil {
        // handle response
    }
}
```

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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Tasks.Run(ctx, operations.RunTaskRequest{
        TaskID: "b77f3a41-0067-44eb-b692-80d1ba77a89e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Task != nil {
        // handle response
    }
}
```

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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
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
                Headers: map[string]interface{}{
                    "asperiores": "nihil",
                    "ipsum": "voluptate",
                    "id": "saepe",
                },
                URL: "http://example.com",
            },
            ScheduledFor: types.MustTimeFromString("2021-10-01T15:53:05Z"),
        },
        TaskID: "4203ce5e-6a95-4d8a-8d44-6ce2af7a73cf",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Task != nil {
        // handle response
    }
}
```
