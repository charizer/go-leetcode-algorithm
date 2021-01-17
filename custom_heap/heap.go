package custom_heap

import "fmt"

// https://www.cnblogs.com/yahuian/p/11945144.html

// 本例为最小堆
// 最大堆只需要修改less函数即可
// 使用slice来存储我们的数据
type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) up(i int) {
	for {
		// 父亲结点
		f := (i - 1) / 2
		// 已经到达堆顶 或者 父亲已经小于自己， 则可以结束
		if i == f || h.less(f, i) {
			break
		}
		// 孩子比父亲小，则交换孩子和父亲的值
		h.swap(f, i)
		// 将父亲的位置赋值给孩子，继续向上调整
		i = f
	}
}

// 注意go中所有参数转递都是值传递
// 所以要让h的变化在函数外也起作用，此处得传指针
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1 // 左孩子
		if l >= len(h) || l < 0 {  // l < 0 整型溢出
			break // i已经是叶子结点了
		}
		j := l // j先赋值为左孩子
		// 如果右孩子比左孩子小，则j赋值为右孩子
		if r := l + 1; r < len(h) && h.less(r, l) {
			j = r // = 2*i + 2， 右孩子
		}
		// j是较小的孩子，将当前节点和较小孩子比较，如果比较小的还小，则已经满足小根堆了
		if h.less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}
		h.swap(i, j) // 交换父结点和子结点，将较小的孩子上移，自己下移
		i = j        //继续向下比较
	}
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}
	n := len(*h) - 1
	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (*h)[n]
	*h = (*h)[0:n]
	// 如果当前元素大于父结点，向下筛选
	if (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else { // 当前元素小于父结点，向上筛选
		h.up(i)
	}
	return x, true
}

// 弹出堆顶的元素，并返回其值
func (h *Heap) Pop() int {
	// 最后一个元素的位置
	n := len(*h) - 1
	// 最后一个元素和堆顶替换
	h.swap(0, n)
	// 占存原堆顶元素，用于返回
	x := (*h)[n]
	// 把原堆顶元素从底层slice中删除
	*h = (*h)[0:n]
	// 新堆顶向下调整
	h.down(0)
	return x
}

// 初始化堆，仅包含叶子结点的子树已经是堆。即在有n个结点的完全二叉树中，当 i>n/2-1 时，以i结点为根的子树已经是堆。
// 从最后一个非叶子节点开始，依次做“下沉”调整
// 时间复杂度O(n)
func (h Heap) Init() {
	n := len(h)
	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

func main()  {
	var args int64= 1
	// args就是实际参数
	printNumber(args)
}

// 这里定义的args就是形式参数
func printNumber(args ...int64)  {
	for _,arg := range args{
		fmt.Println(arg)
	}
}
