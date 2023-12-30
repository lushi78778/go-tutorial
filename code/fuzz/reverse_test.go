package main

import (
    "testing"
    "unicode/utf8"
)

// TestReverse是一个标准的单元测试函数
func TestReverse(t *testing.T) {
    // 测试用例集合
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    // 遍历测试用例
    for _, tc := range testcases {
        // 调用Reverse函数获取反转的字符串
        rev := Reverse(tc.in)
        // 检查结果是否符合预期
        if rev != tc.want {
            t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
}

// // FuzzReverse是一个模糊测试函数
// func FuzzReverse(f *testing.F) {
//     // 模糊测试用例集合
//     testcases := []string{"Hello, world", " ", "!12345"}
//     // 添加模糊测试用例到种子集合中
//     for _, tc := range testcases {
//         f.Add(tc)  // 使用f.Add提供种子集合
//     }
//     // 执行模糊测试
//     // f.Fuzz(func(t *testing.T, orig string) {
//     //     // 调用Reverse函数获取反转的字符串
//     //     rev := Reverse(orig)
//     //     // 再次调用Reverse函数获取双重反转的字符串
//     //     doubleRev := Reverse(rev)
//     //     // 检查原始字符串和双重反转后的字符串是否相等
//     //     if orig != doubleRev {
//     //         t.Errorf("Before: %q, after: %q", orig, doubleRev)
//     //     }
//     //     // 检查原始字符串是否是有效的UTF-8字符串
//     //     // 并且反转后的字符串是否无效
//     //     if utf8.ValidString(orig) && !utf8.ValidString(rev) {
//     //         t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
//     //     }
//     // })

    // 执行模糊测试
    f.Fuzz(func(t *testing.T, orig string) {
        // 调用Reverse函数获取反转的字符串
        rev := Reverse(orig)
        // 再次调用Reverse函数获取双重反转的字符串
        doubleRev := Reverse(rev)

        // 打印原始字符串、反转后的字符串和双重反转后的字符串的字符数
        t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))

        // 检查原始字符串和双重反转后的字符串是否相等
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }

        // 检查原始字符串是否是有效的UTF-8字符串
        // 并且反转后的字符串是否无效
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}