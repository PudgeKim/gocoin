package main

import (
	"github.com/pudgekim/gocoin/explorer"
	"github.com/pudgekim/gocoin/rest"
)

func main() {
	go rest.Start(3030)
	explorer.Start(4040)
}
