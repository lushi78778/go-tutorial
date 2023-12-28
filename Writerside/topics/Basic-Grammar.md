# Basic Grammar

本教程的示例代码为 basic

在学习上节时，可能有些概念对初学者是难以理解的，在这里我们专门新开一节来介绍一些基础语法与数据结构。

这些内容基本上看看就可以了，有些该概念似乎是不好理解的，但这不影响读者跳过本篇继续向下学习。

注意我们的目的是使用，而不是理论精通，建议用时三十分钟，回过头来多看是必要的！

## 前节回顾

在上一节中，我们简单的学习了 Go 语言的一些基本知识，现在我们来简单整理一下。

1. Go 语言的基础组成有以下几个部分

- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

2. 基本概念

- 在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。
- 注释不会被编译，每一个包应该有相关注释。 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾。
- 标识符用来命名变量、类型等程序实体。第一个字符必须是字母或下划线而不能是数字。不能用Go 语言的关键字。
- Go 语言的字符串连接可以通过 + 实现

3. 25 个关键字或保留字

| break       | default     | func      | interface | select     | <br>
| case        | defer       | go        | map       | struct     | <br>
| chan        | else        | goto      | package   | switch     | <br>
| const       | fallthrough | if        | range     | type       | <br>
| continue    | for         | import    | return    | var        |

4. 36 个预定义标识符：

| append   | bool    | byte    | cap     | close   | complex    | complex64   | complex128  | uint16  | <br>
| copy     | false   | float32 | float64 | imag    | int        | int8        | int16       | uint32  | <br>
| int32    | int64   | iota    | len     | make    | new        | nil         | panic       | uint64  | <br>
| print    | println | real    | recover | string  | true       | uint        | uint8       | uintptr |

5. 惯常操作

```Bash
# 初始化项目
go mod init hello      

# 定义系统                                         
$env:GOOS="windows"
```

```Go
package main // 声明一个主包（包是对函数进行分组的一种方式，它由同一目录中的所有文件组成）

import "fmt"  // 导入fmt包，其中包含格式化文本的功能，包括打印到控制台。此程序包是安装Go时获得的标准库程序包之一。

func main() { 
    //实现一个主要功能，将消息打印到控制台。当您运行主程序包时，默认情况下会执行一个主函数。
    fmt.Println("Hello, World!")
}
```

```Bash
# 运行 hello.go
go run hello.go

# 添加新的模块要求到 go.sum（同步模块的依赖项）
go mod tidy

# <prefix>/<descriptive-text> eg:github.com/<project-name>
go mod init example.com/greetings 

# 初始化
go mod init example.com/greetings

# 指定本项目的依赖使用本地代码代替
go mod edit -replace example.com/greetings=../greetings

# 以 `_test.go` 结尾的文件名会告诉命令 此文件包含测试函数。
go test

# 详细测试报告
go test -v

# 编译
go build

# 发现安装路径
go list -f '{{.Target}}'
```

## 数据类型

- 布尔型：只可以是常量 true 或者 false。
- 数字类型 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
- 字符串类型:字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
- 派生类型
    - 指针类型（Pointer）
    - 数组类型
    - 结构化类型(struct)
    - Channel 类型
    - 函数类型
    - 切片类型
    - 接口类型（interface）
    - Map 类型

## 数字类型

- 基于架构的类型
    - uint8：无符号 8 位整型 (0 到 255)
    - uint16：无符号 16 位整型 (0 到 65535)
    - uint32：无符号 32 位整型 (0 到 4294967295)
    - uint64：无符号 64 位整型 (0 到 18446744073709551615)
    - int8：有符号 8 位整型 (-128 到 127)
    - int16：有符号 16 位整型 (-32768 到 32767)
    - int32：有符号 32 位整型 (-2147483648 到 2147483647)
    - int64：有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
- 浮点型
    - float32：IEEE-754 32位浮点型数
    - float64：IEEE-754 64位浮点型数
    - complex64：32 位实数和虚数
    - complex128：64 位实数和虚数
