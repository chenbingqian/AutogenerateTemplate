package main

import (
	"autogenerate/opt"
)

func main() {
	fileinfo := opt.FileInfo{}
	// 链接信息
	fileinfo.DbUrl = "root:123456@tcp(localhost:3306)/kiki_product"
	// 存放路径
	fileinfo.FileUrl = "/Users/Bing/Documents/gocoding/CommodityService/entity"
	// 包名
	fileinfo.PackageName = "entity"
	// 需要自动生成的表名称
	fileinfo.TableNames = []string{"kiki_attribute"}
	//用于格式化的表头
	fileinfo.ReplacePrx = "kiki_"
	opt.Generate(fileinfo)

}
