package entity

type Game struct {
	ID          int
	CategoryID  int
	QuestionIDs []int
	PlayerIDs   []Player
}

type Player struct {
	ID      int
	UserID  int
	GameID  int
	Score   int
	Answers []PlayerAnswer
}

type PlayerAnswer struct {
	ID         int
	PlayerID   int
	QuestionID int
	Choice     PossibleAnswerChoiceType
}
