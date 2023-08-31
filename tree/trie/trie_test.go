package trie

import (
	"fmt"
	"testing"
)

func TestHeight(t *testing.T) {

	testcases := []struct {
		stringSet  []string
		trieHeight int
	}{
		{
			stringSet:  []string{},
			trieHeight: 0,
		},
		{
			stringSet:  []string{"foo"},
			trieHeight: 3,
		},
		{
			stringSet:  []string{"pot", "bowl"},
			trieHeight: 4,
		},
		{
			stringSet:  []string{"paper", "crayons", "notebook"},
			trieHeight: 8,
		},
		{
			stringSet:  []string{"chair", "desk", "calculator", "pencil"},
			trieHeight: 10,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("StringSet_%d", i), func(t *testing.T) {
			trie := New()
			for _, s := range tc.stringSet {
				trie.AddString(s)
			}
			trieHeight := trie.Height()
			if trieHeight != tc.trieHeight {
				t.Errorf("Incorrect trie size; got %d, want %d", trieHeight, tc.trieHeight)
			}
		})
	}
}

func TestCounter(t *testing.T) {

	testcases := []struct {
		stringSet []string
		wordCount int
	}{
		{
			stringSet: []string{},
			wordCount: 0,
		},
		{
			stringSet: []string{"foo"},
			wordCount: 1,
		},
		{
			stringSet: []string{"a", "b", "c"},
			wordCount: 3,
		},
		{
			stringSet: []string{"foo", "boo", "foo", "bang"},
			wordCount: 3,
		},
		{
			stringSet: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
			wordCount: 6,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("StringSet_%d", i), func(t *testing.T) {
			trie := New()
			for _, s := range tc.stringSet {
				trie.AddString(s)
			}
			if trie.WordCount() != tc.wordCount {
				t.Errorf("Incorrect trie word count; got %d, want %d", trie.WordCount(), tc.wordCount)
			}
		})
	}
}
