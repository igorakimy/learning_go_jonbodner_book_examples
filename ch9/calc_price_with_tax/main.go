package main

import (
	"fmt"
	"log"
	"os"

	"github.com/learning-go-book/formatter"
	"github.com/shopspring/decimal"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need two parameters: amount and percent")
		os.Exit(1)
	}
	// Получить новое дробное число из переданной строки
	amount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	percent, err := decimal.NewFromString(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	// Вернуть результат деления значения процента на 100
	percent = percent.Div(decimal.NewFromInt(100))
	// Вычислить увеличение текущего кол-ва на заданный процент, а также округлить
	// значение до двух знаков после точки
	total := amount.Add(amount.Mul(percent)).Round(2)
	fmt.Println(formatter.Space(80, os.Args[1], os.Args[2],
		total.StringFixed(2)))
}
