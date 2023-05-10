package config

import (
	"MiniProject_Rahman/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "123Rahman",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "manage",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Siswa{})
	DB.AutoMigrate(&models.Guru{})
	DB.AutoMigrate(&models.Kelas{})
	DB.AutoMigrate(&models.Mapel{})
	DB.AutoMigrate(&models.Soal{})
}
