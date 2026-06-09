package main

import (
	"fmt"
	"mvg-fontifier/internal/transformer"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type flashMsg struct{}

type model struct {
	textarea textarea.Model
	err      error
	copied   bool
	flashing bool
}

func initialModel() model {
	ti := textarea.New()
	ti.Placeholder = "Type something to fontify..."
	ti.Focus()

	return model{
		textarea: ti,
		err:      nil,
		copied:   false,
		flashing: false,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyCtrlS:
			// Copy transformed text to clipboard
			output := transformer.Transform(m.textarea.Value())
			err := clipboard.WriteAll(output)
			if err != nil {
				m.err = err
			} else {
				m.copied = true
				m.flashing = true
				// Reset flashing after 500ms
				cmds = append(cmds, tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
					return flashMsg{}
				}))
			}
			return m, tea.Batch(cmds...)
		}

	case flashMsg:
		m.flashing = false
		return m, nil

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	// Reset copied status if text changes
	if m.textarea.Value() != "" {
		m.copied = false
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	transformed := transformer.Transform(m.textarea.Value())

	var status string
	if m.err != nil {
		status = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(fmt.Sprintf("Error: %v", m.err))
	} else if m.copied {
		status = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render("Copied to clipboard!")
	} else {
		status = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render("Ctrl+S: Copy | Esc: Quit")
	}

	// Update textarea prompt/bar color if flashing
	if m.flashing {
		m.textarea.FocusedStyle.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF007F"))
	} else {
		m.textarea.FocusedStyle.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	}

	return fmt.Sprintf(
		"Enter text to transform:\n\n%s\n\nResult:\n\n%s\n\n%s",
		m.textarea.View(),
		lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Render(transformed),
		status,
	) + "\n"
}

func runTUI() error {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
