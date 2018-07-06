package bowlingo

import (
	"fmt"
)

// Frame defines a single frame of two rolls from the user.
type Frame struct {
	FirstRoll  int `json:"first"`
	SecondRoll int `json:"second"`
	BonusRoll  int `json:"third"`
}

// Score takes a slice of frames and returns the current score
// or an error.
func Score(frames []Frame) (*int, error) {
	var score int
	framesLen := len(frames)
	if framesLen > 10 {
		return nil, fmt.Errorf("there should be maximum 10 frames")
	}
	rolls := framesToRolls(frames)
	var frameIndex int
	for i := 0; i < framesLen; i++ {
		if isStrike(frameIndex, rolls) {
			score += 10 + strikeBonus(frameIndex, rolls)
			frameIndex++
		} else if isSpare(frameIndex, rolls) {
			score += 10 + spareBonus(frameIndex, rolls)
			frameIndex += 2
		} else {
			score += sumOfBallsInFrame(frameIndex, rolls)
			frameIndex += 2
		}
	}
	return &score, nil
}

func isStrike(frameIndex int, rolls []int) bool {
	return rolls[frameIndex] == 10
}

func isSpare(frameIndex int, rolls []int) bool {
	return rolls[frameIndex]+rolls[frameIndex+1] == 10
}

func sumOfBallsInFrame(frameIndex int, rolls []int) int {
	return rolls[frameIndex] + rolls[frameIndex+1]
}

func spareBonus(frameIndex int, rolls []int) int {
	if frameIndex+2 >= len(rolls) {
		return 0
	}
	return rolls[frameIndex+2]
}

func strikeBonus(frameIndex int, rolls []int) int {
	if frameIndex+2 >= len(rolls) {
		return 0
	}
	return rolls[frameIndex+1] + rolls[frameIndex+2]
}

func framesToRolls(frames []Frame) []int {
	var rolls []int
	for _, f := range frames {
		rolls = append(rolls, f.FirstRoll)
		if f.FirstRoll == 10 && f.SecondRoll == 0 {
			continue
		}
		rolls = append(rolls, f.SecondRoll)
	}
	if len(frames) == 10 {
		rolls = append(rolls, frames[9].BonusRoll)
	}
	return rolls
}
