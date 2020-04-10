package ac_automation

import (
	"fmt"
	"log"
)

type TrieTree struct {
	root *TrieNode
	cnt int
}

func NewTireTree() TrieTree {
	return TrieTree{root:NewTrieNode(0, 0)}
}

func (t *TrieTree) Query(sentence string) []string{
	rs := make([]string, 0, 8)

	words := []rune(sentence)
	pointer := t.root
	cnt := 0
	for i := 0; i < len(words); {
		w := words[i]
		nextNode, nextExist := pointer.subNodes[w]

		if nextExist {
			pointer = nextNode
			i ++
		} else {
			if pointer == t.root {
				i ++
			} else {
				pointer = pointer.failNode
			}
		}

		if pointer.wholeWord != "" {
			rs = append(rs, pointer.wholeWord)
		}
		cnt ++ //执行计数
	}
	log.Print("目标长度:", len(sentence), " 执行次数", cnt)
	return rs
}

func (t *TrieTree) AddWord(word string) *TrieTree {
	tmpWord := ([]rune)(word)
	now := t.root
	for _, r := range tmpWord {
		//fmt.Println(string(r))
		next, exist := now.subNodes[r]
		if now.subNodes == nil || !exist {
			t.cnt ++
			nt := NewTrieNode(r, t.cnt)
			now.addSubNode(r, nt)
			//nt.wholeWord = string(r)
			now = nt
		} else {
			now = next
		}
	}
	now.wholeWord = word
	return t;
}

func (t *TrieTree) FailLink() {
	que := make([]*TrieNode, 0, 2)
	cnt := 0
	head := 0
	for _, node := range t.root.subNodes {
		que = append(que, node)
		cnt ++
	}

	for head < cnt  {
		now := que[head]
		if now.parentNode != nil &&
			now.parentNode.failNode != nil {
			if fn, exist := now.parentNode.failNode.subNodes[now.Char]; exist {
				now.failNode = fn
			} else {
				now.failNode = t.root
			}
		} else {
			now.failNode = t.root
		}

		for _, node := range que[head].subNodes {
			que = append(que, node)
			cnt ++
		}
		head ++
	}
}

func (t *TrieTree) ShowTree() {
	fmt.Println("start...")
	dfs(t.root)
}

func dfs(node *TrieNode) {
	if node.subNodes == nil || len(node.subNodes) == 0 {
		fmt.Println(node.wholeWord)
		return;
	}
	node.toString()
	for _, trieNode := range node.subNodes {
		dfs(trieNode)
	}

}


