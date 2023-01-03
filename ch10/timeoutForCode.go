package main

import (
	"errors"
	"fmt"
	"time"
)

func doSomeWork() (int, error) {
	time.Sleep(2 * time.Second)
	return 10, nil
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	// Считать значение time.Duration по прошествию указанного времени из канала
	// time.After
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func main() {
	res, err := timeLimit()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
