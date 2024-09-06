package api

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/dilshodforever/nasiya-savdo/api/handler"
	"github.com/dilshodforever/nasiya-savdo/api/middleware"
	_ "github.com/dilshodforever/nasiya-savdo/docs"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title auth service API
// @version 1.0
// @description auth service API
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy: ", err)
		panic(err)
	}

	r := gin.Default()

	r.Use(middleware.NewAuth(e))

	u := r.Group("/user")
	{
		u.POST("/register", h.Register)
		u.PUT("/update/:id", h.UpdateUser)
		u.DELETE("/delete/:id", h.DeleteUser)
		u.GET("/getall", h.GetAllUser)
		u.GET("/getbyid/:id", h.GetbyIdUser)
		u.POST("/login", h.LoginUser)
	}

	{
		u.POST("/change-password", h.ChangePassword)
		u.POST("/forgot-password", h.ForgotPassword)
		u.POST("/reset-password", h.ResetPassword)
	}

	{
		u.GET("/get_profil", h.GetProfil)
		u.PUT("/update_profil", h.UpdateProfil)
		u.DELETE("/delete_profil", h.DeleteProfil)
	}

	{
		u.GET("/refresh-token", h.RefreshToken)
	}

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	return r
}
