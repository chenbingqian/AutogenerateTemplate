package mysqlutil

// 文件操作：创建文件，将数据写入txt文件

import (
	"bufio" // 缓存io
	"fmt"
	"os"
)

// 判断文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

type TableWriteData struct {
	TableName string
	DataList  []string
}

// 写文件
func Writefile(url string, writecontext []TableWriteData) {
	var wfile *os.File
	var err error
	// 创建文件夹目录
	os.MkdirAll(url, os.ModePerm)

	for _, dataItem := range writecontext {
		filename := url + "/" + dataItem.TableName + ".go"

		if checkFileIsExist(filename) {
			fmt.Println("文件已存在，删除文件：")
			os.Remove(filename)
		}
		// 创建文件
		fmt.Println("creaet file: ")
		wfile, err = os.Create(filename)

		fmt.Println("file url:", filename)

		if err != nil {
			fmt.Println("error : ", err)
		}

		w := bufio.NewWriter(wfile)

		for _, item := range dataItem.DataList {
			w.WriteString(item)
			w.WriteString("\t\n")
		}
		w.Flush()
	}

	wfile.Close()

}
