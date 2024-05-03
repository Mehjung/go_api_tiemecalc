package timemodule

import (
	"reflect"
	"testing"
)

func TestExtractNumbers(t *testing.T) {
    extractor := &SimpleNumberExtractor{} // Verwendung der implementierten Struct
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
            got, err := extractor.ExtractNumbers(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("ExtractNumbers() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ExtractNumbers() = %v, want %v", got, tt.want)
            }
        })
    }
}
func TestConvertToDecimal(t *testing.T) {
    converter := &SimpleTimeConverter{} // Verwendung der implementierten Struct
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
            output := converter.ConvertToDecimal(tt.input)
            if output != tt.expected {
                t.Errorf("ConvertToDecimal(%v) = %v; want %v", tt.input, output, tt.expected)
            }
        })
    }
}


func TestConvertToTime(t *testing.T) {
    converter := &SimpleTimeConverter{} // Verwendung der implementierten Struct
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
            output := converter.ConvertToTime(tt.input)
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
        {
            name:     "Standard-Ersetzung",
            input:    "12.34 und 56,78",
            numbers:  []float64{1.11, 2.22},
            expected: "1.11 und 2.22",
            err:      "",
        },
        {
            name:     "Keine Übereinstimmungen im String",
            input:    "Keine Zahlen hier",
            numbers:  []float64{},
            expected: "Keine Zahlen hier",
            err:      "",
        },
        {
            name:     "Mehr Zahlen als Muster",
            input:    "12,34",
            numbers:  []float64{1.23, 4.56, 7.89},
            expected: "",
            err:      "fehler: Anzahl der Übereinstimmungen (1) stimmt nicht mit der Anzahl der Zahlen (3) überein",
        },
        {
            name:     "Weniger Zahlen als Muster",
            input:    "12.34 und 56.78 sowie 90,12",
            numbers:  []float64{1.23},
            expected: "",
            err:      "fehler: Anzahl der Übereinstimmungen (3) stimmt nicht mit der Anzahl der Zahlen (1) überein",
        },
        {
            name:     "Negative Zahlen ersetzen",
            input:    "12,34 - 56.78",
            numbers:  []float64{-1.23, -0.56},
            expected: "-1.23 - -0.56",
            err:      "",
        },
        {
            name:     "Null als Eingabe",
            input:    "0.00 und 00,00",
            numbers:  []float64{0, 0},
            expected: "0 und 0",
            err:      "",
        },
        {
            name:     "Ungültige Eingaben ignoriert",
            input:    "Dies ist ein Test, keine Zahlen",
            numbers:  []float64{1.23, 4.56}, // Zahlen bereitgestellt, aber keine Muster im String
            expected: "",
            err:      "fehler: Anzahl der Übereinstimmungen (0) stimmt nicht mit der Anzahl der Zahlen (2) überein",
        },        
        {
            name:     "Präzise Ersetzung in komplexen Strings",
            input:    "Der Wert 123.456 soll geändert werden",
            numbers:  []float64{78.90},
            expected: "Der Wert 78.9 soll geändert werden",
            err:      "",
        },
        {
            name:     "Mischung aus kompatiblen und inkompatiblen Zahlen",
            input:    "123.456 und 789",
            numbers:  []float64{12.34},
            expected: "",
            err:      "fehler: Anzahl der Übereinstimmungen (2) stimmt nicht mit der Anzahl der Zahlen (1) überein", // Beispiel zeigt fehlerhafte Logik im Testbeschreibung
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            output, err := ReplaceTimes(tt.input, tt.numbers)
            if err == nil && tt.err != "" {
                t.Errorf("%s: erwarteter Fehler wurde nicht erhalten.", tt.name)
            } else if err != nil && err.Error() != tt.err {
                t.Errorf("%s: falsche Fehlermeldung erhalten: %v, erwartet: %v", tt.name, err, tt.err)
            }
            if output != tt.expected {
                t.Errorf("%s: Erhalten: %v, Erwartet: %v", tt.name, output, tt.expected)
            }
        })
    }
}





func TestEvaluateExpression(t *testing.T) {
    evaluator := &SimpleExpressionEvaluator{} // Verwendung der implementierten Struct
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
            got, err := evaluator.EvaluateExpression(tt.expr)
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
