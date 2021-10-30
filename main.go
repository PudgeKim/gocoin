package main

import (
	"github.com/pudgekim/gocoin/cli"
	"github.com/pudgekim/gocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
