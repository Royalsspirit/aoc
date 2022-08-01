package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type neighbor struct {
	cave     string
	large    bool
	neighbor []*neighbor
}

/**
* find a way to check if the path is already visited
* store the current cave in a array to build the path
* if end clean the current path and store it the global paths
*
 */
func toTheEnd(n *neighbor, path []string, revisitedCave bool) [][]string {
	// the end of the path
	// return the final path
	if n.cave == "end" {
		path = append(path, n.cave)
		return [][]string{path}
	}
	// small cave can be visited only once
	if n.large == false {
		for _, v := range path {
			if v == n.cave {
				// return an empty double dimensional array to
				// indicate that is not a valid path
				// it will be removed by the spread

				// revisited part 2
				if revisitedCave {
					return [][]string{}
				}
				// start can be visited only once part 2
				if n.cave == "start" {
					return [][]string{}
				}
				revisitedCave = true

			}

		}
	}
	var paths [][]string

	path = append(path, n.cave)
	for i := 0; i < len(n.neighbor); i++ {
		newPath := append([]string{}, path...)
		paths = append(paths, toTheEnd(n.neighbor[i], newPath, revisitedCave)...)
	}
	return paths
}
func dfs(tree map[string]*neighbor, cave string, nbrPath *int, visited map[string]int, path string) {
}
func findNeighbor(from string, to string, paths map[string]*neighbor) *neighbor {
	var n *neighbor
	if _, foundFirst := paths[from]; !foundFirst {
		n = &neighbor{
			cave: from,
		}
		if from == strings.ToUpper(from) {
			n.large = true
		}
		paths[from] = n
	} else {
		n = paths[from]

	}
	var n2 *neighbor
	if _, foundUnder := paths[to]; !foundUnder {
		n2 = &neighbor{
			cave: to,
		}
		if to == strings.ToUpper(to) {
			n2.large = true
		}
		paths[to] = n2
	} else {
		n2 = paths[to]
	}
	n.neighbor = append(n.neighbor, n2)

	return n
}

/**
* a - b
* b - c
* d - e
* e - b
*
*
 */
func buildTree(caves []string) map[string]*neighbor {
	paths := make(map[string]*neighbor)
	for _, v := range caves {
		cave := strings.Split(v, "-")
		paths[cave[0]] = findNeighbor(cave[0], cave[1], paths)
		paths[cave[1]] = findNeighbor(cave[1], cave[0], paths)

	}
	return paths
}
func part1(caves []string) {

	tree := buildTree(caves)
	n := tree["start"]
	result := toTheEnd(n, []string{}, false)
	fmt.Println("result", len(result))

}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	var caves []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		currentValue := scanner.Text()
		caves = append(caves, currentValue)

	}
	if err != nil {
		fmt.Println("error", err)
	}
	part1(caves)
	//fmt.Println(firstPartResult)
	//	fmt.Println(part2(rowData))
}
