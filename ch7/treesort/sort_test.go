// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort

import (
	"testing"
)

func TestString(t *testing.T) {
    root := &tree{value: 3}
    root = add(root, 2)
    root = add(root, 4)
    if root.String() != "[2 3 4]" {
        t.Log(root)
        t.Fail()
    }
}
