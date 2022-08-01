package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(fishes *[]int) {
}

func part2(fishes *map[int]int) {
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var uniqTab [][]string
	var proposal [][]string
	for scanner.Scan() {

		currentValue := scanner.Text()
		lines := strings.Split(currentValue, " | ")
		secLeft := strings.Fields(lines[0])
		secRight := strings.Fields(lines[1])
		uniqTab = append(uniqTab, secLeft)
		proposal = append(proposal, secRight)

	}
	count := 0
	//var matrix [][]int = [][]int{
	//	[]int{1, 1, 1, 1, 1, 1, 0},
	//	[]int{0, 1, 1, 0, 0, 0, 0},
	//	[]int{1, 1, 0, 1, 1, 0, 1},
	//	[]int{1, 1, 1, 1, 0, 0, 1},
	//	[]int{0, 1, 1, 0, 0, 1, 1},
	//	[]int{1, 0, 1, 1, 0, 1, 1},
	//	[]int{1, 0, 1, 1, 1, 1, 1},
	//	[]int{1, 1, 1, 0, 0, 0, 0},
	//	[]int{1, 1, 1, 1, 1, 1, 1},
	//	[]int{1, 1, 1, 1, 0, 1, 1},
	//}
	for _, v := range uniqTab {
		for i := 0; i < len(v); i++ {
			tab := strings.Split(v[i], "")
			sort.Strings(tab)
			v[i] = strings.Join(tab, "")
		}
	}
	for _, v := range proposal {
		for i := 0; i < len(v); i++ {
			tab := strings.Split(v[i], "")
			sort.Strings(tab)
			v[i] = strings.Join(tab, "")
		}
	}
	max := 0
	for k, v := range uniqTab {

		number := make(map[int]string)
		var fiveNumber []string
		var sixNumber []string

		for i := 0; i < len(v); i++ {
			if len(v[i]) == 2 {
				number[1] = v[i]
				count++
			}
			if len(v[i]) == 3 {
				number[7] = v[i]
				count++
			}

			if len(v[i]) == 7 {
				number[8] = v[i]
			}
			if len(v[i]) == 4 {

				number[4] = v[i]

				count++

			}
			if len(v[i]) == 5 {
				fiveNumber = append(fiveNumber, v[i])
			}

			if len(v[i]) == 6 {
				sixNumber = append(sixNumber, v[i])
			}

		}
		/**
		For the 5-segement numbers:

		3 is the only one that includes all of 1

		5 is the one that doesn't but is included in 9

		2 is left

		And for the 6-segment numbers

		6 is the only one that does not contain 1

		9 is the one containing both 1 and 4

		0 is left
		* we have 3 fix len segments and 3 six len segments
		*
		* we assume for five len segment
		* digit number 1:     is the only one which is totally in 3 -
				   |                                         |
		*                                                           -
		*                  |                                         |
		*                                                           -
		*
		* digit number 2: -    has 2 segment which are not in 4    | |
		*                  |                                        -
		*		  -                                          |
		*		 |
		*		  -
		* for 6 len segments
		* digit numnber 4      has all its segments in 9   -
		                  | |                             | |
		*                  -	                           -
		*                   |                               |
		                                                   -
		*
		* digit number 1   |  has 1 segment which is not in 6   -
		                   |                           |
		*                                               -
							       | |
							        -


						**/
		for _, v := range fiveNumber {
			var diff int
			var diffWith4 int
			for _, vv := range number[1] {
				if strings.Index(string(v), string(vv)) == -1 {
					diff++
				}
			}
			for _, vv := range number[4] {
				if strings.Index(string(v), string(vv)) == -1 {
					fmt.Println("five number not found", string(vv), "in", v)
					diffWith4++
				}
			}

			fmt.Println("diff", diff)
			if diff == 0 {
				number[3] = v
			} else if diffWith4 == 2 {
				number[2] = v
			} else {
				number[5] = v
			}
		}
		for _, v := range sixNumber {
			var diff int
			var diffWith4 int
			for _, vv := range number[1] {
				if strings.Index(string(v), string(vv)) == -1 {
					diff++
				}
			}
			for _, vv := range number[4] {
				if strings.Index(string(v), string(vv)) == -1 {
					diffWith4++
				}
			}
			if diffWith4 == 0 {
				number[9] = v
			} else if diff == 1 {
				number[6] = v
			} else {
				number[0] = v
			}
		}
		sum := ""
		for _, vv := range proposal[k] {
			for kk, vvv := range number {
				if vvv == vv {
					sum += strconv.Itoa(kk)
				}
			}
		}
		sumInt, _ := strconv.Atoi(sum)
		max += sumInt

	}
	fmt.Println("max", max)
	if err != nil {
		fmt.Println("error", err)
	}
	// fmt.Println("max", max)
	// part1(card, numbers)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
