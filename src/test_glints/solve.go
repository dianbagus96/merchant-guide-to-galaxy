package main

import (
	"fmt"
	"io/ioutil"
	"roman"
	"strconv"
	"strings"
)

type isinilai struct {
	value string
}

var people = map[string]*isinilai{}

func main() {
	fmt.Println("======================================")
	fmt.Println("Welcome to Merchant Guide To Galaxy")
	fmt.Println("======================================")

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	fmt.Println("=================")
	fmt.Println("Question is")
	fmt.Println("=================")

	hasilSplit := strings.Split(str, "\r")
	fmt.Println(str)

	fmt.Println("=========================")
	fmt.Println("The Value of Variable is")
	fmt.Println("=========================")

	//membaca per line
	for _, item := range hasilSplit {
		//membaca per sentence
		if !strings.Contains(item, "how many") {
			if !strings.Contains(item, "how much") {
				//mencari nilai
				if strings.Contains(item, "is") && !strings.Contains(item, "Credits") {
					//nilai var
					value := strings.Split(item, " ")
					r := strings.TrimSpace(string(value[0]))
					people[r] = &isinilai{value: value[2]}
				}
			}
		}
	}

	for _, item := range hasilSplit {
		//membaca per sentence
		if !strings.Contains(item, "how many") {
			if !strings.Contains(item, "how much") {
				//mencari nilai
				if strings.Contains(item, "is") && strings.Contains(item, "Credits") {
					//nilai credits
					value := strings.Split(item, " ")
					var1 := strings.TrimSpace(string(value[0]))
					var2 := strings.TrimSpace(string(value[1]))
					nilai := people[var1].value + people[var2].value
					s := roman.Arabic(nilai)
					credits, _ := strconv.ParseFloat(value[4], 64)
					semua := credits / float64(s)
					h := fmt.Sprintf("%f", semua)
					people[value[2]] = &isinilai{value: h}
				}

			}
		}

	}
	for key, items := range people {
		fmt.Println(key, items.value)
	}
	fmt.Println("=============")
	fmt.Println("The Answer is")
	fmt.Println("=============")

	for _, item := range hasilSplit {
		if strings.Contains(item, "is") && (strings.Contains(item, "how much") || strings.Contains(item, "how many")) {
			//game dimulai :D
			var nilai string
			if strings.Contains(item, "how much") {
				value := strings.Split(item, " ")
				for i := 3; i <= len(value)-2; i++ {
					nilai = nilai + people[value[i]].value
				}
				fmt.Print(item + " ")
				fmt.Println(roman.Arabic(nilai))

			} else if strings.Contains(item, "how many") {
				value := strings.Split(item, " ")
				er := len(value)
				for i := 4; i <= len(value)-3; i++ {
					nilai = nilai + people[value[i]].value
				}
				hsl, _ := strconv.ParseFloat(people[value[er-2]].value, 64)
				fmt.Print(item + " ")
				fmt.Println(float64(roman.Arabic(nilai)) * hsl)
			}
		} else {
			if !strings.Contains(item, "is") && (strings.Contains(item, "how much") || strings.Contains(item, "how many")) {
				fmt.Println("I have no idea what you are talking about")
			}
		}
	}

}
