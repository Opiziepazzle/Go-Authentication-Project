package controllers

import (
	"github.com/gofiber/fiber/v2"
"github.com/opiziepazzle/golang-auth/initializers"
	"github.com/opiziepazzle/golang-auth/models"
	
)




func IndexController(c *fiber.Ctx) error{
	return c.Render("index", fiber.Map{

	})
}

func AboutController(c *fiber.Ctx) error{
	return c.Render("about", fiber.Map{

	})
}



//if testing on postman use this

// func UsersCommentController(c *fiber.Ctx) error{
// 	user := new(models.User)
// 	if err:= c.BodyParser(user); err !=nil {
// 		return c.Status(500).SendString(err.Error())
// 	}
// initializers.DB.Save(&user)
// return c.JSON(&user)

//  // Redirect to the root route
// //  return c.Redirect("/")
// }



// webApp use this
func UsersCommentController(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	initializers.DB.Save(&user)
	return c.SendStatus(fiber.StatusOK) // Return a success status without any content
	
}







func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	initializers.DB.Find(&user, id)
	return c.JSON(&user)	
	}
	


func GetUsers(c *fiber.Ctx) error {
var users []models.User
initializers.DB.Find(&users)
return c.JSON(&users)	
}



func SaveUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err:= c.BodyParser(user); err !=nil {
		return c.Status(500).SendString(err.Error())
	}
initializers.DB.Save(&user)
return c.JSON(&user)	
}



func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	//First function is used to first get the user from the database
	initializers.DB.First(&user, id)
//we check if the user have an email
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	initializers.DB.Delete(&user)
	return c.SendString("User is Deleted")
}


func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
    user := new(models.User)
	initializers.DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	initializers.DB.Save(&user)
	return c.JSON("User is Deleted")
}




