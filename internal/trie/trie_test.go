package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieBasic(t *testing.T) {
	trie := NewTrie()
	trie.Insert("a")
	trie.Insert("ab")
	trie.Insert("abc")
	trie.Insert("abcd")

	a := trie.AutoComplete("a")
	ab := trie.AutoComplete("ab")
	abc := trie.AutoComplete("abc")
	abcd := trie.AutoComplete("abcd")
	abcde := trie.AutoComplete("abcde")
	other := trie.AutoComplete("other")

	assert.Equal(t, []string{"a", "ab", "abc", "abcd"}, a)
	assert.Equal(t, []string{"ab", "abc", "abcd"}, ab)
	assert.Equal(t, []string{"abc", "abcd"}, abc)
	assert.Equal(t, []string{"abcd"}, abcd)
	assert.Equal(t, []string{}, abcde)
	assert.Equal(t, []string{}, other)
}