- 其他数字类型
    - byte：类似 uint8
    - rune：类似 int32
    - uint：32 或 64 位
    - int：与 uint 一样大小
    - uintptr：无符号整型，用于存放一个指针

## 基础语法

```Go
// 1 指定变量类型，如果没有初始化，则变量默认为零值

// 声明一个变量并初始化
var a = "lushi"

// 没有初始化就为零值
var b int

// bool 零值为 false
var c bool

/**
  *数值类型（包括complex64/128）为 0
  *布尔类型为 false
  *字符串为 ""（空字符串）
  *以下几种类型为 nil
  */
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口

// 2 根据值自行判定变量类型(类型推断)
var d = true

// 如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误
e := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
// 上面这种不带声明格式的只能在函数体中出现


// 显式类型定义常量
const b string = "abc"

// 隐式类型定义常量
const b = "abc"

// 多重赋值
const a, b, c = 1, false, "str" 

// 常量还可以用作枚举
const (
    Unknown = 0
    Female = 1
    Male = 2
)

// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
const (
    a = iota // ota 在 const关键字出现时将被重置为 0
    b = iota //const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
    c = iota
)

// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式
const (
    a = iota
    b
    c
)

const (
    //0 1 2 ha ha 100 100 7 8
    a = iota   // 0
    b          // 1
    c          // 2
    d = "ha"   // 独立值，iota += 1
    e          // "ha"   iota += 1
    f = 100    // iota +=1
    g          // 100  iota +=1
    h = iota   // 7,恢复计数
    i          // 8
)
```

## 值类型和引用类型

- 所有像 `int`、`float`、`bool` 和 `string` 这些基本类型都属于`值类型`，使用这些类型的变量直接指向存在内存中的值
- 当使用等号 = 将一个变量的值赋值给另一个变量时，如：`j = i`，实际上是在内存中将 i 的值进行了`拷贝`
- 通过 `&i` 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）
- 更复杂的数据通常会需要使用多个字，这些数据一般使用`引用类型`保存
- 一个`引用类型`的变量 r1 存储的是 r1 的`值所在的内存地址（数字）`，或`内存地址中第一个字所在的位置`。(指针)
- 当使用赋值语句`r2 = r1` 时，只有`引用（地址）`被复制。如果 r1 的值被改变了，那么这个值的`所有引用都会指向被修改后的内容`

## 变量赋值简短形式，使用 := 赋值操作符

- 使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
- 使用操作符 := 可以高效地创建一个新的变量，称之为`初始化声明`。
- 在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明
- 声明了一个`局部变量`却没有在相同的代码块中使用它，得到编译错误
- `全局变量`是允许声明但不使用的
- 多变量可以在同一行进行赋值（这被称为 并行 或 同时 赋值）
- 空白标识符 _ 可被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。（_ 实际上是一个只写变量，你不能得到它的值）
- 并行赋值用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)

## 运算符

```
算术运算符
假定 A 值为 10，B 值为 20。
+	相加	A + B 输出结果 30
-	相减	A - B 输出结果 -10
*	相乘	A * B 输出结果 200
/	相除	B / A 输出结果 2
%	求余	B % A 输出结果 0
++	自增	A++ 输出结果 11
--	自减	A-- 输出结果 9

关系运算符
假定 A 值为 10，B 值为 20。
==	检查两个值是否相等，如果相等返回 True 否则返回 False。	(A == B) 为 False
!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。	(A != B) 为 True
>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。	(A > B) 为 False
<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。	(A < B) 为 True
>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	(A <= B) 为 True

逻辑运算符
假定 A 值为 True，B 值为 False。
&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。	(A && B) 为 False
||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	(A || B) 为 True
!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。   !(A && B) 为 True

位运算符
假定 A 为60，B 为13
&	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。	(A & B) 结果为 12, 二进制为 0000 1100
|	按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或	(A | B) 结果为 61, 二进制为 0011 1101
^	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。	(A ^ B) 结果为 49, 二进制为 0011 0001
<<	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。	A << 2 结果为 240 ，二进制为 1111 0000
>>	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。	A >> 2 结果为 15 ，二进制为 0000 1111

赋值运算符
=	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
+=	相加后再赋值	C += A 等于 C = C + A
-=	相减后再赋值	C -= A 等于 C = C - A
*=	相乘后再赋值	C *= A 等于 C = C * A
/=	相除后再赋值	C /= A 等于 C = C / A
%=	求余后再赋值	C %= A 等于 C = C % A
<<=	左移后赋值	C <<= 2 等于 C = C << 2
>>=	右移后赋值	C >>= 2 等于 C = C >> 2
&=	按位与后赋值	C &= 2 等于 C = C & 2
^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
|=	按位或后赋值	C |= 2 等于 C = C | 2

其他运算符
&	返回变量存储地址	&a; 将给出变量的实际地址。
*	指针变量。	*a; 是一个指针变量

运算符优先级
由上至下代表优先级由高到低,可以通过使用括号来临时提升某个表达式的整体运算优先级
5	* / % << >> & &^
4	+ - | ^
3	== != < <= > >=
2	&&
1	||
```

