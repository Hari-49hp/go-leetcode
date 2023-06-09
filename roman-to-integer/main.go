package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {

	arr := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if 1 <= len(s) && len(s) <= 15 {

		if len(s) == 1 {
			return arr[s]
		}

		result := 0
		var lastLetter string
		var lastInt int

		for _, j := range s {
			if string(j) == "V" && lastLetter == "I" {
				result = result - lastInt
				result += 4
			} else if string(j) == "L" && lastLetter == "X" {
				result = result - lastInt
				result += 40
			} else if string(j) == "D" && lastLetter == "C" {
				result = result - lastInt
				result += 400
			} else if string(j) == "X" && lastLetter == "I" {
				result = result - lastInt
				result += 9
			} else if string(j) == "M" && lastLetter == "C" {
				fmt.Println("minus--", result, "-", lastInt)
				result = result - lastInt
				result += 900
			} else if string(j) == "C" && lastLetter == "X" {
				result = result - lastInt
				result += 90
			} else {
				result += arr[string(j)]
			}

			lastLetter = string(j)
			lastInt = arr[string(j)]

		}

		return result

	}

	return 0
}
