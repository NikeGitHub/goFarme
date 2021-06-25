package model

//生成
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type TableColumn struct {
	Name    string
	Type    string
	Comment string
}

func GenModel(db *sql.DB) {
	rows, err := db.Query("show tables")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		tableName := ""
		if err := rows.Scan(&tableName); err != nil {
			panic(err)
		}
		if tableName == "phinxlog" {
			continue
		}
		tmpSql := "select column_name,column_type,column_comment from information_schema.columns where table_schema='smartgy' and table_name = '%s'"
		rows, err := db.Query(fmt.Sprintf(tmpSql, tableName))
		if err != nil {
			panic(err)
		}
		columns := make([]*TableColumn, 0)
		for rows.Next() {
			t := TableColumn{}
			if err := rows.Scan(&t.Name, &t.Type, &t.Comment); err != nil {
				panic(err)
			}
			columns = append(columns, &t)
		}
		genFile(tableName, columns)
	}

}

//_ =>驼峰
func Change(name string, first bool) string {
	n := []byte(name)
	if first {
		if 'a' <= n[0] && n[0] <= 'z' {
			n[0] -= 32
		}
	}
	r, err := regexp.Compile("_[a-z]")
	if err != nil {
		panic(err)
	}
	replace := r.ReplaceAllFunc(n, func(data []byte) []byte {
		data = data[1:]
		if 'a' <= data[0] && data[0] <= 'z' {
			data[0] -= 32
		}
		return data
	})
	return string(replace)
	//return strings.Replace(string(replace), "Id", "Id", -1)
}
func genFile(table string, columns []*TableColumn) {
	wd, _ := os.Getwd()
	var path string
	if strings.Contains(wd, "model") == false {
		path = wd + "/model/"
	} // 兼容脚本

	file, err := os.OpenFile(path+Change(table, false)+".go", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if os.IsExist(err) {
		log.Println(table + "存在跳过")
		return
	}
	defer file.Close()
	file.WriteString("package model\n\n")
	file.WriteString(fmt.Sprintf("type %s struct {\n", Change(table, true)))
	//	Id int32 `json:"id"` //id
	for _, column := range columns {
		if strings.Contains(column.Type, "int") {
			file.WriteString(fmt.Sprintf("\t%s int32 `json:\"%s\"` //%s\n", Change(column.Name, true), column.Name, column.Comment))
		} else if strings.Contains(column.Type, "float") || strings.Contains(column.Type, "decimal") || strings.Contains(column.Type, "double") {
			file.WriteString(fmt.Sprintf("\t%s float64 `json:\"%s\"` //%s\n", Change(column.Name, true), column.Name, column.Comment))
		} else if strings.Contains(column.Type, "varchar") || strings.Contains(column.Type, "json") {
			file.WriteString(fmt.Sprintf("\t%s string `json:\"%s\"` //%s\n", Change(column.Name, true), column.Name, column.Comment))
		} else if strings.Contains(column.Type, "text") {
			file.WriteString(fmt.Sprintf("\t%s string `json:\"%s\"` //%s\n", Change(column.Name, true), column.Name, column.Comment))
		} else {
			log.Printf("table%s Name=> %s no!!\n", table, column.Type)
		}
	}
	file.WriteString("}")
}
