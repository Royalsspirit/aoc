package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fold struct {
	position string
	value    string
}
type dot struct {
	x     string
	y     string
	value string
}

func findMaxX(dots map[string]*dot) int {
	max := 0
	for k, _ := range dots {
		x := k[:2]
		kk, _ := strconv.Atoi(x)
		if kk > max {
			max = kk
		}
	}
	return max
}
func findMaxY(dots map[string]*dot) int {
	max := 0
	for k, _ := range dots {
		y := k[2:]
		kk, _ := strconv.Atoi(y)
		if kk > max {
			max = kk
		}
	}
	return max
}
func decryptMess(dots map[string]*dot) {
	x := findMaxX(dots)
	y := findMaxY(dots)
	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			xStr := strconv.Itoa(j)
			if len(xStr) == 1 {
				xStr = "0" + xStr
			}
			yStr := strconv.Itoa(i)
			if len(yStr) == 1 {
				yStr = "0" + yStr
			}
			if _, found := dots[xStr+yStr]; found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}

}
func foldToY(dots map[string]*dot, y string) {
	if len(y) == 1 {
		y = "0" + y
	}

	yy, _ := strconv.Atoi(y)
	for _, v := range dots {
		vy, _ := strconv.Atoi(v.y)
		if vy > yy {
			d := vy - yy
			newValue := yy - d
			yStr := strconv.Itoa(newValue)
			if len(yStr) == 1 {
				yStr = "0" + yStr
			}
			delete(dots, v.x+v.y)
			dots[v.x+yStr] = &dot{x: v.x, y: yStr, value: "#"}
		}
	}
}
func foldToX(dots map[string]*dot, x string) {
	if len(x) == 1 {
		x = "0" + x
	}

	xx, _ := strconv.Atoi(x)
	for _, v := range dots {
		vx, _ := strconv.Atoi(v.x)
		if vx > xx {
			d := vx - xx
			newValue := xx - d

			xStr := strconv.Itoa(newValue)
			if len(xStr) == 1 {
				xStr = "0" + xStr
			}
			delete(dots, v.x+v.y)
			dots[xStr+v.y] = &dot{x: xStr, y: v.y, value: "#"}
		}
	}
}
func part1(dots map[string]*dot, f []*fold) {
	for _, v := range f {
		// part 1 need to run only first fold
		// if k == 0 {
		if v.position == "y" {
			foldToY(dots, v.value)
			//fmt.Println("dots after y", len(dots))
			//fmt.Println("dots", dots)

		}
		if v.position == "x" {
			foldToX(dots, v.value)
			//fmt.Println("final", len(dots))

		}
		//}

	}

}

func part2(dots map[string]*dot) {
	fmt.Println("len", len(dots))
	decryptMess(dots)
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	dots := make(map[string]*dot)
	var folds []*fold
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		currentValue := scanner.Text()

		if currentValue != "" && strings.Index(currentValue, ",") > 0 {
			coord := strings.Split(currentValue, ",")
			var x, y string = coord[0], coord[1]
			if len(coord[0]) == 1 {
				x = "0" + x
			}
			if len(coord[1]) == 1 {
				y = "0" + y
			}
			dots[x+y] = &dot{x: x, y: y, value: "#"}

		}
		if strings.Index(currentValue, "y") > 0 {
			foldAttr := strings.Split(currentValue, "=")
			foldY := foldAttr[1]
			folds = append(folds, &fold{position: "y", value: foldY})
		}
		if strings.Index(currentValue, "x") > 0 {
			foldAttr := strings.Split(currentValue, "=")
			foldX := foldAttr[1]

			folds = append(folds, &fold{position: "x", value: foldX})
		}

	}
	if err != nil {
		fmt.Println("error", err)
	}
	part1(dots, folds)
	fmt.Println("result", len(dots))
	part2(dots)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
