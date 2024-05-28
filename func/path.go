package lem

func Edmonds(farm *AntFarm) []*Path {
    start := []*Room{farm.StartRoom}
    end := farm.EndRoom
    visited := make(map[*Room]bool)
    queue := []*Path{{Rooms: start}}
    var paths []*Path

    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]
        currentRoom := path.Rooms[len(path.Rooms)-1]

        if currentRoom == end {
            newPath := &Path{Rooms: make([]*Room, len(path.Rooms))}
            copy(newPath.Rooms, path.Rooms)
            paths = append(paths, newPath)
            continue
        }
		if visited[currentRoom] {
			continue
		}

        visited[currentRoom] = true

        for _, link := range currentRoom.Links {
            nextRoom := link.Room
            if !visited[nextRoom] {
                newPath := &Path{Rooms: make([]*Room, len(path.Rooms), len(path.Rooms)+1)}
                copy(newPath.Rooms, path.Rooms)
                newPath.Rooms = append(newPath.Rooms, nextRoom)
                queue = append(queue, newPath)
            }
        }
    }

    return paths
}