package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat positif (> 10): ")
	fmt.Scan(&input)

	if input <= 10 {
		fmt.Println("Bilangan harus lebih besar dari 10.")
		return
	}

	inputStr := strconv.Itoa(input)
	n := len(inputStr)
	var left, right string

	if n%2 == 0 {
		left = inputStr[:n/2]
		right = inputStr[n/2:]
	} else {
		left = inputStr[:n/2+1]
		right = inputStr[n/2+1:]
	}

	leftNum, _ := strconv.Atoi(left)
	rightNum, _ := strconv.Atoi(right)

	fmt.Println(leftNum, rightNum)

	sum := leftNum + rightNum
	fmt.Println(sum)
}
