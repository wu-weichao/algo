package hashmap

import (
	"fmt"
	"testing"
)

//func TestAddressHashMap(t *testing.T) {
//	hm := NewAddressHashMap(5)
//	hm.Put(5, "golang")
//	hm.Put(98, "php")
//	hm.Put(64, "java")
//	hm.Put(64, "java")
//	hm.Put(1111, "js")
//	hm.Print()
//	fmt.Println("-----------------------")
//	t.Log(hm.Get(29))
//	t.Log(hm.Get(64))
//	fmt.Println("-----------------------")
//	t.Log(hm.Delete(15))
//	t.Log(hm.Delete(5))
//	hm.Print()
//	hm.Put(55, "vue")
//	hm.Print()
//}

func TestListHashMap(t *testing.T) {
	hm := NewListHashMap(5)
	hm.Put(5, "golang")
	hm.Put(98, "php")
	hm.Put(64, "java")
	hm.Put(64, "java")
	hm.Put(1111, "js")
	hm.Print()
	fmt.Println("-----------------------")
	t.Log(hm.Get(29))
	t.Log(hm.Get(64))
	fmt.Println("-----------------------")
	t.Log(hm.Delete(15))
	t.Log(hm.Delete(5))
	hm.Print()
	hm.Put(55, "vue")
	hm.Print()
}
