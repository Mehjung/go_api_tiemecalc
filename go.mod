module go_api_tiemecalc

go 1.22.2

require go_api_timecalc/timemodule v0.0.0

require github.com/Knetic/govaluate v3.0.0+incompatible // indirect

replace go_api_timecalc/timemodule => ./timemodule

replace go_api_timecalc/calendarmodule => ./calendarmodule
