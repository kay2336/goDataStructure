package main

import "fmt"

/*
	手写数据结构————线性表
	初始化
	求长度
	判断是否为空
	输出线性表

	删除元素
	查询元素
	修改元素
	在某一位置插入元素
*/

func main() {
	//测试函数是否正确 (偷懒就没写在测试文件里了)
	TList.InitList(arr)
	fmt.Println(TList.ListLength())
	fmt.Println(TList.ListEmpty())
	TList.PrintlnAll()

	TList.UpdateByIndex(6, 200)
	TList.UpdateByIndex(0, 100)
	TList.UpdateByIndex(10, 999)
	TList.PrintlnAll()

	fmt.Println(TList.JudgeByElem(200))
	fmt.Println(TList.JudgeByElem(4))
	TList.PrintlnAll()

	TList.DeleteByElem(3)
	TList.DeleteByElem(5)
	TList.PrintlnAll()

	TList.InsertByIndex(0, -1)
	TList.InsertByIndex(12, -100)
	TList.PrintlnAll()

}
