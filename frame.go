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
func CalculateScore(frames []Frame) (*int, error) {
	var score int
	framesCount := len(frames)
	if framesCount > 10 {
		return nil, fmt.Errorf("there should be maximum 10 frames")
	}
	for i := 1; i < framesCount; i++ {
		if IsStrike(frames[i-1]) {
			score += s
		}
	}
	return &score, nil
}

func IsStrike(frame Frame) bool {
	return frame.FirstRoll == 10
}

func IsSpare(frame Frame) bool {
	return frame.FirstRoll+frame.SecondRoll == 10
}
