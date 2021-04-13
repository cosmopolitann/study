# Go 语言 goto 语句

Go 语言的 goto 语句可以无条件地转移到过程中指定的行。

goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。

但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。

### 语法

goto 语法格式如下：

```
goto label;
..
.
label: statement;
```

在变量 a 等于 15 的时候跳过本次循环并回到循环的开始语句 LOOP 处：

## 实例

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		y:=i+1
		fmt.Println("外面的",i)
			for j := 0; j < 8; j++ {

				if y == 2 {
					fmt.Println("要跳出循环了~~~1",i,j)
					goto s
					fmt.Println("要跳出循环了~~~2")
				}
		}

	s:
	}
}

```

//  当满足 条件的时候    y==2  时  会直接 从 s 的地方 调到 下面的 s 的地方 



需要 注意的地方  就是  goto  s  下面的语句 不会执行