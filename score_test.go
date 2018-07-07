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
	expectedRolls := make([]int, 21)
	copy(expectedRolls[:6], []int{10, 7, 3, 8, 2, 10})
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
		assert.Equal(t, 3, spareBonus(0, rolls), "spare bonus should be equal to first roll in the following frame")
	})
	t.Run("spare bonus does not exist", func(t *testing.T) {
		rolls := make([]int, 21)
		rolls[0], rolls[1] = 5, 5
		assert.Zero(t, spareBonus(0, rolls),
			"spare bonus should be 0 as we cannot calculate it due to no extra frame being available")
	})
}

func TestStrikeBonus(t *testing.T) {
	t.Run("next 2 rolls are in the following frame", func(t *testing.T) {
		rolls := []int{10, 3, 4}
		assert.Equal(t, 7, strikeBonus(0, rolls), "bonus should be calculated with both rolls from following frame")
	})
	t.Run("next roll strike, second one in the third frame", func(t *testing.T) {
		rolls := []int{10, 10, 2, 4}
		assert.Equal(t, 12, strikeBonus(0, rolls), "bonus should be calculated with a strike and the first roll from third frame")
	})
	t.Run("next 2 rolls are strikes", func(t *testing.T) {
		rolls := []int{10, 10, 10}
		assert.Equal(t, 20, strikeBonus(0, rolls), "bonus should be calculated with both strikes as they are the next 2 rolls")
	})
	t.Run("cannot calculate strike", func(t *testing.T) {
		rolls := make([]int, maxThrowsPerGame)
		rolls[0] = 10
		assert.Zero(t, strikeBonus(0, rolls), "cannot calculate bonus as there is no extra frame")
	})
}

func TestScoreOnePinOnEveryFrame(t *testing.T) {
	frames := newFrameSlice(1, 10)
	score, _ := Score(frames)
	assert.Equal(t, 10, score, "final score should be 10")
}

func TestScorePerfectGame(t *testing.T) {
	frames := newFrameSlice(10, 10)
	bonus := 10
	frames[9].BonusRoll = &bonus
	frames[9].SecondRoll = 10
	score, _ := Score(frames)
	assert.Equal(t, 300, score, "final score should be 300")
}

func TestScoreOneSpare(t *testing.T) {
	frames := newFrameSlice(5, 1)
	frames = append(frames, Frame{
		FirstRoll:  5,
		SecondRoll: 5,
	}, Frame{
		FirstRoll:  2,
		SecondRoll: 0,
	})
	score, _ := Score(frames)
	assert.Equal(t, 19, score, "final score with one spare should be 19")
}

func TestScoreAllGutter(t *testing.T) {
	frames := newFrameSlice(0, 10)
	score, _ := Score(frames)
	assert.Zero(t, score, "final score should have been 0")
}

func BenchmarkScore(b *testing.B) {
	frames := newFrameSlice(5, 10)
	for n := 0; n < b.N; n++ {
		Score(frames)
	}
}
