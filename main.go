package main

import (
	"./trie"
	"./util"
	"log"
	"time"
)

func main() {
	log.Printf("NodePrime Compound-word Detector by Iyobo Eki \n")
	start := time.Now()

	//Load data from file into Array
	words, err := util.LoadStrings("word.list")
	if err != nil {
		log.Fatalf("LoadStrings: %s", err)
		panic(err)
	}

	//Begin Operation
	var longestCompoundWord string

	//Init Trie
	trie := trie.NewTrie()

	//Run compound-word detection algorithm
	longestCompoundWord = trie.DeriveLongestCompoundWord(words)
	//trie.DisplayToConsole()

	elapsed := time.Since(start)
	log.Printf("Longest Compound Word: %s", longestCompoundWord)
	log.Printf("Duration: %s", elapsed)

}
