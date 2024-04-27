package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	defer file.Close()
	var content []byte
	var tmp = make([]byte, 128)
	//循环读取文件内存只到文件结束
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取结束")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

func readByBufio() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Println(line)
	}
}

func readByIoutil() {
	content, err := os.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Println(string(content))
}

func writeFile() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()
	str := "你好golang\r\n"
	file.Write([]byte(str))
	file.WriteString("直接写入的字符串数据")
}

func writeByBufio() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		//先将数据写入缓冲区
		writer.WriteString("你好golang\r\n")
	}
	//刷新缓冲区
	writer.Flush()
}

func writeByIoutil() {
	str := "hello golang"
	err := os.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err: ", err)
		return
	}
}

func fileRename() {
	err := os.Rename("./test.txt", "./test")
	if err != nil {
		fmt.Println(err)
	}
}

func copyFile(dst, src string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func cpyFile(dst, src string) error {
	source, _ := os.Open(src)
	destination, _ := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0666)
	buf := make([]byte, 128)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

func makeDir() {
	err := os.Mkdir("./abc", 0666)
	if err != nil {
		fmt.Println(err)
	}
}

// 创建多级目录
func makeDirs() {
	err := os.MkdirAll("dir2/dir2", 0666)
	if err != nil {
		fmt.Println(err)
	}
}

// 删除目录和文件
func removeAll(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// readFile()
	// readByBufio()
	// readByIoutil()
	// writeFile()
	// writeByBufio()
	// writeByIoutil()
	// fileRename()
	// copyFile("./text", "./test")
	// cpyFile("./text", "./test")
	// makeDir()
	// makeDirs()
	removeAll("dir2")
}
