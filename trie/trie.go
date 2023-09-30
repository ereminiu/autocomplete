package trie

import (
	"github.com/ereminiu/autocomplete/models"
	"maps"
	"math"
	"sort"
)

type cacheType map[string]int

const K = 5
const INF = math.MaxInt

type TrieNode struct {
	kids       map[rune]*TrieNode
	cache      map[string]int
	isTerminal bool
	freq       int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{kids: make(map[rune]*TrieNode), cache: make(map[string]int, 5), isTerminal: false, freq: 0}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) addWord(node *TrieNode, word string, freq int, k int) {
	if k == len(word)-1 {
		node.isTerminal = true
		node.freq += freq
		node.cache[word] = node.freq
		// relax cache by deleting the least popular word
		if len(node.cache) > K {
			var least string
			minVal := INF
			for w := range node.cache {
				if node.cache[w] < minVal {
					least = w
					minVal = node.cache[w]
				}
			}
			delete(node.cache, least)
		}
		return
	}
	next := rune(word[k+1])
	if node.kids[next] == nil {
		node.kids[next] = NewTrieNode()
	}
	t.addWord(node.kids[next], word, freq, k+1)
	update(node)
}

func update(node *TrieNode) {
	res := make(cacheType)
	maps.Copy(res, node.cache)
	for c := range node.kids {
		res = merge(res, node.kids[c].cache)
	}
	// don't use node.cache = res
	node.cache = make(cacheType, len(res))
	maps.Copy(node.cache, res)
}

func merge(ca, cb cacheType) cacheType {
	both := make([]double, 0, len(ca)+len(cb))
	for key, val := range ca {
		both = append(both, double{key, val})
	}
	for key, val := range cb {
		both = append(both, double{key, val})
	}
	sort.Slice(both, func(i, j int) bool {
		return both[i].freq > both[j].freq
	})
	res := make(cacheType)
	cnt := 0
	for _, pair := range both {
		res[pair.query] = pair.freq
		cnt++
		if cnt == K {
			break
		}
	}
	return res
}

func (t *Trie) Rebuild(records []models.Record) {
	for _, rec := range records {
		p := t.root
		word := rec.Query
		for _, c := range word {
			if p.kids[c] == nil {
				p.kids[c] = NewTrieNode()
			}
			p = p.kids[c]
		}
		p.freq = rec.Freq
		p.cache[word] = p.freq
	}
	rebuildCache(t.root)
}

func rebuildCache(node *TrieNode) {
	for next := range node.kids {
		rebuildCache(node.kids[next])
	}
	update(node)
}

func (t *Trie) AddWord(word string) {
	if len(word) == 0 {
		return
	}
	t.addWord(t.root, word, 1, -1)
	update(t.root)
}

type double struct {
	query string
	freq  int
}

func (t *Trie) GetTopFive(prefix string) []string {
	p := t.root
	for _, c := range prefix {
		if p.kids[c] == nil {
			return make([]string, 0)
		}
		p = p.kids[c]
	}
	res := make([]string, 0, len(p.cache))
	for key := range p.cache {
		res = append(res, key)
	}
	sort.Slice(res, func(i, j int) bool {
		return p.cache[res[i]] > p.cache[res[j]]
	})
	return res[:min(K, len(res))]
}
