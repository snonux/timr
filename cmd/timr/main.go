package main

import (
	"flag"
	"fmt"
	"os"

	"codeberg.org/snonux/timr/internal"
)

func main() {
	versionFlag := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println(internal.Version)
		os.Exit(0)
	}
}
