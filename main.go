package main

import "time"

func printNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		println(i)
	}
}
func printLetters() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for _, letter := range letters {
		time.Sleep(1 * time.Second)
		println(letter)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(11 * time.Second)
}
