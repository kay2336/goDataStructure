package main

const MaxSize = 10

type List struct {
	length int
	data   [MaxSize]int
}

var (
	TList List
	arr   = [MaxSize]int{1, 2, 3, 3, 3, 3, 3, 3, 3, 3}
)
