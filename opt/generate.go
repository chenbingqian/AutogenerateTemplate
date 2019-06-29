package opt

import (
	"autogenerate/mysqlutil"
	"fmt"
	"strings"
)

type FileInfo struct {
	PackageName string
	FileUrl     string
	DbUrl       string
	TableNames  []string
	ReplacePrx  string
}

func Generate(info FileInfo) {

	writecontext := make([]mysqlutil.TableWriteData, 0)
	if info.TableNames != nil {
		context := mysqlutil.TableWriteData{}
		for _, tableName := range info.TableNames {
			data, _ := mysqlutil.TableColumnList(info.DbUrl, tableName)
			context.TableName, context.DataList = analysisFiledData(tableName, info, data)
			println(info.FileUrl)
			writecontext = append(writecontext, context)
		}
	}
	mysqlutil.Writefile(info.FileUrl, writecontext)
}

// 解析 table column
func analysisFiledData(tableName string, info FileInfo, data []mysqlutil.Tableinfo) (string, []string) {
	dataItem := make([]string, 0)
	var structName string
	if info.ReplacePrx != "" {
		structName = strings.Replace(tableName, info.ReplacePrx, "", 1)
	} else {
		structName = tableName
	}

	temp := strings.Split(structName, "_")
	structName = ""
	for _, str := range temp {
		// 首字母大些：截取第一个字符,转大些
		structName = structName + strings.ToUpper(str[0:1]) + str[1:len(str)]
	}
	dataItem = append(dataItem, "package "+info.PackageName)
	dataItem = append(dataItem, "")
	dataItem = append(dataItem, "type "+structName+" struct{")
	for _, item := range data {
		// 字段名
		fileItem := " '" + item.ColumnName.String + "' "
		var filed string = item.ColumnName.String
		filedSplic := strings.Split(filed, "_")
		filed = ""
		for _, filedStr := range filedSplic {
			filed = filed + strings.ToUpper(filedStr[0:1]) + filedStr[1:len(filedStr)]
		}

		// 字段类型
		dataType := typeString(item.DataType.String)
		var coumnType string = item.ColumnType.String
		// 是否为空
		nullItem := ""
		if item.IsNullable.String == "NO" {
			nullItem = " not null "
		}
		str := "	" + filed + "    " + dataType + " " + "`xorm:\"" + coumnType + " " + nullItem + fileItem + "\" json:\"" + item.ColumnName.String + "\"`"
		dataItem = append(dataItem, str)
	}
	dataItem = append(dataItem, "}")
	dataItem = append(dataItem, "\n")
	dataItem = append(dataItem, "func (self *"+structName+") TableName() string{")
	dataItem = append(dataItem, "   return \""+tableName+"\"")
	dataItem = append(dataItem, "}")
	return structName, dataItem
}

// 处理mysql字断类型
func typeString(dataType string) string {

	returnString := ""
	switch dataType {
	case "varchar":
		returnString = "string"
	case "int":
		returnString = "int64"
	case "text":
		returnString = "string"
	}
	if returnString == "" {
		fmt.Println("type not difan:" + dataType)
	}
	return returnString

}
