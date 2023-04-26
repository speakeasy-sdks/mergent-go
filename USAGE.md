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
            APIKey: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.TaskNewInput{
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
    }

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