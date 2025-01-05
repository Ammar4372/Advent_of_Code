package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "strconv"
)

type rule struct {
	first_page_number  string
	second_page_number string
}

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
	rules := parseRules(bytes.Split(buf, []byte("\n\n"))[0])
	queue := parsePrintQueue(bytes.Split(buf, []byte("\n\n"))[1])
	sum := 0
	for i := 0; i < len(queue)-1; i++ {
		if check(rules, queue[i], 0) {
			sum += getMiddlePageNumber(queue[i])
		}
	}
	fmt.Printf("Sum of middle pages: %d\n", sum)
}

func parseRules(buf []byte) []rule {
	rules := []rule{}
	for _, item := range bytes.Split(buf, []byte("\n")) {
		order := bytes.Split(item, []byte("|"))
		if len(order) == 2 {
			r := rule{}
			r.first_page_number = string(order[0])
			r.second_page_number = string(order[1])
			rules = append(rules, r)
		}
	}
	return rules
}

func parsePrintQueue(buf []byte) []string {
	queue := []string{}
	for _, item := range bytes.Split(buf, []byte("\n")) {
		job := string(item)
		queue = append(queue, job)
	}
	return queue
}

func check(rules []rule, job string, i int) bool {
	if i == len(rules) {
		return true
	}
	first_page_index := strings.Index(job, rules[i].first_page_number)
	second_page_index := strings.Index(job, rules[i].second_page_number)
	if first_page_index >= second_page_index && first_page_index >= 0 && second_page_index >= 0 {
		return false
	}
	return check(rules, job, i+1)
}

func getMiddlePageNumber(job string) int {
	queue := strings.Split(job, ",")
	length := len(queue)
	mid := 0
	if length%2 == 0 {
		mid = (length - 1) / 2
	} else {
		mid = length / 2
	}
	n, _ := strconv.ParseInt(queue[mid], 10, 64)
	return int(n)
}
