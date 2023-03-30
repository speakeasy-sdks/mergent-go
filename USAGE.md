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