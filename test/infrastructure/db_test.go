package infra

import "fmt"
import "animapi/infrastructure"
import "testing"

func TestGetDb(t *testing.T) {
    db := infra.GetDB("test", "000")

    res := db.CreateAnimeTable()
    fmt.Printf("%+v\n", res)
    res = db.DropAnimeTable()
    fmt.Printf("%+v\n", res)
    res = db.CreateAnimeTable()
    fmt.Printf("%+v\n", res)

    res = db.InsertAnime()
    res = db.InsertAnime()
    res = db.InsertAnime()
    res = db.InsertAnime()
    fmt.Printf("%+v\n", res)

    rows := db.FindAllAnime()
    for rows.Next() {
        fmt.Println("あっっっったああああああああ")
        var title string
        var id    string
        e := rows.Scan(&id, &title)
        if e != nil {
            fmt.Println(e)
        }
        fmt.Println(
            title,
            id,
        )
    }
}
