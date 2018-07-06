package bowlingo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFramesToRolls(t *testing.T) {
	var frames []Frame
	frames = append(frames, newFrame(10, 0), newFrame(7, 3),
		newFrame(8, 2), newFrame(10, 0))
	rolls := framesToRolls(frames)
	expectedRolls := []int{10, 7, 3, 8, 2, 10}
	assert.Equal(t, len(expectedRolls), len(rolls),
		"the length of actual and expected rolls should be equal")
	assert.EqualValues(t, expectedRolls, rolls,
		"the actual and expected rolls should have the same values")
}

func TestValidateFrame(t *testing.T) {
	t.Run("valid frame", func(t *testing.T) {
		f := Frame{
			FirstRoll:  5,
			SecondRoll: 4,
		}
		assert.Nil(t, validateFrame(0, f), "validateFrame should not return an error")
	})
	t.Run("invalid first roll", func(t *testing.T) {
		f := Frame{
			FirstRoll:  -1,
			SecondRoll: 6,
		}
		assert.NotNil(t, validateFrame(0, f), "validateFrame should return an error")
	})
	t.Run("invalid second roll", func(t *testing.T) {
		f := Frame{
			FirstRoll:  8,
			SecondRoll: 11,
		}
		assert.NotNil(t, validateFrame(0, f), "validateFrame should return an error")
	})
	t.Run("invalid sum of rolls", func(t *testing.T) {
		f := Frame{
			FirstRoll:  6,
			SecondRoll: 5,
		}
		assert.NotNil(t, validateFrame(0, f), "validateFrame should return an error")
	})
	t.Run("bonus roll for frame with index < 9", func(t *testing.T) {
		bonus := 5
		f := Frame{
			FirstRoll:  5,
			SecondRoll: 6,
			BonusRoll:  &bonus,
		}
		assert.NotNil(t, validateFrame(3, f), "validateFrame should return an error")
	})
	t.Run("bonus roll awarded incorrectly", func(t *testing.T) {
		bonus := 2
		f := Frame{
			FirstRoll:  3,
			SecondRoll: 3,
			BonusRoll:  &bonus,
		}
		assert.NotNil(t, validateFrame(9, f), "validateFrame should return an error")
	})
}

func TestIsStrike(t *testing.T) {
	rolls := []int{10, 0, 5, 6}
	t.Run("strike", func(t *testing.T) {
		assert.True(t, isStrike(0, rolls), "current frame does contain a strike")
	})
	t.Run("no strike", func(t *testing.T) {
		assert.False(t, isStrike(2, rolls), "current frame does not contain a strike")
	})
}

func TestIsSpare(t *testing.T) {
	rolls := []int{5, 5, 5, 3}
	t.Run("spare", func(t *testing.T) {
		assert.True(t, isSpare(0, rolls), "current frame does contain a spare")
	})
	t.Run("no spare", func(t *testing.T) {
		assert.False(t, isSpare(2, rolls), "current frame does not contain a spare")
	})
}

func TestFrameScore(t *testing.T) {
	rolls := []int{3, 4, 5, 2}
	assert.Equal(t, 7, frameScore(0, rolls), "first frame score should be equal to 7")
}

func TestSpareBonus(t *testing.T) {
	t.Run("spare bonus exists", func(t *testing.T) {
		rolls := []int{5, 5, 3, 2}
		assert.Equal(t, 3, *spareBonus(0, rolls), "spare bonus should be equal to first roll in the following frame")
	})
	t.Run("spare bonus does not exist", func(t *testing.T) {
		rolls := []int{5, 5}
		assert.Nil(t, spareBonus(0, rolls),
			"spare bonus should be nil as we cannot calculate it due to no extra frame being available")
	})
}

func TestStrikeBonus(t *testing.T) {
	t.Run("next 2 rolls are in the following frame", func(t *testing.T) {
		rolls := []int{10, 3, 4}
		assert.Equal(t, 7, *strikeBonus(0, rolls), "bonus should be calculated with both rolls from following frame")
	})
	t.Run("next roll strike, second one in the third frame", func(t *testing.T) {
		rolls := []int{10, 10, 2, 4}
		assert.Equal(t, 12, *strikeBonus(0, rolls), "bonus should be calculated with a strike and the first roll from third frame")
	})
	t.Run("next 2 rolls are strikes", func(t *testing.T) {
		rolls := []int{10, 10, 10}
		assert.Equal(t, 20, *strikeBonus(0, rolls), "bonus should be calculated with both strikes as they are the next 2 rolls")
	})
	t.Run("cannot calculate strike", func(t *testing.T) {
		rolls := []int{10}
		assert.Nil(t, strikeBonus(0, rolls), "cannot calculate bonus as there is no extra frame")
	})
}

func TestScoreOnePinOnEveryFrame(t *testing.T) {
	frames := newFrameSlice(2, 10)
	score, _ := Score(frames)
	assert.Equal(t, 10, score, "final score should be 10")
}

func TestScorePerfectGame(t *testing.T) {
	frames := newFrameSlice(10, 9)
	bonusRoll := 10
	frames = append(frames, Frame{
		FirstRoll:  10,
		SecondRoll: 10,
		BonusRoll:  &bonusRoll,
	})
	score, _ := Score(frames)
	assert.Equal(t, 300, score, "final score should be 300")
}
