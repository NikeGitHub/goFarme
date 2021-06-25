package model

//生成
import (
	"database/sql"
	"os"
	"testing"
)

func TestChange(t *testing.T) {
	a := Change("a_user", true)
	t.Log(a)
}

func TestGenModel(t *testing.T) {
	db, err := sql.Open("mysql", "epuser:epuser@123-New@tcp(106.75.138.97:3306)/smartgy?charset=utf8")
	//db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/smartgy?charset=utf8")
	if err != nil {
		panic(err)
	}
	GenModel(db)
}

func TestOpenFile(t *testing.T) {
	f, err := os.OpenFile("accessLog.go", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		t.Log(os.IsExist(err), os.IsNotExist(err))
		return
	}
	defer f.Close()
	f.WriteString("123")
	t.Log(f.Name())
}
