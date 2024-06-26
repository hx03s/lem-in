package lem

import "fmt"

func MoveAnt(farm *AntFarm) {
	moved := make(map[*Ant]bool)
	full := make(map[*Room]bool)
	direct := false
	for {
		if AllEnd(farm.Paths) {
			break
		}
		for _, CurrentAnt := range farm.Ants {
				if moved[CurrentAnt] {
					continue
				}
				full[CurrentAnt.CurrentRoom] = false
				if GetRoomIndex(CurrentAnt.CurrentRoom, CurrentAnt.Path) == len(CurrentAnt.Path.Rooms)-1 {
                    // Ant has reached the end room, no need to move
                    moved[CurrentAnt] = true
                    continue
                }
				if len(CurrentAnt.Path.Rooms) == 2 && CurrentAnt.Path.Rooms[0] == farm.StartRoom && direct == false{
					direct = true
					CurrentAnt.CurrentRoom = CurrentAnt.Path.Rooms[GetRoomIndex(CurrentAnt.CurrentRoom, CurrentAnt.Path)+1]
					fmt.Printf("L%v-%v ", CurrentAnt.Id, CurrentAnt.CurrentRoom.Name)
					continue
				} else if len(CurrentAnt.Path.Rooms) == 2 && CurrentAnt.Path.Rooms[0] == farm.StartRoom && direct == true{
					continue
				}
				if full[CurrentAnt.Path.Rooms[GetRoomIndex(CurrentAnt.CurrentRoom, CurrentAnt.Path)+1]] == true && CurrentAnt.Path.Rooms[GetRoomIndex(CurrentAnt.CurrentRoom, CurrentAnt.Path)+1] != farm.EndRoom {
					continue
				}
				CurrentAnt.CurrentRoom = CurrentAnt.Path.Rooms[GetRoomIndex(CurrentAnt.CurrentRoom, CurrentAnt.Path)+1]
				fmt.Printf("L%v-%v ", CurrentAnt.Id, CurrentAnt.CurrentRoom.Name)
				full[CurrentAnt.CurrentRoom] = true
				moved[CurrentAnt] = true
		}
		fmt.Println()
		Reset(farm, moved)
		direct = false
	}
}

func AssignAnts(farm *AntFarm) *AntFarm {
    i := 0
    currentPathIndex := 0

    for {
        if i == farm.NumAnts {
            break
        }

        if currentPathIndex >= len(farm.Paths) {
            currentPathIndex = 0
        }

        lowestCostPath := farm.Paths[currentPathIndex]
        lowestCost := PathCost(lowestCostPath)
        isAllEqual := true

        for j, path := range farm.Paths {
            if PathCost(path) < lowestCost {
                lowestCostPath = path
                lowestCost = PathCost(path)
                currentPathIndex = j
                isAllEqual = false
            } else if PathCost(path) != lowestCost {
                isAllEqual = false
            }
        }

        if isAllEqual && currentPathIndex == len(farm.Paths)-1 {
            currentPathIndex = 0
        }
currentant :=  &Ant{Id: i+1, Path: lowestCostPath,CurrentRoom: lowestCostPath.Rooms[0]}
farm.Ants = append(farm.Ants, currentant)
        lowestCostPath.Queue = append(lowestCostPath.Queue,currentant)
        lowestCostPath.NumAnts++
        i++
        currentPathIndex++
    }

    return farm
}


func PathCost(Current *Path) int {
    return len(Current.Rooms) + Current.NumAnts
}

// func Lowest(paths []*Path, path *Path) int {
//     lowestCost := PathCost(path)
//     lowestIndex := 0
//     allEqual := true

//     for i, current := range paths {
//         if PathCost(current) < lowestCost {
//             lowestCost = PathCost(current)
//             lowestIndex = i
//             allEqual = false
//         } else if PathCost(current) != lowestCost {
//             allEqual = false
//         }
//     }

//     if allEqual {
//         return 0
//     }

//     return lowestIndex
// }


func GetRoomIndex(current *Room,path *Path) int {
for i, room := range path.Rooms {
	if room == current {
		return i
	}
}
return 0
}

func Reset(farm *AntFarm, moved map[*Ant]bool) {
		 for _, path := range farm.Paths {
		for _, ant := range path.Queue {
			 moved[ant] = false
		 }
	 }
}

func AllEnd(paths []*Path) bool {
	allend := true
for _, path := range paths {
	if !allend {
		break
	}
	for _, ant := range path.Queue {
if ant.CurrentRoom != path.Rooms[len(path.Rooms)-1] {
	allend = false
	break
}
	}
}
return allend
}