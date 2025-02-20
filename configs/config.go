package configs

import (
	"fmt"
	"gin-gorm-example/internal/model"
	"log"

	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config 存储应用的配置项
type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	ImagePath  string
	CreateDB   bool
}

// GlobalConfig 存储全局的配置
var GlobalConfig Config
var DB *gorm.DB

// InitConfig 初始化配置
func InitConfig() {
	// 读取配置文件
	cfg, err := ini.Load("../configs/config.ini")
	if err != nil {
		log.Fatalf("Fail to read the file: %v", err)
	}

	// 加载数据库配置
	GlobalConfig.DBUser = cfg.Section("database").Key("user").String()
	GlobalConfig.DBPassword = cfg.Section("database").Key("password").String()
	GlobalConfig.DBHost = cfg.Section("database").Key("host").String()
	GlobalConfig.DBPort, _ = cfg.Section("database").Key("port").Int()
	GlobalConfig.DBName = cfg.Section("database").Key("dbname").String()

	// 加载图片路径
	GlobalConfig.ImagePath = cfg.Section("image").Key("path").String()
	GlobalConfig.CreateDB, _ = cfg.Section("database_settings").Key("createdb").Bool()
}

// InitDB 初始化数据库连接
func InitDB() error {
	// 构建 MySQL 连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GlobalConfig.DBUser,
		GlobalConfig.DBPassword,
		GlobalConfig.DBHost,
		GlobalConfig.DBPort,
		GlobalConfig.DBName)

	// 初始化数据库连接
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	fmt.Println("Database connected")

	if GlobalConfig.CreateDB {
		if err := migrateDB(); err != nil {
			return err
		}
	}
	return nil
}

func migrateDB() error {
	// 自动迁移所有模型
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("error migrating database: %v", err)
	}

	// 此处可以检查更多模型，并添加需要的迁移
	// 比如：DB.AutoMigrate(&model.OtherModel{})
	fmt.Println("Database migration completed successfully!")
	return nil
}
