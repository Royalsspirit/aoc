package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type neighbor struct {
	x        int
	y        int
	value    int
	neighbor []*neighbor
}

func buildNeighbor(o [][]string, n *neighbor, x int, y int, positionNeighbor map[string]*neighbor) *neighbor {

	valueInt, _ := strconv.Atoi(o[y][x])
	xStr := strconv.Itoa(x)
	yStr := strconv.Itoa(y)

	if len(xStr) == 1 {
		xStr = "0" + xStr
	}
	if len(yStr) == 1 {
		yStr = "0" + yStr
	}
	var child *neighbor
	if _, found := positionNeighbor[xStr+yStr]; found {
		child = positionNeighbor[xStr+yStr]
	} else {
		child = &neighbor{
			x:     x,
			y:     y,
			value: valueInt,
		}

		positionNeighbor[xStr+yStr] = child
	}
	n.neighbor = append(n.neighbor, child)
	return n
}
func buildTree(o [][]string) map[string]*neighbor {
	positionNeighbor := make(map[string]*neighbor)
	for y := 0; y < len(o); y++ {
		for x := 0; x < len(o[y]); x++ {
			valueInt, _ := strconv.Atoi(o[y][x])
			xStr := strconv.Itoa(x)
			yStr := strconv.Itoa(y)
			if len(xStr) == 1 {
				xStr = "0" + xStr
			}
			if len(yStr) == 1 {
				yStr = "0" + yStr
			}
			var parent *neighbor
			if _, found := positionNeighbor[xStr+yStr]; found {
				parent = positionNeighbor[xStr+yStr]
			} else {
				parent = &neighbor{
					x:     x,
					y:     y,
					value: valueInt,
				}

				positionNeighbor[xStr+yStr] = parent
			}

			// top left
			if x == 0 && y == 0 {
				parent = buildNeighbor(o, parent, x+1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y+1, positionNeighbor)

				parent = buildNeighbor(o, parent, x+1, y+1, positionNeighbor)
			}
			// top right
			if x == len(o[y])-1 && y == 0 {
				parent = buildNeighbor(o, parent, x-1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y+1, positionNeighbor)

				parent = buildNeighbor(o, parent, x-1, y+1, positionNeighbor)

			}
			// bottom right
			if x == len(o[y])-1 && y == len(o)-1 {
				parent = buildNeighbor(o, parent, x-1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y-1, positionNeighbor)

				parent = buildNeighbor(o, parent, x-1, y-1, positionNeighbor)

			}
			// bottom left
			if x == 0 && y == len(o)-1 {
				parent = buildNeighbor(o, parent, x+1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y-1, positionNeighbor)

				parent = buildNeighbor(o, parent, x+1, y-1, positionNeighbor)
			}
			// border left
			if x == 0 && y > 0 && y < len(o)-1 {
				parent = buildNeighbor(o, parent, x+1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y-1, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y+1, positionNeighbor)

				parent = buildNeighbor(o, parent, x+1, y+1, positionNeighbor)
				parent = buildNeighbor(o, parent, x+1, y-1, positionNeighbor)

			}
			// border top
			if x > 0 && y == 0 && x < len(o[y])-1 {
				parent = buildNeighbor(o, parent, x+1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y+1, positionNeighbor)

				parent = buildNeighbor(o, parent, x+1, y+1, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y+1, positionNeighbor)

			}
			// border right
			if y > 0 && x == len(o[y])-1 && y < len(o)-1 {
				parent = buildNeighbor(o, parent, x, y-1, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y+1, positionNeighbor)

				parent = buildNeighbor(o, parent, x-1, y+1, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y-1, positionNeighbor)
			}
			// border bottom
			if x > 0 && y > 0 && x < len(o[y])-1 && y == len(o)-1 {
				parent = buildNeighbor(o, parent, x, y-1, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x+1, y, positionNeighbor)

				parent = buildNeighbor(o, parent, x+1, y-1, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y-1, positionNeighbor)
			}
			// classic
			if x > 0 && y > 0 && x < len(o[y])-1 && y < len(o)-1 {
				parent = buildNeighbor(o, parent, x, y-1, positionNeighbor)
				parent = buildNeighbor(o, parent, x, y+1, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y, positionNeighbor)
				parent = buildNeighbor(o, parent, x+1, y, positionNeighbor)

				parent = buildNeighbor(o, parent, x+1, y+1, positionNeighbor)
				parent = buildNeighbor(o, parent, x+1, y-1, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y+1, positionNeighbor)
				parent = buildNeighbor(o, parent, x-1, y-1, positionNeighbor)
			}

		}
	}

	return positionNeighbor
}
func flash(n *neighbor, flashes *int, flashing bool, flashed map[string]*neighbor, position map[string]*neighbor) {
	xStr := strconv.Itoa(n.x)
	yStr := strconv.Itoa(n.y)
	if len(xStr) == 1 {
		xStr = "0" + xStr
	}
	if len(yStr) == 1 {
		yStr = "0" + yStr
	}
	_, isFlashed := flashed[xStr+yStr]
	if !isFlashed {
		if n.value == 9 {
			n.value = 0
			*flashes++
			flashed[xStr+yStr] = n
			for _, v := range n.neighbor {
				flash(v, flashes, true, flashed, position)
			}
		} else if n.value == 0 {
			n.value++

		} else {
			n.value++
		}

	}

	fmt.Println("after value", n.value, "from x=", n.x, "y=", n.y)
}
func bfs(position map[string]*neighbor, flashes *int) {
	var queue []*neighbor
	root := position["0000"]
	queue = append(queue, root)
	visited := make(map[string]*neighbor)
	flashed := make(map[string]*neighbor)
	for len(queue) > 0 {
		v := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		xStr := strconv.Itoa(v.x)
		yStr := strconv.Itoa(v.y)
		if len(xStr) == 1 {
			xStr = "0" + xStr
		}
		if len(yStr) == 1 {
			yStr = "0" + yStr
		}
		if _, found := visited[xStr+yStr]; !found {
			flash(v, flashes, false, flashed, position)

			queue = append(queue, v.neighbor...)
			visited[xStr+yStr] = v
		}
	}
}
func part1(o []string) {
	var position [][]string
	for _, v := range o {
		tab := strings.Split(v, "")
		position = append(position, tab)
	}
	var flashes int = 0
	tree := buildTree(position)
	all := false
	step := 0
	for all == false {

		bfs(tree, &flashes)
		all = true
		fmt.Println("analysis")
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				xStr := strconv.Itoa(x)
				yStr := strconv.Itoa(y)
				if len(xStr) == 1 {
					xStr = "0" + xStr
				}
				if len(yStr) == 1 {
					yStr = "0" + yStr
				}
				fmt.Print(tree[xStr+yStr].value)
				if tree[xStr+yStr].value != 0 {
					all = false
				}

			}
			fmt.Println("")
		}
		if all {
			fmt.Println("break !!!!!!!!!!!!!!!!", step)
		}
		step++

	}
	fmt.Println("flashes", flashes)

}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	var octopus []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		currentValue := scanner.Text()
		octopus = append(octopus, currentValue)

	}
	if err != nil {
		fmt.Println("error", err)
	}
	part1(octopus)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
