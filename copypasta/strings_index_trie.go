package copypasta

// CF101628K https://codeforces.com/gym/101628/submission/68323182

type iTrieNode struct {
	sonIndexes     [26]int
	fa             *iTrieNode
	curIndexes     *sTreap
	subTreeIndexes *sTreap
}

func (o *iTrieNode) pushUpAdd(idx int) {
	for ; o.fa != nil; o = o.fa {
		o.subTreeIndexes.put(idx)
		// 其余统计量
	}
}

func (o *iTrieNode) pushUpDel(idx int) {
	for ; o.fa != nil; o = o.fa {
		o.subTreeIndexes.delete(idx)
		// 其余统计量
	}
}

type iTrie struct {
	nodes []*iTrieNode
}

func newIndexTrie() *iTrie {
	return &iTrie{
		nodes: []*iTrieNode{{}}, // init with a root (empty node)
	}
}

func (t *iTrie) ord(c byte) byte { return c - 'a' }

func (t *iTrie) add(s string, idx int) {
	o := t.nodes[0]
	for i := range s {
		c := t.ord(s[i])
		if o.sonIndexes[c] == 0 {
			o.sonIndexes[c] = len(t.nodes)
			t.nodes = append(t.nodes, &iTrieNode{
				fa:             o,
				curIndexes:     &sTreap{seed: 1},
				subTreeIndexes: &sTreap{seed: 1},
			})
		}
		o = t.nodes[o.sonIndexes[c]]
	}
	o.curIndexes.put(idx)
	o.pushUpAdd(idx)
}

func (t *iTrie) del(s string, idx int) {
	o := t.nodes[0]
	for i := range s {
		o = t.nodes[o.sonIndexes[t.ord(s[i])]]
	}
	o.curIndexes.delete(idx)
	o.pushUpDel(idx)
}

// 在 trie 中找字符串 s 的前缀
func (t *iTrie) hasPrefixOfText(s string, l, r int) bool {
	o := t.nodes[0]
	for i := range s {
		idx := o.sonIndexes[t.ord(s[i])]
		if idx == 0 {
			return false
		}
		o = t.nodes[idx]
		if o.curIndexes.hasValueInRange(l, r) {
			return true
		}
	}
	return false
}

// 在 trie 中找前缀为 p 的字符串
func (t *iTrie) hasTextOfPrefix(p string, l, r int) bool {
	o := t.nodes[0]
	for i := range p {
		idx := o.sonIndexes[t.ord(p[i])]
		if idx == 0 {
			return false
		}
		o = t.nodes[idx]
	}
	return o.subTreeIndexes.hasValueInRange(l, r)
}
