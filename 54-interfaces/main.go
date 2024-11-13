package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	fmt.Println("printing on stdout")

	fr := &FileWriter{FileName: "data.txt"}
	fmt.Fprintln(fr, "Hello World")

	//fs := &FileScanner{FileName: "readdata.txt"}
	//Read(p []byte) (n int, err error)
	// var d string
	// _, err := fmt.Fscan(fs, &d)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(d)
	// }

	f, err := os.Open("abcd.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	nums := make([]int, 0)

	for {
		var num int
		time.Sleep(time.Second * 5)
		_, err := fmt.Fscan(f, &num)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				// if err.Error() == "EOF" {
				// 	break
				// } else {
				fmt.Println(err)
				return
			}

		}

		nums = append(nums, num)

		//fmt.Println("Read num", num)
	}
	f.Close()

	fmt.Println(nums)
}

type FileWriter struct {
	FileName string
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	if fw == nil {
		return 0, errors.New("nil FileWriter")
	}
	f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	return f.Write(p)
}

type FileScanner struct {
	FileName string
	//Bytes    []byte
}

// func (fs *FileScanner) Read(p []byte) (n int, err error) {
// 	if fs == nil {
// 		return 0, errors.New("nil FileScanner")
// 	}
// 	f, err := os.Open(fs.FileName)

// 	if err != nil {
// 		return 0, err
// 	}

// 	defer f.Close()
// 	for {
// 		_, err := f.Read(p)
// 		if err != nil {
// 			return 0, nil
// 		}
// 	}
// }

func (fs *FileScanner) Read() func(p []byte) (n int, err error) {

	return func(p []byte) (n int, err error) {
		return 0, nil
	}
}
