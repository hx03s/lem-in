package lem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLine(filename string) (*AntFarm, []*Room, error, map[string]*Room) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %v", err),nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rooms []*Room
	var Farm = &AntFarm{}
	roomMap := make(map[string]*Room)
    var roomCounter int
	lineNum := 1

	// Read the number of ants
	if scanner.Scan() {
		Farm.NumAnts, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, nil, fmt.Errorf("line %d: invalid number of ants", lineNum),nil
		}
		lineNum++
	} else {
		return nil, nil, fmt.Errorf("empty file"),nil
	}

	// Read rooms and links
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 1 { // Special rooms (##start, ##end)
			name := fields[0]
            //if start or end already exits and name is equal ##start or ##end error more than start point or end point
			if name == "##start" || name == "##end" {
                if _, exists := roomMap[name]; exists {
                    return nil, nil, fmt.Errorf("line %d: %s already exists", lineNum, name),nil
                }
				room := &Room{
					Name: name,
				}
                if roomCounter > 0 {
                    rooms[len(rooms)-1].Next = room
                }
				rooms = append(rooms, room)
                roomCounter++
				roomMap[name] = room
			} else if strings.Contains(name, "-") {
				room1, room2, err := ParseLink(fields[0], roomMap)
				if err != nil {
					return nil, nil, fmt.Errorf("line %d: %v", lineNum, err),nil
				}

				room1.Links = append(room1.Links, &Link{Room: room2})
				// room2.Links = append(room2.Links, &Link{Room: room1})
			} else if strings.HasPrefix(name,"#") {
				continue
			} else {
				return nil, nil, fmt.Errorf("line %d: invalid format", lineNum),nil
			}
		} else if len(fields) == 3 { // Regular room
			name, x, y, err := ParseRoom(fields)
			if err != nil {
				return nil, nil, fmt.Errorf("line %d: %v", lineNum, err),nil
			}

			if _, exists := roomMap[name]; exists {
				return nil, nil, fmt.Errorf("line %d: duplicate room '%s'", lineNum, name),nil
			}

			room := &Room{
				Name:   name,
				CoordX: x,
				CoordY: y,
			}
            if roomCounter > 0 {
                rooms[len(rooms)-1].Next = room
            }
			rooms = append(rooms, room)
            roomCounter++
			roomMap[name] = room
		} else {
			return nil, nil, fmt.Errorf("line %d: invalid format", lineNum),nil
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err),nil
	}

	if _, exists := roomMap["##start"]; !exists {
		return nil, nil, fmt.Errorf("start room doesn't exists!"),nil
	}

	if _, exists := roomMap["##end"]; !exists {
		return nil, nil, fmt.Errorf("end room doesn't exists!"),nil
	}
    if rooms[len(rooms)-1].Name == "##start" || rooms[len(rooms)-1].Name == "##end" {
        return nil,nil,fmt.Errorf("No room assigned to %s", rooms[len(rooms)-1].Name),nil
    }
    if roomMap["##start"].Next != nil {
        roomMap["##start"].Next.IsStart = true
		Farm.StartRoom = roomMap["##start"].Next
    }
    
    if roomMap["##end"].Next != nil {
         roomMap["##end"].Next.IsEnd = true
		 Farm.EndRoom = roomMap["##end"].Next
    }
   
	return Farm, rooms, nil,roomMap
}

func ParseRoom(fields []string) (string, int, int, error) {
    name := fields[0]
    if strings.HasPrefix(name, "L") || strings.Contains(name, "-") || strings.Contains(name, "#") {
        return "", 0, 0, fmt.Errorf("invalid room name '%s'", name)
    }

    x, err := strconv.Atoi(fields[1])
    if err != nil {
        return "", 0, 0, fmt.Errorf("invalid x coordinate '%s'", fields[1])
    }

    y, err := strconv.Atoi(fields[2])
    if err != nil {
        return "", 0, 0, fmt.Errorf("invalid y coordinate '%s'", fields[2])
    }

    return name, x, y, nil
}

func ParseLink(line string, roomMap map[string]*Room) (*Room, *Room, error) {
    fields := strings.Split(line, "-")
    if len(fields) != 2 {
        return nil, nil, fmt.Errorf("invalid tunnel format: %s", line)
    }

    room1Name := strings.TrimSpace(fields[0])
    room2Name := strings.TrimSpace(fields[1])

    room1, exists := roomMap[room1Name]
    if !exists {
        return nil, nil, fmt.Errorf("unknown room '%s'", room1Name)
    }

    room2, exists := roomMap[room2Name]
    if !exists {
        return nil, nil, fmt.Errorf("unknown room '%s'", room2Name)
    }

    if room1 == room2 {
        return nil, nil, fmt.Errorf("room '%s' linked to itself", room1Name)
    }

    return room1, room2, nil
}
