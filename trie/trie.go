package trie

import (
	"container/list"
	"fmt"
	"strings"
)

type Trie struct {
	root Node
}

func NewTrie() *Trie {
	var trie = Trie{
		root: Node{letter: rune(0)},
	}

	return &trie
}

func (this *Trie) Insert(word string) {

	currentNode := &this.root
	word = strings.ToLower(word)

	for _, l := range word {

		if existing := currentNode.FindChild(l); existing == nil {
			currentNode = currentNode.AddChild(l, false) //insert new node
			//fmt.Printf(" '%c' ", currentNode.letter)
		} else {

			currentNode = existing //use existing node
			//fmt.Printf(" \\'%c' ", l)
		}
	}

	currentNode.isTerminal = true
	//fmt.Print("\n")
}

func (this *Trie) FindAllPrefixes(word string) []string {
	var prefixrunes []rune
	var prefixes []string

	currentNode := &this.root

	for _, l := range word {

		if existing := currentNode.FindChild(l); existing == nil {
			return prefixes
		} else {
			currentNode = existing //use existing node
		}

		prefixrunes = append(prefixrunes, currentNode.letter)

		if currentNode.isTerminal == true {
			prefixes = append(prefixes, string(prefixrunes))
		}
	}

	return prefixes
}

func (this *Trie) HasWord(word string) bool {

	currentNode := &this.root

	for _, l := range word {

		if existing := currentNode.FindChild(l); existing == nil {
			return false
		} else {
			currentNode = existing //use existing node
		}
	}

	return currentNode.isTerminal
}

type CompoundQueueData struct {
	word   string
	suffix string
}

func (this *Trie) LongestCompoundWord(words []string) string {
	queue := list.New()

	//Insertion-loop
	for _, word := range words {
		word = strings.ToLower(word)
		this.Insert(word)
	}

	/*
		Initial suffix extraction loop. Done here So I don't need to sort the data first, hence
		reducing the time complexity and increasing the speed of the algorithm.
	*/
	for _, word := range words {
		lword := strings.ToLower(word)
		prefixes := this.FindAllPrefixes(lword)
		for _, prefix := range prefixes {
			queue.PushBack(CompoundQueueData{word: word, suffix: strings.Replace(lword, prefix, "", 1)})
		}

	}

	//Further Processing
	longestCompoundWord := ""
	maxLength := 0

	for e := queue.Front(); e != nil; e = e.Next() {
		pair := e.Value.(CompoundQueueData)
		//fmt.Println(pair)
		//shorting make-shift queue as we run to conserve memory....and to act like a queue.
		if prev := e.Prev(); prev != nil {
			queue.Remove(prev)
		}

		if this.HasWord(pair.suffix) && len(pair.word) > maxLength {
			longestCompoundWord = pair.word
			maxLength = len(longestCompoundWord)
			//fmt.Println(pair)
		} else {
			prefixes := this.FindAllPrefixes(pair.suffix)
			for _, prefix := range prefixes {
				queue.PushBack(CompoundQueueData{word: pair.word, suffix: strings.Replace(pair.suffix, prefix, "", 1)})
			}
		}
	}

	return longestCompoundWord
}

func (this *Trie) DisplayToConsole() {
	fmt.Println("\n Displaying Tree...\n")
	this.root.Displayf(0)
}
