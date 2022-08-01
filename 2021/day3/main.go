package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func flatTheArray(rowData []string) map[int]string {
	bits := make(map[int]string)
	for j := 0; j < len(rowData); j++ {
		for i := 0; i < len(rowData[j]); i++ {
			bits[i] += string(rowData[j][i])

		}

	}

	return bits

}
func commonBits(myArray map[int]string, common bool) (string, string) {
	commonBits := ""
	leastCommonBits := ""
	r := regexp.MustCompile(`0`)
	for i := 0; i < len(myArray); i++ {
		result := r.FindAllIndex([]byte(myArray[i]), -1)
		if len(result) > (len(myArray[i]) / 2) {
			commonBits += "0"
			fmt.Println("more 0 than 1")
			leastCommonBits += "1"
		} else {
			commonBits += "1"
			leastCommonBits += "0"
		}

	}
	return commonBits, leastCommonBits

}
func part1(myArray map[int]string) int64 {

	commonBits, leastCommonBits := commonBits(myArray, true)

	commonInteger, _ := strconv.ParseInt(commonBits, 2, 64)
	leastCommonInteger, _ := strconv.ParseInt(leastCommonBits, 2, 64)

	return commonInteger * leastCommonInteger
}

func filterDigit(data []string, index int, common bool) string {
	if len(data) == 1 {
		return data[0]
	}
	var zeros []string
	var ones []string
	for _, v := range data {
		if v[index] == '0' {
			zeros = append(zeros, v)
		} else {
			ones = append(ones, v)
		}
	}

	if len(zeros) > len(ones) == common {
		return filterDigit(zeros, index+1, common)
	}
	return filterDigit(ones, index+1, common)

}
func part2(rowData []string) int64 {
	oxygen, _ := strconv.ParseInt(filterDigit(rowData, 0, true), 2, 64)
	co2, _ := strconv.ParseInt(filterDigit(rowData, 0, false), 2, 64)
	return oxygen * co2
}

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	bits := make(map[int]string)
	var rowData []string
	for scanner.Scan() {
		currentValue := scanner.Text()
		rowData = append(rowData, currentValue)
		/**
		* put every bits with common array index in the same index
		**/
		for i := 0; i < len(currentValue); i++ {
			bits[i] += string(currentValue[i])

		}
	}
	if err != nil {
		fmt.Println("error", err)
	}
	firstPartResult := part1(bits)
	fmt.Println(firstPartResult)
	fmt.Println(part2(rowData))
}
