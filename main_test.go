package main

import (
	"go_api_timecalc/timemodule"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      string
	}{
        {
            name:     "Addition von Zeiten",
            input:    "7:48 + 6:30",
            expected: "14:18",
            err:      "",
        },
        {
            name:     "Subtraktion und komplexe Operation",
            input:    "20:36 - 6:30 + 1:12",
            expected: "15:18",
            err:      "",
        },
        {
            name:     "1/261",
            input:    "2036/261",
            expected: "7:48",
            err:      "",
        },
		{
            name:     "Einfache Addition von Zeiten",
            input:    "3:15 + 2:45",
            expected: "6:00",
            err:      "",
        },
        {
            name:     "Subtraktion von Zeiten mit Überlauf",
            input:    "1:30 - 2:00",
            expected: "-0:30",
            err:      "",
        },
        {
            name:     "Kombination von Addition und Subtraktion",
            input:    "12:15 + 3:45 - 1:00",
            expected: "15:00",
            err:      "",
        },
        {
            name:     "Multiplikation von Zeit mit Skalar",
            input:    "2:00 * 3",
            expected: "6:00",
            err:      "",
        },
        {
            name:     "Division von Zeit durch Skalar",
            input:    "6:00 / 2",
            expected: "3:00",
            err:      "",
        },
        {
            name:     "Komplexe mathematische Operation",
            input:    "(5:00 + 1:15) * 2 - 0:30 / 1:30",
            expected: "12:10",  // Ergebnis abhängig von genauer Interpretation und Implementierung
            err:      "",
        },

        {
            name:     "Zeitüberschreitung bei Minuten",
            input:    "0:65",
            expected: "1:05",
            err:      "",
        },
        {
            name:     "Negative Zeit in Ausdruck",
            input:    "1:00 - 1:30",
            expected: "-0:30",
            err:      "",
        },
        {
            name:     "Division von Zeit durch Zeit",
            input:    "2:00 / 0:30",
            expected: "4:00",
            err:      "",
        },
		{
            name:     "eT Test 1",
            input:    "-39:00 / 5",
            expected: "-7:48",
            err:      "",
        },
		{
            name:     "eT Test 1",
            input:    "7:48 / (-2)",
            expected: "-3:54",
            err:      "",
        },
	}

	sp := timemodule.NewStringProcessor(
		new(timemodule.SimpleNumberExtractor),
		new(timemodule.SimpleTimeConverter),
		new(timemodule.SimpleExpressionEvaluator),
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := sp.ProcessString(tt.input)
			if (err != nil && tt.err == "") || (err == nil && tt.err != "") {
				t.Errorf("%s: Fehlerstatus inkorrekt. Erhalten: %v, Erwartet: %v", tt.name, err, tt.err)
			} else if err != nil && err.Error() != tt.err {
				t.Errorf("%s: falsche Fehlermeldung erhalten: %v, erwartet: %v", tt.name, err, tt.err)
			}
			if result != tt.expected {
				t.Errorf("%s: Erhalten: %v, Erwartet: %v", tt.name, result, tt.expected)
			}
		})
	}
}
