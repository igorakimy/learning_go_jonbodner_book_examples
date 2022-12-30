// Пакет money предоставляет различные утилиты с целью облегчить
// управление денежными средствами
package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// Money содержит информацию о размере денежной суммы и о том,
// в какой валюте она исчисляется
type Money struct {
	Value    decimal.Decimal
	Currency string
}

// Convert преобразует размер суммы из одной валюты в другую.
//
// Эта функция принимает два параметра: экземпляр структуры Money,
// содержащий преобразуемый размер суммы, и строку с названием той валюты,
// в которую преобразуется денежная сумма.
// Convert возвращает сумму в указанной валюте или ошибку в том случае,
// если валюта будет неизвестной или в нее нельзя будет преобразовать
// денежную сумму.
// При возвращении ошибки экземпляр структуры Money устанавливается
// равным нулевому значению.
//
// Поддерживаются следующие валюты:
// USD — доллар США
// CAD — канадский доллар
// EUR — евро
// INR — индийская рупия
//
// Более подробные сведения о курсах обмена валют можно найти
// по адресу https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
	return Money{
		Value:    decimal.NewFromFloat(30.40),
		Currency: to,
	}, nil
}

func main() {
	money := Money{decimal.NewFromInt(10), "EUR"}
	_, err := Convert(money, "USD")
	if err != nil {
		fmt.Println("Invalid convert operation")
	}
}
