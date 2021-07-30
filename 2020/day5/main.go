package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var myArray [][]string
	var card []int
	for scanner.Scan() {
		currentValue := scanner.Text()
		fmt.Println("current", currentValue)
		tmp := strings.Split(currentValue, "")
		myArray = append(myArray, tmp)
	}
	initRow := make([]int, 128)
	for k, _ := range initRow {
		initRow[k] = k
	}
	initColumn := make([]int, 8)
	for k, _ := range initColumn {
		initColumn[k] = k
	}
	for _, v := range myArray {
		row := initRow
		column := initColumn
		fmt.Println("seat managing", v)
		for _, vv := range v {
			fmt.Println("character managing", vv)
			fmt.Println("current row", row)
			fmt.Println("current column", column)

			if vv == "F" {
				midRow := math.Ceil(float64(len(row) / 2))
				fmt.Println("midrow", midRow)
				row = row[0:int(midRow)]
			} else if vv == "B" {

				midRow := math.Ceil(float64(len(row) / 2))
				row = row[int(midRow):]
			}

			if vv == "L" {

				midColumn := math.Ceil(float64(len(column) / 2))
				column = column[0:int(midColumn)]
			} else if vv == "R" {

				midColumn := math.Ceil(float64(len(column) / 2))
				column = column[int(midColumn):]
			}
			fmt.Println("end row", row)
			fmt.Println("end column", column)
		}
		fmt.Println("final slice row", row)
		fmt.Println("final slice column", column)
		card = append(card, row[0]*8+column[0])
		row = initRow
		column = initColumn
	}
	sort.Ints(card)
	fmt.Println("card", card)
}
