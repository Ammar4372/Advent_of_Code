package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	matrix := bytes.Split(buf[:], []byte("\n"))
	pat := [4]byte{'X', 'M', 'A', 'S'}
	count := 0
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// check horizontally
			if check(matrix, pat, 0, i, j, 0, 1) {
				count += 1
			}
			if check(matrix, pat, 0, i, j, 0, -1) {
				count += 1
			}
			// check vertically
			if check(matrix, pat, 0, i, j, 1, 0) {
				count += 1
			}
			if check(matrix, pat, 0, i, j, -1, 0) {
				count += 1
			}
			// check diagonally
			if check(matrix, pat, 0, i, j, 1, 1) {
				count += 1
			}
			if check(matrix, pat, 0, i, j, -1, -1) {
				count += 1
			}
			// check cross diagonally
			if check(matrix, pat, 0, i, j, -1, 1) {
				count += 1
			}
			if check(matrix, pat, 0, i, j, 1, -1) {
				count += 1
			}
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func check(matrix [][]byte, pat [4]byte, index, x, y, dirx, diry int) bool {
	if index == len(pat) {
		return true
	}
	if x < 0 || y < 0 || x >= len(matrix)-1 || y >= len(matrix[x]) {
		return false
	}
	if pat[index] != matrix[x][y] {
		return false
	}
	return check(matrix, pat, index+1, x+dirx, y+diry, dirx, diry)
}
