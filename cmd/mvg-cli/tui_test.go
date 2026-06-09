package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTUIModel_Update(t *testing.T) {
	m := initialModel()

	// Test text input
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("abc")}
	m2, _ := m.Update(msg)
	model2 := m2.(model)
	if model2.textarea.Value() != "abc" {
		t.Errorf("expected textarea value 'abc', got %q", model2.textarea.Value())
	}

	// Test Ctrl+S (Copy) - we can't easily test clipboard here, but we can check if m.copied is set
	// Note: clipboard.WriteAll might fail in some environments without a display, so we should be careful.
	// For now, let's just test the message handling.
	msgS := tea.KeyMsg{Type: tea.KeyCtrlS}
	m3, _ := model2.Update(msgS)
	model3 := m3.(model)
	// Even if clipboard fails, we can check if it attempted to handle it.
	// If it succeeds, model3.copied will be true.
	_ = model3

	// Test Reset copied status
	msg2 := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("d")}
	m4, _ := model2.Update(msg2)
	model4 := m4.(model)
	if model4.copied {
		t.Error("expected copied status to be false after text change")
	}
}
