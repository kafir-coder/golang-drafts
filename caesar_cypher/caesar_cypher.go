package main

import (
	"fmt"
	"strings"
)
func main() {
	fmt.Println("Welcome");
	decoded_text := decoder("L fdph, L vdz, L frqtxhuhg", 3)
	encoded_text := encoder(decoded_text, 3);


	if test(encoded_text, decoded_text) {
		fmt.Println("[OK] ALL TESTS PASSED")
	}else {
		fmt.Println("[ERROR] TESTS FAILED")
	}
	
}

func decoder(encoded_text string, hop uint8) string {
  decoded_text := strings.Map(func(r rune) rune {
		if r == ',' || r == ' ' {
			return r
		}
		return r-rune(hop)
	}, encoded_text)
	return decoded_text;
}

func encoder(decoded_text string, hop uint8) string {
	encoded_text := strings.Map(func(r rune) rune {
		if r == ',' || r == ' ' {
			return r
		}
		return r+rune(hop)
	}, decoded_text)
	return encoded_text
}

func test(encoded_text string, decoded_text string) bool{
	truthy := 0
	if encoded_text == "L fdph, L vdz, L frqtxhuhg" {
		truthy++
	} 
	if decoded_text == "I came, I saw, I conquered" {
		truthy++
	}
	return truthy == 2;
}