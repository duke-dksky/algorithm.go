package main

import "fmt"

/*
	https://www.cnblogs.com/xenny/p/9739600.html
	基于区间上的更新以及求和问题
	修改和查询的复杂度都是O(logN)
	C[i]=A[i−2k+1]+A[i−2k+2]+...+A[i]
	k为i的二进制中从最低位到高位连续零的长度
*/

type Bit struct {
	arr []int
}

func InitBit(n int) *Bit {
	arr := make([]int, n)
	bit := &Bit{arr}
	for i := 2; i < n; i++ {
		bit.Update(i, 1)
	}
	return bit
}

// 该位置之前所有元素的和
func (tree *Bit) Query(pos int) int {
	ret := 0
	for pos > 0 {
		ret += tree.arr[pos]
		pos -= lowbit(pos)
	}
	return ret
}

func (tree *Bit) Update(pos, k int) {
	for pos < len(tree.arr) {
		tree.arr[pos] += k
		pos += lowbit(pos)
	}
}

// 二进制中从最低位到高位连续零的长度k(lowbit=2^k)
func lowbit(x int) int {
	return x & (-x)
}

func main() {
	fmt.Println("vim-go")
}
