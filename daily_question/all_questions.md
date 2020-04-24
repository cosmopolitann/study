# All Questions

###### 1.第一题

```go
package main
import "fmt"
func main() {
	slice := []int{0, 1, 2, 3}
	//myMap := make(map[int]*int)
   fmt.Println("这是slice的地址",&slice)
	for index, value := range slice {
		myMap[index] = &value
		fmt.Printf("%d% d\n",index,&value)

	}
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
```







###### 输出：

```go
map[0]=3
map[1]=3
map[2]=3
map[3]=3
```

###### 为什么？

```go
这是slice的地址 &[0 1 2 3]
0 824634941480
1 824634941480
2 824634941480
3 824634941480

```

###### 指针指向了 同一个地址，所以遍历出来的时候 都是3.



###### 2.第二题

```go
package main
import (
     "fmt"
 )
 
func main() {
     defer_call()
}

func defer_call() {
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印后") }()

    panic("触发异常")
}


```

###### 







###### 3.第三题

```go
package	main

import "fmt"

	func main() {
		a := []int{7, 8, 9}
		fmt.Printf("%+v\n", a)
		ap(a)
		fmt.Printf("%+v\n", a)
		app(a)
		fmt.Printf("%+v\n", a)

	}

	func ap(a []int) {
		//a[0]=2

		a = append(a, 10)
		fmt.Println("ap=a",a)
	}

	func app(a []int) {
		a[0] = 1
	}

```











###### 输出什么。

```go
[7,8,9]
[7,8,9]
[1,8,9]
```

###### 因为 append 导致底层数组重新分配内存了，ap 中的 a 这个slice 的底层数组和外面的不是一个，并没有改变外面的。











###### 4.第四题

