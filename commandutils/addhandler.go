package commandutils

import (
	"encoding/csv"
	"fmt"

	"example.org/to-do-app/fileutils"
)

func AddTask(taskName string) {
	file := fileutils.GetPage()

	writer := csv.NewWriter(file)
	newRow := []string{taskName}

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
