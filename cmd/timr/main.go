package main

import (
	"flag"
	"fmt"
	"os"

	"codeberg.org/snonux/timr/internal/version"
)

func main() {
	versionFlag := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version.Version)
		os.Exit(0)
	}
}
