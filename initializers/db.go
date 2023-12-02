package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"fmt"
	"github.com/opiziepazzle/golang-auth/models"

)

var DB *gorm.DB

func ConnectToDatabase(){
	var err error
	dsn := os.Getenv("DB_URL")

  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
	fmt.Println(err.Error())
	panic("Unable to Connect to Database")
  }
}

func SyncDB(){
	DB.AutoMigrate(&models.User{})
	//DB.Migrator().DropTable(&models.User{})
    
}


