package transformer

import "testing"

func TestTransform(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Full alphabet substitution",
			input:    "ADEFJNST acdfjmnosuvz",
			expected: "ΔĐΣϜJͶϞͲ αϲɗϝʝϻηoϟυvʐ",
		},
		{
			name:     "Mixed characters",
			input:    "Hello World!",
			expected: "Hello Worlɗ!",
		},
		{
			name:     "No substitution",
			input:    "12345!@#",
			expected: "12345!@#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transform(tt.input); got != tt.expected {
				t.Errorf("Transform() = %v, want %v", got, tt.expected)
			}
		})
	}
}
