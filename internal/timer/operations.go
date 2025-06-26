package timer

import (
	"fmt"
	"os"
	"time"
)

func StartTimer() (string, error) {
	state, err := LoadState()
	if err != nil {
		return "", fmt.Errorf("error loading state: %w", err)
	}

	if state.Running {
		return "Timer is already running.", nil
	}

	state.Running = true
	state.StartTime = time.Now()
	if err := state.Save(); err != nil {
		return "", fmt.Errorf("error saving state: %w", err)
	}
	return "Timer started.", nil
}

func StopTimer() (string, error) {
	state, err := LoadState()
	if err != nil {
		return "", fmt.Errorf("error loading state: %w", err)
	}

	if !state.Running {
		return "Timer is not running.", nil
	}

	state.Running = false
	state.ElapsedTime += time.Since(state.StartTime)
	if err := state.Save(); err != nil {
		return "", fmt.Errorf("error saving state: %w", err)
	}
	return "Timer stopped.", nil
}

func GetStatus() (string, error) {
	state, err := LoadState()
	if err != nil {
		return "", fmt.Errorf("error loading state: %w", err)
	}

	if state.Running {
		elapsed := (state.ElapsedTime + time.Since(state.StartTime)).Round(time.Second)
		return fmt.Sprintf("Status: Running\nElapsed Time: %s", elapsed), nil
	} else {
		elapsed := state.ElapsedTime.Round(time.Second)
		return fmt.Sprintf("Status: Stopped\nElapsed Time: %s", elapsed), nil
	}
}

func ResetTimer() (string, error) {
	stateFile, err := GetStateFile()
	if err != nil {
		return "", fmt.Errorf("error getting state file path: %w", err)
	}
	if err := os.Remove(stateFile); err != nil {
		if !os.IsNotExist(err) {
			return "", fmt.Errorf("error resetting timer: %w", err)
		}
	}
	state := State{}
	if err := state.Save(); err != nil {
		return "", fmt.Errorf("error saving state: %w", err)
	}
	return "Timer reset.", nil
}

func GetRawStatus() (string, error) {
	state, err := LoadState()
	if err != nil {
		return "", fmt.Errorf("error loading state: %w", err)
	}

	elapsed := state.ElapsedTime
	if state.Running {
		elapsed += time.Since(state.StartTime)
	}

	return fmt.Sprintf("%d", int(elapsed.Seconds())), nil
}

func GetRawMinutesStatus() (string, error) {
	state, err := LoadState()
	if err != nil {
		return "", fmt.Errorf("error loading state: %w", err)
	}

	elapsed := state.ElapsedTime
	if state.Running {
		elapsed += time.Since(state.StartTime)
	}

	return fmt.Sprintf("%d", int(elapsed.Minutes())), nil
}

func GetPromptStatus() (string, error) {
	state, err := LoadState()
	if err != nil {
		return "", fmt.Errorf("error loading state: %w", err)
	}

	elapsed := state.ElapsedTime
	if state.Running {
		elapsed += time.Since(state.StartTime)
	}

	if elapsed == 0 {
		return "", nil
	}

	icon := "⏸"
	if state.Running {
		icon = "▶"
	}

	return fmt.Sprintf("%s%s", icon, elapsed.Round(time.Second)), nil
}
