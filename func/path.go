package lem

func Edmonds(farm *AntFarm) *AntFarm {
	// create a edmonds karp algo to find all possible paths and save them in Antfarm.paths if the path is viable
	start := farm.StartRoom
	current := start
	current.Visited = true
	start.Visited = true
	path := path {}
	path.Rooms = append(path.Rooms,start)
	pathsfound := false

	for !pathsfound {
		for _, startlink := range start.Links {
			if startlink.Room.Visited == false {
				startlink.Room.Visited = true
				current = startlink.Room
				path.Rooms = append(path.Rooms, current)
				for !current.IsEnd {
					for _, currentlink := range current.Links {
						if currentlink.Room.Visited == false {
							currentlink.Room.Visited = true
							current = currentlink.Room
							path.Rooms = append(path.Rooms, current)
							break
						}
					}
				}
				farm.Paths = append(farm.Paths, path.Rooms)
				path.Rooms = nil
				current.Visited = false
				path.Rooms = append(path.Rooms, start)
			}
			pathsfound = true
		}
	}

	return farm
}