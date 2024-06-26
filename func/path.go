package lem

import "math"

func Edmonds(farm *AntFarm) []*Path {
    direct := false
    start := []*Room{farm.StartRoom}
    end := farm.EndRoom
    visited := make(map[*Room]bool)
    queue := []*Path{{Rooms: start}}
    var paths []*Path
    visited[farm.StartRoom] = true

    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]
        currentRoom := path.Rooms[len(path.Rooms)-1]

        if currentRoom == end {
            newPath := &Path{Rooms: make([]*Room, len(path.Rooms))}
            copy(newPath.Rooms, path.Rooms)
            paths = append(paths, newPath)
            if len(newPath.Rooms) == 2 {
                direct = true
            }

            // Reset visited status of rooms not part of finalized path
            for room := range visited {
                if room == farm.StartRoom {
                    continue
                }
                visited[room] = false
            }
            for _, currentpath := range paths {
                for _, currentroom := range currentpath.Rooms {
                    if currentroom == farm.EndRoom {
                        continue
                    }
                visited[currentroom] = true
                } 
            }
            queue = queue[:0]
            queue = []*Path{{Rooms: start}}
            continue
        }
        // if visited[currentRoom] {
        //     continue
        // }

        visited[currentRoom] = true

        for _, link := range currentRoom.Links {
            if direct == true && currentRoom == farm.StartRoom && link.Room == farm.EndRoom {
                    continue
            }
            nextRoom := link.Room
            if !visited[nextRoom] {
                // visited[nextRoom] = true
                newPath := &Path{Rooms: make([]*Room, len(path.Rooms), len(path.Rooms)+1)}
                copy(newPath.Rooms, path.Rooms)
                newPath.Rooms = append(newPath.Rooms, nextRoom)
                queue = append(queue, newPath)
            }
        }
    }

    // paths = chooseOptimalPaths(paths, farm.StartRoom)

    return paths
}

func chooseOptimalPaths(paths []*Path, startRoom *Room) []*Path {
	// Group paths by their second room
	groups := make(map[*Room][]*Path)
	for _, path := range paths {
		secondRoom := path.Rooms[1]
		groups[secondRoom] = append(groups[secondRoom], path)
	}

	// Find the optimal path for each group
	var optimalPaths []*Path
	for _, pathsInGroup := range groups {
		minSharedRooms := math.MaxFloat64
		minPathLength := math.MaxFloat64
		var optimalPath *Path

		for _, path := range pathsInGroup {
			sharedRooms := 0.0
			for _, otherPaths := range groups {
				if &otherPaths != &pathsInGroup {
					for _, otherPath := range otherPaths {
						if hasSharedRooms(path, otherPath) {
							sharedRooms++
							break
						}
					}
				}
			}

			pathLength := float64(len(path.Rooms))

			if sharedRooms < minSharedRooms || (sharedRooms == minSharedRooms && pathLength < minPathLength) {
				minSharedRooms = sharedRooms
				minPathLength = pathLength
				optimalPath = path
			}
		}

		optimalPaths = append(optimalPaths, optimalPath)
	}

	// now i want to filter optimalPaths to choose one path that has the shortest length between paths that has shared rooms with other paths in other groups
	var filteredPaths []*Path
	for _, path := range optimalPaths {
		hasSharedRoomsWithOthers := false
		for _, otherPath := range optimalPaths {
			if path != otherPath && hasSharedRooms(path, otherPath) {
				hasSharedRoomsWithOthers = true
				break
			}
		}

		if !hasSharedRoomsWithOthers {
			filteredPaths = append(filteredPaths, path)
		}
	}

	if len(filteredPaths) < len(optimalPaths) {
		minPathLength := math.Inf(1)
		var selectedPath *Path

		for _, path := range optimalPaths {
			if !contains(filteredPaths, path) {
				pathLength := float64(len(path.Rooms))
				if pathLength < minPathLength {
					minPathLength = pathLength
					selectedPath = path
				}
			}
		}

		filteredPaths = append(filteredPaths, selectedPath)
	}

	return filteredPaths
}

func hasSharedRooms(path1, path2 *Path) bool {
    rooms1 := make(map[*Room]bool)
    for _, room := range path1.Rooms {
        if room != path1.Rooms[0] && room != path1.Rooms[len(path1.Rooms)-1] {
            rooms1[room] = true
        }
    }

    for _, room := range path2.Rooms {
        if room != path2.Rooms[0] && room != path2.Rooms[len(path2.Rooms)-1] {
            if rooms1[room] {
                return true
            }
        }
    }

    return false
}

func contains(paths []*Path, path *Path) bool {
    for _, p := range paths {
        if p == path {
            return true
        }
    }
    return false
}
