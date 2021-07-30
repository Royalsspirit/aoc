package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(cmd [][]string) int {
	var acc int = 0
	index := make(map[int]int)
	var i int = 0
	for i < len(cmd) {
		if index[i] != 0 {
			return -1
		}
		index[i] = i
		nbr, _ := strconv.Atoi((cmd)[i][1])
		if (cmd)[i][0] == "acc" {
			acc += nbr
		} else if (cmd)[i][0] == "jmp" {
			i = i + nbr - 1
		}

		i++
	}
	return acc

}
func part2(currentCmd string) string {
	if currentCmd == "jmp" {
		currentCmd = "nop"
	} else if currentCmd == "nop" {
		currentCmd = "jmp"
	}
	return currentCmd
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var cmd [][]string
	for scanner.Scan() {
		currentValue := scanner.Text()
		cmd = append(cmd, strings.Split(currentValue, " "))
	}
	start := 0
	var acc int
	for start < len(cmd) {
		if cmd[start][0] == "jmp" || cmd[start][0] == "nop" {
			cmdBackup := make([][]string, len(cmd))
			for k := range cmd {
				cmdBackup[k] = make([]string, len(cmd[k]))
				copy(cmdBackup[k], cmd[k])
			}

			cmdBackup[start][0] = part2(cmdBackup[start][0])
			acc = part1(cmdBackup)
			if acc != -1 {
				fmt.Println("final acc", acc)
				break
			}
		}

		start++
	}
}
