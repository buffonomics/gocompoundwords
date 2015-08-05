package main

import (
	"./trie"
	"./util"
	"flag"
	"log"
	"time"
)

func main() {
	log.Printf("NodePrime Compound-word Detector by Iyobo Eki \n")

	//Determine input data file
	datafile := flag.String("file", "word.list", "Path to a file containing new-line seperated strings")
	flag.Parse()

	start := time.Now()

	//Load data from file into Array
	words, err := util.LoadStrings(*datafile)
	if err != nil {
		log.Fatalf("%s", err)
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
