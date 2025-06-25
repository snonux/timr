package main

import (
	"fmt"
	"os"

	"timr/internal/timer"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	var err error
	var output string

	switch os.Args[1] {
	case "start":
		output, err = timer.StartTimer()
	case "stop", "pause":
		output, err = timer.StopTimer()
	case "status":
		output, err = timer.GetStatus()
	case "reset":
		output, err = timer.ResetTimer()
	default:
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(output)
}

func printUsage() {
	fmt.Println("Usage: timr <start|stop|pause|status|reset>")
}
