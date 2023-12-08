package controllers

import (
	"fmt"
	"log"
	"net/http"

	Dto "gitub.com/RomainC75/biblio/api/dto"
	Service "gitub.com/RomainC75/biblio/api/services"
	"gitub.com/RomainC75/biblio/pkg/errors"
	utils "gitub.com/RomainC75/biblio/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService Service.UserServiceInterface
}

func NewAuthCtrl() *AuthController {
	return &AuthController{
		userService: Service.NewUserSrv(),
	}
}

func (controller *AuthController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "register form"})
	// html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
	// 	"title": "Register",
	// })

}

func (controller *AuthController) HandleRegister(c *gin.Context) {
	fmt.Println("--> POST new register")
	// validate request
	var request Dto.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		// 	// see also ShouldBindJSON()
		errors.Init()
		errors.SetFromErrors(err)

		fmt.Println("-> err:  ", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "body needs a valid email + password"})
		return
	}

	if controller.userService.CheckIfUserExists(request.Email) {
		// see also ShouldBindJSON()
		errors.Init()
		errors.Add("email", "email address already used")

		c.JSON(http.StatusConflict, gin.H{"message": "user already exists"})
		return
	}

	// create the user
	user, err := controller.userService.Create(request)

	// check if any error at user creation
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// //redirect to homepage
	log.Printf("user successfully created with email %s \n", user.Email)
	c.JSON(http.StatusOK, gin.H{"message": "user created", "user": user})
}

func (controller *AuthController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "register form"})
	// html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
	// 	"title": "Login",
	// })
}

func (controller *AuthController) HandleLogin(c *gin.Context) {

	var request Dto.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		// see also ShouldBindJSON()

		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	user, err := controller.userService.HandleUserLogin(request)
	if err != nil {
		// see also ShouldBindJSON()

		c.Redirect(http.StatusFound, "/login")
		return
	}
	fmt.Println("found user : ", user)

	token, err := utils.Generate(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (controller *AuthController) HandleWhoAmI(c *gin.Context) {
	id, exists := c.Get("user_id")
	if exists {
		data := map[string]interface{}{
			"id": id,
			"email": c.GetString("user_email"),
		}
		c.JSON(http.StatusOK, gin.H{
			"user": data,
		})
	}
}
