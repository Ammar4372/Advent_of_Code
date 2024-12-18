package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func mergesort(list []int) []int {
	lenght := len(list)
	if lenght <= 1 {
		return list
	}
	pivot := 0
	if lenght%2 == 1 {
		pivot = (lenght + 1) / 2
	} else {
		pivot = lenght / 2
	}
	return merge(mergesort(list[0:pivot]), mergesort(list[pivot:]))
}
func merge(left []int, right []int) []int {
	i, j, merged := 0, 0, []int{}
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			merged = append(merged, right[j])
			j++
		} else {
			merged = append(merged, left[i])
			i++
		}
	}
	for _, v := range left[i:] {
		merged = append(merged, v)
	}
	for _, v := range right[j:] {
		merged = append(merged, v)
	}
	return merged
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
	fmt.Printf("lenght of list1 : %v, lenght of list2: %v\n", len(list1), len(list2))
	list1 = mergesort(list1)
	list2 = mergesort(list2)
	i, j, sum_of_distance := 0, 0, 0
	for i < len(list1) && j < len(list2) {
		if list1[i] > list2[j] {
			sum_of_distance += list1[i] - list2[j]
		} else {
			sum_of_distance += list2[j] - list1[i]
		}
		i += 1
		j += 1
	}
	fmt.Printf("sum of distance: %d\n", sum_of_distance)
}
