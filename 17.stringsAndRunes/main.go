package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	// 字符串等同与[]byte
	fmt.Println("len:", len(s)) //放回该字符串的原始字节的长度

	// 索引到字符串中会在每个索引处生成原始字节值,此循环生成构成 s 中码位的所有字节的十六进制值
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	//要计算一个字符串中有多少个符文，我们可以使用 utf8 包。请注意，RuneCountInString 的运行时间取决于字符串的大小，因为它必须按顺序解码每个 UTF-8 rune。一些泰语字符由可以跨越多个字节的 UTF-8 码位表示，因此此计数的结果可能会令人惊讶
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// 范围循环专门处理字符串，并解码每个 rune 及其在字符串中的偏移量

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// 也可以通过使用 utf8 来实现相同的迭代。DecodeRuneInString 函数
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
