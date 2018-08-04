//App project main.go
package main

import (
	"fmt"
	"os"
	"vcomp"
	"vrdbms"
)

func main() {
	host, _ := os.Hostname()
	fmt.Println("Running on host : ", host)
	//bytes := getText()
	//symDict := vcomp.Count(bytes)

	//testCompression()
	//vcomp.BitReadWriteTest()
}

func testCompression() {
	fileName := "/home/venugopal/Documents/Code/GoData/sample.txt"
	file, err := os.Open(fileName)
	if err == nil {
		var b []byte = make([]byte, 310, os.Getpagesize())
		fmt.Println("Size:", len(b), " Capacity:", cap(b))
		size, err := file.Read(b)
		fmt.Println("Opened file & no of bytes read :", size)
		if err == nil {
			compressed := vcomp.Hcompress(b)
			vcomp.Hdecode(compressed)
		} else {
			fmt.Println("Error in opening file", fileName)
		}

	} else {
		fmt.Println("Error in opening file", fileName)
	}
}

func getText() []byte {
	fileName := "/home/venugopal/Documents/Code/GoData/sample.txt"
	file, err := os.Open(fileName)
	if err == nil {
		var b []byte = make([]byte, 310, os.Getpagesize())
		fmt.Println("Size:", len(b), " Capacity:", cap(b))
		size, _ := file.Read(b)
		fmt.Println("Number of bytes read", size)
		return b
	} else {
		fmt.Println("Error in opening file", fileName)
		return nil
	}
	return nil
}

func tstreadwrite() {
	sysInfo()
	for i := 0; i < 5; i++ {
		testWrite(i)
	}
	for i := 0; i < 5; i++ {
		testRead(i)
	}
}

func testWrite(pageNum int) {
	page := new(vrdbms.Page)
	page.Num = pageNum
	pageSize := os.Getpagesize()
	bytes := make([]byte, 0, pageSize)
	for i := 0; i < pageSize; i++ {
		b := byte(i / 255)
		bytes = append(bytes, b)
	}
	fmt.Println("Slice Size is : ", len(bytes))
	page.Bytes = bytes
	vrdbms.Write(page)
}

func testRead(pageNum int) {
	page := vrdbms.Read(pageNum)
	fmt.Printf("Page number is : ", page.Num)
}

func sysInfo() {
	host, _ := os.Hostname()
	fmt.Println("Running on :", host)
	fmt.Println("Page size : ", os.Getpagesize())
}
