package main

import (
	"fmt"
)

// uint 溢出示例
func main() {
	// 1. 加法溢出:uint8 最大值是 255,再加 1 会回绕到 0
	var a uint8 = 255
	a++
	fmt.Printf("uint8 255 + 1 = %d (溢出回绕到 0)\n", a)

	// 2. 乘法溢出
	var b uint8 = 16
	b *= 16 // 16*16 = 256,超出 uint8 范围
	fmt.Printf("uint8 16 * 16 = %d (256 溢出后为 0)\n", b)

	// 3. 减法下溢:无符号数 0 减 1 变成最大值
	var c uint8 = 0
	c--
	fmt.Printf("uint8 0 - 1 = %d (下溢变成最大值 255)\n", c)

	// 4. uint 的零值以及循环中的溢出陷阱
	var d uint8 = 0
	for i := d - 1; i >= 0; i-- { // i 是 uint8,永远不会小于 0,死循环!
		fmt.Printf("i = %d\n", i)
		if i > 250 { // 防止真的死循环,跑几步就退出
			fmt.Println("uint8  i >= 0 恒成立,这是死循环陷阱,手动退出")
			break
		}
	}
}
