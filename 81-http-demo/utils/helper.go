package utils

import (
	"context"
	"log"
	"os"
)

var (
	ChFileWrite chan []byte
	FileName    string
)

func init() {
	ChFileWrite = make(chan []byte)
	//go Write()
}

func Write(ctx context.Context) {
	f, err := os.OpenFile(FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err.Error())
	}
	//defer f.Close()

	for {
		select {
		case data, ok := <-ChFileWrite:
			if ok {
				_, err = f.Write(data)
				if err != nil {
					log.Println(err.Error())
				}
			} else {
				log.Println("cloing the file")
				f.Close()
			}

		case <-ctx.Done():
			log.Println("cloing the file")
			f.Close()
			return
		}
	}
}

// syscall.O_WRONLY 0000 0001
// syscall.O_APPEND 0000 0101
// syscall.O_CREATE 0001 0101
// 1111 1111 // this is bitwise operation numbers are not real numbers , just to explain

func WriteToFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
