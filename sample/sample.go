package sample

import "fmt"

// 構造体
type Person struct {
	Id      int
	Name    string
	Reading string
}

// レシーバ関数
func (p *Person) Greet() {
	fmt.Println("こんにちは")
}
