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

func main() {
	obj := NewTrie()
	obj.Insert("duke")
	fmt.Println(obj.Search("duke"))
	fmt.Println(obj.Search("dukk"))
	fmt.Println(obj.StartWith("du"))
}
