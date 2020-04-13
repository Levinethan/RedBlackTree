package main

import "fmt"

func main()  {
	tree := NewRBtree()
	for i := 0 ; i< 1000000 ; i++{
		tree.Insert(Int(i))
	}

	for i := 0; i <900000;i++{
		tree.Delete(Int(i))

	}

	fmt.Println(tree.GetDepth())
	//100  12层
	//1000个数据 18层
	//10000000 45层 内存消耗500M
}