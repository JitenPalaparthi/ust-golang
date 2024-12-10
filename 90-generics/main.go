package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	sum1, err := addAny(10, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sum1:", sum1)
	}

	sum2, err := addAnyE(10.123, 20.123)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sum2:", sum2)
	}

	var a1, b1 uint8
	a1, b1 = 10, 12

	sum3, err := addAnyE(a1, b1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sum3:", sum3)
	}

	sum4, err := addAny(nil, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sum4:", sum4)
	}

	sum5, err := addAny(1212, 21312.123)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sum5:", sum5)
	}

	sum6 := addGenerics(123.123, 123) // -> addGenerics_1(float64,int)float64
	fmt.Println("sum6:", sum6)

	sum7 := addGenericsI(a1, 123) // addGenericsI_2(uint8,int)int
	fmt.Println("sum7:", sum7)

}

// IDT(Interface Definition Table) like VTable in c++ or rust
func addAny(a, b any) (any, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.New("a and b must be of same type")
	}
	if a == nil || b == nil {
		return nil, errors.New("a or b is nil")
	}
	sum := 0.0
	switch a.(type) {
	case int:
		sum = float64(a.(int) + b.(int))
		return sum, nil
	case uint:
		sum = float64(a.(uint) + b.(uint))
		return sum, nil
	case int8:
		sum = float64(a.(int8) + b.(int8))
		return sum, nil
	case int16:
		sum = float64(a.(int16) + b.(int16))
		return sum, nil
	case int32:
		sum = float64(a.(int32) + b.(int32))
		return sum, nil
	case int64:
		sum = float64(a.(int64) + b.(int64))
		return sum, nil
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
		return sum, nil
	case uint16:
		sum = float64(a.(uint16) + b.(uint16))
		return sum, nil
	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
		return sum, nil
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
		return sum, nil
	case float32:
		sum = float64(a.(float32) + b.(float32))
		return sum, nil
	case float64:
		sum = a.(float64) + b.(float64)
		return sum, nil
	default:
		return nil, errors.ErrUnsupported
	}
}

func addAnyE(a, b Empty) (Empty, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.New("a and b must be of same type")
	}
	if a == nil || b == nil {
		return nil, errors.New("a or b is nil")
	}
	sum := 0.0
	switch a.(type) {
	case int:
		sum = float64(a.(int) + b.(int))
		return sum, nil
	case uint:
		sum = float64(a.(uint) + b.(uint))
		return sum, nil
	case int8:
		sum = float64(a.(int8) + b.(int8))
		return sum, nil
	case int16:
		sum = float64(a.(int16) + b.(int16))
		return sum, nil
	case int32:
		sum = float64(a.(int32) + b.(int32))
		return sum, nil
	case int64:
		sum = float64(a.(int64) + b.(int64))
		return sum, nil
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
		return sum, nil
	case uint16:
		sum = float64(a.(uint16) + b.(uint16))
		return sum, nil
	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
		return sum, nil
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
		return sum, nil
	case float32:
		sum = float64(a.(float32) + b.(float32))
		return sum, nil
	case float64:
		sum = a.(float64) + b.(float64)
		return sum, nil
	default:
		return nil, errors.ErrUnsupported
	}
}

//Monomorphization

func addGenerics[T int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int32 | int64 | float32 | float64](a, b T) T {
	return a + b
}

// 0000 0000
// os.O_CREATE | os.O_APPEND
/* 0001 1000
// os.O_CREATE | os.O_APPEND | O_WRONLY
// 0001 1010

// 0644 ->
// R+W+E-> 4+2+1=>7
// R+W->6
// R->4
// 0644 (File or Directory CurrentUser Group Others)

const(
O_RDONLY int = syscall.O_RDONLY // open the file read-only.
O_WRONLY int = syscall.O_WRONLY // open the file write-only.
O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
O_APPEND int = syscall.O_APPEND // append data to the file when writing.
O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)
*/

func addGenericsI[T Generic](a, b T) T {
	return a + b
}

type Generic interface {
	int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int32 | int64 | float32 | float64
}

type Empty interface {
}

const (
	Active = iota
	_
	Pending
	Inactive
	_
	Created
)

const (
	READ = 1 << (iota * 2)
	WRITE
	EXECUTE
)

// calc function with generics
// add
// sub
// mul
// div
