# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`timr` is a simple command-line time tracking tool written in Go. It provides a minimalist stopwatch-style timer that persists state across sessions, allowing users to track time spent on tasks.

## Common Development Commands

### Build & Run
- `task` or `go build -o timr ./cmd/timr` - Build the executable
- `task run` or `go run ./cmd/timr` - Run without building
- `task install` or `go install ./cmd/timr` - Install to GOPATH/bin

### Testing
- `task test` or `go test ./...` - Run all tests
- `go test ./internal/timer -v` - Run specific package tests with verbose output

## Architecture

The codebase follows a clean modular structure:

- **Entry Point**: `/cmd/timr/main.go` - CLI command handling
- **Core Logic**: `/internal/timer/` - Timer state management and operations
  - `timer.go` - State persistence (JSON file at `~/.config/timr/.timr_state`)
  - `operations.go` - Core timer operations (start, stop, status, reset)
- **Interactive UI**: `/internal/live/` - Full-screen timer using Bubble Tea TUI framework
- **Version Info**: `/internal/version.go` - Version constants

## Key Implementation Details

1. **State Persistence**: Timer state is stored as JSON in `~/.config/timr/.timr_state`, containing:
   - `running`: boolean indicating if timer is active
   - `startTime`: timestamp when timer was last started
   - `elapsedTime`: accumulated time in nanoseconds

2. **Command Structure**: The CLI supports these subcommands:
   - `start` - Start/resume timer
   - `stop`/`pause` - Stop timer (synonyms)
   - `status` - Show formatted status
   - `status raw` - Show elapsed seconds (raw number)
   - `status rawm` - Show elapsed minutes (raw number)
   - `reset` - Reset timer to zero
   - `prompt` - Special format for shell prompt integration (returns empty if not running)
   - `live` - Interactive TUI mode with keyboard controls

3. **TUI Framework**: Uses Bubble Tea for the interactive live mode, with dependencies on various charmbracelet libraries for terminal styling.

4. **Error Handling**: Operations that modify state (start, stop, reset) save state immediately and return errors if persistence fails.

## Development Guidelines

- Timer operations should maintain atomic state updates
- All state modifications must persist immediately to handle unexpected exits
- The live mode runs in a separate Bubble Tea program and doesn't modify timer state directly
- Shell prompt integration returns empty string when timer is not running to avoid cluttering prompts