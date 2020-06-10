// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"
import "testing"

func TestLenZeroInitially(t *testing.T) {
	s := &IntSet{}
	if s.Len() != 0 {
		t.Logf("%d != 0", s.Len())
		t.Fail()
	}
}

func TestLenAfterAddingElements(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Add(2000)
	if s.Len() != 2 {
		t.Logf("%d != 2", s.Len())
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Remove(0)
	if s.Has(0) {
		t.Log(s)
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Add(1000)
	s.Clear()
	if s.Has(0) || s.Has(1000) {
		t.Log(s)
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	orig := &IntSet{}
	orig.Add(1)
	copy := orig.Copy()
	copy.Add(2)
	if !copy.Has(1) || orig.Has(2) {
		t.Log(orig, copy)
		t.Fail()
	}
}

func TestAddAll(t *testing.T) {
    s := &IntSet{}
    s.AddAll(0, 2, 4)
    if !s.Has(0) || !s.Has(2) || !s.Has(4) {
        t.Log(s)
        t.Fail()
    }
}

func TestIntersectWith(t *testing.T) {
    s1 := &IntSet{}
    s2 := &IntSet{}
    s1.AddAll(0,2,4)
    s2.AddAll(1,2,3)
    s1.IntersectWith(s2)
    if !s1.Has(2) || s1.Len() != 1 {
        t.Log(s1)
        t.Fail()
    }
}

func TestDifferenceWith(t *testing.T) {
    s := &IntSet{}
    s.AddAll(0,2,4)
    u := &IntSet{}
    u.AddAll(1,2,3)
    s.DifferenceWith(u)
    expected := &IntSet{}
    expected.AddAll(0,4)
    if s.String() != expected.String() {
        t.Log(s)
        t.Fail()
    }
}

func TestSymmetricDifferenceWith(t *testing.T) {
    s := &IntSet{}
    s.AddAll(0,2,4)
    u := &IntSet{}
    u.AddAll(1,2,3)
    s.SymmetricDifferenceWith(u)
    expected := &IntSet{}
    expected.AddAll(0,1,3,4)
    if s.String() != expected.String() {
        t.Log(s)
        t.Fail()
    }
}

func TestElem(t *testing.T) {
    elems := []int{0,2,4,6}
    s := &IntSet{}
    s.AddAll(elems...)
    for i, n := range s.Elems() {
        if elems[i] != n{
            t.Log(s.Elems())
            t.Fail()
        }
    }
}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
