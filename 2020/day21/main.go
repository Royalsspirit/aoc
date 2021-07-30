package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	/**
	* create a dict of allergens and their ingredients
	**/
	allergens := make(map[string]map[string]int)
	/**
	* store initial fields
	**/
	originAllergens := make(map[string][]string)

	r, _ := regexp.Compile(`\(contains ([a-z, ]+)\)`)
	for scanner.Scan() {
		currentValue := scanner.Text()
		matches := r.FindAll([]byte(currentValue), -1)
		rawCurrentMatch := string(matches[0])
		// get right side to get allergens
		currentAllergens := strings.Replace(rawCurrentMatch, "contains ", "", -1)
		currentAllergens = currentAllergens[1 : len(currentAllergens)-1]
		allergensAsArray := strings.Split(string(currentAllergens), ", ")
		// get left side to get ingredients
		ingredientsClean := strings.Replace(currentValue, string(rawCurrentMatch), "", -1)
		ingredientsArray := strings.Fields(ingredientsClean)
		// store inital fields
		originAllergens[currentAllergens] = append(originAllergens[currentAllergens], ingredientsArray...)
		for i := 0; i < len(allergensAsArray); i++ {
			// set state of allergens child. If it's empty, first equal true otherwise false
			first := false
			if allergens[allergensAsArray[i]] == nil {
				allergens[allergensAsArray[i]] = make(map[string]int)
				first = true
			}
			/**
			* can be optimize by adding in allergens map only common value
			**/
			for _, vv := range ingredientsArray {
				if !first {
					if _, found := allergens[allergensAsArray[i]][vv]; found {
						allergens[allergensAsArray[i]][vv] += 1
					}
				} else {

					allergens[allergensAsArray[i]][vv] += 1
				}

			}
		}

	}
	allergenDone := make(map[string]string)
	var allergenToIngredient []map[string]string
	var allergenKeys []string
	ingredientKey := make(map[string][]string)
	for k, _ := range allergens {
		allergenKeys = append(allergenKeys, k)
		for kk, _ := range allergens[k] {
			ingredientKey[k] = append(ingredientKey[k], kk)
		}
	}
	sort.Strings(allergenKeys)
	for i := 0; i < len(allergenKeys); i++ {
		allergenKey := allergenKeys[i]
		var maxField string
		var associatedAllergen string
		// try to find the case where one allergen has a unique ingredient
		var candidateCount, maxV int = 0, 0
		for _, ing := range ingredientKey[allergenKey] {
			vv := allergens[allergenKey][ing]
			if _, found := allergenDone[ing]; !found {
				if vv > maxV {
					maxV = vv
					maxField = ing
					associatedAllergen = allergenKey
					candidateCount = 1
				} else if vv == maxV {
					candidateCount++
				}
			}
		}
		// if unique ingredient is found. Start again from top of the list
		if candidateCount == 1 {
			allergenDone[maxField] = maxField
			allergenKeys = append(allergenKeys[:i], allergenKeys[i+1:]...)
			var currentAllergenDetails map[string]string = map[string]string{
				"ingredient": associatedAllergen,
				"allergen":   maxField,
			}
			allergenToIngredient = append(allergenToIngredient, currentAllergenDetails)
			// set -1 to start at 0 in next loop
			i = -1
		}
	}
	/**
	  	  get keys of origin allergens
	**/
	var originAllergensKeys []string
	for originKey, _ := range originAllergens {
		originAllergensKeys = append(originAllergensKeys, originKey)
	}
	var count int = 0
	for _, originKey := range originAllergensKeys {
		for _, vv := range originAllergens[originKey] {
			if _, found := allergenDone[vv]; !found {
				count++
			}
		}
	}

	sort.Slice(allergenToIngredient, func(i, j int) bool {
		ingA := allergenToIngredient[i]["ingredient"]
		ingB := allergenToIngredient[j]["ingredient"]
		result := ingA < ingB
		return result
	})
	var allergenString []string
	for _, v := range allergenToIngredient {
		allergenString = append(allergenString, v["allergen"])
	}
	fmt.Println("part 1", count)
	fmt.Println("part 2", strings.Join(allergenString, ","))
	//fmt.Println("final response part 1", numbersWithoutBraket)
	/**

		sorted allergen to ingredient [map[allergen:gpgrb ingredient:dairy] map[allergen:tjlz ingredient:eggs] map[allergen:gtjmd ingredient:fish] map[allergen:spbxz ingredient:nuts] map[allergen:pfdkkzp ingredient:peanuts] map[allergen:xcfpc ingredient:shellfish] map[allergen:txzv ingredient:soy] map[allergen:znqbr ingredient:wheat]]
		part 1 1913 allergens done gpgrb,tjlz,gtjmd,spbxz,pfdkkzp,xcfpc,txzv,znqbr

	part 1 1913 allergens done gpgrb,tjlz,gtjmd,xcfpc,spbxz,pfdkkzp,txzv,znqbr

	*/
	/*
		sorted [('dairy', ['gpgrb']), ('eggs', ['tjlz']), ('fish', ['gtjmd']), ('nuts', ['spbxz']), ('peanuts', ['pfdkkzp']), ('shellfish', ['xcfpc']), ('soy', ['txzv']), ('wheat', ['znqbr'])]
		Part Two: gpgrb,tjlz,gtjmd,spbxz,pfdkkzp,xcfpc,txzv,znqbr

	*/
	/**
		allergen map[
		dairy:map[gpgrb:10]
	        eggs:map[tjlz:6]
		fish:map[gtjmd:5]
		nuts:map[spbxz:10]
		peanuts:map[pfdkkzp:12]
		shellfish:map[xcfpc:11]
		soy:map[txzv:7]
		wheat:map[znqbr:9]
		]
		**/
}
