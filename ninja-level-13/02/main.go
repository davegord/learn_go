package main

import (
	"fmt"

	"github.com/davegord/learn_go/ninja-level-13/02/quote"
	"github.com/davegord/learn_go/ninja-level-13/02/word"
)

func main() {

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	fmt.Println(word.Count(quote.SunAlso))

}
