package main

import (
"fmt"
  "os"

	"github.com/gofiber/fiber/v2"
	
    "github.com/opiziepazzle/golang-auth/initializers"
     "github.com/opiziepazzle/golang-auth/controllers"
    "github.com/opiziepazzle/golang-auth/middleware"

    // "github.com/opiziepazzle/golang-auth/handler" 
)

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDatabase()
    initializers.SyncDB()
}





func main() {



    
    
   
    app := fiber.New()
  
    
    
    
    
//set up app 
    // app := fiber.New(fiber.Config{
    //     Views: engine,
    // })


// Apply the RequireAuth middleware globally
//app.Use(middleware.RequireAuth)
    
    
    // Configure App
 app.Static("/", "./public")

   
 
 //routes
 Routes(app)









// Debug print statement
fmt.Println("Starting the application...")

    //start app
    app.Listen(":" + os.Getenv("PORT"))




    
}



func Routes(app *fiber.App){
	app.Get("/", controllers.IndexController)
	app.Get("/about", controllers.AboutController)
	




    app.Get("/user/:id", controllers.GetUser)
    app.Get("/users", controllers.GetUsers)
    app.Post("/user", controllers.SaveUser)
    app.Delete("/user/:id", controllers.DeleteUser)
    app.Put("/user/:id", controllers.UpdateUser)
    app.Post("/signup", controllers.Signup)
    app.Post("/login", controllers.Login)
    app.Get("/validate", middleware.RequireAuth, controllers.Validate)
    app.Post("/upload", controllers.UploadFile)
}
