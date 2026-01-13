package entity

type Game struct {
	ID          int
	Category    string
	QuestionIDs []int
	PlayerIDs   []int
}
