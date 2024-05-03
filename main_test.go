package main

import (
	"testing"
	"go_api_timecalc/timemodule"
)

func TestEndToEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Test 1", 7.48, 7.8},
		{"Test 2", 7.30, 7.5},
		// Fügen Sie hier weitere Testfälle hinzu
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := timemodule.ConvertToDecimal(tt.input)
			if output != tt.expected {
				t.Errorf("ConvertToDecimal(%v) = %v; want %v", tt.input, output, tt.expected)
			}
		})
	}
}