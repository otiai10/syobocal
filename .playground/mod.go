package main

import "fmt"

var Hoge = map[string]string{
	"hoge": "fuga",
}

func main() {
	fmt.Printf("%+v", Hoge)
}
