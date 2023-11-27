package say

import "fmt"

// SayHi -.
func SayHi() {
	fmt.Println("Hi!")
}

// 如果被其他包引用，首先会执行 init()
func init() {
	fmt.Println("init func execute")
}
