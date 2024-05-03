package main

import (
	"fmt"
	"go_api_timecalc/timemodule"
)

func main() {
    // Verwenden Sie Funktionen aus dem timemodule Modul
    result,err := timemodule.ProcessString("7.48 + 6.30 - ( 5.20)")
	if err != nil {
		fmt.Println(err)
	}

    fmt.Println(result)
}