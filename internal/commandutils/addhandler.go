package commandutils

import (
	"encoding/csv"
	"example.org/to-do-app/internal/fileutils"
	"fmt"
	"strconv"
)

func AddTask(taskName string) {
	file := fileutils.GetFile()
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
			break
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
			fmt.Println("Fehler beim Konvertieren der ID:", err)
		}

		id = strconv.Itoa(tmpId + 1)
	}

	return id
}
