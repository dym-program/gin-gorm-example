package configs

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gin-gorm-example/internal/model"
)

var (
	DB           *gorm.DB
	GlobalConfig Config
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	ImagePath  string // 如："/static/images/"
	UploadDir  string // 如："./uploads/"
	CreateDB   bool
}

func InitConfig() {
	cfg, err := ini.Load("../configs/config.ini")
	if err != nil {
		log.Fatalf("Fail to read config.ini: %v", err)
	}

	GlobalConfig = Config{
		DBUser:     cfg.Section("database").Key("user").String(),
		DBPassword: cfg.Section("database").Key("password").String(),
		DBHost:     cfg.Section("database").Key("host").String(),
		DBPort:     cfg.Section("database").Key("port").MustInt(3306),
		DBName:     cfg.Section("database").Key("dbname").String(),
		ImagePath:  cfg.Section("image").Key("path").String(),
		UploadDir:  cfg.Section("image").Key("upload_dir").String(),
		CreateDB:   cfg.Section("database_settings").Key("createdb").MustBool(false),
	}
}

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GlobalConfig.DBUser,
		GlobalConfig.DBPassword,
		GlobalConfig.DBHost,
		GlobalConfig.DBPort,
		GlobalConfig.DBName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if GlobalConfig.CreateDB {
		// 自动迁移 User 和 Image 模型
		err = DB.AutoMigrate(&model.User{}, &model.Image{})
		if err != nil {
			return err
		}
	}
	return nil
}
