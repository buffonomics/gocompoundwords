package trie

import (
	"bytes"
)

type Node struct {
	letter     rune
	children   []Node
	isTerminal bool
}

func (this *Node) HasChild(letter rune) bool {
	for _, child := range this.children {
		if child.letter == letter {
			return true
		}
	}
	return false
}

func (this *Node) AddChild(letter rune, isTerminal bool) {
	newnode := Node{
		letter:     letter,
		isTerminal: isTerminal,
	}

	this.children = append(this.children, newnode)

}

func (this *Node) Display(out bytes.Buffer) {
	out.WriteString("  ")
	out.WriteRune(this.letter)
	for _, c := range this.children {
		c.Display(out)
	}
}
