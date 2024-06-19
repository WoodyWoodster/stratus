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
	var architecture string
	var name string
	var resource string
	var runtime string

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
					Value(&resource).
					Options(huh.NewOptions("Lambda", "API Gateway", "DynamoDB")...).
					Title("What serverless resource would you like to create?"),
			),
			huh.NewGroup(
				huh.NewInput().
					Key("name").
					Title("Resource name").
					Value(&name).
					Validate(
						func(s string) error {
							if s == "" {
								return fmt.Errorf("resource name is required")
							}
							return nil
						},
					),
			).WithHideFunc(func() bool {
				return resource == ""
			}),
			huh.NewGroup(
				huh.NewSelect[string]().
					Key("architecture").
					Value(&architecture).
					Options(huh.NewOptions("x86_64", "arm64")...).
					Title("What architecture would you like to use?"),
			),
			huh.NewGroup(
				huh.NewSelect[string]().
					Key("runtime").
					Value(&runtime).
					Options(huh.NewOptions(
						"nodejs20.x",
						"nodejs18.x",
						"nodejs16.x",
						"python3.12",
						"python3.11",
						"python3.10",
						"python3.9",
						"python3.8",
						"provided.al2023",
					)...).
					Title("What runtime would you like to use?"),
			).WithHideFunc(func() bool {
				return resource != "Lambda"
			}),
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
		return createResource(m)
	}

	return m.form.View()
}

func createResource(m Model) string {
	architecture := m.form.GetString("architecture")
	name := m.form.GetString("name")
	resource := m.form.GetString("resource")
	runtime := m.form.GetString("runtime")
	s := fmt.Sprintf("Creating %s named %s\n", resource, name)

	if resource == "Lambda" {
		s += fmt.Sprintf("Using runtime %s on %s\n", runtime, architecture)
	}
	s += "Done!"
	return s
}

func main() {
	p := tea.NewProgram(NewModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Shit! It broke: %v", err)
		os.Exit(1)
	}
}
