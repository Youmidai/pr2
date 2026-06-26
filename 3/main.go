package main

import (
	"fmt"
	"math/rand"
)

func Shuffle(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	rows, cols := 3, 3
	total := rows * cols

	pool := make([]int, total)
	for i := 0; i < total; i++ {
		pool[i] = i + 1
	}

	Shuffle(pool)

	matrix := make([][]int, rows)
	k := 0

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = pool[k]
			k++
		}
	}

	for _, row := range matrix {
		fmt.Println(row)
	}
}
