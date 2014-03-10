package main

import "fmt"
import "animapi/infrastructure"

func main() {
    db := infra.GetDB("sample")
    fmt.Println(db)
}
