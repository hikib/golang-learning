package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	yearNum, weekNum := time.Now().ISOWeek()
	year := strconv.Itoa(yearNum)
	fmt.Println(year)
	week := strconv.Itoa(weekNum)
	fmt.Println(week)

	f := fmt.Sprintf("%02d", weekNum)
	r := fmt.Sprintf("%02d", 11)
	// fmt.Printf("%02d", 12)
	fmt.Println(f)
	fmt.Println(r)

}
