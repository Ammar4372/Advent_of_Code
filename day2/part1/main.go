package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err.Error())
	}
	buf := make([]byte, 191858)
	for {
		_, err = file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
	}
	s := string(buf)
	safe_reports := 0
	for _, value := range strings.Split(s, "\n") {
		if safe := check_report(value); safe {
			safe_reports += 1
		}
	}
	fmt.Printf("num of safe reports: %d\n", safe_reports)
}

func check_report(s string) bool {
	numbers := strings.Split(s, " ")
	if len(numbers) == 1 {
		return false
	}
	nums := []int64{}
	for i := range numbers {
		number, _ := strconv.ParseInt(numbers[i], 10, 32)
		nums = append(nums, number)
	}
	diff := nums[0] - nums[1]
	for i := 0; i < len(nums)-1; i++ {
		var difference int64
		if diff > 0 {
			difference = nums[i] - nums[i+1]
		} else {
			difference = nums[i+1] - nums[i]
		}
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}
