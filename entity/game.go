package entity

type Game struct {
	ID       int
	Category string
	Question []Question
	Players  []User
}
