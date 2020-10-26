package main

import "fmt"

// O(1)
func test() {
	x := 1
	y := 2
	fmt.Printf("The sum: %d\n", x+y)
}

// O(logn)
func testlogn(n int) {
	i := 1
	for i <= n {
		i = i * 2
	}
	fmt.Printf("The i: %d\n", i)
}

// O(n)
func testn(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Printf("The sum: %d\n", sum)
}

// nO(logn)
func testnlogn(n int) {
	for j := 0; j < n; j++ {
		i := 1
		for i <= n {
			i = i * 2
		}
		fmt.Printf("The i: %d\n", i)
	}
}

// O(n^2)
func testn2(n int) {
	sum := 0
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			sum += i
		}
	}
	fmt.Printf("The sum: %d\n", sum)
}

// O(m+n)
func testmaddn(m, n int) {
	for i := 0; i < m; i++ {
		// ...
	}
	for j := 0; j < n; j++ {
		// ...
	}
}

// O(m*n)
func testmmcln(m, n int) {
	for i := 0; i < m; i++ {
		// ...
		for j := 0; j < n; j++ {
			// ...
		}
	}
}

func main() {
	// O(1)
	// test()

	// O(logn)
	//testlogn(10)

	// O(n)
	//testn(100)

	// O(nlogn)
	//testnlogn(10)

	// O(n^2)
	//testn2(10)

	// O(m+n)
	//testmn(10, 100)

	// O(m*n)
	//testmmcln(10, 100)
}
