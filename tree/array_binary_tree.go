package tree

type ArrayBinaryTree struct {
	data     []interface{}
	capacity int
	length   int
}

func NewArrayBinaryTree(cap int) *ArrayBinaryTree {
	return &ArrayBinaryTree{
		data:     make([]interface{}, cap),
		capacity: cap,
		length:   0,
	}
}

func (this *ArrayBinaryTree) Insert(v int) {

}

func (this *ArrayBinaryTree) Delete(v int) {

}

func (this *ArrayBinaryTree) Get(v int) {

}

func (this *ArrayBinaryTree) Print() {

}
