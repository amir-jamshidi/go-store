package entity

import "time"

type Game struct {
	ID          int
	CategoryID  int
	QuestionIDs []int
	PlayerIDs   []int
}

type Player struct {
	ID        int
	UserID    int
	GameID    int
	Score     int
	Answers   []PlayerAnswer
	StartTime time.Time
}

type PlayerAnswer struct {
	ID         int
	PlayerID   int
	QuestionID int
	Choice     PossibleAnswerChoiceType
}