## 语句

### 条件判断

```Go
package main

import "fmt"

func main() {

   /* 局部变量定义 */
   var a int = 100
   var b int = 200
   
   /* 判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   } else {
       /* 如果条件为 false 则执行以下语句 */
       fmt.Printf("a 不小于 20\n" )
       
       // 嵌套
       if b == 200 {
          /* if 条件语句为 true 执行 */
          fmt.Printf("a 的值为 100 ， b 的值为 200\n" )
       }
   }
   fmt.Printf("a 的值为 : %d\n", a)


   /* 定义局部变量 */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

   switch {
      case grade == "A" :
         fmt.Printf("优秀!\n" )    
      case grade == "B", grade == "C" :
         fmt.Printf("良好\n" )      
      case grade == "D" :
         fmt.Printf("及格\n" )      
      case grade == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
}
```

### select 语句

```Go
package main

import "fmt"

func main() {
  // 定义两个通道  c1 和 c2
  ch1 := make(chan string)
  ch2 := make(chan string)

  // 启动两个 goroutine，分别从两个通道中获取数据
  go func() {
    for {
      ch1 <- "from 1"
    }
  }()
  go func() {
    for {
      ch2 <- "from 2"
    }
  }()

  // 使用 select 语句非阻塞地从两个通道中获取数据
  for {
    select 语句等待两个通道的数据。如果接收到 c1 的数据，就会打印 "from 1"；如果接收到 c2 的数据，就会打印 "from 2"。
    select {
    case msg1 := <-ch1:
      fmt.Println(msg1)
    case msg2 := <-ch2:
      fmt.Println(msg2)
    default:
      // 如果两个通道都没有可用的数据，则执行这里的语句
      fmt.Println("no message received")
    }
  }
}
```

### 循环

```Go
package main

import "fmt"

func main() {

   /* ===================================================== */  
   /* 定义局部变量 */
   var i, j int

   for i=2; i < 100; i++ {
      for j=2; j <= (i/j); j++ {
         if(i%j==0) {
            break; // 如果发现因子，则不是素数
         }
      }
      if(j > (i/j)) {
         fmt.Printf("%d  是素数\n", i);
      }
   }  
   
   /* ===================================================== */  
   //0 1 2 3 4
   for i := 0; i < 10; i++ {
      if i == 5 {
          break // 当 i 等于 5 时跳出循环
      }
      fmt.Println(i)
   }
   
   /* ===================================================== */  
   /* 定义局部变量 */
   var a int = 10

   /* for 循环 */
   for a < 20 {
      if a == 15 {
         /* 跳过此次循环 */
         a = a + 1;
         continue;
      }
      fmt.Printf("a 的值为 : %d\n", a);
      a++;    
   }  
   
   /* ===================================================== */  
   /* 定义局部变量 */
   var a int = 10

   /* 循环 */
   LOOP: for a < 20 {
      if a == 15 {
         /* 跳过迭代 */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", a)
      a++    
   }  
}

```

## 函数

- 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
- 函数可以作为实参

