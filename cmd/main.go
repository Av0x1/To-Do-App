package main

import (
	"example.org/to-do-app/internal/commandutils"
	"example.org/to-do-app/internal/ioutils"
	"fmt"
	"os"
)

func main() {
	ioutils.GreetUser()

	if len(os.Args) < 3 {
		fmt.Println("Usage: todo <add|complete|list|delete> <text>")
	}

	processAction(os.Args)

	fmt.Scanf("h")
}

func processAction(args []string) {
	command := args[2]

	switch command {
	case "add":
		fmt.Println("Füge Aufgabe hinzu: ", args[3])
		commandutils.AddTask(args[3])
		fmt.Println("Aufgabe hinzugefügt.")
	case "complete":
		fmt.Println("Aufgabe erledigt:", args[3])
		//commandutils.SetTaskDone(args[3])
		fmt.Println("Aufgabe erledigt.")
	case "list":
		fmt.Println("Anzeigen der To-Dos.")
	case "delete":
		fmt.Println("Aufgabe gelöscht:", args[3])
	}
}
