package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	record := `{"id":"12345","date_ordered":"2020-05-01T13:01:02Z","customer_id":"3","items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]}`
	fmt.Println(record)
	var newOrder Order
	err := json.Unmarshal([]byte(record), &newOrder)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range newOrder.Items {
		fmt.Println(item.ID)
	}
	rawJson, err := json.Marshal(newOrder)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rawJson))
	fmt.Println(record == string(rawJson))
}
