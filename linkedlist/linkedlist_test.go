package linkedlist

import (
	"testing"
)

func TestMakeLinkList(t *testing.T) {
	testcases := [][]int{
		{},
		{1},
		{1, 2, 3},
	}
	for _, sl := range testcases {
		l := MakeLinkedList(sl)
		t.Logf("l = [%v]", l)
	}

}
func TestListNodeEquals(t *testing.T) {
	testcases := []struct {
		l1     []int
		l2     []int
		equals bool
	}{
		{
			// 1
			l1:     []int{},
			l2:     []int{},
			equals: true,
		},
		{
			// 2
			l1:     []int{},
			l2:     []int{1},
			equals: false,
		},
		{
			// 3
			l1:     []int{1},
			l2:     []int{1},
			equals: true,
		},
		{
			// 4
			l1:     []int{1},
			l2:     []int{1, 2, 3},
			equals: false,
		},
		{
			// 5
			l1:     []int{1, 2, 3},
			l2:     []int{1, 2, 3},
			equals: true,
		},
		{
			// 6
			l1:     []int{1, 2},
			l2:     []int{},
			equals: false,
		},
	}
	for i, tcase := range testcases {
		l1 := MakeLinkedList(tcase.l1)
		l2 := MakeLinkedList(tcase.l2)
		equals := ListNodeEquals(l1, l2)
		if equals == tcase.equals {
			t.Logf("%d/%d PASSED ", i+1, len(testcases))
		} else {
			t.Errorf("%d/%d FAILED %v,%v ", i+1, len(testcases), l1, l2)
		}

	}

}
