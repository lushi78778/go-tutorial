package main

import (
    // "errors"
    "fmt"
    // "unicode/utf8"
)

// // Reverse函数接受一个字符串并返回其反转版本
// func Reverse(s string) string {
//     // 将字符串转换为字节切片以便进行修改
//     b := []byte(s)
//     // 使用双指针法反转字节切片中的元素
//     for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
//         b[i], b[j] = b[j], b[i]
//     }
//     // 将字节切片转换回字符串并返回
//     return string(b)
// }

// Reverse函数接受一个字符串并返回其反转版本
func Reverse(s string) string {
    // if !utf8.ValidString(s) {
    //     return s, errors.New("input is not valid UTF-8")
    // }

    // 将字符串转换为rune切片以便进行修改
    r := []rune(s)
    fmt.Printf("runes: %q\n", r)
    // 使用双指针法反转rune切片中的元素
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    // 将rune切片转换回字符串并返回
    return string(r)
}


func main() {
    // 输入字符串
    input := "The quick brown fox jumped over the lazy dog"
    // 调用Reverse函数获取反转的字符串
    rev := Reverse(input)
    // 再次调用Reverse函数获取双重反转的字符串
    doubleRev := Reverse(rev)
    
    // 打印原始字符串
    fmt.Printf("original: %q\n", input)
    // 打印反转后的字符串
    fmt.Printf("reversed: %q\n", rev)
    // 打印双重反转后的字符串
    fmt.Printf("reversed again: %q\n", doubleRev)
}
