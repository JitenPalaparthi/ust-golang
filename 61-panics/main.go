package main

import (
	"fmt"
	"os"
)

func main() {
	//fatal("I am exsting due to some error")
	//var num *int
	//ptr := square1(num)
	//fmt.Println("sq:", ptr)
	///var num int = 1

	//fmt.Println(100 / num)

	fmt.Println("hello world, I am getting executed here?")

	sq := square2(10001)
	fmt.Println(sq)
}

func square1(num *int) *int {
	var sq *int
	*sq = (*num * *num)
	return sq
}

func square2(num uint16) uint16 {
	// v := math.MaxInt16
	// var r int16 = num * num
	// fmt.Println(r, v)
	// if r > v {
	// 	panic("result is more than its range")
	// }
	// return (num * num)

	if num > 10000 {
		panic("result is more than its range")
	}
	return (num * num)
}

func fatal(msg ...any) {
	fmt.Println(msg...)
	os.Exit(1)
}

// panic is a build in stuff that , the error is so severe that the program cannot
// run normally

// error
// panic  -> There is a recovery for panic
// fatal  -> ther eis no recovery mechanisam for fatal.fatal is not a construct but just os.exit(1) with some  print statement

// create your own error interface
// impllement it on a simple function may be on a struct like rest
