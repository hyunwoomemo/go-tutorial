package main

import (
	"fmt"

	mydict "github.com/hyunwoomemo/dict/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	word := "hello"
	dictionary.Add(word, "First")

	def, _ := dictionary.Search(word)
	fmt.Println(def)

	dictionary.Delete(word)

	def, err2 := dictionary.Search(word)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(def)
	}

}