package main

import (
	"fmt"
	"os"

	"example.org/to-do-app/commandutils"
	"example.org/to-do-app/ioutils"
)

func main() {
	ioutils.GreetUser()

	if len(os.Args) < 3 {
		fmt.Println("Usage: todo <add|delete> <text>")
	}

	processAction(os.Args)

	fmt.Scanf("h")
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
