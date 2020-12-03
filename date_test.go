package date

import (
	"bytes"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	testScenarios := []struct {
		year     int
		month    time.Month
		day      int
		expected string
	}{
		{-1, time.January, 1, "-0001-01-01"},
		{2020, time.January, 1, "2020-01-01"},
		{2020, time.January, 32, "2020-02-01"},
		{2020, time.January, 0, "2019-12-31"},
		{2020, time.January, -1, "2019-12-30"},
	}

	for _, scenario := range testScenarios {
		d := NewDate(scenario.year, scenario.month, scenario.day)

		if d.String() != scenario.expected {
			t.Errorf("Expected %s, got %s", scenario.expected, d.String())
		}
	}
}

func TestParse(t *testing.T) {
	testScenarios := []struct {
		value     string
		expectErr bool
	}{
		// invalid
		{"2020/01/02", true},
		{"01/2020/02", true},
		{"01-02-2020", true},
		{"01-02", true},
		{"2020 01 02", true},
		{"-0001-01-01", true},
		{"2020-13-01", true},
		{"2020-01-50", true},
		{"2020-01-02 10:00:00", true},
		// valid
		{"0001-01-01", false},
		{"2020-01-31", false},
	}

	for _, scenario := range testScenarios {
		d, err := Parse(scenario.value)

		if !scenario.expectErr && err != nil {
			t.Errorf("Wasn't expecting error, got %v", err)
		} else if scenario.expectErr && err == nil {
			t.Error("Expected error, got nil")
		}

		var expectedVal string
		if scenario.expectErr {
			expectedVal = "0001-01-01"
		} else {
			expectedVal = scenario.value
		}

		if d.String() != expectedVal {
			t.Errorf("Expected %s, got %s", expectedVal, d.String())
		}
	}
}

func TestFormat(t *testing.T) {
	testScenarios := []struct {
		date     Date
		layout   string
		expected string
	}{
		{NewDate(2020, time.January, 1), "January 02, 2006", "January 01, 2020"},
		{NewDate(2020, time.December, 1), "Mon, 02 Jan 2006", "Tue, 01 Dec 2020"},
	}

	for _, scenario := range testScenarios {
		if scenario.date.Format(scenario.layout) != scenario.expected {
			t.Errorf("Expected %s, got %s", scenario.expected, scenario.date.String())
		}
	}
}

func TestString(t *testing.T) {
	testScenarios := []struct {
		date     Date
		expected string
	}{
		{NewDate(-1, time.January, 1), "-0001-01-01"},
		{NewDate(2020, time.January, 1), "2020-01-01"},
		{NewDate(2020, time.January, 32), "2020-02-01"},
		{NewDate(2020, time.January, 0), "2019-12-31"},
		{NewDate(2020, time.January, -1), "2019-12-30"},
	}

	for _, scenario := range testScenarios {
		if scenario.date.String() != scenario.expected {
			t.Errorf("Expected %s, got %s", scenario.expected, scenario.date.String())
		}
	}
}

func TestIsZero(t *testing.T) {
	testScenarios := []struct {
		date       Date
		expectZero bool
	}{
		{NewDate(1, time.January, 1), true},
		{NewDate(-1, time.January, 1), false},
		{NewDate(2020, time.January, 1), false},
	}

	for _, scenario := range testScenarios {
		if scenario.date.IsZero() != scenario.expectZero {
			t.Errorf("Expected %v, got %v", scenario.expectZero, scenario.date.IsZero())
		}
	}
}

func TestDay(t *testing.T) {
	testScenarios := []struct {
		date     Date
		expected int
	}{
		{NewDate(-1, time.January, 2), 2},
		{NewDate(2020, time.January, 32), 1},
		{NewDate(2020, time.January, 0), 31},
		{NewDate(2020, time.January, -1), 30},
	}

	for _, scenario := range testScenarios {
		if day := scenario.date.Day(); day != scenario.expected {
			t.Errorf("Expected %d, got %d for date %v", scenario.expected, day, scenario.date)
		}
	}
}

func TestMonth(t *testing.T) {
	testScenarios := []struct {
		date     Date
		expected time.Month
	}{
		{NewDate(-1, time.January, 2), time.January},
		{NewDate(2020, time.January, 32), time.February},
		{NewDate(2020, time.January, 0), time.December},
		{NewDate(2020, time.January, -1), time.December},
	}

	for _, scenario := range testScenarios {
		if month := scenario.date.Month(); month != scenario.expected {
			t.Errorf("Expected %v, got %v for date %v", scenario.expected, month, scenario.date)
		}
	}
}

func TestYear(t *testing.T) {
	testScenarios := []struct {
		date     Date
		expected int
	}{
		{NewDate(0, time.January, 1), 0},
		{NewDate(-1, time.January, 1), -1},
		{NewDate(2020, time.December, 32), 2021},
		{NewDate(2020, time.January, 0), 2019},
		{NewDate(2020, time.January, -1), 2019},
	}

	for _, scenario := range testScenarios {
		if year := scenario.date.Year(); year != scenario.expected {
			t.Errorf("Expected %d, got %d for date %v", scenario.expected, year, scenario.date)
		}
	}
}

