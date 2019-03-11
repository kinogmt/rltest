package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"time"
)

func main() {
	rl, _ := readline.New("> ")
	defer rl.Close()

	fmt.Printf("This line will be cleaned up after 10 seconds...")

	time.Sleep(10 * time.Second)
	_, _ = rl.Readline()
}
