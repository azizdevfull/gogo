package main

import (
	"fmt"
	"sync"
	"time"
)

func generateNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine tugashini bildiradi
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i // Goroutine’dan ma'lumotni channel orqali yuborish
	}
	close(ch) // Channelni yopamiz, boshqa ma'lumot kelmasligi uchun
}

func printLetters(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine tugashini bildiradi
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, letter := range letters {
		time.Sleep(1 * time.Second)
		ch <- letter
	}
	close(ch) // Channelni yopamiz
}

func main() {
	var wg sync.WaitGroup // WaitGroup yaratamiz

	ch := make(chan int)     // Channel yaratish
	ch2 := make(chan string) // Ikkinchi channel

	// Goroutines ishga tushirildi
	wg.Add(2) // Ikki goroutine qo'shish
	go generateNumbers(ch, &wg)
	go printLetters(ch2, &wg)

	// Channel'dan qiymatlarni olish
	go func() {
		for num := range ch {
			fmt.Println(num)
		}
	}()

	go func() {
		for letter := range ch2 {
			fmt.Println(letter)
		}
	}()

	// Goroutines tugashini kutish
	wg.Wait()
}
