package main

import (
	"./trie"
	"./util"
	"fmt"
	"log"
	"sort"
	"time"
)

func main() {
	fmt.Println("NodePrime Quiz")
	//Load strings from file into Array
	words, err := util.LoadStrings("word.list")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		panic(err)
	}

	//Begin Operation
	var longestCompoundWord string
	start := time.Now()

	//sort array of lines
	sort.Sort(util.ByLength(words))

	//Init Trie
	trie := trie.NewTrie()

	testwords := []string{"Howdy", "Yam", "Orange", "How", "Sundry", "Suntan", "saxophone", "YamOranges", "howyam", "superfly"}
	sort.Sort(util.ByLength(testwords))
	longestCompoundWord = trie.LongestCompoundWord(testwords)
	//trie.Display()

	elapsed := time.Since(start)
	log.Printf("Longest Compound Word: %s", longestCompoundWord)
	log.Printf("Duration: %s", elapsed)

}
