package helpers

import (
	"fmt"
	"log"
	"math"
	"time"
)

// PrintError func to print errors
func PrintError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Timer waiting 2s function
func Timer() {
	time.Sleep(100 * time.Millisecond)
}

// Similarity calculate similariti finded anime, Api: "similarity" - value form 0.00 to 1.00
func Similarity(s string) {
	fmt.Print(s)
}

// FromTo calculate the time in which the selected scene appeared, APi: "form", "to" - value in seconds
// x is value from APi
func FromTo(x float64) {
	// calculate hours - 1h is 3600s
	h := math.Floor(x / 3600)
	m := math.Floor((x - h*3600) / 60)
	s := x - (x*3600 + m*60)

	_, err := fmt.Println(h, m, s)
	PrintError(err)
}
