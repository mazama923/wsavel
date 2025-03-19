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
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	str := fmt.Sprintf("\n\n   %s %s\n\n", m.spinner.View(), m.message)
	if m.quitting {
		return str + "\n"
	}
	return str
}

func StartSpinner(message string) {
	done = make(chan bool)
	m := initialModel(message)

	p := tea.NewProgram(m)
	go func() {
		p.Run()
	}()
}

func TestSpinner(message string) {
	done = make(chan bool)
	m := initialModel(message)

	p := tea.NewProgram(m)
	p.Run()
}

func StopSpinner() {
	if done != nil {
		close(done)
	}
}
