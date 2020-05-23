# Go实现数据结构线性表—顺序表

##### 代码

```go
package main

import "fmt"

// 数据结构之线性表--顺序表
type List struct {
    Len      int            //线性表长度
    Capacity int            // 表容量
    Prt      *[]interface{} // 指向线性表空间指针
}

// 初始化
func (l *List) ListInit(capacity int) {
    l.Capacity = capacity
    l.Len = 0
    m := make([]interface{}, capacity)
    l.Prt = &m
}

// 判空
func (l *List) ListIsEmpty() bool {
    if l.Len == 0 {
        return true
    } else {
        return false
    }
}

// 判满
func (l *List) ListIsFul() bool {
    if l.Len == l.Capacity {
        return true
    } else {
        return false
    }
}

// 根据下标Get元素
func (l *List) ListGet(index int) (interface{}, bool) {
    if index < 0 || index > l.Len {
        return nil, false
    } else {
        return (*l.Prt)[index], true
    }
}

// 根据传入的值，返回第一个匹配的元素下标
func (l *List) ListLocal(elem interface{}) (int, bool) {
    for i, _ := range *l.Prt {
        if elem == (*l.Prt)[i] {
            return i, true
        }
    }
    return -1, false
}

// 寻找元素的前驱（当前元素的前一个元素）
func (l *List) ListElemPre(elem interface{}) (interface{}, bool) {
    i, _ := l.ListLocal(elem)
    // 顺序表中不存在该元素，或者元素为第一个元素，无前驱元素
    if i == -1 || i == 0 {
        return nil, false
    } else {
        pre := (*l.Prt)[i-1]
        return pre, true
    }
}

// 寻找元素的后驱（当前元素的后一个元素）
func (l *List) ListElemNext(elem interface{}) (interface{}, bool) {
    i, _ := l.ListLocal(elem)
    // 顺序表中不存在该元素，或者元素为最后一个元素，无后驱元素
    if i == -1 || i == l.Len-1 {
        return nil, false
    } else {
        N := (*l.Prt)[i+1]
        return N, true
    }
}

// 插入元素,index为插入的位置，elem为插入值
func (l *List) ListInsert(index int, elem interface{}) bool {
    // 判断下标有效性，以及表是否满
    if index < 0 || index > l.Capacity || l.ListIsFul() {
        return false
    } else {
        // 先将index位置元素以及之后的元素后移一位
        for i := l.Len - 1; i >= index; i-- {
            (*l.Prt)[i+1] = (*l.Prt)[i]
        }
        // 插入元素
        (*l.Prt)[index] = elem
        l.Len++
        return true
    }
}

// 删除元素
func (l *List) ListDelete(index int) bool {
    // 判断下标有效性，以及表是否空
    if index < 0 || index > l.Capacity || l.ListIsEmpty() {
        return false
    } else {
        // 注意边界
        for i := index; i < l.Len-1; i++ {
            (*l.Prt)[i] = (*l.Prt)[i+1]
        }
        l.Len--
        return true
    }
}

// 遍历
func (l *List) ListTraverse() {
    for i := 0; i < l.Len; i++ {
        fmt.Println((*l.Prt)[i])
    }
}

// 清空
func (l *List) ListClear() {
    l.Len = 0
    // 指针为空
    l.Prt=nil
}

func main() {
    var li List
    li.ListInit(4)
    // true
    fmt.Println(li.ListIsEmpty())
    // false
    fmt.Println(li.ListIsFul())
    // 定义一个Struct类型
    type s struct {
        name string
        age  int
    }
    student1 := s{name: "abc", age: 10}
    student2 := s{name: "efg", age: 10}
    li.ListInsert(0, student1)
    li.ListInsert(1, student2)
    // false
    fmt.Println(li.ListIsEmpty())
    li.ListInsert(2, 1000)
    li.ListInsert(3, "GOGO")
    // {abc 10}
    // {efg 10}
    // 1000
    // GoGO
    li.ListTraverse()
    // true
    fmt.Println(li.ListIsFul())
    // false,已满插入失败
    fmt.Println(li.ListInsert(4, "jjj"))
    li.ListDelete(2)
    // {abc 10}
    // {efg 10}
    // GoGO
    li.ListTraverse()
    el, _ := li.ListGet(1)
    // {efg 10}
    fmt.Println(el)
    b, b1 := li.ListLocal(student2)
    // 1 true
    fmt.Println(b, b1)
    n1, n2 := li.ListElemNext(student2)
    // GOGO true
    fmt.Println(n1, n2)
    p1, p2 := li.ListElemPre("GOGO")
    // {efg 10} true
    fmt.Println(p1, p2)
    li.ListClear()
    li.ListTraverse()
}
```

