package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// JSON, readers and writers
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}
	// Создаем временный файл
	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	// Гарантируем выполнение удаления временного файла после использования
	defer func() {
		err = os.Remove(tmpFile.Name())
		if err != nil {
			panic(err)
		}
	}()
	// Создаем кодировщик и кодируем структуру
	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}
	// Закрываем временный файл
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}
	// Считывание формата JSON. Открываем временный файл, возвращая io.Reader
	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}
	// Создаем пустой экземпляр структуры
	var fromFile Person
	// Создаем декодировщик из временного файла и декодируем закодированные
	// данные в экземпляр структуры
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}
	// Закрываем временный файл
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	// Выводим на экран содержимое экземпляра структуры
	fmt.Printf("%+v\n", fromFile)
	// {Name:Fred Age:40}
}
