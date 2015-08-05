package trie

import (
	"fmt"
)

//A Node to be used within a Trie
type Node struct {
	letter     rune
	children   []*Node
	isTerminal bool
}

//Checks if letter exists in direct child nodes.
func (this *Node) HasChild(letter rune) bool {
	if child := this.FindChild(letter); child != nil {
		return true
	} else {
		return false
	}

}

//Searches direct children for a node with the letter.
func (this *Node) FindChild(letter rune) *Node {
	for _, child := range this.children {
		if child.letter == letter {
			return child
		}
	}
	return nil
}

//Create a new child node with letter. Set isTerminal to true is this node ends a word.
func (this *Node) AddChild(letter rune, isTerminal bool) *Node {
	newnode := &Node{
		letter:     letter,
		isTerminal: isTerminal,
	}

	this.children = append(this.children, newnode)
	return newnode
}

//For Debugging. Use to recursively display all children of this node depth-first.
func (this *Node) Displayf(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("...")
	}
	fmt.Printf("%c \n", this.letter)

	for _, c := range this.children {
		c.Displayf(indent + 1)
	}
}
