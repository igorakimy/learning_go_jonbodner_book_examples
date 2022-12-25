package main

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// бизнес-логика
	return "logic"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// получение данных
	c.L.Process("logic")
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
