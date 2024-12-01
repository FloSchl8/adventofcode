package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input/2024/day01/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(file)

	fmt.Println("day01(input): ", day01(input))
	fmt.Println("day02(input): ", day02(input))
}

func day01(input string) int {
	lines := strings.Split(input, "\n")

	ns := make([]int, 0)
	ms := make([]int, 0)

	sum := 0
	for _, line := range lines {
		if line != "" {
			numbers := strings.Split(line, "   ")
			n, _ := strconv.Atoi(numbers[0])
			m, _ := strconv.Atoi(numbers[1])

			ns = append(ns, n)
			ms = append(ms, m)
		}
	}

	slices.Sort(ns)
	slices.Sort(ms)

	for i := range len(ns) {
		sum += int(math.Abs(float64(ms[i] - ns[i])))
	}
	return sum
}

func day02(input string) int {
	lines := strings.Split(input, "\n")

	ns := make([]int, 0)
	ms := make([]int, 0)
	for _, line := range lines {
		if line != "" {
			numbers := strings.Split(line, "   ")
			n, _ := strconv.Atoi(numbers[0])
			m, _ := strconv.Atoi(numbers[1])
			ns = append(ns, n)
			ms = append(ms, m)

		}
	}

	ncount := make(map[int]int)
	for _, v := range ns {
		ncount[v]++
	}

	mcount := make(map[int]int)
	for _, v := range ms {
		mcount[v]++
	}

	sum := 0

	for k, v := range ncount {
		if m, ok := mcount[k]; ok {
			sum += k * v * m
		}
	}

	return sum
}
