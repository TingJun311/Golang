package main

import(
	"fmt"
)

type TrieNode struct {
    value    rune
    isEnd    bool
    children map[rune]*TrieNode
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{&TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, r := range word {
        child, ok := node.children[r]
        if !ok {
            child = &TrieNode{value: r, children: make(map[rune]*TrieNode)}
            node.children[r] = child
        }
        node = child
    }
    node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *TrieNode {
    node := t.root
    for _, r := range prefix {
        child, ok := node.children[r]
        if !ok {
            return nil
        }
        node = child
    }
    return node
}

func (t *Trie) AutoComplete(prefix string) []string {
    results := []string{}
    node := t.SearchPrefix(prefix)
    if node == nil {
        return results
    }
    t.collectWords(node, prefix, &results)
    return results
}

func (t *Trie) collectWords(node *TrieNode, word string, results *[]string) {
    if node.isEnd {
        *results = append(*results, word)
    }
    for _, child := range node.children {
        t.collectWords(child, word+string(child.value), results)
    }
}


func main() {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("orange")
	trie.Insert("abc")

	completions := trie.AutoComplete("ap")
	fmt.Println(completions) // ["apple"]

	completions = trie.AutoComplete("b")
	fmt.Println(completions) // ["banana"]

	completions = trie.AutoComplete("c")
	fmt.Println(completions) // []
}
