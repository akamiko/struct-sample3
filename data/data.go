package data

import (
	"fmt"
)

type Mydata struct {
	Num int
	Str string
}

func (m *Mydata) Print() {
	fmt.Println("test")
}
