package main

import "fmt"

/*
	https://www.cnblogs.com/xenny/p/9739600.html
	解决基于区间上的更新以及求和问题
	修改和查询的复杂度都是O(logN)
	C[i]=A[i−2^k+1]+A[i−2^k+2]+...+A[i]
	k为i的二进制中从最低位到高位连续零的长度
*/

type Bit struct {
	arr []int
}

func (tree *Bit) Query(pos int) int {
	ret := 0
	for 0 < pos && pos < len(tree.arr) {
		ret += tree.arr[pos]
		pos -= lowbit(pos)
	}
	return ret
}

func (tree *Bit) Update(pos, k int) {
	for 0 < pos && pos < len(tree.arr) {
		tree.arr[pos] += k
		pos += lowbit(pos)
	}
}

// 二进制中从最低位到高位连续零的长度为k(lowbit=2^k)
func lowbit(x int) int {
	return x & (-x)
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}

	// 1. 传统数组, 单点更新, 单点查询

	// 2. 差分建树, 区间更新, 单点查询
	//t := &Bit{
	//	arr: make([]int, len(a)+5),
	//}
	//for i := 1; i < len(a); i++ {
	//	t.Update(i, a[i]-a[i-1])
	//}
	//t.Update(2, 10)
	//t.Update(4+1, -10)
	//fmt.Println(t.Query(2))

	// 3. 普通建树, 单点更新, 区间查找
	t1 := &Bit{
		arr: make([]int, len(a)+5),
	}
	for i := 1; i < len(a); i++ {
		t1.Update(i, a[i])
	}
	fmt.Println(t1.Query(2))
	fmt.Println(t1.Query(3))

}
