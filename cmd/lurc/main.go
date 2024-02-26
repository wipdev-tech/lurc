package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	l "github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

var styles = struct {
	header l.Style
	prompt l.Style
}{
	header: l.NewStyle().
		// Background(l.Color("36")).
		Bold(true).
		Foreground(l.Color("36")),
	prompt: l.NewStyle().Foreground(l.Color("248")),
}

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	// ti.CharLimit = 156
	// ti.Width = 20
	ti.Placeholder = "http://localhost:8080"

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return styles.header.Render(cursiveLogo) + "\n\n" +
		fmt.Sprintf(
			"URL:\n%s\n\n%s",
			m.textInput.View(),
			"(esc/ctrl+c to quit)",
		) + "\n"
}

func main() {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	text := m.(model).textInput.Value()
	if text == "" {
		text = m.(model).textInput.Placeholder
	}

	method := "GET"
	var body io.Reader
	request, err := http.NewRequest(method, text, body)
	if err != nil {
		fmt.Println("Error forming request:", err)
	}
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error getting response:", err)
		return
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error forming response body:", err)
	}
	fmt.Println(string(responseBody))
}
