package core

import (
	"testing"
)

func TestGameEnd(t *testing.T) {
	var tests = []struct {
		name     string
		input    Board
		expected bool
	}{
		{
			name: "Check on end table",
			input: Board{
				{2, 4, 2, 4},
				{4, 2, 4, 2},
				{2, 4, 2, 4},
				{4, 2, 4, 2},
			},
			expected: true,
		},
		{
			name: "Check with not-full board",
			input: Board{
				{2, 4, 2, 4},
				{4, 2, 4, 2},
				{2, 4, 2, 4},
				{4, 2, 4, 0},
			},
			expected: false,
		},
		{
			name: "Check with repeat values in row",
			input: Board{
				{2, 4, 2, 2},
				{4, 2, 4, 2},
				{2, 4, 2, 4},
				{4, 2, 4, 2},
			},
			expected: false,
		},
		{
			name: "Check with repeat values in columt",
			input: Board{
				{16, 4, 2, 4},
				{2, 8, 4, 2},
				{2, 4, 2, 4},
				{4, 2, 4, 2},
			},
			expected: false,
		},
		{
			name: "Check with real table",
			input: Board{
				{4, 2, 8, 2},
				{2, 32, 16, 4},
				{2, 32, 16, 2},
				{2, 8, 2, 4},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := GameEnd(&tt.input)

			if res != tt.expected {
				t.Errorf("Invalid result of game end, expected %t, but got %t", tt.expected, res)
			}
		})
	}
}
