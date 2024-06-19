package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form *huh.Form
}

func NewModel() Model {
	return Model{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Stratus").
					Description("Stratus simplifies serverless app development on AWS.\n\n").
					Next(true).
					NextLabel("Start"),
			),
			huh.NewGroup(
				huh.NewSelect[string]().
					Key("resource").
					Options(huh.NewOptions("Lambda", "API Gateway", "DynamoDB")...).
					Title("What serverless resource would you like to create?"),
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}
	return m, cmd
}

func (m Model) View() string {
	if m.form.State == huh.StateCompleted {
		resource := m.form.GetString("resource")
		return fmt.Sprintf("You chose: %s\n", resource)
	}

	return m.form.View()
}

func main() {
	p := tea.NewProgram(NewModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Shit! It broke: %v", err)
		os.Exit(1)
	}
}
