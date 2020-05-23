# Go 单链表

```go
package main

import "fmt"

//1.初始化
//2.判空
//3.判满
//4.获取元素
//5.添加元素
//6.删除元素
//7.获取长度
//8.遍历
//9.清空
//10. 获取元素 的前一个元素
//11. 获取后一个 元素
//12. 传入一个值 返回匹配的下标
type List struct {
	Len  int
	Cap  int
	Data *[]interface{}
}

func (l *List) Init(c int) {
	l.Len = 0
	l.Cap = c
	m := make([]interface{}, c)
	l.Data = &m

}
func (l *List) Empty() bool {
	if l.Len == 0 {
		return true
	}

	return false

}
func (l *List) Funll() bool {
	if l.Len == l.Cap {
		return true

	}
	return false
}

func (l *List) GetData(i int) bool {

	if i < 0 || i > l.Len {
		return false
	}

	fmt.Println((*l.Data)[i-1])

	return true

}
func (l *List) Insert(index int, elem interface{}) bool {

	if index < 0 || l.Len == l.Cap || index > l.Cap {
		return false
	}
	//在指定地方插入元素
	for i := l.Len - 1; i >= index; i-- {
		(*l.Data)[i+1] = (*l.Data)[i]

	}

	(*l.Data)[index] = elem
	l.Len++
	return true
}

//删除元素

func (l *List) Delete(index int, elem interface{}) bool {

	if index < 0 || index > l.Cap || l.Len == 0 {
		return false
	}
	//删除元素  循环遍历

	for i := index; i < l.Len-1; i++ {

		(*l.Data)[i-1] = (*l.Data)[i]

	}
	l.Len--

	return true
}

//遍历
func (l *List) Range() {

	fmt.Println("开始遍历链表===")
	for i := 0; i < l.Cap; i++ {
		fmt.Println((*l.Data)[i])

	}
}

//清空

func (l *List) Clear() {
	l.Len = 0
	l.Data = nil
}

//返回长度
func (l *List) LenthAndCap() (a, b int) {

	return l.Cap, l.Len
}

func main() {

	//实现顺序表
	//初始化链表

	var l List
	l.Init(10)
	//打印一下 l的长度
	a, b := l.LenthAndCap()
	fmt.Println("len=", a, "cap=", b)

	//判空
	m := l.Empty()
	fmt.Println("是否为空=", m)

	//添加元素

	result := l.Insert(0, "lily")
	fmt.Println("插入结果", result)

	l.Insert(4, "天下第一")
	l.Range()
	fmt.Println("遍历链表")

	//清空

	l.Clear()
	t := l.Empty()
	fmt.Println("是否为空=", t)


	i,o:=l.LenthAndCap()
	fmt.Println(i,o)
}

```

