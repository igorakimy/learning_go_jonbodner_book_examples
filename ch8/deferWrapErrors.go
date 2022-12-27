package main

import (
	"fmt"
)

func doThing1(val int) (int, error) {
	return 123, nil
}

func doThing2(val string) (string, error) {
	return "hello", nil
}

func doThing3(val1 int, val2 string) (string, error) {
	return "result", nil
}

// Иногда нужно обернуть в одно и то же сообщение несколько ошибок

func DoSomeThings(val1 int, val2 string) (string, error) {
	val3, err := doThing1(val1)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	result, err := doThing3(val3, val4)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	return result, nil
}

// Код выше можно упростить с помощью оператора defer

func DoSomeThings2(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}

func main() {

}
