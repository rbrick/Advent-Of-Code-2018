package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	scanner *bufio.Scanner
	inputs  = []string{}
)

func init() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	scanner = bufio.NewScanner(f)
}
func part1() {
	threeCount := 0
	twoCount := 0
	for scanner.Scan() {
		text := scanner.Text()
		instances := map[int32]int32{}

		for _, x := range text {
			if v, ok := instances[x]; ok {
				instances[x] = v + 1
			} else {
				instances[x] = 1
			}
		}

		seen3, seen2 := false, false
		for _, v := range instances {
			if v == 3 && !seen3 {
				threeCount++
				seen3 = true
			}
			if v == 2 && !seen2 {
				twoCount++
				seen2 = true
			}
		}
		inputs = append(inputs, text)
	}

	fmt.Println("checksum:", twoCount*threeCount)
}

func part2() {
	loop:
	for i, v := range inputs {
		for x := range inputs {
			if x != i {
				v2 := []rune(inputs[x])
				sameCount := 1
				same := ""
				for rIdx, r := range v {
					if v2[rIdx] == r {
						// same string
						sameCount++
						same += string(v2[rIdx])
					} else {
						sameCount--
					}
				}

				if sameCount == len(v)-1 {
					fmt.Println("boxid:", same)
					break loop
				}
			}
			continue
		}
	}
}

func main() {
	part1()
	part2()
}
