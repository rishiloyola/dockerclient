package main

import (
	"fmt"
	"zkclient"
)

func main() {
	c := zkclient.New()
	c.Init("127.0.0.1:2181")
	c.Connect("pushpin1", "/pushpin/server")
	fmt.Println("started")
}
