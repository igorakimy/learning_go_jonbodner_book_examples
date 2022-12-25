package main

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// Встраивание интерфейсов в интерфейс

type ReadCloser interface {
	Reader
	Closer
}

func main() {

}
