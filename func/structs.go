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
}

// type Node struct {
// 	Name      string
// 	Neighbors []*Node
// }

// type Ant struct {
// 	Id          int
// 	Farm        *AntFarm
// 	CurrentNode *Node
// 	NextNode    *Node
// 	Visited     bool
// }
