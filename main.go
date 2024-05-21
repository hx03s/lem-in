package main

import (
    "fmt"
    lem "lem/func"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a file path as an argument.")
        return
    } else if len(os.Args) > 2 {
		fmt.Println("Please provide only one argument.")
		return
	}
    
	fmt.Println("Reading file...")
    filename := os.Args[1]
    numAnts, rooms, err := lem.ReadLine(filename)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Number of ants: %d\n", numAnts.NumAnts)
    fmt.Println("Rooms:")
    for _, room := range rooms {
        fmt.Printf("  %s (%d, %d) - Start: %t, End: %t\n", room.Name, room.CoordX, room.CoordY, room.IsStart, room.IsEnd)
        fmt.Print("    Links: ")
        for i, link := range room.Links {
            if i > 0 {
                fmt.Print(", ")
            }
            fmt.Print(link.Room.Name)
        }
        fmt.Println()
    }
}

