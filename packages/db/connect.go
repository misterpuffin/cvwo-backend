package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"log"
	"server/packages/config"
)


func Connect() (*gorm.DB, error) {
	user := config.Config[config.MYSQL_USER]
	password := config.Config[config.MYSQL_PASSWORD]
	database := config.Config[config.MYSQL_DB]
	// host := config.Config[config.MYSQL_SERVER_HOST]

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, database)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Println("Connection Failed to Open")
		return nil, err
	} else { 
		log.Println("Connection Established")
		return db, err
	}


// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//   	if err != nil {
//     	panic("failed to connect database")
//   	}

//   // Migrate the schema
//   db.AutoMigrate(&Product{})

//   // Create
//   db.Create(&Product{Code: "D42", Price: 100})

//   // Read
//   var product Product
//   db.First(&product, 1) // find product with integer primary key
//   db.First(&product, "code = ?", "D42") // find product with code D42

//   // Update - update product's price to 200
//   db.Model(&product).Update("Price", 200)
//   // Update - update multiple fields
//   db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
//   db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

//   // Delete - delete product
//   db.Delete(&product, 1)
}