package calendarmodule

import (
	"fmt"
	"go_api_timecalc/timemodule"
	"time"
)

func NetWorkdays(start, end string, max ...int) (int, error) {
	totalDays := 0
	maxDays := 261 // Default-Wert

	if len(max) > 0 {
		maxDays = max[0] // Wenn ein max-Wert angegeben ist, verwenden Sie diesen
	}

	// Parse deutsche Datumsstrings
	startDate, err := time.Parse("02.01.2006", start)
	if err != nil {
		return 0, err
	}
	endDate, err := time.Parse("02.01.2006", end)
	if err != nil {
		return 0, err
	}

	if startDate.After(endDate) {
		return 0, fmt.Errorf("startdatum darf nicht nach dem Enddatum liegen")
	}

	for d := startDate; d.Before(endDate) || d.Equal(endDate); {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			totalDays++
			if totalDays >= maxDays {
				return maxDays, nil
			}
		}
		d = d.AddDate(0, 0, 1)
	}
	return totalDays, nil
}

func TotalDays(start, end string) (int, error) {
	// Parse deutsche Datumsstrings
	startDate, err := time.Parse("02.01.2006", start)
	if err != nil {
		return 0, err
	}
	endDate, err := time.Parse("02.01.2006", end)
	if err != nil {
		return 0, err
	}

	// Berechne die Differenz in Stunden und teile durch 24, um die Differenz in Tagen zu erhalten
	diff := endDate.Sub(startDate).Hours() / 24

	return int(diff) + 1, nil // Addiere 1, um das Enddatum einzuschließen
}

func getIJAZbyDate(start, end , jaz  string) (string, error){
	sp := timemodule.NewDefaultStringProcessor()
    result, err := sp.ProcessString("7.48 + 6.30 - ( 5.20)")	// Berechnen Sie die Anzahl der JAZ-Tage, die auf die Anzahl der Tage verteilt werden
	if err != nil {
		return "", err
	}

	return result, nil
}