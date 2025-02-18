package commandutils

import (
	"encoding/csv"
	"fmt"
	"strconv"

	"example.org/to-do-app/fileutils"
)

func AddTask(taskName string) {
	file := fileutils.GetPage()
	writer := csv.NewWriter(file)
	reader := csv.NewReader(file)

	lastTodo := readLastLine(reader)
	id := getNextId(lastTodo)

	newRow := []string{id, taskName, "Aktiv"}

	err := writer.Write(newRow)

	if err != nil {
		fmt.Println("Fehler beim Schreiben der Zeile:", err)
		return
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		fmt.Println("Fehler beim Flushen:", err)
	}
}

func readLastLine(reader *csv.Reader) []string {
	var lastTodo []string

	for {
		record, err := reader.Read()
		if err != nil {
			break // EOF oder Fehler
		}
		lastTodo = record
	}

	return lastTodo
}

func getNextId(lastTodo []string) string {
	var id string

	if len(lastTodo) == 0 {
		id = "1"
	} else {
		tmpId, err := strconv.Atoi(lastTodo[0])

		if err != nil {
			fmt.Println("Fehler biem Konvertieren der ID:", err)
		}

		id = strconv.Itoa(tmpId + 1)
	}

	return id
}
