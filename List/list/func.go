package main

import "fmt"

// InitList
// 转存数组至顺序表内
func (list *List) InitList(t [MaxSize]int) {
	//记录转存数组长度
	cnt := 0
	for i, _ := range t {
		(*list).data[i] = t[i]
		cnt++
	}
	(*list).length = cnt
}

// ListLength
// 返回顺序表的的长度
func (list *List) ListLength() int {
	return (*list).length
}

// ListEmpty
// 判断顺序表是否为空
func (list *List) ListEmpty() bool {
	return (*list).length == 0
}

// PrintlnAll
// 输出顺序表内所有数组元素
func (list *List) PrintlnAll() {
	//遍历以顺序表内长度变量为准
	for i := 0; i < (*list).length; i++ {
		fmt.Println("顺序表序号为", i, "的元素为:", (*list).data[i])
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

}

// UpdateByIndex
// 更新顺序表序号为index的元素,将值修改为e
// 特判序号是否合理
func (list *List) UpdateByIndex(i, t int) {
	if i < 0 || i >= (*list).length {
		return
	}
	(*list).data[i] = t
}

// JudgeByElem
// 查询顺序表中是否存在值为e的元素
// 如果有，则返回true，如果没有，则返回false
func (list *List) JudgeByElem(e int) bool {
	for _, v := range (*list).data {
		if v == e {
			return true
		}
	}
	return false
}

// DeleteByElem
// 删除值为e的元素
// 删除一个或者多个
// 删除的是第一个，最后一个，中间个
func (list *List) DeleteByElem(t int) {
	cnt := 0
	for i, v := range (*list).data {
		//如果元素值与t相同,则记录删除元素个数
		//如果不同,则移动当前元素位置至 i-cnt 处
		if v == t {
			cnt++
		} else {
			(*list).data[i-cnt] = (*list).data[i]
		}
	}
	//修正顺序表数组长度
	(*list).length = (*list).length - cnt
}

// InsertByIndex
// 向顺序表插入指定序号的元素
// 特判插入序号是否合理
// 插入指定序号位置后,[index:(*list).length]内的所有元素向后移动一位
func (list *List) InsertByIndex(index, e int) {
	if index < 0 || index >= (*list).length || (*list).length+1 > 50 {
		return
	}
	//从数组尾开始向后移动
	for i := (*list).length; i >= index; i-- {
		(*list).data[i+1] = (*list).data[i]
	}
	(*list).data[index] = e
	(*list).length++
}
