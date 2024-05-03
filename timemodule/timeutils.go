package timemodule

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

// NumberExtractor definiert ein Interface für das Extrahieren von Zahlen aus Strings.
type NumberExtractor interface {
    ExtractNumbers(string) ([]float64, error)
}

// SimpleNumberExtractor implementiert das NumberExtractor Interface.
type SimpleNumberExtractor struct{}

func (ne *SimpleNumberExtractor) ExtractNumbers(s string) ([]float64, error) {
    re := regexp.MustCompile(`\d+([.,:]\d+)?|\d+`)
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

// TimeConverter definiert ein Interface für die Konvertierung zwischen Dezimal- und Zeitformaten.
type TimeConverter interface {
    ConvertToDecimal(time float64) float64
    ConvertToTime(decimal float64) string
}

// SimpleTimeConverter implementiert das TimeConverter Interface.
type SimpleTimeConverter struct{}

func (tc *SimpleTimeConverter) ConvertToDecimal(time float64) float64 {
    hours := math.Floor(time)
    minutes := (time - hours) * 100
    decimal := hours + minutes / 60
    return math.Round(decimal * 100) / 100
}

func (tc *SimpleTimeConverter) ConvertToTime(decimal float64) string {
    sign := ""
    if decimal < 0 {
        sign = "-"
        decimal = -decimal
    }
    hours := math.Floor(decimal)
    minutes := math.Round((decimal - hours) * 60)
    return fmt.Sprintf("%s%d:%02d", sign, int(hours), int(minutes))
}

// ExpressionEvaluator definiert ein Interface für das Auswerten mathematischer Ausdrücke.
type ExpressionEvaluator interface {
    EvaluateExpression(expr string) (float64, error)
}

// SimpleExpressionEvaluator implementiert das ExpressionEvaluator Interface.
type SimpleExpressionEvaluator struct{}

func (ee *SimpleExpressionEvaluator) EvaluateExpression(expr string) (float64, error) {
    expression, err := govaluate.NewEvaluableExpression(expr)
    if err != nil {
        return 0, err
    }

    result, err := expression.Evaluate(nil)
    if err != nil {
        return 0, err
    }

    floatResult, ok := result.(float64)
    if !ok {
        return 0, fmt.Errorf("result is not a float64")
    }

    return floatResult, nil
}

func ReplaceTimes(s string, numbers []float64) (string, error) {
    re := regexp.MustCompile(`\d+([.,:]\d+)?|\d+`)
    matches := re.FindAllStringIndex(s, -1)
    if len(matches) != len(numbers) {
        return "", fmt.Errorf("fehler: Anzahl der Übereinstimmungen (%d) stimmt nicht mit der Anzahl der Zahlen (%d) überein", len(matches), len(numbers))
    }
    s = replaceNumbersInString(s, matches, numbers)
    return s, nil
}

// replaceNumbersInString führt die eigentliche Ersetzung der Zahlen im String durch.
func replaceNumbersInString(s string, matches [][]int, numbers []float64) string {
    for i := len(matches) - 1; i >= 0; i-- {
        start, end := matches[i][0], matches[i][1]
        numberStr := strconv.FormatFloat(numbers[i], 'f', -1, 64)
        s = s[:start] + numberStr + s[end:]
    }
    return s
}


// StringProcessor beinhaltet alle Abhängigkeiten, die zum Verarbeiten von Strings benötigt werden.
type StringProcessor struct {
    NumberExtractor
    TimeConverter
    ExpressionEvaluator
}

// NewStringProcessor erstellt eine neue Instanz von StringProcessor mit den gegebenen Abhängigkeiten.
func NewStringProcessor(ne NumberExtractor, tc TimeConverter, ee ExpressionEvaluator) *StringProcessor {
    return &StringProcessor{
        NumberExtractor:    ne,
        TimeConverter:      tc,
        ExpressionEvaluator: ee,
    }
}

// ProcessString verarbeitet den gegebenen String und gibt das Ergebnis zurück.
func (sp *StringProcessor) ProcessString(s string) (string, error) {
    numbers, err := sp.ExtractNumbers(s)
    if err != nil {
        return "", err
    }

    for i, number := range numbers {
        numbers[i] = sp.ConvertToDecimal(number)
    }

    s, err = ReplaceTimes(s, numbers)
    if err != nil {
        return "", err
    }

    result, err := sp.EvaluateExpression(s)
    if err != nil {
        return "", err
    }

    return sp.ConvertToTime(result), nil
}

func NewDefaultStringProcessor() *StringProcessor {
    return NewStringProcessor(
        new(SimpleNumberExtractor),
        new(SimpleTimeConverter),
        new(SimpleExpressionEvaluator),
    )
}