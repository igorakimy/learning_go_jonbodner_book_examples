package main

import (
	"encoding/binary"
	"math/bits"
	"unsafe"
)

// Проверка работы на "младшеконечной" платформе
var isLE bool

func init() {
	var x uint16 = 0xFF00
	// Использовать небезопасный указатель, чтобы преобразовать число в массив байтов
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	// Затем проверить, чему равен первый байт,
	// если FF - "старшеконечная", 00 - "младшеконечная"
	isLE = xb[0] == 0x00
}

type Data struct {
	Value  uint32
	Label  [10]byte
	Active bool
	// Go дополняет эти данные 1 байтом до "круглого" числа
}

// Сопоставление с помощью безопасного Go-кода

func DataFromBytes(b [16]byte) Data {
	d := Data{}
	d.Value = binary.BigEndian.Uint32(b[:4])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

// Сопоставление с помощью unsafe.Pointer

func DataFromBytesUnsafe(b [16]byte) Data {
	// Берем указатель на байтовый массив и преобразуем его в небезопасный указатель
	// unsafe.Pointer. Затем преобразуем небезопасный указатель в указатель *Data.
	// Поскольку нужно возвращать структуру, а не указатель, затем разыменовываем его.
	data := *(*Data)(unsafe.Pointer(&b))
	// Проверить, является ли текущая платформа "младшеконечной"
	if isLE {
		// Если это так, то инвертируем порядок байтов в поле Value
		data.Value = bits.ReverseBytes32(data.Value)
	}
	return data
}

// Запись содержимого структуры Data обратно в сеть безопасным Go-кодом

func BytesFromData(d Data) [16]byte {
	out := [16]byte{}
	binary.BigEndian.PutUint32(out[:4], d.Value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

// Запись содержимого структуры Data с помощью пакета unsafe

func BytesFromDataUnsafe(d Data) [16]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	b := *(*[16]byte)(unsafe.Pointer(&d))
	return b
}
