package bowlingo

import (
	"fmt"
)

const (
	maxThrowsPerGame  = 21
	totalNumberOfPins = 10
	maxNumberOfFrames = 10
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
	if framesLen > maxNumberOfFrames {
		return score, fmt.Errorf("there should be maximum 10 frames")
	}
	rolls := framesToRolls(frames)
	for frameIndex, rollIndex := 0, 0; frameIndex < framesLen; frameIndex++ {
		if err := validateFrame(frameIndex, frames[frameIndex]); err != nil {
			return score, fmt.Errorf("frame %d is invalid: %v", frameIndex, err)
		}
		if isStrike(rollIndex, rolls) {
			score += totalNumberOfPins + strikeBonus(rollIndex, rolls)
			rollIndex++
		} else if isSpare(rollIndex, rolls) {
			score += totalNumberOfPins + spareBonus(rollIndex, rolls)
			rollIndex += 2
		} else {
			score += frameScore(rollIndex, rolls)
			rollIndex += 2
		}
	}
	return score, nil
}

// isStrike checks whether the current frame contains a strike.
func isStrike(frameIndex int, rolls []int) bool {
	return rolls[frameIndex] == totalNumberOfPins
}

// isSpare checks whether the current frame contains a spare.
func isSpare(frameIndex int, rolls []int) bool {
	return rolls[frameIndex]+rolls[frameIndex+1] == totalNumberOfPins
}

// frameScore returns the sum of the two rolls in the current frame.
func frameScore(frameIndex int, rolls []int) int {
	return rolls[frameIndex] + rolls[frameIndex+1]
}

// spareBonus calculates the score to be added to 10 when the
// player hits a spare in the current frame.
func spareBonus(frameIndex int, rolls []int) int {
	return rolls[frameIndex+2]
}

// strikeBonus calculates the score to be added to 10 when the
// player hits a strike in the current frame.
func strikeBonus(frameIndex int, rolls []int) int {
	return rolls[frameIndex+1] + rolls[frameIndex+2]
}

// framesToRolls converts a slice of frames to a slice of rolls. This is needed
// because, if we only used frames, there would have been code smells, such as
// nested if/else if/else blocks. This way, the code is much cleaner.
func framesToRolls(frames []Frame) []int {
	rolls := make([]int, maxThrowsPerGame)
	var rollIndex int
	for _, f := range frames {
		rolls[rollIndex] = f.FirstRoll
		if f.FirstRoll == totalNumberOfPins && f.SecondRoll == 0 {
			rollIndex++
			continue
		}
		rolls[rollIndex+1] = f.SecondRoll
		rollIndex += 2
	}
	if len(frames) == maxNumberOfFrames && frames[9].BonusRoll != nil {
		rolls[rollIndex] = *frames[9].BonusRoll
	}
	return rolls
}

// validateFrame takes a frame and its index and performs different validations
// on the frame and its rolls.
func validateFrame(frameIndex int, frame Frame) error {
	if frame.FirstRoll < 0 || frame.FirstRoll > totalNumberOfPins {
		return fmt.Errorf("invalid value for first roll of the frame")
	}
	if frame.SecondRoll < 0 || frame.SecondRoll > totalNumberOfPins {
		return fmt.Errorf("invalid value for second roll of the frame")
	}
	if frame.FirstRoll+frame.SecondRoll > totalNumberOfPins && frameIndex < 9 {
		return fmt.Errorf("sum of rolls must be less than 10")
	}
	if frameIndex < 9 && frame.BonusRoll != nil {
		return fmt.Errorf("only last frame can have a bonus roll")
	}
	if frameIndex == 9 && frame.BonusRoll != nil && frame.FirstRoll != 10 &&
		frame.FirstRoll+frame.SecondRoll != totalNumberOfPins {
		return fmt.Errorf("bonus roll should not be awarded as the first 2 rolls are neither a spare nor a strike")
	}
	return nil
}

// newFrame is used for testing purposes in order to create a frame with
// specified first and second roll.
func newFrame(firstRoll, secondRoll int) Frame {
	return Frame{
		FirstRoll:  firstRoll,
		SecondRoll: secondRoll,
	}
}

// newFrameSlice is used for testing purposes in order to create a slice of frames
// with the first roll equal to the totalPinsPerFrame defined by the test creator.
func newFrameSlice(totalPinsPerFrame, noFrames int) []Frame {
	frames := make([]Frame, noFrames)
	for i := 0; i < noFrames; i++ {
		firstRoll := totalPinsPerFrame
		secondRoll := 0
		frames[i] = Frame{
			FirstRoll:  firstRoll,
			SecondRoll: secondRoll,
		}
	}
	return frames
}
