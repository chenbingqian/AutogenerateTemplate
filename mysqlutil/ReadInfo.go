package mysqlutil

// 操作mysql.db:
// 1、引入库
// 2、连接mysql信息
// 3、操作数据：新增，修改，删除，条件查询，列表查询
import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Tableinfo struct {
	ColumnName    sql.NullString
	ColumnType    sql.NullString
	IsNullable    sql.NullString
	ColumnComment sql.NullString
	ColumnKey     sql.NullString
	ColumnDefault sql.NullString
	DataType      sql.NullString
}

var Db *sql.DB

// 入口函数
func ReadData(url string) {
	var err error
	Db, err = sql.Open("mysql", url)
	if err != nil {
		log.Println(err)
		return
	}
}

// 查询列表
func TableColumnList(url string, tableName string) ([]Tableinfo, error) {
	ReadData(url)
	sql := "select column_name columnName ,column_type columnType,is_nullable isNullable,column_comment columnComment,column_key columnKey,column_default columnDefault,data_type dataType from information_schema.COLUMNS where  TABLE_NAME=?"
	stmt, erritem := Db.Prepare(sql)
	if erritem != nil {
		return nil, erritem
	}
	rows, e := stmt.Query(tableName)
	if e != nil {
		return nil, e
	}
	tableitem := Tableinfo{}

	list := make([]Tableinfo, 0)
	for rows.Next() {
		rows.Scan(&tableitem.ColumnName, &tableitem.ColumnType, &tableitem.IsNullable, &tableitem.ColumnComment, &tableitem.ColumnKey, &tableitem.ColumnDefault, &tableitem.DataType)
		list = append(list, tableitem)
	}
	return list, nil
}
