package gormtest

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestGormCreate(t *testing.T)  {
	db, err := gorm.Open(sqlite.Open("./test.db"),&gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	newProduct :=Product{
		Code: "G01",
		Price: 300,
	}
	db.Create(&newProduct)
	newProduct1 :=Product{
		Code: "G02",
		Price: 301,
	}
	db.Select("Code","Price").Create(&newProduct1)
}

//func TestGorm()  {
	
//}