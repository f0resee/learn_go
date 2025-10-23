package gormtest

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"math"
	"math/rand"
	"testing"
	"time"
)

type ConfigProcess struct {
	ID          int64  `gorm:"id"`
	NamespaceID int64  `gorm:"namespace_id"`
	ConfigItems []byte `gorm:"config_items"`
	Comment     string `gorm:"comment"`
	WorkflowID  int64  `gorm:"workflow_id"`
	Status      string `gorm:"status"`

	CreatedAt *time.Time `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
	CreatedBy string     `gorm:"created_by"`
	UpdatedBy string     `gorm:"updated_by"`
}

func (c *ConfigProcess) TableName() string {
	return "config_process"
}

func TestConfigProcessCreate(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var n = rand.Int63n(math.MaxInt32)
	for n < 0 {
		n = rand.Int63()
	}
	cp := ConfigProcess{
		NamespaceID: n,
		ConfigItems: []byte("config items" + utils.ToString(n)),
		Comment:     "create config " + utils.ToString(n),
		WorkflowID:  rand.Int63n(math.MaxInt32),
		Status:      "pending",

		CreatedBy: "jack",
		UpdatedBy: "jack",
	}

	dsn := "root:196322@tcp(127.0.0.1:3306)/testgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	} else {
		var data ConfigProcess
		res := db.Where("workflow_id = ?", cp.WorkflowID).First(&data)
		t.Log("query rows affected: ", res.RowsAffected)
		t.Log("query error: ", res.Error)
		if res.RowsAffected == 0 {
			ret := db.Create(&cp)
			t.Log("create rows affected: ", ret.RowsAffected)
			t.Log("create error: ", ret.Error)
			t.Log("cp ", cp)
		} else {
			t.Logf("[%v] already exist", data)
		}
	}
}

func TestConfigProcessQuery(t *testing.T) {
	t.Run("query by workflow id", func(t *testing.T) {
		dsn := "root:196322@tcp(127.0.0.1:3306)/testgorm?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			t.Fatal(err)
		} else {
			var data ConfigProcess
			res := db.Where("workflow_id = ?", 1641408836).First(&data)
			t.Log("query rows affected: ", res.RowsAffected)
			t.Log("query error: ", res.Error)
			if res.RowsAffected == 0 {
				t.Logf("rows affected: %v", res.RowsAffected)
			} else {
				t.Logf("find: %v", data)
			}
		}
	})

	t.Run("query by create time", func(t *testing.T) {
		dsn := "root:196322@tcp(127.0.0.1:3306)/testgorm?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			t.Fatal(err)
		} else {
			var data ConfigProcess
			res := db.Order("created_at desc").First(&data)
			t.Log("query rows affected: ", res.RowsAffected)
			t.Log("query error: ", res.Error)
			if res.RowsAffected == 0 {
				t.Logf("rows affected: %v", res.RowsAffected)
			} else {
				t.Logf("find: %v", data)
			}
		}
	})
}

func TestConfigProcessUpdate(t *testing.T) {
	t.Run("query by workflow id", func(t *testing.T) {
		dsn := "root:196322@tcp(127.0.0.1:3306)/testgorm?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			t.Fatal(err)
		} else {
			var data ConfigProcess
			res := db.Model(&ConfigProcess{}).Where("workflow_id = ?", 1641408836).Update("status", "pass")

			t.Log("query rows affected: ", res.RowsAffected)
			t.Log("query error: ", res.Error)
			if res.RowsAffected == 0 {
				t.Logf("rows affected: %v", res.RowsAffected)
			} else {
				t.Logf("find: %v", data)
			}
		}
	})
}

func TestConfigProcessDelete(t *testing.T) {
	t.Run("query by workflow id", func(t *testing.T) {
		dsn := "root:196322@tcp(127.0.0.1:3306)/testgorm?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			t.Fatal(err)
		} else {
			var data ConfigProcess
			data.NamespaceID = 335510041
			res := db.Where("namespace_id = ?", data.NamespaceID).Delete(&ConfigProcess{})

			t.Log("query rows affected: ", res.RowsAffected)
			t.Log("query error: ", res.Error)
			if res.RowsAffected == 0 {
				t.Logf("rows affected: %v", res.RowsAffected)
			} else {
				t.Logf("find: %v", data)
			}
		}
	})
}
