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
	for i := 0; i < framesCount-1; i++ {
		if isStrike(frames[i]) {
			score += 10 + frameScore(frames[i+1])
		} else if isSpare(frames[i]) {
			score += 10 + frames[i+1].FirstRoll
		} else {
			score += frameScore(frames[i])
		}
	}
	return &score, nil
}

func isStrike(frame Frame) bool {
	return frame.FirstRoll == 10
}

func isSpare(frame Frame) bool {
	return frame.FirstRoll+frame.SecondRoll == 10
}

func frameScore(frame Frame) int {
	return frame.FirstRoll + frame.SecondRoll
}
