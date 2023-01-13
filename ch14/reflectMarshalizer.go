package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type MyData struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	HasPet bool   `csv:"has_pet"`
}

func marshalHeader(vt reflect.Type) []string {
	var row []string
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			row = append(row, curTag)
		}
	}
	return row
}

func marshalOne(vv reflect.Value) ([]string, error) {
	// Создать пустой строковый срез
	var row []string
	// Получить тип значения
	vt := vv.Type()
	// Обойти все поля типа
	for i := 0; i < vt.NumField(); i++ {
		// Получить поле структуры по индексу
		fieldVal := vv.Field(i)
		// Если тег поля не содержит указанного ключа, продолжать цикл
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		// Проверить вид значения поля
		switch fieldVal.Kind() {
		case reflect.Int:
			row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
		// Если строковый тип
		case reflect.String:
			// Добавляем в итоговый строковый срез строку,
			// которую получаем вызывая метод String
			row = append(row, fieldVal.String())
		// Если булев тип
		case reflect.Bool:
			// Добавляем в итоговый срез булево значение преобразованное к строке
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		// В остальных случаях возвращаем ошибку
		default:
			return nil, fmt.Errorf("cannot handle field of kind %v", fieldVal.Kind())
		}
	}
	// Вернуть итоговый строковый срез
	return row, nil
}

func unmarshalOne(row []string, namePos map[string]int, vv reflect.Value) error {
	vt := vv.Type()
	// Обойти все поля вновь созданного экземпляра типа reflect.Value
	for i := 0; i < vv.NumField(); i++ {
		typeField := vt.Field(i)
		// Получить позицию поля, используя тег структуры csv
		pos, ok := namePos[typeField.Tag.Get("csv")]
		// Если поле не найдено, продолжить обход
		if !ok {
			continue
		}
		// Получить значение поля по его позиции
		val := row[pos]
		// Получить поле по его индексу
		field := vv.Field(i)
		// Преобразовать значение из типа string в подходящий тип и присвоить текущему полю
		switch field.Kind() {
		case reflect.Int:
			i, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(i)
		case reflect.String:
			field.SetString(val)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			field.SetBool(b)
		default:
			return fmt.Errorf("cannot handle field of kind %v", field.Kind())
		}
	}
	return nil
}

func Marshal(v interface{}) ([][]string, error) {
	// Получить экземпляр reflect.Value значения v
	sliceVal := reflect.ValueOf(v)
	// Если тип значения не равен срезу, тогда возвращаем ошибку
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	// Получить тип элемента внутри составного типа данных
	structType := sliceVal.Type().Elem()
	// Если вид типа не равен структуре, тогда возвращаем ошибку
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must me a slice of structs")
	}
	var out [][]string
	// Передать тип структуры для маршализации заголовка и сохранить его в переменной header
	header := marshalHeader(structType)
	// Добавить заголовок в конечный срез
	out = append(out, header)
	// Обойти все элементы среза структур, используя рефлексию
	for i := 0; i < sliceVal.Len(); i++ {
		// Передать экземпляр типа reflect.Value в функцию marshalOne
		row, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}
		// Добавить новый строковый срез
		out = append(out, row)
	}
	// Вернуть срез строковых срезов
	return out, nil
}

func Unmarshal(data [][]string, v interface{}) error {
	// Преобразовать указатель на срез структуры в экземпляр reflect.Value
	sliceValPtr := reflect.ValueOf(v)
	errMsg := "must be a pointer to a slice of structs"
	if sliceValPtr.Kind() != reflect.Ptr {
		return errors.New(errMsg)
	}
	// Получить базовый срез
	sliceVal := sliceValPtr.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New(errMsg)
	}
	// Получить тип содержащихся в срезе структур
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New(errMsg)
	}

	// Предполагается, что первая строка представляет собой заголовок с именами столбцов
	header := data[0]
	// Создать карту, которая позволяет ассоциировать значение тега
	// структуры csv с правильным элементом данных
	namePos := make(map[string]int, len(header))
	for k, v := range header {
		namePos[v] = k
	}

	// Затем обходим все оставшиеся строковые срезы
	for _, row := range data[1:] {
		// Создаем новый экземпляр типа reflect.Value,
		// использующий тип reflect.Type каждой структуры
		newVal := reflect.New(structType).Elem()
		// Вызываем функцию unmarshalOne, чтобы скопировать данные
		// текущего строкового среза в структуру
		err := unmarshalOne(row, namePos, newVal)
		if err != nil {
			return err
		}
		// Затем добавляем структуру в общий срез
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}
	return nil
}

func main() {
	data := `name,age,has_pet
Jon,"100",true
"Fred ""The Hammer"" Smith",42,false
Martha,37,"true"`

	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var entries []MyData
	err = Unmarshal(allData, &entries)
	if err != nil {
		panic(err)
	}
	fmt.Println(entries)
	// [{Jon 100 true} {Fred "The Hammer" Smith 42 false} {Martha 37 true}]

	// Преобразование записей в конечный результат
	out, err := Marshal(entries)
	if err != nil {
		panic(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	err = w.WriteAll(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(sb)
}
