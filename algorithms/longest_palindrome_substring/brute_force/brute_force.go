package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n-- Longest Palindrome Substring Brute Force Technique --")

	var str string
	var txt []string

	fmt.Print("\nEnter a text string: ")
	fmt.Scanf("%s", &str)

	txt = strings.Split(str, "")
	str = "#" + strings.Join(txt, "#") + "#"
	txt = strings.Split(str, "")

	fmt.Println("\nPalindrome\\s are:")
	p := BruteForce(txt)
	for i := range p {
		fmt.Println("'"+i+"'", "of length", p[i])
	}
}

// BruteForce finds Longest Palindrome Substring in a given string using Brute Force technique
func BruteForce(txt []string) map[string]int {

	pMap := make(map[string]int)

	for i := 0; i < len(txt); i++ {

		k := i + 1

		for j := i; j >= 0 && k <= len(txt); j-- {

			subStr := txt[j:k]
			flag := 0

			for p := 0; p < len(subStr); p++ {
				if subStr[p] != subStr[len(subStr)-1-p] {
					flag = 1
					break
				}
			}

			if flag == 1 {
				flag = 0
			} else {
				pMap[strings.Join(subStr, "")] = len(subStr)
			}
			k++

		}

	}

	for i := range pMap {
		temp := strings.Split(i, "#")
		str := strings.Join(temp, "")
		delete(pMap, i)
		pMap[str] = len(str)
	}

	// key "" generated for # as palindrome in hash map
	delete(pMap, "")

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
