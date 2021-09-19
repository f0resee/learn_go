package gormtest

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestQuickStart(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D44", Price: 106})
	db.Create(&Product{Code: "D45", Price: 107})
	//
	var res []Product
	db.First(&res)
	fmt.Println(res)
	// Read
	var product Product
	db.First(&product, "price = ?", 100)  // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(product)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Println("run here")
	// Delete - delete product
	db.Delete(&product, 1)
}
