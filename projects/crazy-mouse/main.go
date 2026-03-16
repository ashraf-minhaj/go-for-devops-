package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for i := 0; i < 10; i++ { // 10 moves
		robotgo.MoveMouse(100+i*10, 100+i*10)
		time.Sleep(1 * time.Second)
	}
}
