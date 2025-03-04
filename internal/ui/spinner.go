package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var done chan bool

type model struct {
	spinner  spinner.Model
	quitting bool
	message  string
}

func initialModel(message string) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		spinner: s,
		message: message,
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	select {
	case <-done:
		m.quitting = true
		return m, tea.Quit
	default:
		return m, m.spinner.Tick
	}
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	return fmt.Sprintf("\n%s %s\n", m.spinner.View(), m.message)
}

func StartSpinner(message string) {
	done = make(chan bool)
	m := initialModel(message)

	p := tea.NewProgram(m)
	go func() {
		p.Run()
	}()
}

func StopSpinner() {
	if done != nil {
		close(done)
	}
}
