package fileutils

import "os"

func GetFile() *os.File {
	file, err := os.OpenFile("todo.csv", os.O_APPEND, 0666)

	return createFile(err, file)
}

func createFile(err error, file *os.File) *os.File {
	if err != nil {
		file, err := os.Create("todo.csv")

		if err != nil {
			panic(err)
		}

		return file
	}

	return file
}
