package main

import (
	"fmt"
	"zkclient"
)

func main() {
	c := zkclient.New()
	c.Init("127.0.0.1")
}
