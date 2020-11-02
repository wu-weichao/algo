package binary

import (
	"testing"
)

//func TestBinarySearch(t *testing.T) {
//	data := []int{1, 2, 4, 5, 6, 7, 9, 11, 13, 15, 16, 17, 18, 19, 25, 35, 42, 47, 55, 65, 77, 78, 79, 82, 88, 99, 100}
//	t.Log(BinarySearch(data, 99))
//	t.Log(BinarySearch(data, 3))
//}

//func TestBinarySearchEqFirst(t *testing.T) {
//	data := []int{1, 2, 4, 5, 6, 7, 9, 11, 13, 13, 13, 17, 18, 19, 25, 35, 42, 47, 77, 77, 77, 78, 79, 82, 99, 99, 100}
//	t.Log(BinarySearchEqFirst(data, 77))
//	t.Log(BinarySearchEqFirst(data, 17))
//}

//func TestBinarySearchEqLast(t *testing.T) {
//	data := []int{1, 2, 4, 5, 6, 7, 9, 11, 13, 13, 13, 17, 18, 19, 25, 35, 42, 47, 77, 77, 77, 78, 79, 82, 99, 99, 100}
//	t.Log(BinarySearchEqLast(data, 77))
//	t.Log(BinarySearchEqLast(data, 17))
//}

//func TestBinarySearchEgtFirst(t *testing.T) {
//	data := []int{1, 2, 4, 5, 6, 7, 9, 11, 13, 13, 13, 17, 18, 19, 25, 35, 42, 47, 77, 77, 77, 78, 79, 82, 99, 99, 100}
//	index := BinarySearchEgtFirst(data, 55)
//	if index >= 0 {
//		t.Log(index, data[index])
//	} else {
//		t.Log("egt 55 is not exists")
//	}
//	index = BinarySearchEgtFirst(data, 8)
//	if index >= 0 {
//		t.Log(index, data[index])
//	} else {
//		t.Log("egt 8 is not exists")
//	}
//}

func TestBinarySearchEltLast(t *testing.T) {
	data := []int{1, 2, 4, 5, 6, 7, 9, 11, 13, 13, 13, 17, 18, 19, 25, 35, 42, 47, 77, 77, 77, 78, 79, 82, 99, 99, 100}
	index := BinarySearchEltLast(data, 77)
	if index >= 0 {
		t.Log(index, data[index])
	} else {
		t.Log("elt 77 is not exists")
	}
	index = BinarySearchEltLast(data, 8)
	if index >= 0 {
		t.Log(index, data[index])
	} else {
		t.Log("elt 8 is not exists")
	}
}
