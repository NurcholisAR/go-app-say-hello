package main

import (
	"fmt"

	gomodulesayhello "github.com/NurcholisAR/go-module-say-hello/v2"
)

func main() {
	app := gomodulesayhello.Say_Hello("oke")
	fmt.Println(app)
}
