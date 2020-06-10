// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
    "bytes"
    "fmt"
)

const wordSize = 32 << (^uint(0) >> 63)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
    words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
    word, bit := x/wordSize, uint(x%wordSize)
    return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
    word, bit := x/wordSize, uint(x%wordSize)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

func (s *IntSet) IntersectWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] &= tword
        }
    }
}

func (s *IntSet) DifferenceWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] &^= tword
        }
    }
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
    tmp := s.Copy()
    s.UnionWith(t)
    tmp.IntersectWith(t)
    s.DifferenceWith(tmp)
}

//!-intset

func popcount(x uint) int {
    count := 0
    for x != 0 {
        count++
        x &= x - 1
    }
    return count
}

// Return the number of elements
func (s * IntSet) Len() int {
    count := 0
    for _, word := range s.words {
        count += popcount(word)
    }
    return count
}

// Remove x from the set
func (s *IntSet) Remove(x int) {
    word, bit := x/wordSize, uint(x%wordSize)
    s.words[word] &^= 1 << bit
}

// Remove all elements from the set 
func (s * IntSet) Clear() {
    for i := range s.words {
        s.words[i] = 0
    }
}

// Return a copy of the set 
func (s *IntSet) Copy() *IntSet {
    tmp := &IntSet{}
    tmp.words = make([]uint, len(s.words))
    copy(tmp.words, s.words)
    return tmp
}

// Add a list of values
func (s *IntSet) AddAll(nums ...int) {
    for _, x := range nums {
        s.Add(x)
    }
}

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < wordSize; j++ {
            if word&(1<<uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", wordSize*i+j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}

//!-string

func (s *IntSet) Elems() []int {
    e := make([]int, 0)
    for i, word := range s.words {
        for j := 0; j < wordSize; j++ {
            if word&(1<<uint(j)) != 0 {
                e = append(e, i*wordSize+j)
            }
        }
    }
    return e
}
