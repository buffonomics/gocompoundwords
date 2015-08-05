package main

import (
	"./trie"
	"./util"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("NodePrime Quiz")
	start := time.Now()
	//Load strings from file into Array
	words, err := util.LoadStrings("word.list")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		panic(err)
	}

	//Begin Operation
	var longestCompoundWord string

	//Init Trie
	trie := trie.NewTrie()

	longestCompoundWord = trie.LongestCompoundWord(words)
	//trie.DisplayToConsole()

	elapsed := time.Since(start)
	log.Printf("Longest Compound Word: %s", longestCompoundWord)
	log.Printf("Duration: %s", elapsed)

}
