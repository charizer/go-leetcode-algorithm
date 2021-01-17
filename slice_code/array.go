package slice_code

import "fmt"

// 在 Go 中，与 C 数组变量隐式作为指针使用不同，Go 数组是值类型，赋值和函数传参操作都会复制整个数组数据。

func ArrayAddress(){
	arrayA := [2]int{100, 200}
	var arrayB [2]int
	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

// arrayA : 0xc0000162d0 , [100 200]
// arrayB : 0xc0000162e0 , [100 200]
// func Array : 0xc0000be010 , [100 200]

// 可以看到，三个内存地址都不同，这也就验证了 Go 中数组赋值和函数传参都是值复制的。那这会导致什么问题呢？

// 假想每次传参都用数组，那么每次数组都要被复制一遍。如果数组大小有 100万，在64位机器上就需要花费大约 800W 字节，即 8MB 内存。
// 这样会消耗掉大量的内存。于是乎有人想到，函数传参用数组的指针。

func ArrayPoint(){
	arrayA := []int{100, 200}
	testArrayPoint(&arrayA)   // 1.传数组指针
	arrayB := arrayA[:]
	testArrayPoint(&arrayB)   // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}

func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}

// func Array : 0xc000062080 , [100 200]
// func Array : 0xc0000620c0 , [100 300]
// arrayA : 0xc000062080 , [100 400]

// 这也就证明了数组指针确实到达了我们想要的效果。现在就算是传入10亿的数组，也只需要再栈上分配一个8个字节的内存给指针就可以了。
// 这样更加高效的利用内存，性能也比之前的好。

// 不过传指针会有一个弊端，从打印结果可以看到，第一行和第三行指针地址都是同一个，万一原数组的指针指向更改了，
// 那么函数里面的指针指向都会跟着更改。

// 切片的优势也就表现出来了。用切片传数组参数，既可以达到节约内存的目的，也可以达到合理处理好共享内存的问题。
// 打印结果第二行就是切片，切片的指针和原来数组的指针是不同的。

// 由此我们可以得出结论：

// 把第一个大数组传递给函数会消耗很多内存，采用切片的方式传参可以避免上述问题。
// 切片是引用传递，所以它们不需要使用额外的内存并且比使用数组更有效率。