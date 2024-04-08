package main

import (
	"fmt"
	"strings"
)

// Work with strings, string builder and runes
func stringRunes() {
	fmt.Println("Rune test")
	var myRune = 'a'
	fmt.Printf("%v %T\n", myRune, myRune)
	fmt.Println()
	fmt.Println("String vs runes")
	var text1 string = "résumé"
	var text2 []rune = []rune(text1)
	fmt.Println("--With string:")
	fmt.Println(text1)
	for _, v := range text1 {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
	fmt.Println("length:", len(text1)) // Count bytes
	fmt.Printf("index [1]: %v %T\n", text1[1], text1[1]) // First byte of 'é'
	fmt.Println("--With runes:")
	fmt.Println(text2)
	for _, v := range text2 {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
	fmt.Println("length:", len(text2)) // Count encoded UTF-8 chars
	fmt.Printf("index [1]: %v %T\n", text2[1], text2[1]) // All bytes of 'é'
	fmt.Println()

	// String concatenation
	fmt.Println("String concatenation")
	letters := []string{"a", "b", "c"} 
	var string1 = ""
	for i := range letters {
		string1 += letters[i]
	}
	fmt.Println(string1)
	fmt.Println()
	
	// Using string builder
	fmt.Println("String builder")
	moreLetters := []string{"a", "b", "c"}
	var builder strings.Builder 
	for j := range moreLetters {
		builder.WriteString(moreLetters[j])
	}
	var string2 = builder.String()
	fmt.Println(string2)
	fmt.Println()
}
