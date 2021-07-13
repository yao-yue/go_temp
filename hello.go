package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/yao-yue/go_temp/morestrings"
)

//在cmp这个文件里面  compare.go 里面的这个文件
// package cmp 声明为cmp包
// 然后有这个func Diff(x, y interface{}, opts ...Option) string {
// 是大写的我们，也就是pubilc的 我们就可以进行引用

//我们可以做一个练习，写一个项目，然后引入这个包的reverse.go

func main() {
	fmt.Println("!oG ,olleH")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
