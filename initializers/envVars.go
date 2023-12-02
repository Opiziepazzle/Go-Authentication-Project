package initializers

import (
	"log"
	"github.com/joho/godotenv"
	
	
)
//func wahatever should balways start with capital letter
func LoadEnvVariables(){
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
}
}