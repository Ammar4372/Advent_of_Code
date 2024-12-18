package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value       int
	count       int
	left_child  *node
	right_child *node
}

func newtree(list []int) *node {
	root := node{
		value: -1,
		count: 0,
	}
	for i := range list {
		root.addNode(list[i])
	}
	return &root
}

func (root *node) addNode(value int) {
	if root.value == value {
		root.count += 1
		return
	}
	if value > root.value {
		if root.left_child != nil {
			root.left_child.addNode(value)
			return
		}
		root.left_child = &node{
			value: value,
			count: 1,
		}
	}
	if value < root.value {
		if root.right_child != nil {
			root.right_child.addNode(value)
			return
		}
		root.right_child = &node{
			value: value,
			count: 1,
		}
	}
	return
}
func (root *node) getCount(value int) int {
	if root.value == value {
		return root.count
	}
	if value > root.value && root.left_child != nil {
		return root.left_child.getCount(value)
	}
	if value < root.value && root.right_child != nil {
		return root.right_child.getCount(value)
	}
	return 0
}
func (root *node) calculateSimilarityScore(list []int) int {
	similarityScore := 0
	for i := range list {
		count := root.getCount(list[i])
		similarityScore += list[i] * count
	}
	return similarityScore
}
func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err.Error())
	}
	buf := make([]byte, 16384)
	for {
		_, err = file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
	}
	list1 := []int{}
	list2 := []int{}
	str := string(buf[:])
	for _, value := range strings.Split(str, "\n") {
		numbers := strings.Split(value, " ")
		if len(numbers) != 1 {
			number1, _ := strconv.ParseInt(numbers[0], 10, 32)
			number2, _ := strconv.ParseInt(numbers[len(numbers)-1], 10, 32)
			list1 = append(list1, int(number1))
			list2 = append(list2, int(number2))
		}
	}
	tree := newtree(list2)
	similarityScore := tree.calculateSimilarityScore(list1)
	fmt.Printf("similarity score: %d\n", similarityScore)
}
