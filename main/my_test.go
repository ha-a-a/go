package main

import (
	"fmt"
	"testing"
)

// 普通测试
func TestIsPalindrome(t *testing.T) {

	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}

}

// 基准测试,测量一个程序在固定工作负载下的性能
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("detartrate")
	}
}

// 示例函数，作为文档，但是也会执行
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
