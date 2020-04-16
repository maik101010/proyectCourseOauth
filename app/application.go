package app

import (
	"github.com/gin-gonic/gin"
	"github.com/maik101010/proyectCourseOauth/src/http"
	"github.com/maik101010/proyectCourseOauth/src/repository/db"
	"github.com/maik101010/proyectCourseOauth/src/repository/rest"
	"github.com/maik101010/proyectCourseOauth/src/service/access_token"
)

var(
	router = gin.Default()
)
func StartApplication()  {
	atHandler := http.NewHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))
	router.GET("oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("oauth/access_token/", atHandler.Create)
	router.Run(":8080")

}
