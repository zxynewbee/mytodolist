package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库配置 - 请根据您的实际配置修改
const (
	DBUser     = "root"
	DBPassword = "root"
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "bubble"
)

func main() {
	// 构建数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser, DBPassword, DBHost, DBPort, DBName)

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	// 检查连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	fmt.Println("成功连接到MySQL数据库")

	// 检查todos表是否存在
	var tableExists bool
	err = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ? AND table_name = 'todos'", DBName).Scan(&tableExists)
	if err != nil {
		log.Fatalf("检查表是否存在时出错: %v", err)
	}

	if tableExists {
		fmt.Println("todos表已存在，无需操作")
	} else {
		fmt.Println("创建todos表...")
		// 创建todos表（只包含三个基本字段）
		createTableSQL := `
			CREATE TABLE todos (
				id INT AUTO_INCREMENT PRIMARY KEY,
				title VARCHAR(255) NOT NULL,
				status TINYINT DEFAULT 0
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
		`

		_, err = db.Exec(createTableSQL)
		if err != nil {
			log.Fatalf("创建表失败: %v", err)
		}
		fmt.Println("成功创建todos表（包含id、title、status三个字段）")
	}

	fmt.Println("数据库表结构准备完成")
}
