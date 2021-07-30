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
	color string
	count int
}

/**
* part 1
**/
func countGold(currentColor string, dic map[string][]bag) int {
	var tmp int = 0
	//	fmt.Println("loop over current color map", currentColor, "which contain", dic[currentColor])
	if len(dic[currentColor]) > 0 {
		for _, v := range dic[currentColor] {
			if v.color == "shiny gold" {
				tmp += v.count
				//				fmt.Println("found !!")
			} else {
				bagUnderCurrentBag := countGold(v.color, dic)
				//				fmt.Println("bag of current color", currentColor, "have this color", v.color, "which contain", bagUnderCurrentBag, "other bags")
				result := v.count * bagUnderCurrentBag
				tmp += result
			}

		}
	}
	return tmp
}

/**
part 2
**/
func countBags(currentColor string, dic map[string][]bag) int {
	var tmp int = 0
	if len(dic[currentColor]) > 0 {
		for _, v := range dic[currentColor] {
			fmt.Println("in dic[currentcolor]", dic[currentColor], "current color is", currentColor)
			bagUnderCurrentBag := countBags(v.color, dic)
			tmp += v.count
			tmp += v.count * bagUnderCurrentBag
		}
	}
	fmt.Println("result tmp", tmp)
	return tmp

}
func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var rules []string
	dic := make(map[string][]bag)
	/**
	* move all line in rules array
	**/
	for scanner.Scan() {
		currentValue := scanner.Text()
		rules = append(rules, currentValue)
	}
	/*
			* each line contain a left part which is a bag container
			* and a right part which is children of left part
			*
		        * for each line we create a dictionnary with a left part color as key
			* and right part colors as value with some quantity
	*/
	re := regexp.MustCompile("[0-9]+")
	for _, v := range rules {
		/**
		* can probably be improved these splits statements by using regex
		**/
		bags := strings.Split(v, "contain ")
		leftColors := strings.Split(bags[0], " bags")
		var c bag
		c.color = leftColors[0]
		dic[c.color] = []bag{}
		/**
		* the if statement means that the right part contain several colors
		**/
		if strings.Index(bags[1], ",") != -1 {
			/*
			* same here, can be improved by using regex
			 */
			bagsHandler := strings.Replace(bags[1], ".", ", ", 1)
			customBags := strings.Replace(bagsHandler, "bag,", "bags,", -1)
			haveBags := strings.Split(customBags, " bags, ")
			haveBags = haveBags[:len(haveBags)-1]

			for _, v := range haveBags {
				numberOfBag := re.FindAllString(v, -1)
				catchNumber := numberOfBag[0] + " "
				cleanColor := strings.Replace(v, catchNumber, "", -1)
				intNumberOfBag, _ := strconv.Atoi(numberOfBag[0])

				dic[c.color] = append(dic[c.color], bag{
					color: cleanColor,
					count: intNumberOfBag,
				})
			}

		} else {
			oneBag := strings.Split(bags[1], " bag")

			numberOfBag := re.FindAllString(oneBag[0], -1)
			if len(numberOfBag) > 0 {
				/*
				* can be a function which clean right part
				 */
				catchNumber := numberOfBag[0] + " "
				cleanColor := strings.Replace(oneBag[0], catchNumber, "", -1)

				currentNumber, _ := strconv.Atoi(numberOfBag[0])
				dic[c.color] = append(dic[c.color], bag{
					color: cleanColor,
					count: currentNumber,
				})

			}

		}
	}
	//	fmt.Println("dic", dic)
	var final int = 0
	for k, _ := range dic {
		var golds int
		golds = countGold(k, dic)
		fmt.Println("current color", k, "have gods", golds)
		if golds > 0 {
			final++
		}
	}
	fmt.Println("len of dic", len(dic))
	fmt.Println("gold", final)
	fmt.Println("shiny", dic["shiny gold"])
	fmt.Println("number of bags", countBags("shiny gold", dic))
}
