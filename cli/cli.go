package cli

import (
	"flag"
	"fmt"
	"github.com/pudgekim/gocoin/explorer"
	"github.com/pudgekim/gocoin/rest"
	"runtime"
)

func usage() {
	fmt.Println("Please use the following commands")
	fmt.Println()
	fmt.Println("-port=4000    Start the HTML Explorer")
	fmt.Println("-mode=rest:   Choose between 'html' and 'rest'")
	runtime.Goexit()
}

func Start() {
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
