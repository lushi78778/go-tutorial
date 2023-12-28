# Generics

本教程的示例代码为 generics

介绍泛型。建议用时十分钟。

```Bash
mkdir generics
cd generics

# 初始化项目
go mod init example/generics
```

现在有一个需求，是求两个数的和，一共有两组，分别是int64 和 float64类型


```Go
//简单实现

func main() {
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}
```

但是，对于我来说，写两个函数 `SumInts` `SumFloats`实在是太麻烦了，要是话有别的类型，不得再搞一边？

于是

```Go
func main() {
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Generic Sums: %v and %v\n",
        SumIntsOrFloats[string, int64](ints),
        SumIntsOrFloats[string, float64](floats))

    //可以类型推断哦~ 更简单了
    fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
        SumIntsOrFloats(ints),
        SumIntsOrFloats(floats))


}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

这样更加的方便了，但是 `SumIntsOrFloats` 函数的参数表达依旧很复杂 `V int64 | float64`,还可以再简化

```Go
package main

import "fmt"

// 本质上是个类型约束 不用再写int64 | float64
type Number interface { // 声明要用作类型约束的接口类型。Number
    int64 | float64 // 声明接口内部的并集。int64 float64
}

func main() {
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Generic Sums with Constraint: %v and %v\n",
        SumNumbers(ints),
        SumNumbers(floats))
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

这次我使用了 `Number` 来代替 `int64 | float64`，抽象出了一个 `Number`接口，定义了具体的数的种类

可以看出来，泛型就是抽象的艺术，让代码更加的简洁优雅。