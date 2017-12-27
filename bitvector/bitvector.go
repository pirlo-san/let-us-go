package bitvector

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

const (
	wordBitCount = 32 << (^uint(0) >> 63)
)

// add x into set s
func (s *IntSet) Add(x int) {
	word, index := wordIndex(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << index
}

// check wether x is in set s
func (s *IntSet) Has(x int) bool {
	word, index := wordIndex(x)
	if word >= len(s.words) {
		return false
	}

	return (s.words[word] & (1 << index)) != 0
}

// set s as the union of set s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for idx, word := range t.words {
		if idx >= len(s.words) {
			s.words = append(s.words, word)
			continue
		}
		s.words[idx] |= word
	}
}

// return the string format of set s
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		for j := 0; j < wordBitCount; j++ {
			if (word & (1 << uint(j))) == 0 {
				continue
			}
			if buf.Len() > len("{") {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(&buf, "%d", i*wordBitCount+j)
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func wordIndex(x int) (int, uint) {
	return x / wordBitCount, uint(x) % wordBitCount
}
