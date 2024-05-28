package lem

type Room struct {
	Id      int
	Name    string
	CoordX  int
	CoordY  int
	IsStart bool
	IsEnd   bool
	Visited bool
	Links   []*Link
	Next *Room
}

type Link struct {
	Room *Room
	Next *Link
}

type AntFarm struct {
	NumAnts   int
	StartRoom *Room
	EndRoom   *Room
	Rooms     []*Room
	Paths []*Path
	Ants []*Ant
}

type Path struct {
	Rooms []*Room
	Queue []*Ant
	NumAnts int
}

type Ant struct {
	Id          int
	Path *Path
	CurrentRoom *Room
	NextRoom   *Room
}
