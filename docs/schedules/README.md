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
        Cron: mergent.String("illum"),
        Description: mergent.String("vel"),
        Dtstart: types.MustTimeFromString("2021-09-16T11:56:06.019Z"),
        Name: mergent.String("Willie Gulgowski DVM"),
        Paused: mergent.Bool(false),
        Queue: mergent.String("tempora"),
        Request: shared.Request{
            Body: mergent.String("Body String"),
            Headers: map[string]interface{}{
                "molestiae": "minus",
                "placeat": "voluptatum",
            },
            URL: "http://example.com",
        },
        Rrule: mergent.String("iusto"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

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
        ScheduleID: "96ed151a-05df-4c2d-9f7c-c78ca1ba928f",
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
        ScheduleID: "c816742c-b739-4205-9293-96fea7596eb1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```

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
        ScheduleID: "0faaa235-2c59-4559-87af-f1a3a2fa9467",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Tasks != nil {
        // handle response
    }
}
```

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
            Cron: mergent.String("molestiae"),
            Description: mergent.String("velit"),
            Dtstart: types.MustTimeFromString("2022-09-06T22:51:09.401Z"),
            Name: mergent.String("Gloria Padberg"),
            Paused: mergent.Bool(false),
            Queue: mergent.String("odit"),
            Request: &shared.Request{
                Body: mergent.String("Body String"),
                Headers: map[string]interface{}{
                    "sequi": "tenetur",
                    "ipsam": "id",
                    "possimus": "aut",
                    "quasi": "error",
                },
                URL: "http://example.com",
            },
            Rrule: mergent.String("temporibus"),
        },
        ScheduleID: "a1ffe78f-097b-4007-8f15-471b5e6e13b9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schedule != nil {
        // handle response
    }
}
```
