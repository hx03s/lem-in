package lem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLine(filename string) (int, []*Room, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, nil, fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var numAnts int
    var rooms []*Room
    roomMap := make(map[string]*Room)
    lineNum := 1

    // Read the number of ants
    if scanner.Scan() {
        numAnts, err = strconv.Atoi(scanner.Text())
        if err != nil {
            return 0, nil, fmt.Errorf("line %d: invalid number of ants", lineNum)
        }
        lineNum++
    } else {
        return 0, nil, fmt.Errorf("empty file")
    }

    // Read rooms and links
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)

        if len(fields) == 1 { // Special rooms (##start, ##end)
            name := fields[0]
            if name == "##start" || name == "##end" {
                room := &Room{
                    Name: name,
                }
                rooms = append(rooms, room)
                roomMap[name] = room
            } else if strings.Contains(name,"-") {
				room1, room2, err := parseLink(fields[0], roomMap)
				if err != nil {
					return 0, nil, fmt.Errorf("line %d: %v", lineNum, err)
				}
	
				room1.Links = append(room1.Links, &Link{Room: room2})
				room2.Links = append(room2.Links, &Link{Room: room1})
			} else {
                return 0, nil, fmt.Errorf("line %d: invalid format", lineNum)
            }
        } else if len(fields) == 3 { // Regular room
            name, x, y, err := parseRoom(fields)
            if err != nil {
                return 0, nil, fmt.Errorf("line %d: %v", lineNum, err)
            }

            if _, exists := roomMap[name]; exists {
                return 0, nil, fmt.Errorf("line %d: duplicate room '%s'", lineNum, name)
            }

            room := &Room{
                Name:   name,
                CoordX: x,
                CoordY: y,
            }
            rooms = append(rooms, room)
            roomMap[name] = room
        } else {
            return 0, nil, fmt.Errorf("line %d: invalid format", lineNum)
        }

        lineNum++
    }

    if err := scanner.Err(); err != nil {
        return 0, nil, fmt.Errorf("error reading file: %v", err)
    }

    return numAnts, rooms, nil
}

func parseRoom(fields []string) (string, int, int, error) {
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

func parseLink(line string, roomMap map[string]*Room) (*Room, *Room, error) {
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
