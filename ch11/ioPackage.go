package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	// Создаем буфер один раз, а затем многократно его используем с каждым
	// вызовом метода r.Read.
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		// Используя возвращаемое методом r.Read значение n, узнаем, сколько
		// байтов было записано в буфер.
		n, err := r.Read(buf)
		// Производим обход части среза buf, чтобы обработать считанные данные.
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		// Сигналом о завершении чтения из переменной r для нас служит
		// возвращение методом r.Read ошибки io.EOF. Это не совсем ошибка,
		// а просто сигнал о том, что в экземпляре интерфейса io.Reader больше
		// нет непрочитанных данных.
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		_ = gr.Close()
		_ = r.Close()
	}, nil
}

func main() {
	s := "hello world"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counts)
	// map[d:1 e:1 h:1 l:3 o:2 r:1 w:1]

	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	counts, err = countLetters(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counts)
	// map[a:2 d:2 e:2 h:2 l:3 n:1 o:4 r:2 u:1 w:2 y:1]
}
