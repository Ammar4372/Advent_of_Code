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
	defer file.Close()
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
	s := string(buf[:])
	safe_reports := 0
	for _, value := range strings.Split(s, "\n") {
		numbers := strings.Split(value, " ")
		if len(numbers) == 1 {
			continue
		}
		nums := []int64{}
		for i := range numbers {
			number, _ := strconv.ParseInt(numbers[i], 10, 32)
			nums = append(nums, number)
		}
		if check_report(nums) {
			safe_reports += 1
		} else {
			fmt.Printf("Report: %v\n", nums)
			for i := 0; i < len(nums); i++ {
				nums_cp := copy_report(nums[0:i], nums[i+1:])
				fmt.Printf("Report: %v, excluding: %d, index: %d\n", nums_cp, nums[i], i)
				if check_report(nums_cp) {
					safe_reports += 1
					break
				}
			}
		}
	}
	fmt.Printf("num of safe reports: %d\n", safe_reports)
}

func check_report(nums []int64) bool {
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

func copy_report(left, right []int64) []int64 {
	cp := []int64{}
	for i := range left {
		cp = append(cp, left[i])
	}
	for i := range right {
		cp = append(cp, right[i])
	}
	return cp
}
