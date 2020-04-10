package ac_automation

import (
	"fmt"
	"sync"
)


type TrieNode struct {

	Char       rune
	wholeWord  string
	val        int
	failNode   *TrieNode
	parentNode *TrieNode
	subNodes   map[rune]*TrieNode
	lock       sync.Mutex

}

func NewTrieNode(char rune, val int) *TrieNode{
	return &TrieNode{
		Char:       char,
		val:        val,
		failNode:   nil,
		parentNode: nil,
		subNodes:   nil,
		lock:       sync.Mutex{},
	}
}

func (n *TrieNode) addSubNode(char rune, sn *TrieNode) *TrieNode {
	if n.subNodes == nil {
		if n.subNodes == nil {
			n.lock.Lock()
			n.subNodes = make(map[rune]*TrieNode)
			n.lock.Unlock()
		}
	}
	n.subNodes[char] = sn
	sn.parentNode = n
	return n
}

func (n *TrieNode) toString() {
	fmt.Print("|[", n.val, "]")
	if n.failNode != nil {
		fmt.Println("[", n.failNode.val, "]")
	}
}