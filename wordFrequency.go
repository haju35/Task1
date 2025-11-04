package main

import(
	"fmt"
 "strings"
)

func wordFrequency(s string) map[string]int{
	s = strings.ToLower(s)

	var sentence string
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || ch == ' '{
			sentence += string(ch)
		}
	}

	words := strings.Fields(sentence)

	freq := make(map[string]int)
	for _, word := range words{
		freq[word]++
	}
	return freq
}

func main(){
	text := "what do you do ? what are you doing?"
	frequencies := wordFrequency(text)

	for word, count := range frequencies {
		fmt.Printf("%s: %d\n", word, count)
	}
}