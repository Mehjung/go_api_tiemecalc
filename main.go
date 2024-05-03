package main

import (
	"fmt"
	"go_api_timecalc/timemodule"
)

func main() {
    sp := timemodule.NewDefaultStringProcessor()
    result, err := sp.ProcessString("7.48 + 6.30 - ( 5.20)")
    if err != nil {
        fmt.Println("Fehler beim Verarbeiten des Strings:", err)
        return
    }

    fmt.Println("Ergebnis:", result)
}