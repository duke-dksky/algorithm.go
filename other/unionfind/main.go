package main

import "fmt"

type unionFind struct {
	parent []int
}

func NewUnionFind(n int) *unionFind {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &unionFind{
		parent: p,
	}
}

func (uf *unionFind) find(x int) int {
	for uf.parent[x] != x {
		x = uf.parent[x]
	}
	return x
}

func (uf *unionFind) Union(x, y int) {
	rootX, rootY := uf.find(x), uf.find(y)
	uf.parent[rootX] = rootY
}

func (uf *unionFind) IsConnected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func main() {
	fmt.Println("vim-go")
}
