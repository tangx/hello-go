package databases

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	gorm.Model
	Username string `gorm: username`
	Password string `gorm: username`
}

func Test_Migrate(t *testing.T) {

	db, err := gorm.Open("mysql", "root:V69Gz4wsysY@(172.17.0.38:3306)/hello_gin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

}

func Test_DB(t *testing.T) {

	db, err := gorm.Open("mysql", "root:V69Gz4wsysY@(172.17.0.38:3306)/hello_gin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	p := Person{
		Username: "tangxin2",
		Password: "GjFYvhtSDeGEVCDrD",
	}

	ok := db.NewRecord(p)
	fmt.Printf("before create ok = %t\n", ok)
	db.Create(&p)
	ok = db.NewRecord(p)
	fmt.Printf("after ok = %t\n", ok)

}
