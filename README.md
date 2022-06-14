## Test results

```bash
$ go test -race -cover ./...
?   	github.com/ariel17/nirvana-assessment	[no test files]
ok  	github.com/ariel17/nirvana-assessment/pkg/configs	(cached)	coverage: 100.0% of statements
ok  	github.com/ariel17/nirvana-assessment/pkg/server	(cached)	coverage: 76.5% of statements
ok  	github.com/ariel17/nirvana-assessment/pkg/services	(cached)	coverage: 100.0% of statements
```

## Local execution

```bash
# terminal 1

$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> github.com/ariel17/nirvana-assessment/pkg/server.CoalesceHandler (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080

[GIN] 2022/06/14 - 03:27:38 | 200 |  785.810674ms |       127.0.0.1 | GET      "/?member_id=1"


# terminal 2

$ curl -s http://localhost:8080/\?member_id\=1 | jq .

{
  "deductible": 1066,
  "stop_loss": 11000,
  "oop_max": 5666
}
```
