package bowlingo

import (
	"fmt"
)

// Frame defines a single frame of two rolls from the user.
type Frame struct {
	FirstRoll  int `json:"first"`
	SecondRoll int `json:"second"`
}

// CalculateScore takes a slice of frames and returns the current score
// or an error.
func CalculateScore(frames []Frame) (int, error) {
	for i := range frames {
		fmt.Println(i)
	}
	return 0, nil
}
