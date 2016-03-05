package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello, we are Holberton School")

	// Get the current time.
	t := time.Now()
	fmt.Println("the date is", t)

	// Print year, month and day of the current time.
	fmt.Println("the year is", t.Year())
	fmt.Println("the month is", t.Month())
	fmt.Println("the day is", t.Day())
}