```Go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Google", "Baidu")
   fmt.Println(a, b)
}

//创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量(匿名函数)
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
   }
}

// 方法
/* 定义结构体 */
type Circle struct {
  radius float64
}
//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}
```

## 变量

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

## 数据结构

```Go

// 以下定义了数组 balance 长度为 10 类型为 float32：
var balance [10]float32
// 初始化列表来初始化数组
var numbers = [5]int{1, 2, 3, 4, 5}
// 使用 := 简短声明语法来声明和初始化数组
numbers := [5]int{1, 2, 3, 4, 5}
//读取了数组 balance 第 10 个元素的值
var salary float32 = balance[9]

//nil 指针称为空指针。nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

var a int= 20   /* 声明实际变量 */
var ip *int        /* 声明指针变量 */
ip = &a  /* 指针变量的存储地址 */
fmt.Printf("a 变量的地址是: %x\n", &a  )
/* 指针变量的存储地址 */
fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
/* 使用指针访问值 */
fmt.Printf("*ip 变量的值: %d\n", *ip )

// 定义结构体
type Books struct {
   title string
   author string
   subject string
   book_id int
}
fmt.Println(Books{title: "Go 语言", author: "lushi", subject: "Go 教程", book_id: 356})

// Go 语言切片是对数组的抽象(动态数组)
var numbers = make([]int,3,5)
printSlice(numbers) //len=3 cap=5 slice=[0 0 0]
func printSlice(x []int){
   // 切片是可索引的，并且可以由 len() 方法获取长度。
   // 切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

// 可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]
/* 创建切片 */
numbers := []int{0,1,2,3,4,5,6,7,8}  
/* 打印子切片从索引1(包含) 到索引4(不包含)*/
fmt.Println("numbers[1:4] ==", numbers[1:4])
/* 默认下限为 0*/
fmt.Println("numbers[:3] ==", numbers[:3])
/* 默认上限为 len(s)*/
fmt.Println("numbers[4:] ==", numbers[4:])

/* 允许追加空切片 */
numbers = append(numbers, 0)
/* 向切片添加一个元素 */
numbers = append(numbers, 1)
/* 同时添加多个元素 */
numbers = append(numbers, 2,3,4)
/* 创建切片 numbers1 是之前切片的两倍容量*/
numbers1 := make([]int, len(numbers), (cap(numbers))*2)
/* 拷贝 numbers 的内容到 numbers1 */
copy(numbers1,numbers)

// 范围  range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
for key, value := range oldMap {
    newMap[key] = value
}
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
  fmt.Printf("2**%d = %d\n", i, v)
}

// Map 是一种无序的键值对的集合
// 使用字面量创建 Map
m := map[string]int{
    "apple": 1,
    "banana": 2,
    "orange": 3,
}
// 获取键值对
v1 := m["apple"]
v2, ok := m["pear"]  // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
// 修改键值对
m["apple"] = 5
// 获取 Map 的长度
len := len(m)
// 遍历 Map
for k, v := range m {
    fmt.Printf("key=%s, value=%d\n", k, v)
}
// 删除键值对
delete(m, "banana")
```

## 递归函数

斐波那契数列

```Go
package main

import "fmt"

func fibonacci(n int) int {
  if n < 2 {
   return n
  }
  return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
    var i int
    for i = 0; i < 10; i++ {
       fmt.Printf("%d\t", fibonacci(i))
    }
}
```

## 类型转换

- 不支持隐式转换类型
- 类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为： `T(value)`

```Go
var a int = 10
var b float64 = float64(a)

var str string = "10"
var num int
num, _ = strconv.Atoi(str) //strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误
```

## 接口

```Go
package main

import (
    "fmt"
)

type Phone interface {
    call() //定义了一个接口 Phone，接口里面有一个方法 call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
    // 面定义了一个 Phone 类型变量，并分别为之赋值为 NokiaPhone 和 IPhone。然后调用 call() 方法
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()

}
```

## goroutine

详细说明：https://www.yuque.com/lushi78778/blog/bluuec

```Go
package main

import (
        "fmt"
        "time"
)

func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}

func main() {
        go say("world")
        say("hello")
}
```
