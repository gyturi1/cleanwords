package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidWord(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"T1", "ablak", true},
		{"T2", "darazs", true},
		{"T3", "cscscscscs", true},
		{"T4", "ágyás", true},
		{"T5", "ááááá", true},
		{"T6", "szabály", true},
		{"T7", "alak", false},
		{"T8", "hadonászott", false},
		{"T9", "*akarmi", false},
		{"T10", "_barmi", false},
		{"T10", "bar_mi", false},
		{"T10", "kölcsön_", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidWord(tt.arg); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidLetter(t *testing.T) {
	tests := []struct {
		name string
		arg  rune
		want bool
	}{
		{"T1", 'a', true},
		{"T2", 'z', true},
		{"T3", '2', false},
		{"T3", '*', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidLetter(tt.arg); got != tt.want {
				t.Errorf("isValidLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possibleLetterCounts(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"T0", "", []int{0}},
		{"T1", "elv", []int{3}},
		{"T2", "arat", []int{4}},
		{"T3", "zseb", []int{4, 3}},
		{"T4", "zsálya", []int{6, 5, 5, 4}},
		{"T5", "kérész", []int{6, 5}},
		{"T5", "dzsúsz", []int{6, 5, 5, 4, 5, 4, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := possibleLetterCounts(tt.s)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
