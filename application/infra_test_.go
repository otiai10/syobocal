package main

import "fmt"
import "animapi/infrastructure"

func main() {
    db := infra.GetDB("test", "000")
    fmt.Println(db)
    // こういうのrepositoryでやろうね本当は
    rows := db.FindAllAnime()
    for rows.Next() {
        var title string
        _ = rows.Scan(&title)
        fmt.Printf(title + "\n")
    }
}
