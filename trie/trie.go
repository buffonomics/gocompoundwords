package trie

import (
	"bytes"
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
	size := len(word)
	fmt.Print(" Inserting ")
	word = strings.ToLower(word)

	for i, l := range word {

		if existing := currentNode.FindChild(l); existing == nil {
			fmt.Printf(" '%c' ", l)
			currentNode = currentNode.AddChild(l, false) //insert new node
			fmt.Printf(" '%c' ", currentNode.letter)
		} else {
			fmt.Printf(" \\'%c' ", l)
			currentNode = existing //use existing node
		}

		//If End of word
		if i == size-1 {
			currentNode.isTerminal = true
			fmt.Print("\n")
		}
	}
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

		//If End of word
		//if i == size-1 {
		//	currentNode.isTerminal = true
		//	fmt.Print("\n")
		//}
	}

	return prefixes
}

func (this *Trie) LongestCompoundWord() string {
	var prefixQueue []string

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
