package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	files "github.com/asyarb/aoc-2024/pkg"
)

func getLists(path string) ([]int, []int) {
	file := files.OpenRelative(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lefts []int
	var rights []int

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		left, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}

		right, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	return lefts, rights
}

func partOne(path string) int {
	sum := 0
	lefts, rights := getLists(path)

	sort.Ints(lefts)
	sort.Ints(rights)

	for i, left := range lefts {
		right := rights[i]

		diff := int(math.Abs(float64(left - right)))
		sum = sum + diff
	}

	return sum
}

func partTwo(path string) int {
	lefts, rights := getLists(path)

	scores := make(map[int]int)

	for _, right := range rights {
		_, exists := scores[right]
		if exists {
			scores[right]++
		} else {
			scores[right] = 1
		}
	}

	similarity := 0
	for _, left := range lefts {
		score, exists := scores[left]
		if !exists {
			continue
		}

		similarity = similarity + (left * score)
	}

	return similarity
}

func main() {
	fmt.Println("Part 1:", partOne("./input.txt"))
	fmt.Println("Part 2:", partTwo("./input.txt"))
}
