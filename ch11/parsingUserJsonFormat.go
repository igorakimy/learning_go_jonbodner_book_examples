package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type NewOrder struct {
	ID          string      `json:"id"`
	DateOrdered RFC822ZTime `json:"date_ordered"`
	CustomerID  string      `json:"customer_id"`
	Items       []NewItem   `json:"items"`
}

type NewItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RFC822ZTime struct {
	time.Time
}

func (rt *RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC822ZTime{t}
	return nil
}

func main() {
	data := `{
		"id":"12345",
		"date_ordered":"04 Jan 22 18:27 +0300",
		"customer_id":"3",
		"items":[
			{"id":"xyz123","name":"Thing 1"},
			{"id":"abc789","name":"Thing 2"}
		]
	}`
	var order NewOrder
	err := json.Unmarshal([]byte(data), &order)
	if err != nil {
		panic(err)
	}
	fmt.Println(order)

	out, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
