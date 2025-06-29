package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"timr/internal/live"
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
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "raw":
				output, err = timer.GetRawStatus()
			case "rawm":
				output, err = timer.GetRawMinutesStatus()
			default:
				printUsage()
				os.Exit(1)
			}
		} else {
			output, err = timer.GetStatus()
		}
	case "reset":
		output, err = timer.ResetTimer()
	case "prompt":
		output, err = timer.GetPromptStatus()
	case "track":
		if len(os.Args) < 3 {
			fmt.Println("Error: track command requires a description")
			printUsage()
			os.Exit(1)
		}
		// Join all arguments after "track" as the description
		description := strings.Join(os.Args[2:], " ")
		output, err = timer.TrackTime(description)
	case "live":
		p := tea.NewProgram(live.NewModel())
		if err := p.Start(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		return
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
	fmt.Println("Usage: timr <start|stop|pause|status|reset|live|prompt|track <description>>")
}
