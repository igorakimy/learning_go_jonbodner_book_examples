package create_file

import (
	"fmt"
	"os"
	"testing"
)

// createFile - это вспомогательная функция, вызываемая из нескольких тестов
func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	// записываем данные в переменную f
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	// выполняем тестирование, не беспокоясь о высвобождении ресурсов
	if fName != "tmpFile" {
		fmt.Printf("Name '%s' is incorrect", fName)
	}
}
