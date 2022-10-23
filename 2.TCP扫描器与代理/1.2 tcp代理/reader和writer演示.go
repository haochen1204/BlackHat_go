package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// FooReader 定义了一个从标准输入(stdin)读取的io.Reader
type FooReader struct{}

// 从标准输入(stdin)读数据
func (FooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// FooWriter 定义里一个从标准输出(stdout)的io.Writer
type FooWriter struct{}

// 将数据写入标准输出
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	// 实例化reader和writer
	var (
		reader FooReader
		writer FooWriter
	)
	/*
		// 创建缓冲区以保存输入输出
		input := make([]byte, 4096)

		// 使用reader读取输入
		s, err := reader.Read(input)
		if err != nil {
			log.Fatalln("Unable to read data")
		}
		fmt.Printf("Read %d bytes from stdin\n", s)

		// 使用writer写入输出
		s, err = writer.Write(input)
		if err != nil {
			log.Fatalln("unable to write data")
		}
		fmt.Printf("Wrote %d byter to stdout\n", s)
	*/
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
