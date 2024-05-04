package calendarmodule

import (
	"testing"
)

func TestNetWorkdays(t *testing.T) {
	cases := []struct {
		start, end string
		max        []int
		want       int
	}{
		{"01.01.2022", "31.12.2022", nil, 260}, // Ein ganzes Jahr, kein max
		{"01.01.2023", "31.12.2023", nil, 260}, // Ein ganzes Jahr, kein max
		{"01.01.2024", "31.12.2024", nil, 261}, // Ein ganzes Jahr, kein max
		{"01.01.2022", "31.12.2022", []int{10}, 10}, // Ein ganzes Jahr, max ist 10
		{"01.01.2022", "07.01.2022", nil, 5}, // Eine Woche, kein max
		{"01.01.2022", "07.01.2022", []int{3}, 3}, // Eine Woche, max ist 3
	}

	for _, c := range cases {
		got, err := NetWorkdays(c.start, c.end, c.max...)
		if err != nil {
			t.Errorf("NetWorkdays(%v, %v, %v) returned error: %v", c.start, c.end, c.max, err)
		} else if got != c.want {
			t.Errorf("NetWorkdays(%v, %v, %v) == %v, want %v", c.start, c.end, c.max, got, c.want)
		}
	}
}

func TestTotalDays(t *testing.T) {
	cases := []struct {
		start, end string
		want       int
	}{
		{"01.01.2022", "31.12.2022", 365}, // Ein ganzes Jahr
		{"01.01.2023", "31.12.2023", 365}, // Ein ganzes Jahr
		{"01.01.2024", "31.12.2024", 366}, // Ein Schaltjahr
		{"01.01.2022", "07.01.2022", 7},  // Eine Woche
	}

	for _, c := range cases {
		got, err := TotalDays(c.start, c.end)
		if err != nil {
			t.Errorf("TotalDays(%v, %v) returned error: %v", c.start, c.end, err)
		} else if got != c.want {
			t.Errorf("TotalDays(%v, %v) == %v, want %v", c.start, c.end, got, c.want)
		}
	}
}