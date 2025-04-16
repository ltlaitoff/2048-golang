package core

import (
	"fmt"
	"testing"
)

type MoveTest struct {
	name     string
	board    Board
	action   string
	expected Board
}

func TestTopMove(t *testing.T) {
	var tests = []MoveTest{
		{
			name: "Base move",
			board: Board{
				{0, 0, 2, 2},
				{2, 2, 2, 0},
				{2, 0, 0, 0},
				{0, 2, 0, 2},
			},
			action: "TOP",
			expected: Board{
				{4, 4, 4, 4},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "Stacking",
			board: Board{
				{2, 0, 2, 2},
				{2, 2, 2, 0},
				{2, 2, 0, 2},
				{0, 2, 2, 2},
			},
			action: "TOP",
			expected: Board{
				{4, 4, 4, 4},
				{2, 2, 2, 2},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "Stacking with different values",
			board: Board{
				{4, 2, 2, 8},
				{2, 2, 2, 4},
				{2, 4, 2, 4},
				{0, 2, 4, 2},
			},
			action: "TOP",
			expected: Board{
				{4, 4, 4, 8},
				{4, 4, 2, 8},
				{0, 2, 4, 2},
				{0, 0, 0, 0},
			},
		},
	}

	moveTest(t, tests)
}

func TestLeftMove(t *testing.T) {
	var tests = []MoveTest{
		{
			name: "Base move",
			board: Board{
				{2, 0, 2, 0},
				{0, 2, 2, 0},
				{0, 2, 0, 2},
				{2, 0, 0, 2},
			},
			action: "LEFT",
			expected: Board{
				{4, 0, 0, 0},
				{4, 0, 0, 0},
				{4, 0, 0, 0},
				{4, 0, 0, 0},
			},
		},
		{
			name: "Stacking",
			board: Board{
				{2, 0, 2, 2},
				{2, 2, 2, 0},
				{2, 2, 0, 2},
				{0, 2, 2, 2},
			},
			action: "LEFT",
			expected: Board{
				{4, 2, 0, 0},
				{4, 2, 0, 0},
				{4, 2, 0, 0},
				{4, 2, 0, 0},
			},
		},
		{
			name: "Stacking with different values",
			board: Board{
				{8, 4, 4, 2},
				{2, 2, 2, 4},
				{2, 2, 4, 2},
				{4, 2, 2, 0},
			},
			action: "LEFT",
			expected: Board{
				{8, 8, 2, 0},
				{4, 2, 4, 0},
				{4, 4, 2, 0},
				{4, 4, 0, 0},
			},
		},
	}

	moveTest(t, tests)
}

func TestRightMove(t *testing.T) {
	var tests = []MoveTest{
		{
			name: "Base move",
			board: Board{
				{2, 0, 2, 0},
				{0, 2, 2, 0},
				{0, 2, 0, 2},
				{2, 0, 0, 2},
			},
			action: "RIGHT",
			expected: Board{
				{0, 0, 0, 4},
				{0, 0, 0, 4},
				{0, 0, 0, 4},
				{0, 0, 0, 4},
			},
		},
		{
			name: "Stacking",
			board: Board{
				{2, 0, 2, 2},
				{2, 2, 2, 0},
				{2, 2, 0, 2},
				{0, 2, 2, 2},
			},
			action: "RIGHT",
			expected: Board{
				{0, 0, 2, 4},
				{0, 0, 2, 4},
				{0, 0, 2, 4},
				{0, 0, 2, 4},
			},
		},
		{
			name: "Stacking with different values",
			board: Board{
				{2, 4, 4, 8},
				{4, 2, 2, 2},
				{2, 4, 2, 2},
				{0, 2, 2, 4},
			},
			action: "RIGHT",
			expected: Board{
				{0, 2, 8, 8},
				{0, 4, 2, 4},
				{0, 2, 4, 4},
				{0, 0, 4, 4},
			},
		},
	}

	moveTest(t, tests)
}

func TestBottomMove(t *testing.T) {
	var tests = []MoveTest{
		{
			name: "Base move",
			board: Board{
				{2, 0, 2, 0},
				{0, 2, 0, 0},
				{0, 2, 2, 2},
				{2, 0, 0, 2},
			},
			action: "BOTTOM",
			expected: Board{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{4, 4, 4, 4},
			},
		},
		{
			name: "Stacking",
			board: Board{
				{2, 0, 2, 2},
				{2, 2, 2, 0},
				{2, 2, 0, 2},
				{0, 2, 2, 2},
			},
			action: "BOTTOM",
			expected: Board{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 2, 2, 2},
				{4, 4, 4, 4},
			},
		},
		{
			name: "Stacking with different values",
			board: Board{
				{2, 4, 2, 2},
				{2, 4, 2, 4},
				{0, 8, 2, 4},
				{4, 2, 4, 8},
			},
			action: "BOTTOM",
			expected: Board{
				{0, 0, 0, 0},
				{0, 8, 2, 2},
				{4, 8, 4, 8},
				{4, 2, 4, 8},
			},
		},
	}

	moveTest(t, tests)
}

func moveTest(t *testing.T, tests []MoveTest) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initial := fmt.Sprint(tt.board)
			MoveCells(&tt.board, tt.action)

			if tt.board != tt.expected {
				t.Errorf("Invalid result!\nExpected: \n%s\nGot:\n%s\nInitial:\n%s", fmt.Sprint(tt.expected), fmt.Sprint(tt.board), initial)
			}
		})
	}
}
