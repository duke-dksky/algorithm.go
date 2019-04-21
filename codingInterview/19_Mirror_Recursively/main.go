package main

import (
	"algorithm.go/utils"
	"fmt"
)

func main() {

	t := utils.NewTree(8)
	t = utils.InsertNodeToTree(t, 10)
	t = utils.InsertNodeToTree(t, 6)
	t = utils.InsertNodeToTree(t, 5)
	t = utils.InsertNodeToTree(t, 7)
	t = utils.InsertNodeToTree(t, 9)
	t = utils.InsertNodeToTree(t, 11)

	utils.LevelOrder(t)
	fmt.Println()
}
