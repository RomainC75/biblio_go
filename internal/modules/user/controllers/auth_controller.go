package controllers

import (
	"net/http"

	UserService "gitub.com/RomainC75/biblio/internal/modules/user/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "register form"})
	// html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
	// 	"title": "Register",
	// })

}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// validate request
	// var request auth.RegisterRequest
	// if err := c.ShouldBind(&request); err != nil {
	// 	// see also ShouldBindJSON()
	// 	errors.Init()
	// 	errors.SetFromErrors(err)
	// 	sessions.Set(c, "errors", converters.MapToString(errors.Get()))

	// 	old.Init()
	// 	old.Set(c)
	// 	sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

	// 	c.Redirect(http.StatusFound, "/register")
	// 	return
	// }

	// if controller.userService.CheckIfUserExists(request.Email) {
	// 	// see also ShouldBindJSON()
	// 	errors.Init()
	// 	errors.Add("email", "email address already used")
	// 	sessions.Set(c, "errors", converters.MapToString(errors.Get()))

	// 	old.Init()
	// 	old.Set(c)
	// 	sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

	// 	c.Redirect(http.StatusFound, "/register")
	// 	return
	// }

	// create the user
	// user, err := controller.userService.Create(request)

	// // check if any error at user creation
	// if err != nil {
	// 	c.Redirect(http.StatusFound, "/register")
	// 	return
	// }

	// sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	// //redirect to homepage
	// log.Printf("user successfully created with name %s \n", user.Name)
	// c.Redirect(http.StatusFound, "/")

}

func (controller *Controller) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "register form"})
	// html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
	// 	"title": "Login",
	// })
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	// validate request
	// var request auth.LoginRequest
	// if err := c.ShouldBind(&request); err != nil {
	// 	// see also ShouldBindJSON()
	// 	errors.Init()
	// 	errors.SetFromErrors(err)
	// 	sessions.Set(c, "errors", converters.MapToString(errors.Get()))

	// 	old.Init()
	// 	old.Set(c)
	// 	sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

	// 	c.Redirect(http.StatusFound, "/login")
	// 	return
	// }

	// user, err := controller.userService.HandleUserLogin(request)
	// if err != nil {
	// 	// see also ShouldBindJSON()
	// 	errors.Init()
	// 	errors.Add("email", err.Error())
	// 	sessions.Set(c, "errors", converters.MapToString(errors.Get()))

	// 	old.Init()
	// 	old.Set(c)
	// 	sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

	// 	c.Redirect(http.StatusFound, "/login")
	// 	return
	// }

	// sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	// //redirect to homepage
	// log.Printf("user successfully logged with name %s \n", user.Name)
	// c.Redirect(http.StatusFound, "/")

	c.JSON(http.StatusOK, gin.H{"message": "register form"})
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	// sessions.Remove(c, "auth")
	// c.Redirect(http.StatusFound, "/")
	c.JSON(http.StatusOK, gin.H{"message": "register form"})
}
