package main

import (
	"fmt"
)

type ChannelScoreboardManager chan func(map[string]int)

func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	var out int
	var ok bool
	done := make(chan struct{})
	csm <- func(m map[string]int) {
		out, ok = m[name]
		close(done)
	}
	<-done
	return out, ok
}

func (csm ChannelScoreboardManager) Update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
}

func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
	// Объявить карту
	scoreboard := map[string]int{}
	// Слушать два канала
	for {
		select {
		// Получить сигнал о завершении работы
		case <-done:
			return
		// Получить функцию для чтения или модификации карты
		case f := <-in:
			f(scoreboard)
		}
	}
}

func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
	ch := make(ChannelScoreboardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func main() {
	ch, fClose := NewChannelScoreboardManager()
	defer fClose()
	ch.Update("hello", 10)
	ch.Update("world", 12)
	if val, ok := ch.Read("hello"); ok {
		fmt.Println(val)
	}
}
