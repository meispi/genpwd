package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var arr []string

func permute(i int, word string, temp string) {
	if i >= len(word) {
		//fmt.Println(temp)
		arr = append(arr, temp)
		return
	}

	c := word[i]
	if !unicode.IsLetter(rune(c)) {
		permute(i+1, word, temp+string(c))
	}

	if unicode.IsUpper(rune(c)) {
		permute(i+1, word, temp+string(c))
		c := unicode.ToLower(rune(c))
		permute(i+1, word, temp+string(c))
	} else {
		permute(i+1, word, temp+string(c))
		c := unicode.ToUpper(rune(c))
		permute(i+1, word, temp+string(c))
	}
}

func simnum(word string) {
	word = strings.Replace(word, "o", "0", -1)
	word = strings.Replace(word, "l", "1", -1)
	word = strings.Replace(word, "e", "3", -1)
	word = strings.Replace(word, "a", "4", -1)
	word = strings.Replace(word, "s", "5", -1)
	word = strings.Replace(word, "t", "7", -1)
	arr = append(arr, word)
}

func addnum() {
	var temp []string
	for i := 0; i < len(arr); i++ {
		temp = append(temp, arr[i]+fmt.Sprint("1"), arr[i]+"12", arr[i]+"123", arr[i]+"1234", arr[i]+"12345")
		for j := 2000; j <= 2021; j++ {
			temp = append(temp, arr[i]+fmt.Sprint(j))
		}
	}
	arr = append(arr, temp...)
}

func main() {
	var word string
	fmt.Scanln(&word)
	permute(0, word, "") // 2^n words
	simnum(word)         // 1 word
	addnum()             // 27*(2^n+1) words
	//fmt.Println(len(arr)) // 28*(2^n + 1) words
	file, err := os.Create(word + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wfile, err := os.OpenFile(word+".txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	for i := range arr {
		if _, err := wfile.WriteString(arr[i] + "\n"); err != nil {
			panic(err)
		}
	}
}
