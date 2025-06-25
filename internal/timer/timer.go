package timer

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const (
	stateFile = ".timr_state"
)

// stateFilePathOverride is used by tests to override the state file path.
var stateFilePathOverride string

type State struct {
	StartTime    time.Time
	ElapsedTime  time.Duration
	Running      bool
}

func GetStateFile() (string, error) {
	if stateFilePathOverride != "" {
		return stateFilePathOverride, nil
	}
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "timr", stateFile), nil
}

func LoadState() (State, error) {
	var state State
	stateFilePath, err := GetStateFile()
	if err != nil {
		return state, err
	}

	data, err := os.ReadFile(stateFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return State{}, nil
		}
		return state, err
	}

	err = json.Unmarshal(data, &state)
	return state, err
}

func (s *State) Save() error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	stateFilePath, err := GetStateFile()
	if err != nil {
		return err
	}

	return os.WriteFile(stateFilePath, data, 0644)
}
