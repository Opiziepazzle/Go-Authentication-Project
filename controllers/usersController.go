package controllers

import(
	 "github.com/gofiber/fiber/v2"



	


	 "github.com/opiziepazzle/golang-auth/initializers"
	 "github.com/opiziepazzle/golang-auth/models"
	
	"golang.org/x/crypto/bcrypt"
	
	 "strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
    "os"
)


func Signup(c *fiber.Ctx) error {
	// Get the email/password of req body
	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("Failed to read body")
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.Status(400).SendString("Failed to hash password")
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return c.Status(500).SendString("Failed to create user")
	}

	// Respond
	return c.Status(200).JSON(fiber.Map{})
}





// func Signup(c *fiber.Ctx) error {
// 	// Get the email/password of the request body
// 	var body struct {
// 		Email    string
// 		Password string
// 	}

// 	if err := c.BodyParser(&body); err != nil {
// 		return c.Status(400).SendString("Failed to read body")
// 	}

// 	// Retrieve the uploaded file
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return c.Status(400).SendString("Upload failed")
// 	}

// 	fileBytes, err := file.Open()
// 	if err != nil {
// 		return c.Status(500).SendString("Failed to process the file")
// 	}
// 	defer fileBytes.Close()

// 	data, err := ioutil.ReadAll(fileBytes)
// 	if err != nil {
// 		return c.Status(500).SendString("Failed to read the file")
// 	}

// 	// Hash the password
// 	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
// 	if err != nil {
// 		return c.Status(400).SendString("Failed to hash password")
// 	}

// 	// Create the user
// 	user := models.User{
// 		Email:    body.Email,
// 		Password: string(hash),
// 		File:     data,
// 	}
// 	result := initializers.DB.Create(&user)

// 	if result.Error != nil {
// 		return c.Status(500).SendString("Failed to create user")
// 	}

// 	// Respond
// 	return c.Status(200).JSON(fiber.Map{})
// }







func Login(c *fiber.Ctx) error {
	// Get the email/password of req body
	var body struct {
		Email    string
		Password string
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read body",
		})
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Compare sent-in password with user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to authenticate",
		})
	}

	// Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	

	// Sign and get the complete encoded token as a string using secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	
	// Send the token back
	cookie := fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: false,
		Secure:   true,
		SameSite: "Lax",
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}





func Validate(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": user,
	})
}


 func UploadFile(c *fiber.Ctx) error {
	// Parse the form file field
	file, err := c.FormFile("file")
	if err != nil {
	   // Handle the error
	   return err
	}
 
	// Save the file to a destination
	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
	   // Handle the error
	   return err
	}
 
	// File uploaded successfully
	return c.SendString("File uploaded")
 }
 





















