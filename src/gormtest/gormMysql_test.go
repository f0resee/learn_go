package gormtest

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)
type User struct {
	ID  uint `gorm:"column:user_id;primaryKey"`
	Name string `gorm:"column:user_name"`
}

func (User)TableName() string {
	return "user"
}
func TestGormMysql(t *testing.T) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:196322@tcp(127.0.0.1:3306)/testgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var user User
	if err !=nil {
		t.Log("open failed")
	}else{
		db.First(&user,1)
		t.Log(user)
	}
}
