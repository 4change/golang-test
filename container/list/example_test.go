// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"container/list"
	"fmt"
)

func Example() {
	// Create a new list and put some numbers in it.
	l := list.New()				// 	Front						<--root-->										Back
	e4 := l.PushBack(4)		//	Front						<--root<-->4-->									Back
	e1 := l.PushFront(1)		//	Front					<--1<-->root<-->4-->								Back
	l.InsertBefore(3, e4)		//	Front					<--1<-->root<-->3<-->４-->							Back
	l.InsertAfter(2, e1)		//	Front				<--1<-->2<-->root<-->3<-->4-->							Back

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
