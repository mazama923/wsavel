package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

type updateMessageMsg string

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
	message  string
}

var activeProgram *tea.Program

func initialModel(message string) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, message: message}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	case updateMessageMsg:
		m.message = string(msg)
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s %s\n\n", m.spinner.View(), m.message)
	if m.quitting {
		return str + "\n"
	}
	return str
}

func StartSpinner(message string) {
	p := tea.NewProgram(initialModel(message))
	activeProgram = p

	// Exécuter le programme dans une goroutine
	go func() {
		if _, err := p.Run(); err != nil {
			fmt.Println(err)
		}
	}()
}

func UpdateSpinnerMessage(newMessage string) {
	if activeProgram != nil {
		activeProgram.Send(updateMessageMsg(newMessage))
	}
}

func StopSpinner() {
	if activeProgram != nil {
		activeProgram.Quit()
		activeProgram = nil
	}
}
