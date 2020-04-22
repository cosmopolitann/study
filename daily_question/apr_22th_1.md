# 每日一题

###### 第一题

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

###### 输出什么？

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