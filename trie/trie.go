package trie

import (
	"bytes"
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
	//fmt.Print(" Inserting ")
	word = strings.ToLower(word)

	for _, l := range word {

		if existing := currentNode.FindChild(l); existing == nil {
			currentNode = currentNode.AddChild(l, false) //insert new node
			fmt.Printf(" '%c' ", currentNode.letter)
		} else {

			currentNode = existing //use existing node
			fmt.Printf(" \\'%c' ", l)
		}
	}

	currentNode.isTerminal = true
	fmt.Print("\n")
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

	//TODO: Instead of enforcing sorting externally, Ensure insertion is seperate from initial compound queue population.
	//Sorting just adds more uneccessary time complexity, as I can easily just do a 2n run.

	//Insertion
	for _, word := range words {
		word = strings.ToLower(word)
		prefixes := this.FindAllPrefixes(word)
		for _, prefix := range prefixes {
			queue.PushBack(CompoundQueueData{word: word, suffix: strings.Replace(word, prefix, "", -1)})
		}
		this.Insert(word)
	}

	//Processing
	longestWord := ""
	maxLength := 0

	for e := queue.Front(); e != nil; e = e.Next() {
		pair := e.Value.(CompoundQueueData)
		fmt.Println(pair)
		//shorting list as we run to conserve memory
		if prev := e.Prev(); prev != nil {
			queue.Remove(prev)
		}

		if this.HasWord(pair.suffix) && len(pair.word) > maxLength {
			longestWord = pair.word
		} else {
			prefixes := this.FindAllPrefixes(pair.suffix)
			for _, prefix := range prefixes {
				queue.PushBack(CompoundQueueData{word: pair.word, suffix: strings.Replace(pair.suffix, prefix, "", -1)})
			}
		}

	}

	return longestWord
}

func (this *Trie) Display() string {
	var out bytes.Buffer
	fmt.Println("\n Displaying Tree...\n")
	this.root.Displayf(0)
	//out.Flush()
	output := out.String()
	fmt.Printf("%s \n", output)

	return output
}
