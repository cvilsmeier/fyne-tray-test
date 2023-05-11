package main

import (
	"os"
	"time"

	"fyne.io/systray"
)

func main() {
	systray.Run(onReady, onExit) // blocking
}

func onReady() {
	// show white icon
	systray.SetIcon(mustLoadIcon("white.png"))
	// sleep 1 second
	time.Sleep(1 * time.Second)
	// show red icon - this does not work
	systray.SetIcon(mustLoadIcon("red.png"))
}

func onExit() {
	// clean up here
}

func mustLoadIcon(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}
