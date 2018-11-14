package logger

import (
	"fmt"

	_ "github.com/sirupsen/logrus"
)

// 構造体
type Logger struct {
	//Id int
}

/*
func NewLogger() *Logger {
	fmt.Println("construct")
	return &Logger{1}
}*/

// レシーバ関数
func (l *Logger) Info() {
	fmt.Println("log write!!")
}
