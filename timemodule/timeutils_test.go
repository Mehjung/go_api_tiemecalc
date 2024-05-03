package timemodule

import (
	"reflect"
	"testing"
)

func TestExtractNumbers(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []float64
		wantErr bool
	}{
		{
			name:    "Test 1",
			input:   "7,48 + 6:30 - ( 5.20)",
			want:    []float64{7.48, 6.3, 5.2},
			wantErr: false,
		},
		{
			name:    "Test 2",
			input:   "1.23 * 4,56 / 7:89",
			want:    []float64{1.23, 4.56, 7.89},
			wantErr: false,
		},
		{
			name:    "Test 3",
			input:   "no numbers here",
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractNumbers(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestConvertToDecimal(t *testing.T) {
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
            output := ConvertToDecimal(tt.input)
            if output != tt.expected {
                t.Errorf("ConvertToDecimal(%v) = %v; want %v", tt.input, output, tt.expected)
            }
        })
    }
}

func TestConvertToTime(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"Test 1", 7.8, "7:48"},
		{"Test 2", 7.5, "7:30"},
		{"Test 3", -7.5, "-7:30"},
		// Fügen Sie hier weitere Testfälle hinzu
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ConvertToTime(tt.input)
			if output != tt.expected {
				t.Errorf("ConvertToTime(%v) = %v; want %v", tt.input, output, tt.expected)
			}
		})
	}
}

func TestReplaceTimes(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        numbers  []float64
        expected string
        err      string
    }{
        {"Test 1", "7,48 + 5.30 - (6.30)", []float64{7.8, 5.5, 6.5}, "7.8 + 5.5 - (6.5)", ""},
        {"Test 2", "07:48", []float64{1.23}, "1.23", ""},
        {"Test 3", "07:48 und 7.48", []float64{1.23}, "", "anzahl der Übereinstimmungen stimmt nicht mit der Anzahl der Zahlen überein"},
        {"Test 4", "7,48 * 5.30 / (6.30)", []float64{1.1, 2.2, 3.3}, "1.1 * 2.2 / (3.3)", ""},
        {"Test 5", "7,48 + 5.30 - (6.30) * 7.48 / 5.30", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, "1.1 + 2.2 - (3.3) * 4.4 / 5.5", ""},
        {"Test 6", "7,48", []float64{1.23, 4.56}, "", "anzahl der Übereinstimmungen stimmt nicht mit der Anzahl der Zahlen überein"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            output, err := ReplaceTimes(tt.input, tt.numbers)
            if output != tt.expected || (err != nil && err.Error() != tt.err) {
                t.Errorf("ReplaceTimes(%v, %v) = %v, %v; want %v, %v", tt.input, tt.numbers, output, err, tt.expected, tt.err)
            }
        })
    }
}


func TestEvaluateExpression(t *testing.T) {
    tests := []struct {
        name    string
        expr    string
        want    float64
        wantErr bool
    }{
        {
            name:    "Test 1: Einfacher Ausdruck",
            expr:    "2+2",
            want:    4,
            wantErr: false,
        },
        {
            name:    "Test 2: Komplexer Ausdruck",
            expr:    "(3+5)*2",
            want:    16,
            wantErr: false,
        },
        {
            name:    "Test 3: Ungültiger Ausdruck",
            expr:    "2+",
            want:    0,
            wantErr: true,
        },
        {
            name:    "Test 4: Nicht numerisches Ergebnis",
            expr:    "'test'+2",
            want:    0,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := EvaluateExpression(tt.expr)
            if (err != nil) != tt.wantErr {
                t.Errorf("EvaluateExpression() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("EvaluateExpression() = %v, want %v", got, tt.want)
            }
        })
    }
}