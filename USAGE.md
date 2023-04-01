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
                "provident": "distinctio",
                "quibusdam": "unde",
                "nulla": "corrupti",
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