package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	files "github.com/asyarb/aoc-2024/pkg"
)

type Direction int

const (
	Increasing Direction = iota
	Decreasing
	Unknown
)

type Report []int

func (r Report) Omit(i int) Report {
	cpy := make(Report, r.Size())
	copy(cpy, r)

	return append(cpy[:i], cpy[i+1:]...)
}

func (r Report) Size() int {
	return len(r)
}

func (r Report) Safe() bool {
	direction := Unknown

	for i, num := range r {
		// if we're on the last element, just move on.
		if (i + 1) == r.Size() {
			continue
		}

		next := r[i+1]

		// must increase by 1 - 3
		diff := int(math.Abs(float64(next - num)))
		if diff > 3 || diff < 1 {
			return false
		}

		if direction == Unknown {
			if num < next {
				direction = Increasing
			} else {
				direction = Decreasing
			}
		}

		if direction == Increasing && num > next {
			return false
		}

		if direction == Decreasing && num < next {
			return false
		}
	}

	return true
}

func getReports(path string) []Report {
	file := files.OpenRelative(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports []Report

	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.Fields(line)

		nums := make([]int, len(strNums))

		for i, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatal(err)
			}

			nums[i] = num
		}

		reports = append(reports, nums)
	}

	return reports
}

func partOne(path string) int {
	reports := getReports(path)
	safe := 0

	for _, r := range reports {
		if r.Safe() {
			safe++
		}
	}

	return safe
}

func partTwo(path string) int {
	reports := getReports(path)
	safe := 0

	for _, r := range reports {
		if r.Safe() {
			safe++
		} else {
			for j := range r {
				if r.Omit(j).Safe() {
					safe++
					break
				}
			}

		}
	}

	return safe
}

func main() {
	fmt.Println("Part 1:", partOne("./input.txt"))
	fmt.Println("Part 2:", partTwo("./input.txt"))
}
