# animapi
is a simple client for SyoboCal.

It can

1. Find programs from syobocal directly 
2. Store animes, anisons and programs in local MySQL.
3. (the combination of 1 and 2) Crawl and store those.
4. Find animes, anisons and programs from local MySQL.

[![Build Status](https://travis-ci.org/otiai10/animapi.svg?branch=ws/v3)](https://travis-ci.org/otiai10/animapi)

# usage
### in go code
```go
package main

import "github.com/otiai10/animapi"
import "fmt"

func main() {
    since := animapi.Since("-1w")    

    programs := animapi.SYOBOCAL.FindPrograms(since)
    fmt.Printf("%+v", programs)

    programs = animapi.LOCAL.FindPrograms(since)
    fmt.Printf("%+v", programs)
}
```
### in command line
```sh
% animapi crawl -cron 24h  -mysql
OK, just starting crawler
[mode]  cron per 24h
[mysql] localhost:3306
```

# test
```
% go test ./...
```

# dependencies
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)


