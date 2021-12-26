package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	ProgressFormat1 = "\b\b%d%%"
	ProgressFormat2 = "\b\b\b%d%%"
	ProgressFormat3 = "\b\b\b%d%%"
)

func progress(ch chan int64) {
	ft := ProgressFormat1
	for c := range ch {
		if c >= 10 && c < 100 {
			ft = ProgressFormat2
		} else if c >= 100 {
			ft = ProgressFormat3
		}
		fmt.Printf(ft, c)
	}
}

var ch = make(chan int64)

func main() {

	f, err := os.Open("./LICENSE")
	if err != nil {
		panic(err)
	}
	fileInfo, _ := f.Stat()

	fmt.Printf("读取进度: 0%%")

	defer close(ch)

	go progress(ch)

	var sum = make([]byte, 0)

	for {
		b := make([]byte, 1024) // 读取文件要存的字节
		n, err := f.Read(b)
		if err != nil && err != io.EOF {
			log.Fatal(err.Error())
			return
		}
		if n == 0 {
			break
		}

		sum = append(sum, b...)

		time.Sleep(time.Second)

		go func() {
			sleight := int64(len(sum)*100) / fileInfo.Size()
			if sleight > 100 {
				sleight = 100
			}
			ch <- sleight
		}()
	}

}
