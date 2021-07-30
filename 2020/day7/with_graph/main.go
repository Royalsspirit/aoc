package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	color   string
	contain []containance
}
type containance struct {
	number int
	bag    []*bag
}

func lookUp(color string, bags []*bag) *bag {
	for _, v := range bags {
		//		fmt.Println("bag color", v.color, "len", len(v.color), "current color", color, "len", len(color))
		if v.color == color {
			return v
		}
	}
	return nil
}

func displayGraph(bags []*bag) {
	for _, v := range bags {
		if v.contain != nil {
			for _, vv := range v.contain {

				fmt.Println("color", v.color, "contain", vv.number, "of", vv.bag)
			}
		}
	}
}

func findGold(bags []*bag) {
	for _, v := range bags {
		if v.contain != nil && v.color != "gold" {
			for _, vv := range v.contain {
				if vv.bag != nil {
					findGold(vv.bag)
				}
			}
		}
	}
}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var rules []string
	var dic []*bag
	for scanner.Scan() {
		currentValue := scanner.Text()
		rules = append(rules, currentValue)
	}
	for _, v := range rules {
		bags := strings.Split(v, "contain")
		leftColors := strings.Split(bags[0], " bags")
		var c bag
		c.color = leftColors[0]
		dic = append(dic, &c)

	}
	var index int = 0

	re := regexp.MustCompile("[0-9]+")
	for _, v := range rules {
		bags := strings.Split(v, " contain ")
		if strings.Index(bags[1], ",") != -1 {
			bagsHandler := strings.Replace(bags[1], ".", ", ", 1)
			customBags := strings.Replace(bagsHandler, "bag,", "bags,", -1)
			haveBags := strings.Split(customBags, " bags, ")
			haveBags = haveBags[:len(haveBags)-1]

			for _, v := range haveBags {
				var bags []*bag
				numberOfBag := re.FindAllString(v, -1)
				catchNumber := numberOfBag[0] + " "
				cleanColor := strings.Replace(v, catchNumber, "", -1)
				lookUpBag := lookUp(cleanColor, dic)
				if lookUpBag != nil {
					currentNumber, _ := strconv.Atoi(numberOfBag[0])
					bags = append(bags, lookUpBag)
					contain := containance{
						number: currentNumber,
						bag:    bags,
					}
					dic[index].contain = append(dic[index].contain, contain)
				}
			}

		} else {
			oneBag := strings.Split(v, "bags")
			fmt.Println("looking for one bag", oneBag[0])
			lookUpBag := lookUp(oneBag[0], dic)
			if lookUpBag != nil {
				var bags []*bag
				numberOfBag := re.FindAllString(v, -1)
				currentNumber, _ := strconv.Atoi(numberOfBag[0])
				bags = append(bags, lookUpBag)
				contain := containance{
					number: currentNumber,
					bag:    bags,
				}
				dic[index].contain = append(dic[index].contain, contain)
			}

		}
		index++
	}
	for _, v := range dic {
		fmt.Printf("%p\n", v)
		fmt.Println("v color", v.color)
		fmt.Println("v container", v.contain)
	}
}