func TestBefore(t *testing.T) {
	testScenarios := []struct {
		dateA    Date
		dateB    Date
		expected bool
	}{
		{NewDate(2020, time.January, 2), NewDate(2020, time.January, 1), false},
		{NewDate(2020, time.February, 1), NewDate(2020, time.January, 2), false},
		{NewDate(2020, time.January, 1), NewDate(2019, time.February, 2), false},
		{NewDate(-1, time.January, 1), NewDate(-2, time.January, 1), false},
		{NewDate(2020, time.January, 1), NewDate(2020, time.January, 2), true},
		{NewDate(2020, time.January, 32), NewDate(2020, time.February, 2), true},
		{NewDate(2019, time.January, 1), NewDate(2020, time.January, 1), true},
		{NewDate(-2, time.January, 1), NewDate(-1, time.January, 1), true},
	}

	for _, scenario := range testScenarios {
		if scenario.dateA.Before(scenario.dateB) != scenario.expected {
			if scenario.expected {
				t.Errorf("Expected %v to be before %v", scenario.dateA, scenario.dateB)
			} else {
				t.Errorf("Wasn't expecting %v to be before %v", scenario.dateA, scenario.dateB)
			}
		}
	}
}

func TestAfter(t *testing.T) {
	testScenarios := []struct {
		dateA    Date
		dateB    Date
		expected bool
	}{
		{NewDate(2020, time.January, 1), NewDate(2020, time.January, 2), false},
		{NewDate(2020, time.January, 2), NewDate(2020, time.February, 1), false},
		{NewDate(2019, time.February, 2), NewDate(2020, time.January, 1), false},
		{NewDate(-2, time.January, 1), NewDate(-1, time.January, 1), false},
		{NewDate(2020, time.January, 2), NewDate(2020, time.January, 1), true},
		{NewDate(2020, time.February, 2), NewDate(2020, time.January, 32), true},
		{NewDate(2020, time.January, 1), NewDate(2019, time.January, 1), true},
		{NewDate(-1, time.January, 1), NewDate(-2, time.January, 1), true},
	}

	for _, scenario := range testScenarios {
		if scenario.dateA.After(scenario.dateB) != scenario.expected {
			if scenario.expected {
				t.Errorf("Expected %v to be after %v", scenario.dateA, scenario.dateB)
			} else {
				t.Errorf("Wasn't expecting %v to be after %v", scenario.dateA, scenario.dateB)
			}
		}
	}
}

func TestEqual(t *testing.T) {
	testScenarios := []struct {
		dateA    Date
		dateB    Date
		expected bool
	}{
		{NewDate(2020, time.January, 1), NewDate(2020, time.January, 2), false},
		{NewDate(2020, time.January, 1), NewDate(2020, time.February, 1), false},
		{NewDate(2019, time.January, 2), NewDate(2020, time.January, 1), false},
		{NewDate(-2, time.January, 1), NewDate(-1, time.January, 1), false},
		{NewDate(2020, time.January, 1), NewDate(2020, time.January, 1), true},
		{NewDate(2020, time.February, 1), NewDate(2020, time.February, 1), true},
		{NewDate(2020, time.February, 1), NewDate(2020, time.January, 32), true},
		{NewDate(2019, time.December, 31), NewDate(2020, time.January, 0), true},
		{NewDate(2019, time.December, 30), NewDate(2020, time.January, -1), true},
		{NewDate(-1, time.January, 1), NewDate(-1, time.January, 1), true},
	}

	for _, scenario := range testScenarios {
		if scenario.dateA.Equal(scenario.dateB) != scenario.expected {
			if scenario.expected {
				t.Errorf("Expected %v to be equal to %v", scenario.dateA, scenario.dateB)
			} else {
				t.Errorf("Wasn't expecting %v to be equal to %v", scenario.dateA, scenario.dateB)
			}
		}
	}
}

func TestSub(t *testing.T) {
	testScenarios := []struct {
		dateA    Date
		dateB    Date
		expected int
	}{
		{NewDate(2020, time.January, 1), NewDate(2020, time.January, 2), -1},
		{NewDate(2020, time.January, 3), NewDate(2020, time.January, 1), 2},
		{NewDate(2020, time.January, 1), NewDate(2020, time.January, 1), 0},
		{NewDate(2021, time.January, 1), NewDate(2020, time.January, 1), 366},
	}

	for _, scenario := range testScenarios {
		if result := scenario.dateA.Sub(scenario.dateB); result != scenario.expected {
			t.Errorf("Expected %v - %v to be %v, got %v", scenario.dateA, scenario.dateB, scenario.expected, result)
		}
	}
}

