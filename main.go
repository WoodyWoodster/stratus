package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var frameworks = []string{"AWS SAM", "Serverless"}

type model struct {
	choice string
	cursor int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(frameworks) - 1
			}
		case "down", "j":
			m.cursor++
			if m.cursor >= len(frameworks) {
				m.cursor = 0
			}
		case "enter", " ":
			m.choice = frameworks[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("What framework would you like to use?\n\n")

	for i := 0; i < len(frameworks); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(frameworks[i])
		s.WriteString("\n")
	}
	s.WriteString("\nPress q to quit.\n")

	return s.String()
}

func main() {
	p := tea.NewProgram(model{})
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Printf("You chose: %s\n", m.choice)
	}
}
