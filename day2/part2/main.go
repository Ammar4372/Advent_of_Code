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
	buf := make([]byte, 19185)
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
	all_inc := false
	all_dec := false
	tolerance_used := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			if all_dec && !tolerance_used {
				tolerance_used = true
				nums = append(nums[:i], nums[i+1:]...)
				continue
			}
			all_inc = true
			all_dec = false
		} else if nums[i] > nums[i+1] {
			if all_inc && !tolerance_used {
				tolerance_used = true
				nums = append(nums[:i], nums[i+1:]...)
				continue
			}
			all_dec = true
			all_inc = false
		} else {
			if !tolerance_used {
				tolerance_used = true
				nums = append(nums[:i], nums[i+1:]...)
				continue
			}
			all_dec = false
			all_inc = false
		}
	}
	if !all_dec && !all_inc {
		fmt.Printf("Report: %v\n", nums)
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		var difference int64
		if all_inc {
			difference = nums[i+1] - nums[i]
		} else {
			difference = nums[i] - nums[i+1]
		}
		if difference < 1 || difference > 3 {
			if tolerance_used {
				break
			}
			tolerance_used = true
			nums = append(nums[:i+1], nums[i+2:]...)
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		var difference int64
		if all_inc {
			difference = nums[i+1] - nums[i]
		} else {
			difference = nums[i] - nums[i+1]
		}
		if difference < 1 || difference > 3 {
			fmt.Printf("Report: %v\n", nums)
			return false
		}
	}
	fmt.Printf("Report: %v\n", nums)
	return true
}
