package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"example.org/to-do-app/commandutils"
	"example.org/to-do-app/fileutils"
	"example.org/to-do-app/ioutils"
)

func main() {
	ioutils.GreetUser()

	if len(os.Args) < 3 {
		fmt.Println("Usage: todo <add|delete> <text>")
	}

	processAction(os.Args)

	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{
		"1", "My task", "today",
	})

	file, err := os.Open("example.csv")

	if err != nil {
		panic(err)
	}

	file = fileutils.GetPage()
	fileContent, err := fileutils.ReadFile(file)

	fmt.Scanf("h")
	fmt.Printf(fileContent[0])
}

func processAction(args []string) {
	command := args[2]

	switch command {
	case "add":
		fmt.Println("Füge Aufgabe hinzu:", args[3])
		commandutils.AddTask(args[3])
	case "done":
		fmt.Println("Aufgabe erledigt:", args[3])
	case "del":
		fmt.Println("Aufgabe gelöscht:", args[3])
	}
}
