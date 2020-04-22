# 每日一题

###### 3.

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

