package main

import (
	"fmt"
	//"log"

	"github.com/akamiko/struct-sample2/book"
	_ "github.com/akamiko/struct-sample2/logger"
	_ "github.com/akamiko/struct-sample2/sample"
)

/*
type mydata struct {
	num int
	str string
}
*/
func main() {
	/*
		var x data.Mydata
		x.Num = 20
		x.Str = "anything"
		fmt.Printf("x.Num=%d, x.Str=%s\n", x.Num, x.Str)
		data.Print(x)
	*/
	/*
		var log logger.Logger
		log.Info()
	*/
	b := book.NewBook("The World")
	//fmt.Printf("%+v\n", b)
	fmt.Println(b.GetTitle())
	//log.Info()
	/*
		var p sample.Person
		//p = sample.Person{Id: 1, Name: "test", Reading: "call me"}
		p.Greet()
	*/
	/*
		var x sample.Sample
		x.X = 1
		x.Y = 2
		fmt.Print(x.X)
		p := sample.Sample{1, 2}
		sample.Sample.Print()
	*/
	/*
		var x mydata
		x.num = 10
		x.str = "something"
		fmt.Printf("x.num=%d, x.str=%s\n", x.num, x.str)
	*/
}
