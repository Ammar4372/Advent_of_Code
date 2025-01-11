package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	floorMap := bytes.Split(buf[0:len(buf)-1], []byte("\n"))
	x, y := findGuardPosition(floorMap)
	pointsVisited := countPointsVisited(floorMap, x, y)
	printMap(floorMap)
	fmt.Printf("Distinct Point Visited:%d\n", pointsVisited)
}
func printMap(floorMap [][]byte) {
	for i := range floorMap {
		fmt.Printf("%s\n", floorMap[i])
	}
}
func findGuardPosition(floorMap [][]byte) (int, int) {
	for i := range floorMap {
		for j := range floorMap[i] {
			if floorMap[i][j] == '^' {
				return i, j
			}
		}
	}
	return 0, 0
}

func countPointsVisited(floorMap [][]byte, x, y int) int {
	pointsVisited := 1
	angle := (3 * math.Pi) / 2
	dirx := int(math.Sin(angle))
	diry := int(math.Cos(angle))
	for {
		if x+dirx >= len(floorMap) || y+dirx >= len(floorMap[0]) || x+dirx < 0 || y+diry < 0 {
			pointsVisited += 1
			break
		}
		if floorMap[x+dirx][y+diry] == '#' {
			angle += math.Pi / 2
			dirx = int(math.Sin(angle))
			diry = int(math.Cos(angle))
		}
		if floorMap[x][y] == '.' {
			pointsVisited += 1
			floorMap[x][y] = 'X'
		}
		x += dirx
		y += diry
	}
	return pointsVisited
}
