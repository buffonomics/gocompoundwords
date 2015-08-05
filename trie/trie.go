package trie

import (
	"bytes"
	"fmt"
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
	currentNode := this.root
	size := len(word)
	fmt.Print(" Inserting ")

	for i, l := range word {
		if !currentNode.HasChild(l) {
			fmt.Printf(" '%c'' ", l)
			if i == size-1 {
				currentNode.AddChild(l, true) //isTemrinal/end of word
				fmt.Println("end of Word")
			} else {
				currentNode.AddChild(l, false)

			}

		}
	}
}

func (this *Trie) Display() string {
	var out bytes.Buffer
	fmt.Println("\n Displaying Tree...\n")
	this.root.Display(out)
	output := out.String()
	fmt.Printf("%s \n", output)

	return output
}
