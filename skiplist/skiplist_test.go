package skiplist

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	skiplist := NewSkipList()
	// insert
	fmt.Println("---------- insert -----------")
	skiplist.Insert("wwc", 5)
	t.Log(skiplist)
	skiplist.Insert("qcq", 20)
	skiplist.Insert("wyh", 15)
	skiplist.Insert("go", 35)
	t.Log(skiplist)
	skiplist.Print()
	// find
	fmt.Println("---------- find -----------")
	t.Log(skiplist.Find("qcq", 20))
	t.Log(skiplist.Find("php", 30))
	// delete
	fmt.Println("---------- delete -----------")
	t.Log(skiplist.Delete("vue", 18))
	t.Log(skiplist)
	t.Log(skiplist.Delete("qcq", 20))
	t.Log(skiplist)
	skiplist.Print()
}
