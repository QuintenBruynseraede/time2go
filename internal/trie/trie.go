package trie

type TrieNode struct {
	Children map[string]*TrieNode
	Value    string
	Terminal bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Children: make(map[string]*TrieNode),
		},
	}
}

func (t *Trie) Insert(value string) {
	current := t.Root
	runes := []rune(value)

	for i, char := range runes {
		if !Contains(current.Children, string(char)) {
			prefix := runes[0 : i+1]
			current.Children[string(char)] = &TrieNode{
				Children: make(map[string]*TrieNode),
				Value:    string(prefix),
			}
		}
		current = current.Children[string(char)]
	}
	current.Terminal = true
}

func (t *Trie) AutoComplete(prefix string) []string {
	words := make([]string, 0)
	current := t.Root

	for _, char := range prefix {
		if !Contains(current.Children, string(char)) {
			return []string{}
		}
		current = current.Children[string(char)]
	}

	return collectWords(current, words)
}

// Recursively fill accumulating list with all words starting from a node
func collectWords(current *TrieNode, words []string) []string {
	if current.Terminal {
		words = append(words, current.Value)
	}
	for _, child := range current.Children {
		words = collectWords(child, words)
	}
	return words
}

func Contains(m map[string]*TrieNode, key string) bool {
	_, ok := m[key]
	return ok
}
