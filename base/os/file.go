package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	/*path, err := filepath.Abs("test/files")
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println(path)
    err2 := os.MkdirAll(path, os.ModePerm)   //建立多级文件夹
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	file, _ := os.Create("test/files/name")
	defer  file.Close()
	fmt.Println(file.Read([]byte("牛逼")))
	fmt.Println(file.Name())
	file.WriteString("卧槽牛逼了")
	fmt.Println()*/
	file, err := os.Open("test/files/name")
	if err != nil {
		fmt.Println(err)
	}
	bs := make([]byte, 4,4)
	n, err2 := file.Read(bs)
	fmt.Println(err2, n, bs)

	for {
		n, err := file.Read(bs)
		fmt.Println(n)
		if n == 0 || err == io.EOF {
			fmt.Println("end....")
			break
		}
		fmt.Println(string(bs[:n]))
	}

}
