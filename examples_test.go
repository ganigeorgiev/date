package date_test

import (
	"fmt"
	"time"

	"github.com/ganigeorgiev/date"
)

func ExampleNewDate() {
	d1 := date.NewDate(2020, time.January, 15)
	d2 := date.NewDate(2020, time.January, 32) // auto normalized

	fmt.Println(d1)
	fmt.Println(d2)
	// Output:
	// 2020-01-15
	// 2020-02-01
}

func ExampleParse() {
	d, _ := date.Parse("2020-01-01")

	fmt.Println(d)
	// Output: 2020-01-01
}

func ExampleDate_IsZero() {
	d1 := date.Date{}
	d2, _ := date.Parse("0001-01-01")
	d3, _ := date.Parse("2020-01-01")

	fmt.Println(d1.IsZero())
	fmt.Println(d2.IsZero())
	fmt.Println(d3.IsZero())
	// Output:
	// true
	// true
	// false
}

func ExampleDate_Day() {
	d := date.NewDate(2020, time.January, 1)

	fmt.Println(d.Day())
	// Output: 1
}

func ExampleDate_Month() {
	d := date.NewDate(2020, time.January, 1)

	fmt.Println(d.Month())
	// Output: January
}

func ExampleDate_Year() {
	d := date.NewDate(2020, time.January, 1)

	fmt.Println(d.Year())
	// Output: 2020
}

func ExampleDate_Before() {
	d1 := date.NewDate(2020, time.January, 1)
	d2 := date.NewDate(2020, time.February, 1)

	fmt.Println(d1.Before(d2))
	// Output: true
}

func ExampleDate_After() {
	d1 := date.NewDate(2020, time.January, 1)
	d2 := date.NewDate(2020, time.February, 1)

	fmt.Println(d2.After(d1))
	// Output: true
}

func ExampleDate_Equal() {
	d1 := date.NewDate(2020, time.February, 1)
	d2 := date.NewDate(2020, time.January, 32)

	fmt.Println(d1.Equal(d2))
	// Output: true
}

func ExampleDate_Sub() {
	d1 := date.NewDate(2020, time.January, 1)
	d2 := date.NewDate(2020, time.January, 15)

	fmt.Println(d2.Sub(d1))
	// Output: 14
}
