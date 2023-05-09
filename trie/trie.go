package trie

import (
	"fmt"
	"strings"
	"unicode"
)

type TrieNode struct {
    value    rune
    isEnd    bool
    children map[rune]*TrieNode
}

type Trie struct {
    root *TrieNode
    size int
}

func NewTrie() *Trie {
    return &Trie{&TrieNode{children: make(map[rune]*TrieNode)}, 0}
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, r := range word {
        child, ok := node.children[r]
        if !ok {
            child = &TrieNode {
                value: r, 
                children: make(map[rune]*TrieNode),
            }
            node.children[r] = child
        }
        node = child
    }
    if !node.isEnd {
		node.isEnd = true
		t.size++
	}
}

func (t *Trie) SearchPrefix(prefix string) *TrieNode {
    node := t.root
    for _, r := range strings.ToLower(prefix) {
        child, ok := node.children[r]
        if !ok {
            child, ok = node.children[unicode.ToUpper(r)]
            if !ok {
                return nil
            }
        }
        node = child
    }
    return node
}

func (t *Trie) Size() int {
	return t.size
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
        t.collectWords(child, word + string(child.value), results)
    }
}


func (t *Trie) Clear() {
	t.root.children = make(map[rune]*TrieNode)
	t.root.isEnd = false
	t.size = 0
}

func (t *Trie) normalize(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func (t *Trie) Delete(word string) bool {
	node := t.root
	parent := node
	var parentRune rune
	for _, r := range word {
		r = t.normalize(r)
		child, ok := node.children[r]
		if !ok {
			return false
		}
		parent = node
		parentRune = r
		node = child
	}
	if !node.isEnd {
		return false
	}
	node.isEnd = false
	t.size--
	if len(node.children) == 0 {
		delete(parent.children, parentRune)
	}
	return true
}

func (t *Trie) Display() {
    fmt.Println("Trie:")
    t.root.display(0)
}

func (n *TrieNode) display(level int) {
    prefix := ""
    for i := 0; i < level; i++ {
        prefix += "  "
    }
    fmt.Printf("%s%c", prefix, n.value)
    if n.isEnd {
        fmt.Println(" *")
    } else {
        fmt.Println()
    }
    for _, child := range n.children {
        child.display(level + 1)
    }
}

func (t *Trie) GetAllWords() []string {
    words := []string{}
    t.root.getAllWords("", &words)
    return words
}

func (n *TrieNode) getAllWords(prefix string, words *[]string) {
    if n.isEnd {
        *words = append(*words, prefix + string(n.value))
    }
    for _, child := range n.children {
        child.getAllWords(prefix + string(n.value), words)
    }
}

func Main2() {
	title := []string {
        "Apparel", 
        "Aparel", 
        "Ai", 
        "Electronics", 
        "Books", 
        "Home & Kitchen", 
        "Beauty & Personal Care", 
        "Sports & Fitness", 
        "Toys & Games", 
        "Automotive", 
        "Grocery",
    }
	trie := NewTrie()

	for _, i := range title {
		trie.Insert(i)
	}

	completions := trie.AutoComplete("T")
	fmt.Println(completions) // ["apple"]
}
