package live

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	timrTimer "timr/internal/timer"
)

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Model is the Bubble Tea model for the live timer view.
type Model struct {
	state       timrTimer.State
	quitting    bool
	helpStyle   lipgloss.Style
	timerStyle  lipgloss.Style
	statusStyle lipgloss.Style
}

// NewModel creates a new Model.
func NewModel() Model {
	state, err := timrTimer.LoadState()
	if err != nil {
		panic(err) // Or handle more gracefully
	}

	return Model{
		state:       state,
		helpStyle:   lipgloss.NewStyle().Faint(true),
		timerStyle:  lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BFFF")),
		statusStyle: lipgloss.NewStyle().Italic(true),
	}
}

// Init is the first function that will be called.
func (m Model) Init() tea.Cmd {
	if m.state.Running {
		return tick()
	}
	return nil
}

// Update is called when a message is received.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if !m.state.Running {
			return m, nil
		}
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			if err := m.state.Save(); err != nil {
				// handle error
			}
			return m, tea.Quit

		case "s":
			if m.state.Running {
				// Stop the timer
				m.state.ElapsedTime += time.Since(m.state.StartTime)
				m.state.Running = false
				return m, nil // Stop ticking
			} else {
				// Start the timer
				m.state.Running = true
				m.state.StartTime = time.Now()
				return m, tick() // Start ticking
			}

		case "r":
			// Reset the timer
			m.state = timrTimer.State{}
			if _, err := timrTimer.ResetTimer(); err != nil {
				// handle error
			}
			return m, nil // Stop ticking
		}
	}

	return m, nil
}

// View renders the model's state to the terminal.
func (m Model) View() string {
	if m.quitting {
		return ""
	}

	var currentSession time.Duration
	if m.state.Running {
		currentSession = time.Since(m.state.StartTime)
	}

	totalElapsed := (m.state.ElapsedTime + currentSession).Round(time.Second)

	status := "Paused"
	if m.state.Running {
		status = "Running"
	}

	return fmt.Sprintf(
		"%s\n%s\n\n%s",
		m.timerStyle.Render(totalElapsed.String()),
		m.statusStyle.Render(status),
		m.helpStyle.Render("q: quit, s: start/stop, r: reset"),
	)
}