func TestScan(t *testing.T) {
	testScenarios := []struct {
		rawValue      interface{}
		expectedValue string
		expectErr     bool
	}{
		{[]byte("invalid"), "0001-01-01", true},
		{"invalid", "0001-01-01", true},
		{"2020-13-01", "0001-01-01", true},
		{"2020-01-32", "0001-01-01", true},
		{"-0001-01-01", "0001-01-01", true},
		{"2020-01-01 10:00:00", "0001-01-01", true},
		{[]int{2020, 1, 2}, "0001-01-01", true},
		{nil, "0001-01-01", false},
		{"", "0001-01-01", false},
		{[]byte(""), "0001-01-01", false},
		{[]byte("2020-01-02"), "2020-01-02", false},
		{"2020-01-02", "2020-01-02", false},
		{time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC), "2020-01-01", false},
	}

	for _, scenario := range testScenarios {
		d := Date{}
		err := d.Scan(scenario.rawValue)

		if !scenario.expectErr && err != nil {
			t.Errorf("Wasn't expecting error, got %v", err)
		} else if scenario.expectErr && err == nil {
			t.Error("Expected error, got nil")
		}

		if d.String() != scenario.expectedValue {
			t.Errorf("Expected %s, got %s for %v", scenario.expectedValue, d.String(), scenario.rawValue)
		}

		if d.t.Hour() != 0 ||
			d.t.Minute() != 0 ||
			d.t.Second() != 0 ||
			d.t.Nanosecond() != 0 ||
			d.t.Location() != time.UTC {
			t.Errorf("Expected the internal time data to be zero with UTC location, got %v", d.t)
		}
	}
}

func TestValue(t *testing.T) {
	testScenarios := []struct {
		date     Date
		expected time.Time
	}{
		{Date{}, time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{NewDate(2020, time.January, 1), time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{NewDate(2020, time.January, 32), time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC)},
		{NewDate(2020, time.January, 0), time.Date(2019, time.December, 31, 0, 0, 0, 0, time.UTC)},
		{NewDate(2020, time.January, -1), time.Date(2019, time.December, 30, 0, 0, 0, 0, time.UTC)},
	}

	for _, scenario := range testScenarios {
		val, err := scenario.date.Value()

		if err != nil {
			t.Errorf("Wasn't expecting an error, got %v", err)
		}

		if val != scenario.expected {
			t.Errorf("Expected %v, got %v", scenario.expected, val)
		}
	}
}

func TestMarshalText(t *testing.T) {
	testScenarios := []struct {
		date     Date
		expected []byte
	}{
		{Date{}, []byte("0001-01-01")},
		{NewDate(-1, time.January, 1), []byte("-0001-01-01")},
		{NewDate(2020, time.January, 1), []byte("2020-01-01")},
		{NewDate(2020, time.January, 32), []byte("2020-02-01")},
		{NewDate(2020, time.January, 0), []byte("2019-12-31")},
		{NewDate(2020, time.January, -1), []byte("2019-12-30")},
	}

	for _, scenario := range testScenarios {
		val, err := scenario.date.MarshalText()

		if err != nil {
			t.Errorf("Wasn't expecting an error, got %v", err)
		}

		if bytes.Compare(val, scenario.expected) != 0 {
			t.Errorf("Expected %v, got %v", scenario.expected, val)
		}
	}
}

func TestUnmarshalText(t *testing.T) {
	testScenarios := []struct {
		bytes         []byte
		expectedValue string
		expectErr     bool
	}{
		// invalid
		{[]byte("2020/01/02"), "0001-01-01", true},
		{[]byte("01/2020/02"), "0001-01-01", true},
		{[]byte("01-02-2020"), "0001-01-01", true},
		{[]byte("01-02"), "0001-01-01", true},
		{[]byte("2020 01 02"), "0001-01-01", true},
		{[]byte("-0001-01-01"), "0001-01-01", true},
		{[]byte("2020-13-01"), "0001-01-01", true},
		{[]byte("2020-01-50"), "0001-01-01", true},
		{[]byte("2020-01-02 10:00:00"), "0001-01-01", true},
		// valid
		{[]byte(""), "0001-01-01", false},
		{nil, "0001-01-01", false},
		{[]byte("0001-01-01"), "0001-01-01", false},
		{[]byte("2020-01-31"), "2020-01-31", false},
	}

	for _, scenario := range testScenarios {
		d := Date{}
		err := d.UnmarshalText(scenario.bytes)

		if !scenario.expectErr && err != nil {
			t.Errorf("Wasn't expecting error, got %v", err)
		} else if scenario.expectErr && err == nil {
			t.Error("Expected error, got nil")
		}

		if d.String() != scenario.expectedValue {
			t.Errorf("Expected %s, got %s for %v", scenario.expectedValue, d.String(), scenario.bytes)
		}

		if d.t.Hour() != 0 ||
			d.t.Minute() != 0 ||
			d.t.Second() != 0 ||
			d.t.Nanosecond() != 0 ||
			d.t.Location() != time.UTC {
			t.Errorf("Expected the internal time data to be zero with UTC location, got %v", d.t)
		}
	}
}
