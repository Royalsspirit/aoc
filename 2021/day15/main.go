package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**************************************************************
*
*         improvement: do not use str as key of ceil
*                  -> create a struct with x and y keys
*
*****************************************************************
**/
type ceil struct {
	name     string
	x        int
	y        int
	neighbor []*ceil
}
type queue struct {
	ceil     *ceil
	distance int
}

type priorityQueue []*queue

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	node := x.(*queue)
	*pq = append(*pq, node)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return node
}

/**
* manage the best distance between nodes
* 116
* 138
* 213
* source = (0,0), 1
* put to queue neighbor = (1,0) 1 and (0,1) 1 with distance 1
*      loop over (1,0) 1 neighbor = (1,1), 3 and (2, 0) 6 with distance 2
*      loop over (0,1) 1 neighbor = (0,2) 2 and (2,0) 3 with distance 2 ===> (lowest path)
*

* try to manage the visited state
**/
func dijkstra(ceils map[string]*ceil, targetX int, targetY int) {
	targetxStr := strconv.Itoa(targetX)
	targetyStr := strconv.Itoa(targetY)
	fmt.Println("target", targetxStr+targetyStr)
	c := ceils["000000"]
	cc := &queue{ceil: c, distance: 0}
	fmt.Println("c", &c)
	distances := make(map[string]int)
	visited := make(map[string]int)
	prev := make(map[string]*ceil)
	distances["000000"] = 0
	/*
		* implemented a priority queue
		* more faster than a manual priority queue:
			var current *ceil
			newQueue := queue

			var minDist = 999
			// manual priority queue which find the ceil with the lowest risk
			// for k, v := range distances {
			// 	for _, q := range newQueue {
			// 		if q == ceils[k] && minDist > v {
			// 			minDist = v
			// 			current = ceils[k]
			// 			break
			// 		}

			// 	}
			// }
			// for i, q := range newQueue {
			// 	if q == current {
			// 		queue = append(queue[:i], queue[i+1:]...)
			// 		break
			// 	}

			// }

	*/
	pq := priorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, cc)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*queue)
		current := cur.ceil
		xStr := strconv.Itoa(current.x)
		yStr := strconv.Itoa(current.y)

		if len(xStr) == 1 {
			xStr = "0" + xStr
		}
		if len(yStr) == 1 {
			yStr = "0" + yStr
		}
		if len(xStr) == 2 {
			xStr = "0" + xStr
		}
		if len(yStr) == 2 {
			yStr = "0" + yStr
		}

		currentDist := distances[xStr+yStr]
		if current.x == targetX && current.y == targetY {
			fmt.Println("distances", distances)
			total := 0
			for current.x != 0 {
				cxStr := strconv.Itoa(current.x)
				cyStr := strconv.Itoa(current.y)
				if len(cxStr) == 1 {
					cxStr = "0" + cxStr
				}
				if len(cyStr) == 1 {
					cyStr = "0" + cyStr
				}
				if len(cxStr) == 2 {
					cxStr = "0" + cxStr
				}
				if len(cyStr) == 2 {
					cyStr = "0" + cyStr
				}

				ctT, _ := strconv.Atoi(current.name)

				total += ctT
				current = prev[cxStr+cyStr]

				fmt.Println("current", current.x, "y=", current.y)
			}
			fmt.Println("total", total)
			break
		}

		if _, ok := visited[xStr+yStr]; !ok {
			visited[xStr+yStr] = 1
			for _, v := range current.neighbor {
				xxStr := strconv.Itoa(v.x)
				yyStr := strconv.Itoa(v.y)
				if len(xxStr) == 1 {
					xxStr = "0" + xxStr
				}
				if len(yyStr) == 1 {
					yyStr = "0" + yyStr
				}
				if len(xxStr) == 2 {
					xxStr = "0" + xxStr
				}
				if len(yyStr) == 2 {
					yyStr = "0" + yyStr
				}

				// loop over neighbor
				// find the lowest risk path
				// put the node to the queue

				cT, _ := strconv.Atoi(v.name)
				distTotal := cT + currentDist
				d := distances[xxStr+yyStr]
				fmt.Println("d", d)
				if distTotal < d || d == 0 {
					heap.Push(&pq, &queue{ceil: v, distance: distTotal})
					prev[xxStr+yyStr] = current
					distances[xxStr+yyStr] = distTotal
				}

			}
		}

	}

	fmt.Println("distances", distances[targetxStr+targetyStr])

}
func buildNeighbor(c *ceil, card [][]string, x int, y int, ceils map[string]*ceil) *ceil {

	xStr := strconv.Itoa(x)
	yStr := strconv.Itoa(y)

	if len(xStr) == 1 {
		xStr = "0" + xStr
	}
	if len(yStr) == 1 {
		yStr = "0" + yStr
	}
	if len(xStr) == 2 {
		xStr = "0" + xStr
	}
	if len(yStr) == 2 {
		yStr = "0" + yStr
	}
	var child *ceil
	if _, found := ceils[xStr+yStr]; found {
		child = ceils[xStr+yStr]
	} else {
		child = &ceil{
			name: card[y][x],
			x:    x,
			y:    y,
		}

		ceils[xStr+yStr] = child
	}
	c.neighbor = append(c.neighbor, child)
	return c
}
func part1(card [][]string) {
	ceils := make(map[string]*ceil)
	targetX := 0
	targetY := 0
	for y := 0; y < len(card); y++ {
		targetX = len(card[y]) - 1
		targetY = len(card) - 1
		for x := 0; x < len(card[y]); x++ {
			xStr := strconv.Itoa(x)
			yStr := strconv.Itoa(y)
			if len(xStr) == 1 {
				xStr = "0" + xStr
			}
			if len(yStr) == 1 {
				yStr = "0" + yStr
			}
			if len(xStr) == 2 {
				xStr = "0" + xStr
			}
			if len(yStr) == 2 {
				yStr = "0" + yStr
			}
			var c *ceil
			if _, found := ceils[xStr+yStr]; found {
				c = ceils[xStr+yStr]
			} else {
				c = &ceil{
					name:     card[y][x],
					x:        x,
					y:        y,
					neighbor: []*ceil{},
				}
				ceils[xStr+yStr] = c
			}
			/******************************************************************
						* following lines can be replaces by
			*******************************************************************
						* if y != 0 {
					        *       c = buildNeighbor(c, card, x, y-1, ceils)
						* }
						* if x != 0 {
				                *       c = buildNeighbor(c, card, x-1, y, ceils)
						* }
						* if x < len(row)-1 {
				                *       c = buildNeighbor(c, card, x+1, y, ceils)
						* }
						* if y < len(fullGrid)-1 {
				                *       c = buildNeighbor(c, card, x, y+1, ceils)
						* }
						*
						*
			*/
			// top left
			if x == 0 && y == 0 {
				c = buildNeighbor(c, card, x+1, y, ceils)
				c = buildNeighbor(c, card, x, y+1, ceils)
			}
			// top right
			if x == len(card[y])-1 && y == 0 {
				c = buildNeighbor(c, card, x-1, y, ceils)
				c = buildNeighbor(c, card, x, y+1, ceils)

			}
			// bottom right
			if x == len(card[y])-1 && y == len(card)-1 {
				c = buildNeighbor(c, card, x-1, y, ceils)
				c = buildNeighbor(c, card, x, y-1, ceils)

			}
			// bottom left
			if x == 0 && y == len(card)-1 {
				c = buildNeighbor(c, card, x+1, y, ceils)
				c = buildNeighbor(c, card, x, y-1, ceils)
			}
			// border left
			if x == 0 && y > 0 && y < len(card)-1 {
				c = buildNeighbor(c, card, x+1, y, ceils)
				c = buildNeighbor(c, card, x, y-1, ceils)
				c = buildNeighbor(c, card, x, y+1, ceils)
			}
			// border top
			if x > 0 && y == 0 && x < len(card[y])-1 {
				c = buildNeighbor(c, card, x+1, y, ceils)
				c = buildNeighbor(c, card, x-1, y, ceils)
				c = buildNeighbor(c, card, x, y+1, ceils)
			}
			// border right
			if y > 0 && x == len(card[y])-1 && y < len(card)-1 {
				c = buildNeighbor(c, card, x, y-1, ceils)
				c = buildNeighbor(c, card, x-1, y, ceils)
				c = buildNeighbor(c, card, x, y+1, ceils)
			}
			// border bottom
			if x > 0 && y > 0 && x < len(card[y])-1 && y == len(card)-1 {
				c = buildNeighbor(c, card, x, y-1, ceils)
				c = buildNeighbor(c, card, x-1, y, ceils)
				c = buildNeighbor(c, card, x+1, y, ceils)
			}
			// classic
			if x > 0 && y > 0 && x < len(card[y])-1 && y < len(card)-1 {
				c = buildNeighbor(c, card, x, y-1, ceils)
				c = buildNeighbor(c, card, x, y+1, ceils)
				c = buildNeighbor(c, card, x-1, y, ceils)
				c = buildNeighbor(c, card, x+1, y, ceils)
			}

		}

	}
	dijkstra(ceils, targetX, targetY)

}
func part2(card [][]string) {
	cCard := card
	fmt.Println("len", len(cCard))
	fmt.Println("len", len(cCard[0]))

	ceils := make(map[string]*ceil)
	lenY := len(cCard)
	lenX := len(cCard[0])
	for z := 1; z < 5; z++ {
		for i := 0; i < lenY; i++ {
			for j := 0; j < lenX; j++ {
				originValue, _ := strconv.Atoi(cCard[i][j])
				newValue := originValue + z
				if newValue > 9 {
					newValue = newValue % 9
				}
				newValueStr := strconv.Itoa(newValue)
				if j == 0 {
					cCard = append(cCard, []string{newValueStr})
				} else {
					cCard[(lenY*z)+i] = append(cCard[(lenY*z)+i], newValueStr)
				}

			}

		}
	}
	lenY = len(cCard)
	lenX = len(cCard[0])

	for z := 1; z < 5; z++ {
		for i := 0; i < lenY; i++ {
			for j := 0; j < lenX; j++ {
				originValue, _ := strconv.Atoi(cCard[i][j])
				newValue := originValue + z
				if newValue > 9 {
					newValue = newValue % 9
				}
				newValueStr := strconv.Itoa(newValue)
				cCard[i] = append(cCard[i], newValueStr)

			}

		}
	}
	/**
	[1 1 6 3 7 5 1 7 4 2 2 2 7 4 8 6 2 8 5 3 3 3 8 5 9 7 3 9 6 4 4 4 9 6 1 8 4 1 7 5 5 5 1 7 2 9 5 2 8 6]
	[1 3 8 1 3 7 3 6 7 2 2 4 9 2 4 8 4 7 8 3 3 5 1 3 5 9 5 8 9 4 4 6 2 4 6 1 6 9 1 5 5 7 3 5 7 2 7 1 2 6]
	[2 1 3 6 5 1 1 3 2 8 3 2 4 7 6 2 2 4 3 9 4 3 5 8 7 3 3 5 4 1 5 4 6 9 8 4 4 6 5 2 6 5 7 1 9 5 5 7 6 3]
	[3 6 9 4 9 3 1 5 6 9 4 7 1 5 1 4 2 6 7 1 5 8 2 6 2 5 3 7 8 2 6 9 3 7 3 6 4 8 9 3 7 1 4 8 4 7 5 9 1 4]
	[7 4 6 3 4 1 7 1 1 1 8 5 7 4 5 2 8 2 2 2 9 6 8 5 6 3 9 3 3 3 1 7 9 6 7 4 1 4 4 4 2 8 1 7 8 5 2 5 5 5]
	[1 3 1 9 1 2 8 1 3 7 2 4 2 1 2 3 9 2 4 8 3 5 3 2 3 4 1 3 5 9 4 6 4 3 4 5 2 4 6 1 5 7 5 4 5 6 3 5 7 2]
	[1 3 5 9 9 1 2 4 2 1 2 4 6 1 1 2 3 5 3 2 3 5 7 2 2 3 4 6 4 3 4 6 8 3 3 4 5 7 5 4 5 7 9 4 4 5 6 8 6 5]
	[3 1 2 5 4 2 1 6 3 9 4 2 3 6 5 3 2 7 4 1 5 3 4 7 6 4 3 8 5 2 6 4 5 8 7 5 4 9 6 3 7 5 6 9 8 6 5 1 7 4]
	[1 2 9 3 1 3 8 5 2 1 2 3 1 4 2 4 9 6 3 2 3 4 2 5 3 5 1 7 4 3 4 5 3 6 4 6 2 8 5 4 5 6 4 7 5 7 3 9 6 5]
	[2 3 1 1 9 4 4 5 8 1 3 4 2 2 1 5 5 6 9 2 4 5 3 3 2 6 6 7 1 3 5 6 4 4 3 7 7 8 2 4 6 7 5 5 4 8 8 9 3 5]
	[2 2 7 4 8 6 2 8 5 3]
	[2 4 9 2 4 8 4 7 8 3]
	[3 2 4 7 6 2 2 4 3 9]
	[4 7 1 5 1 4 2 6 7 1]
	[8 5 7 4 5 2 8 2 2 2]
	[2 4 2 1 2 3 9 2 4 8]
	[2 4 6 1 1 2 3 5 3 2]
	[4 2 3 6 5 3 2 7 4 1]
	[2 3 1 4 2 4 9 6 3 2]
	[3 4 2 2 1 5 5 6 9 2]
	[3 3 8 5 9 7 3 9 6 4]
	[3 5 1 3 5 9 5 8 9 4]
	[4 3 5 8 7 3 3 5 4 1]
	[5 8 2 6 2 5 3 7 8 2]
	[9 6 8 5 6 3 9 3 3 3]
	[3 5 3 2 3 4 1 3 5 9]
	[3 5 7 2 2 3 4 6 4 3]
	[5 3 4 7 6 4 3 8 5 2]
	[3 4 2 5 3 5 1 7 4 3]
	[4 5 3 3 2 6 6 7 1 3]
	[4 4 9 6 1 8 4 1 7 5]
	[4 6 2 4 6 1 6 9 1 5]
	[5 4 6 9 8 4 4 6 5 2]
	[6 9 3 7 3 6 4 8 9 3]
	[1 7 9 6 7 4 1 4 4 4]
	[4 6 4 3 4 5 2 4 6 1]
	[4 6 8 3 3 4 5 7 5 4]
	[6 4 5 8 7 5 4 9 6 3]
	[4 5 3 6 4 6 2 8 5 4]
	[5 6 4 4 3 7 7 8 2 4]
	[5 5 1 7 2 9 5 2 8 6]
	[5 7 3 5 7 2 7 1 2 6]
	[6 5 7 1 9 5 5 7 6 3]
	[7 1 4 8 4 7 5 9 1 4]
	[2 8 1 7 8 5 2 5 5 5]
	[5 7 5 4 5 6 3 5 7 2]
	[5 7 9 4 4 5 6 8 6 5]
	[7 5 6 9 8 6 5 1 7 4]
	[5 6 4 7 5 7 3 9 6 5]
	[6 7 5 5 4 8 8 9 3 5]		**/

	fmt.Println("ccard", cCard)
	for y := 0; y < len(cCard); y++ {
		for x := 0; x < len(cCard[y]); x++ {
			xStr := strconv.Itoa(x)
			yStr := strconv.Itoa(y)
			if len(xStr) == 1 {
				xStr = "0" + xStr
			}
			if len(yStr) == 1 {
				yStr = "0" + yStr
			}
			if len(xStr) == 2 {
				xStr = "0" + xStr
			}
			if len(yStr) == 2 {
				yStr = "0" + yStr
			}
			var c *ceil
			if _, found := ceils[xStr+yStr]; found {
				c = ceils[xStr+yStr]
			} else {
				c = &ceil{
					name:     cCard[y][x],
					x:        x,
					y:        y,
					neighbor: []*ceil{},
				}
				ceils[xStr+yStr] = c
			}
			// top left
			if x == 0 && y == 0 {
				c = buildNeighbor(c, cCard, x+1, y, ceils)
				c = buildNeighbor(c, cCard, x, y+1, ceils)
			}
			// top right
			if x == len(cCard[y])-1 && y == 0 {
				c = buildNeighbor(c, cCard, x-1, y, ceils)
				c = buildNeighbor(c, cCard, x, y+1, ceils)

			}
			// bottom right
			if x == len(cCard[y])-1 && y == len(cCard)-1 {
				c = buildNeighbor(c, cCard, x-1, y, ceils)
				c = buildNeighbor(c, cCard, x, y-1, ceils)

			}
			// bottom left
			if x == 0 && y == len(cCard)-1 {
				c = buildNeighbor(c, cCard, x+1, y, ceils)
				c = buildNeighbor(c, cCard, x, y-1, ceils)
			}
			// border left
			if x == 0 && y > 0 && y < len(cCard)-1 {
				c = buildNeighbor(c, cCard, x+1, y, ceils)
				c = buildNeighbor(c, cCard, x, y-1, ceils)
				c = buildNeighbor(c, cCard, x, y+1, ceils)
			}
			// border top
			if x > 0 && y == 0 && x < len(cCard[y])-1 {
				c = buildNeighbor(c, cCard, x+1, y, ceils)
				c = buildNeighbor(c, cCard, x-1, y, ceils)
				c = buildNeighbor(c, cCard, x, y+1, ceils)
			}
			// border right
			if y > 0 && x == len(cCard[y])-1 && y < len(cCard)-1 {
				c = buildNeighbor(c, cCard, x, y-1, ceils)
				c = buildNeighbor(c, cCard, x-1, y, ceils)
				c = buildNeighbor(c, cCard, x, y+1, ceils)
			}
			// border bottom
			if x > 0 && y > 0 && x < len(cCard[y])-1 && y == len(cCard)-1 {
				c = buildNeighbor(c, cCard, x, y-1, ceils)
				c = buildNeighbor(c, cCard, x-1, y, ceils)
				c = buildNeighbor(c, cCard, x+1, y, ceils)
			}
			// classic
			if x > 0 && y > 0 && x < len(cCard[y])-1 && y < len(cCard)-1 {
				c = buildNeighbor(c, cCard, x, y-1, ceils)
				c = buildNeighbor(c, cCard, x, y+1, ceils)
				c = buildNeighbor(c, cCard, x-1, y, ceils)
				c = buildNeighbor(c, cCard, x+1, y, ceils)
			}

		}

	}
	lenY = len(cCard)
	lenX = len(cCard[0])
	fmt.Println("lenY", lenY, "lenX", lenX)
	dijkstra(ceils, lenX-1, lenY-1)

}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var card [][]string
	for scanner.Scan() {

		currentValue := scanner.Text()
		card = append(card, strings.Split(currentValue, ""))

	}
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("card", card)
	part2(card)
	//	part2(template, pairs)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
