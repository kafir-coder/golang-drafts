package main

import (
	"fmt"
	"strings"
)

func main() {
	plaintext := "mocada"
	keyword := "bela"
	s := encode(strings.ToUpper(plaintext) ,strings.ToUpper(keyword))
	a := decode(s, keyword)
	fmt.Println("encoded text =>", s)
	fmt.Println("decoded text =>", a)
}

func transform_keyword(keyword string, text_size int) string {
	keyword_size := len(keyword)
	if len(keyword) > text_size {
		panic("keyword shouldn't be greater than text")
	}
	difference := text_size-keyword_size
	counter := 0
	for i := 0; i < int(difference); i++ {
		if counter == keyword_size {
			counter = 0
		}
		join := []string {keyword, string(keyword[counter])}
		keyword = strings.Join(join, "");
		counter++
	}
	return keyword
}

func encode(plaintext string, keyword string) string {
	has_same_size := len(plaintext) == len(keyword)
	if !has_same_size  {
		keyword = transform_keyword(keyword, len(plaintext))
	}
	counter := 0
	encoded_text := strings.Map(func(r rune) rune {
		hop := character_position(r)
		encryted_rune := retrieve_encoded_rune(int(plaintext[counter]), hop)
		counter++
		return rune(encryted_rune)	
	}, keyword)
	return encoded_text
}

func decode(encoded_text string, keyword string) string {
	has_same_size := len(encoded_text) == len(keyword)
	if !has_same_size  {
		keyword = transform_keyword(keyword, len(encoded_text))
	}
	counter := 0
	decoded_text := strings.Map(func (r rune) rune {
		hop := character_position(r)
		encryted_rune := retrieve_decoded_rune(int(encoded_text[counter]), hop)
		counter++
		return rune(encryted_rune)
	}, strings.ToUpper(keyword))
	return decoded_text
}

func character_position(character rune) int {
	alphabet := []rune{ 
			'A', 'B', 'C', 'D', 'E', 'F','G', 
			'H', 'I', 'J', 'K','L', 'M', 'N', 'O', 
			'P','Q', 'R', 'S', 'T', 'U','V', 'W', 'X', 
			'Y', 'Z',
	}
	return strings.IndexRune(string(alphabet), character)
}

func retrieve_encoded_rune(start int, add int) rune {
	position := start
	for counter := start; counter < start+add; counter++ {
		if counter == 90 {
			position = 65
		}
		position++
	}
	return rune(position)
}

func retrieve_decoded_rune(start int, add int) rune {
	position := start
	for counter := start; counter < start+add; counter++ {
		fmt.Println(counter)
		if position == 65 {
			position = 90
		}
		position--
	}
	return rune(position)
}