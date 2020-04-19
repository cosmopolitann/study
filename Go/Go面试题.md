# Go面试题

## 1.Go struct能不能比较？

###### 答：可以比较，也不可以比较.

* ###### 同一个struct 能不能比较.

* ###### 不相同struct能不能比较.

* ###### struct可以做map的键吗？

##### 同一个struct 能不能比较.

##### 可以比较，也不可以比较.

```go
type T1 struct {
    Name  string
    Age   int
    Arr   [2]bool
    ptr   *int
    slice []int
    map1  map[string]string
}

func main() {
    t1 := T1{
        Name:  "yxc",
        Age:   1,
        Arr:   [2]bool{true, false},
        ptr:   new(int),
        slice: []int{1, 2, 3},
        map1:  make(map[string]string, 0),
    }
    t2 := T1{
        Name:  "yxc",
        Age:   1,
        Arr:   [2]bool{true, false},
        ptr:   new(int),
        slice: []int{1, 2, 3},
        map1:  make(map[string]string, 0),
    }

    // 报错 实例不能比较 Invalid operation: t1 == t2 (operator == not defined on T1)
    // fmt.Println(t1 == t2)
    // 指针可以比较
    fmt.Println(&t1 == &t2) // false

    t3 := &T1{
        Name:  "yxc",
        Age:   1,
        Arr:   [2]bool{true, false},
        ptr:   new(int),
        slice: []int{1, 2, 3},
        map1:  make(map[string]string, 0),
    }

    t4 := &T1{
        Name:  "yxc",
        Age:   1,
        Arr:   [2]bool{true, false},
        ptr:   new(int),
        slice: []int{1, 2, 3},
        map1:  make(map[string]string, 0),
    }

    fmt.Println(t3 == t4)                  // false
    fmt.Println(reflect.DeepEqual(t3, t4)) // true
    fmt.Printf("%p, %p \n", t3, t4)        // 0xc000046050, 0xc0000460a0
    fmt.Printf("%p, %p \n", &t3, &t4)      // 0xc000006030, 0xc000006038

    // 前面加*，表示指针指向的值，即结构体实例，不能用==
    // Invalid operation: *t3 == *t4 (operator == not defined on T1)
    // fmt.Println(*t3 == *t4)

    t5 := t3
    fmt.Println(t3 == t5)                  // true
    fmt.Println(reflect.DeepEqual(t3, t5)) // true
    fmt.Printf("%p, %p \n", t3, t5)        // 0xc000046050, 0xc000046050
    fmt.Printf("%p, %p \n", &t3, &t5)      // 0xc000006030, 0xc000006040

}

```

- t1, t2是同一个struct两个赋值相同的实例，因为成员变量带有了不能比较的成员，所以只要写 == 就报错

- t3 t4 虽然能用 == ，但是本质上是比较的指针类型，*t3 == *t4 一样的一写就报错

  ##### 两个不同的struct的实例能不能比较

  ##### 可以比较，也不可以比较.

  ```go
  type T2 struct {
      Name  string
      Age   int
      Arr   [2]bool
      ptr   *int
  }
  
  type T3 struct {
      Name  string
      Age   int
      Arr   [2]bool
      ptr   *int
  }
  
  func main() {
  
      var ss1 T2
      var ss2 T3
      // Cannot use 'ss2' (type T3) as type T2 in assignment
      //ss1 = ss2
      ss3 := T2(ss2)
      fmt.Println(ss3==ss1) // true
  }
  ```

  T2和T3是不同的结构体，但可以强制转换，所以强转之后可以比较

  ```go
  type T2 struct {
      Name  string
      Age   int
      Arr   [2]bool
      ptr   *int
      map1  map[string]string
  }
  
  type T3 struct {
      Name  string
      Age   int
      Arr   [2]bool
      ptr   *int
      map1  map[string]string
  }
  
  func main() {
      var ss1 T2
      var ss2 T3
      // Cannot use 'ss2' (type T3) as type T2 in assignment
      //ss1 = ss2
      ss3 := T2(ss2)
      // Invalid operation: ss3==ss1 (operator == not defined on T2)
      // fmt.Println(ss3==ss1)   含有不可比较成员变量
  }
  ```

  ##### 可排序、可比较和不可比较

  - 可排序的数据类型有三种，Integer，Floating-point，和String
  - 可比较的数据类型除了上述三种外，还有Boolean，Complex，Pointer，Channel，Interface和Array
  - 不可比较的数据类型包括，Slice, Map, 和Function.

  ##### struct可以作为map的key吗？

  可以，也不可以.

  ```go
  type T1 struct {
      Name  string
      Age   int
      Arr   [2]bool
      ptr   *int
      slice []int
      map1  map[string]string
  }
  
  type T2 struct {
      Name  string
      Age   int
      Arr   [2]bool
      ptr   *int
  }
  
  func main() {
      n := make(map[T2]string, 0)   // 无报错
      fmt.Print(n)   // map[]
      // lnvalid map key type: the comparison operators == and != must be fully defined for key type
      // m := make(map[T1]string, 0)
      // fmt.Println(m)
  }
  ```

  struct必须是可比较的，才能作为key，否则编译时报错。

  ## 2.Map如何顺序读取？

  ###### 答：map不能顺序读取，是因为他是无序的，想要有序读取，首先的解决的问题就是，把ｋｅｙ变为有序，所以可以把key放入切片，对切片进行排序，遍历切片，通过key取值。

  https://segmentfault.com/a/1190000011873706

  ## 3.为什么遍历 Go map 是无序的？

  https://www.jianshu.com/p/2fd7064bbe44

  