package timemodule

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

func ExtractNumbers(s string) ([]float64, error) {
	re := regexp.MustCompile(`\d+[\.,:]\d+`)
	matches := re.FindAllString(s, -1)
	var numbers []float64
	for _, match := range matches {
		match = strings.ReplaceAll(match, ",", ".")
		match = strings.ReplaceAll(match, ":", ".")
		number, err := strconv.ParseFloat(match, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func ConvertToDecimal(time float64) float64 {
    hours := math.Floor(time)
    minutes := (time - hours) * 100
    decimal := hours + minutes/60
    return math.Round(decimal*100) / 100 // Runden auf 2 Dezimalstellen
}

func ConvertToTime(decimal float64) string {
    sign := ""
    if decimal < 0 {
        sign = "-"
        decimal = -decimal
    }
    hours := math.Floor(decimal)
    minutes := math.Round((decimal - hours) * 60)
    return fmt.Sprintf("%s%d:%02d", sign, int(hours), int(minutes))
}

func ReplaceTimes(s string, numbers []float64) (string, error) {
    re := regexp.MustCompile(`\d+[\.,:]\d+`)
    matches := re.FindAllStringIndex(s, -1)
    if len(matches) != len(numbers) {
        return "", errors.New("anzahl der Übereinstimmungen stimmt nicht mit der Anzahl der Zahlen überein")
    }
    for i := len(matches) - 1; i >= 0; i-- {
        start, end := matches[i][0], matches[i][1]
        s = s[:start] + strconv.FormatFloat(numbers[i], 'f', -1, 64) + s[end:]
    }
    return s, nil
}

func EvaluateExpression(expr string) (float64, error) {
    expression, err := govaluate.NewEvaluableExpression(expr)
    if err != nil {
        return 0, err
    }

    result, err := expression.Evaluate(nil)
    if err != nil {
        return 0, err
    }

    // Überprüfen Sie, ob das Ergebnis ein float64 ist
    floatResult, ok := result.(float64)
    if !ok {
        return 0, fmt.Errorf("result is not a float64")
    }

	//floatResult = math.Round(floatResult*100) / 100

    return floatResult, nil
}

func ProcessString(s string) (string, error) {
    // Schritt 1: Zahlen extrahieren
    numbers, err := ExtractNumbers(s)
    if err != nil {
        return "", err
    }

    // Schritt 2: Zahlen in Dezimal umwandeln
    for i, number := range numbers {
        numbers[i] = ConvertToDecimal(number)
    }

    // Schritt 3: Zahlen im String ersetzen
    s, err = ReplaceTimes(s, numbers)
    if err != nil {
        return "", err
    }

    // Schritt 4: Ausdruck auswerten
    result, err := EvaluateExpression(s)
    if err != nil {
        return "", err
    }

    // Schritt 5: Ergebnis in Zeit umwandeln
    return ConvertToTime(result), nil
}