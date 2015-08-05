package trie

import (
	"bytes"
	"fmt"
	"strconv"
)

type Node struct {
	letter rune
	//parent     *Node
	children   []*Node
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

func (this *Node) FindChild(letter rune) *Node {
	for _, child := range this.children {
		if child.letter == letter {
			return child
		}
	}
	return nil
}

func (this *Node) AddChild(letter rune, isTerminal bool) *Node {
	newnode := &Node{
		letter:     letter,
		isTerminal: isTerminal,
		//children:   make([]Node, 3),
	}

	this.children = append(this.children, newnode)
	return newnode
}

func (this *Node) Display(out *bytes.Buffer, indent int) {
	out.WriteString(strconv.Itoa(indent))
	out.WriteRune(this.letter)
	out.WriteString("-->\n")

	for _, c := range this.children {
		c.Display(out, indent+1)
	}
	out.WriteString("\n")
}

func (this *Node) Displayf(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("...")
	}
	fmt.Printf("%c \n", this.letter)

	for _, c := range this.children {
		c.Displayf(indent + 1)
	}
	//fmt.Print("\n")
}