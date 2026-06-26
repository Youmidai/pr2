package main

import (
	"fmt"
	"strings"
)

func RomanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	s = strings.ToUpper(s)

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}

	return result
}

func main() {
	fmt.Println("III =", RomanToInt("III"))
	fmt.Println("IX =", RomanToInt("IX"))
	fmt.Println("LVIII =", RomanToInt("LViii"))
	fmt.Println("MCMXCIV =", RomanToInt("McMxCiV"))
}
