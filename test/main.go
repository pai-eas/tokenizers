package main

import (
	"fmt"

	"github.com/daulet/tokenizers"
)

func main() {
	// load tokenizer from file
	tk, err := tokenizers.FromFile("./data/bert-base-uncased.json")
	if err != nil {
		panic(err)
	}
	fmt.Print(tk.Encode("abc", false))
	// release native resources
	defer tk.Close()
}
