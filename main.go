package main

import (
	"fmt"
	"go_api_timecalc/timemodule"
	"net/http"
	"sync"
)

func main() {
    http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
        var wg sync.WaitGroup
        wg.Add(1)

        go func() {
            defer wg.Done()

            sp := timemodule.NewDefaultStringProcessor()
            result, err := sp.ProcessString("7.48 + 6.30 - ( 5.20)")
            if err != nil {
                http.Error(w, "Fehler beim Verarbeiten des Strings: "+err.Error(), http.StatusInternalServerError)
                return
            }
            fmt.Fprintf(w, "Ergebnis: %v", result)
        }()

        wg.Wait()
    })

    fmt.Println("Server läuft auf Port 8080")
    http.ListenAndServe(":8080", nil)
}