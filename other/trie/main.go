package main

import "fmt"

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		if t.next[word[i]-'a'] == nil {
			t.next[word[i]-'a'] = new(Trie)
		}
		t = t.next[word[i]-'a']
	}
	t.isEnd = true
}

func (t *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if t.next[word[i]-'a'] == nil {
			return false
		}
		t = t.next[word[i]-'a']
	}
	return t.isEnd
}

func (t *Trie) StartWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if t.next[prefix[i]-'a'] == nil {
			return false
		}
		t = t.next[prefix[i]-'a']
	}
	return true
}

type IntTrie struct {
	next [2]*IntTrie
}

func NewIntTrie() *IntTrie {
	return &IntTrie{}
}

func (t *IntTrie) Insert(x int) {
	for i := 63; i >= 0; i-- {
		p := (x >> i) & 0x01
		if t.next[p] == nil {
			t.next[p] = new(IntTrie)
		}
		t = t.next[p]
	}
}

// trie树的应用场景都需要改写search的功能
func (t *IntTrie) Search(x int) bool {
	for i := 63; i >= 0; i-- {
		p := (x >> i) & 0x01
		if t.next[p] == nil {
			return false
		}
		t = t.next[p]
	}
	return true
}

func main() {
	obj := NewTrie()
	fmt.Println(obj.next)
	obj.Insert("a")
	fmt.Println(obj.next)
}
