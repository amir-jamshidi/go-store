package entity

type Question struct {
	ID              int
	Text            string
	PossibleAnswer  []PossibleAnswer
	CorrectAnswerID int
	Difficulty      QuestionDifficultyType
	CategoryID      int
}

type PossibleAnswer struct {
	ID     int
	Text   string
	Choice PossibleAnswerChoiceType
}

type PossibleAnswerChoiceType uint8

func (p PossibleAnswerChoiceType) IsValid() bool {
	if p >= PossibleAnswerA && p <= PossibleAnswerD {
		return true
	}
	return false
}

const (
	PossibleAnswerA PossibleAnswerChoiceType = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

type QuestionDifficultyType uint8

const (
	QuestionDifficultyEasy QuestionDifficultyType = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (q QuestionDifficultyType) IsValid() bool {
	if q >= QuestionDifficultyEasy && 1 <= QuestionDifficultyHard {
		return true
	}
	return false
}
