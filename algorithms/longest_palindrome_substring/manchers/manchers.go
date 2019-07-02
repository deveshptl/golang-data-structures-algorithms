package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("\n-- Longest Palindrome Substring Mancher's Algorithm --")

	var str string
	var txt []string

	fmt.Print("\nEnter a text string: ")
	fmt.Scanf("%s", &str)

	txt = strings.Split(str, "")
	str = "^" + "#" + strings.Join(txt, "#") + "#" + "$"
	txt = strings.Split(str, "")

	fmt.Println("\nPalindrome\\s are:")
	p := Manchers(txt)
	for i := range p {
		fmt.Println("'"+i+"'", "of length", p[i])
	}
}

// Manchers performs mancher's algorithm which finds longest palindrome substring in a given string
func Manchers(txt []string) map[string]int {
	subStr := make([]int, len(txt))
	c := 0
	r := 0

	for i := 1; i < len(txt)-1; i++ {
		mirror := 2*c - i
		if i < r {
			subStr[i] = int(math.Min(float64(r-i), float64(subStr[mirror])))
		}

		for txt[i+(1+subStr[i])] == txt[i-(1+subStr[i])] {
			subStr[i]++
		}

		if i+subStr[i] > r {
			c = i
			r = i + subStr[i]
		}

	}

	return findPalindromes(txt, subStr)
}

func findPalindromes(txt []string, subStr []int) map[string]int {
	pMap := make(map[string]int)

	for i := 0; i < len(subStr); i++ {
		if subStr[i] == 0 {
			continue
		}
		j := i - 1
		k := i + 2
		jLen := i - subStr[i]
		kLen := i + subStr[i] + 1
		for j >= jLen && k <= kLen {
			temp := strings.Join(txt[j:k], "")
			pMap[temp] = 0
			j--
			k++
		}
	}

	for i := range pMap {
		temp := strings.Split(i, "#")
		str := strings.Join(temp, "")

		// delete first then insert new map value
		delete(pMap, i)
		pMap[str] = len(str)
	}

	return pMap
}

// Test strings
//
// ABABABA
// [BABAB:5 ABABABA:7 B:1 A:1 BAB:3 ABABA:5 ABA:3]
//
// ABBAABBA
// [A:1 BAAB:4 ABBA:4 AA:2 BBAABB:6 ABBAABBA:8 B:1 BB:2]
//
// FORGEEKSSKEEGFOR
// [SS:2 R:1 G:1 O:1 EE:2 EEKSSKEE:8 S:1 GEEKSSKEEG:10 KSSK:4 F:1 K:1 EKSSKE:6 E:1]
//
// ABCBABCBABCBA
// [BCBABCBABCB:11 BAB:3 B:1 A:1 BCB:3 C:1 ABCBA:5 BABCBAB:7 CBABC:5 BCBABCB:7 CBABCBABC:9 ABCBABCBABCBA:13 ABCBABCBA:9]
//
// ABAABA
// [ABA:3 ABAABA:6 B:1 BAAB:4 AA:2 A:1]
