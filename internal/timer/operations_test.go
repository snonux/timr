package timer

import (
	"path/filepath"
	"testing"
	"time"
)

// setup sets up a temporary state file for testing.
func setup(t *testing.T) {
	t.Helper()
	tempDir := t.TempDir()
	stateFilePathOverride = filepath.Join(tempDir, ".timr_state")
	t.Cleanup(func() {
		stateFilePathOverride = ""
	})
}

func TestStartTimer(t *testing.T) {
	setup(t)

	// Start the timer
	msg, err := StartTimer()
	if err != nil {
		t.Fatalf("StartTimer() error = %v", err)
	}
	if msg != "Timer started." {
		t.Errorf("StartTimer() msg = %v, want %v", msg, "Timer started.")
	}

	// Check the state
	state, err := LoadState()
	if err != nil {
		t.Fatalf("LoadState() error = %v", err)
	}
	if !state.Running {
		t.Error("state.Running = false, want true")
	}

	// Try to start again
	msg, err = StartTimer()
	if err != nil {
		t.Fatalf("StartTimer() error = %v", err)
	}
	if msg != "Timer is already running." {
		t.Errorf("StartTimer() msg = %v, want %v", msg, "Timer is already running.")
	}
}

func TestStopTimer(t *testing.T) {
	setup(t)

	// Stop before starting
	msg, err := StopTimer()
	if err != nil {
		t.Fatalf("StopTimer() error = %v", err)
	}
	if msg != "Timer is not running." {
		t.Errorf("StopTimer() msg = %v, want %v", msg, "Timer is not running.")
	}

	// Start and then stop the timer
	_, _ = StartTimer()
	time.Sleep(10 * time.Millisecond) // Simulate work
	msg, err = StopTimer()
	if err != nil {
		t.Fatalf("StopTimer() error = %v", err)
	}
	if msg != "Timer stopped." {
		t.Errorf("StopTimer() msg = %v, want %v", msg, "Timer stopped.")
	}

	// Check the state
	state, err := LoadState()
	if err != nil {
		t.Fatalf("LoadState() error = %v", err)
	}
	if state.Running {
		t.Error("state.Running = true, want false")
	}
	if state.ElapsedTime == 0 {
		t.Error("state.ElapsedTime = 0, want > 0")
	}
}

func TestGetStatus(t *testing.T) {
	setup(t)

	// Status when stopped
	msg, err := GetStatus()
	if err != nil {
		t.Fatalf("GetStatus() error = %v", err)
	}
	want := "Status: Stopped\nElapsed Time: 0s"
	if msg != want {
		t.Errorf("GetStatus() msg = %q, want %q", msg, want)
	}

	// Status when running
	_, _ = StartTimer()
	msg, err = GetStatus()
	if err != nil {
		t.Fatalf("GetStatus() error = %v", err)
	}
	want = "Status: Running\nElapsed Time: 0s"
	if msg != want {
		t.Errorf("GetStatus() msg = %q, want %q", msg, want)
	}
}

func TestGetRawStatus(t *testing.T) {
	setup(t)

	// Raw status when stopped
	msg, err := GetRawStatus()
	if err != nil {
		t.Fatalf("GetRawStatus() error = %v", err)
	}
	want := "0"
	if msg != want {
		t.Errorf("GetRawStatus() msg = %q, want %q", msg, want)
	}

	// Raw status when running
	state, err := LoadState()
	if err != nil {
		t.Fatalf("LoadState() error = %v", err)
	}
	state.Running = true
	state.StartTime = time.Now().Add(-2 * time.Second) // Set start time 2 seconds ago
	state.ElapsedTime = 0                               // Reset elapsed time for this specific test
	if err := state.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	msg, err = GetRawStatus()
	if err != nil {
		t.Fatalf("GetRawStatus() error = %v", err)
	}
	want = "2"
	if msg != want {
		t.Errorf("GetRawStatus() msg = %q, want %q", msg, want)
	}
}

func TestGetRawMinutesStatus(t *testing.T) {
	setup(t)

	// Raw minutes status when stopped
	msg, err := GetRawMinutesStatus()
	if err != nil {
		t.Fatalf("GetRawMinutesStatus() error = %v", err)
	}
	want := "0"
	if msg != want {
		t.Errorf("GetRawMinutesStatus() msg = %q, want %q", msg, want)
	}

	// Raw minutes status when running (simulating 2 minutes)
	state, err := LoadState()
	if err != nil {
		t.Fatalf("LoadState() error = %v", err)
	}
	state.Running = true
	state.StartTime = time.Now().Add(-2 * time.Minute) // Set start time 2 minutes ago
	state.ElapsedTime = 0                               // Reset elapsed time for this specific test
	if err := state.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	msg, err = GetRawMinutesStatus()
	if err != nil {
		t.Fatalf("GetRawMinutesStatus() error = %v", err)
	}
	want = "2"
	if msg != want {
		t.Errorf("GetRawMinutesStatus() msg = %q, want %q", msg, want)
	}
}

func TestResetTimer(t *testing.T) {
	setup(t)

	// Start timer to create a state file
	_, _ = StartTimer()

	// Reset the timer
	msg, err := ResetTimer()
	if err != nil {
		t.Fatalf("ResetTimer() error = %v", err)
	}
	if msg != "Timer reset." {
		t.Errorf("ResetTimer() msg = %v, want %v", msg, "Timer reset.")
	}

	// Check the state
	state, err := LoadState()
	if err != nil {
		t.Fatalf("LoadState() error = %v", err)
	}
	if state.Running {
		t.Error("state.Running = true, want false")
	}
	if state.ElapsedTime != 0 {
		t.Errorf("state.ElapsedTime = %v, want 0", state.ElapsedTime)
	}
}
