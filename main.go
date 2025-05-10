package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0 ga bo'lish mumkun emas")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Xato:", err)
	} else {
		fmt.Println("Natija:", result)
	}
}
