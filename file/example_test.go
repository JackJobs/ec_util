package file_test

import (
	"fmt"
	"os"
)

//由于test会自动删除创建的文件，因此放到一个测试来执行
func Example() {
	/*目录操作*/
	//创建目录
	os.Mkdir("temp", 0777)
	//创建多级子目录
	os.MkdirAll("temp/test", 0777)
	//删除目录，当目录下有文件或者其他目录时会出错
	err := os.Remove("temp")
	if err != nil {
		fmt.Println(err)
	}
	//删除多级子目录
	os.RemoveAll("temp")

	/*文件操作*/
	//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的
	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//写入byte类型的信息到文件
	file.Write([]byte("hello byte\r\n"))
	//写入string信息到文件
	file.WriteString("hello string\r\n")
	//只读方式打开文件
	readFile, err := os.Open("temp.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024)
	for {
		//读取数据到byte中
		n, _ := readFile.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

	// Output:
	// remove temp: The directory is not empty.
	// hello byte
	// hello string

}


