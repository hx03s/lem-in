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
	} else {
		fmt.Println("Reading file...")
		lem.ReadLine(os.Args[1])
	}
}
