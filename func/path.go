package lem

func Edmonds(farm *AntFarm) []*Path {
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

        // visited[currentRoom] = true

        for _, link := range currentRoom.Links {
            nextRoom := link.Room
            if !visited[nextRoom] {
                visited[nextRoom] = true
                newPath := &Path{Rooms: make([]*Room, len(path.Rooms), len(path.Rooms)+1)}
                copy(newPath.Rooms, path.Rooms)
                newPath.Rooms = append(newPath.Rooms, nextRoom)
                queue = append(queue, newPath)
            }
        }
    }

    return paths
}
