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
	fmt.Print(" Inserting ")
	word = strings.ToLower(word)

	for _, l := range word {

		if existing := currentNode.FindChild(l); existing == nil {
			fmt.Printf(" '%c' ", l)
			currentNode = currentNode.AddChild(l, false) //insert new node
			fmt.Printf(" '%c' ", currentNode.letter)
		} else {
			fmt.Printf(" \\'%c' ", l)
			currentNode = existing //use existing node
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

type CompoundQueueData struct {
	word   string
	prefix string
	length int
}

func (this *Trie) LongestCompoundWord(words []string) string {
	queue := list.New()

	//Insertion
	for _, word := range words {
		prefixes := this.FindAllPrefixes(word)
		for _, prefix := range prefixes {
			queue.PushBack(CompoundQueueData{word: word, prefix: prefix, length: len(word)})
		}
		this.Insert(word)
	}

	//Processing
	longestWord := ""
	//maxLength := 0

	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(CompoundQueueData))
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
