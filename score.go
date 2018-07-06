package main

import (
	"fmt"
)

// Frame defines a single frame of two rolls from the user.
type Frame struct {
	FirstRoll  int  `json:"firstRoll"`
	SecondRoll int  `json:"secondRoll"`
	BonusRoll  *int `json:"bonusRoll"`
}

// Score takes a slice of frames and returns the current score
// or an error.
func Score(frames []Frame) (int, error) {
	var score int
	framesLen := len(frames)
	if framesLen > 10 {
		return score, fmt.Errorf("there should be maximum 10 frames")
	}
	rolls := framesToRolls(frames)
	var frameIndex int
	for i := 0; i < framesLen; i++ {
		if err := validateFrame(i, frames[i]); err != nil {
			return score, err
		}
		if isStrike(frameIndex, rolls) {
			bonus := strikeBonus(frameIndex, rolls)
			// if we do not have 1 or 2 more frames to calculate
			// the score, we return the current score as we cannot
			// compute the final one until the other frame(s) are rolled as well
			if bonus == nil {
				return score, nil
			}
			score += 10 + *bonus
			frameIndex++
		} else if isSpare(frameIndex, rolls) {
			bonus := spareBonus(frameIndex, rolls)
			// same applies to spare: if we do not have any extra frame to add the first roll
			// to 10, we cannot compute the final score, so we return the current score
			if bonus == nil {
				return score, nil
			}
			score += 10 + *bonus
			frameIndex += 2
		} else {
			score += frameScore(frameIndex, rolls)
			frameIndex += 2
		}
	}
	return score, nil
}

// isStrike checks whether the current frame contains a strike.
func isStrike(frameIndex int, rolls []int) bool {
	return rolls[frameIndex] == 10
}

// isSpare checks whether the current frame contains a spare.
func isSpare(frameIndex int, rolls []int) bool {
	return rolls[frameIndex]+rolls[frameIndex+1] == 10
}

// frameScore returns the sum of the two rolls in the current frame.
func frameScore(frameIndex int, rolls []int) int {
	return rolls[frameIndex] + rolls[frameIndex+1]
}

// spareBonus calculates the score to be added to 10 when the
// player hits a spare in the current frame.
func spareBonus(frameIndex int, rolls []int) *int {
	if frameIndex+2 >= len(rolls) {
		return nil
	}
	return &rolls[frameIndex+2]
}

// strikeBonus calculates the score to be added to 10 when the
// player hits a strike in the current frame.
func strikeBonus(frameIndex int, rolls []int) *int {
	if frameIndex+2 >= len(rolls) {
		return nil
	}
	bonus := rolls[frameIndex+1] + rolls[frameIndex+2]
	return &bonus
}

// framesToRolls converts a slice of frames to a slice of rolls. This is needed
// because, if we only used frames, there would have been code smells, such as
// nested if/else if/else blocks. This way, the code is much cleaner.
func framesToRolls(frames []Frame) []int {
	var rolls []int
	for _, f := range frames {
		rolls = append(rolls, f.FirstRoll)
		if f.FirstRoll == 10 && f.SecondRoll == 0 {
			continue
		}
		rolls = append(rolls, f.SecondRoll)
	}
	if len(frames) == 10 && frames[9].BonusRoll != nil {
		rolls = append(rolls, *frames[9].BonusRoll)
	}
	return rolls
}

// validateFrame takes a frame and its index and performs different validations
// on the frame and its rolls.
func validateFrame(frameIndex int, frame Frame) error {
	if frame.FirstRoll < 0 || frame.FirstRoll > 10 {
		return fmt.Errorf("invalid value for first roll of the frame")
	}
	if frame.SecondRoll < 0 || frame.SecondRoll > 10 {
		return fmt.Errorf("invalid value for second roll of the frame")
	}
	if frame.FirstRoll+frame.SecondRoll > 10 {
		return fmt.Errorf("sum of rolls must be less than 10")
	}
	if frameIndex < 9 && frame.BonusRoll != nil {
		return fmt.Errorf("only last frame can have a bonus roll")
	}
	if frameIndex == 9 && frame.BonusRoll != nil && frame.FirstRoll != 10 &&
		frame.FirstRoll+frame.SecondRoll != 10 {
		return fmt.Errorf("bonus roll should not be awarded as the first 2 rolls are neither a spare nor a strike")
	}
	return nil
}
