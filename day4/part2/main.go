package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type point struct {
	x int
	y int
}

type line struct {
	p1 point
	p2 point
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	matrix := bytes.Split(buf[:], []byte("\n"))
	pat := [3]byte{'M', 'A', 'S'}
	lines := []line{}
	count := 0
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			p1 := point{x: i, y: j}
			// check horizontally
			if ok, p2 := check(matrix, pat, 0, i, j, 0, 1); ok {
				l := line{p1: p1, p2: p2}
				lines = append(lines, l)
			}
			if ok, p2 := check(matrix, pat, 0, i, j, 0, -1); ok {
				l := line{p1: p1, p2: p2}
				lines = append(lines, l)
			}
			// check vertically
			if ok, p2 := check(matrix, pat, 0, i, j, 1, 0); ok {
				l := line{p1: p1, p2: p2}
				lines = append(lines, l)
			}
			if ok, p2 := check(matrix, pat, 0, i, j, -1, 0); ok {
				l := line{p1: p1, p2: p2}
				lines = append(lines, l)
			}
			// check diagonally
			if ok, p2 := check(matrix, pat, 0, i, j, 1, 1); ok {
				l := line{p1: p1, p2: p2}
				lines = append(lines, l)
			}
			if ok, p2 := check(matrix, pat, 0, i, j, -1, -1); ok {
				l := line{p1: p1, p2: p2}
				lines = append(lines, l)
			}
			// check cross diagonally
			if ok, p2 := check(matrix, pat, 0, i, j, -1, 1); ok {
				l := line{p1: p1, p2: p2}
				lines = append(lines, l)
			}
			if ok, p2 := check(matrix, pat, 0, i, j, 1, -1); ok {
				l := line{p1: p1, p2: p2}
				lines = append(lines, l)
			}
		}
	}
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			if check_points(matrix, lines[i], lines[j]) {
				count += 1
			}
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func check(matrix [][]byte, pat [3]byte, index, x, y, dirx, diry int) (bool, point) {
	if index == len(pat) {
		p2 := point{x: x - dirx, y: y - diry}
		return true, p2
	}
	if x < 0 || y < 0 || x >= len(matrix)-1 || y >= len(matrix[x]) {
		return false, point{}
	}
	if pat[index] != matrix[x][y] {
		return false, point{}
	}
	return check(matrix, pat, index+1, x+dirx, y+diry, dirx, diry)
}
func check_points(mat [][]byte, l1, l2 line) bool {
	mid1 := calculate_midpoint(l1)
	mid2 := calculate_midpoint(l2)
	return mid1.x == mid2.x && mid1.y == mid2.y && mat[mid1.x][mid1.y] == 'A' && is_digonal(l1) && is_digonal(l2)
}

func calculate_midpoint(l line) point {
	x := (l.p1.x + l.p2.x) / 2
	y := (l.p1.y + l.p2.y) / 2
	return point{x: x, y: y}
}
func is_digonal(l line) bool {
	return l.p1.x != l.p2.x && l.p1.y != l.p2.y
}
