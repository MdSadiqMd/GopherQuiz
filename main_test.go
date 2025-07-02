package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestParseLines(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected []problem
	}{
		{
			name: "normal input",
			input: [][]string{
				{"2+2", "4"},
				{"capital of France", " Paris "},
			},
			expected: []problem{
				{question: "2+2", answer: "4"},
				{question: "capital of France", answer: "Paris"},
			},
		},
		{
			name:     "empty input",
			input:    [][]string{},
			expected: []problem{},
		},
		{
			name: "whitespace trimming",
			input: [][]string{
				{"Q", "  a  "},
			},
			expected: []problem{
				{question: "Q", answer: "a"},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseLines(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestCheckAnswer(t *testing.T) {
	p := problem{question: "2+2", answer: "4"}
	if !checkAnswer(p, "4") {
		t.Error("Expected answer to be correct")
	}
	if checkAnswer(p, "5") {
		t.Error("Expected answer to be incorrect")
	}
}

func checkAnswer(p problem, answer string) bool {
	return p.answer == answer
}
