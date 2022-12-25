package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func dataStub() error {
	// Одна пара фигурных скобок используется для типа interface{},
	// вторая пара служит для создания экземпляра карты
	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile("testdata/sample.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, &data)
	// содержимое contents теперь содержится в карте data
	if err != nil {
		return err
	}
	return nil
}

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

func main() {
	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}

	fmt.Println(i)
	// {Fred Fredson}

}
