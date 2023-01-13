package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	s := []int{10, 20, 30}
	sHdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len)
	// 3
	fmt.Println(sHdr.Cap)
	// 3

	// Поскольку размер типа int может составлять 32 или 64 бита, нужно
	// использовать функцию unsafe.Sizeof, чтобы узнать, сколько байтов
	// занимает каждое значение в том блоке памяти, на который указывает поле Data.
	intByteSize := unsafe.Sizeof(s[0])
	fmt.Println(intByteSize)
	// 8
	for i := 0; i < sHdr.Len; i++ {
		// Привести к типу uintptr значение i и умножаем его на размер типа int, добавляем
		// результат к содержимому пола Data, преобразовываем значение типа uintptr в небезопасный
		// указатель(unsafe.Pointer), а затем в указатель на тип int и, наконец, производим
		// разыменовываем указатель на тип int, чтобы получить значение.
		intVal := *(*int)(unsafe.Pointer(sHdr.Data + intByteSize*uintptr(i)))
		fmt.Println(intVal)
	}
	// 10
	// 20
	// 30
	runtime.KeepAlive(s)
}
