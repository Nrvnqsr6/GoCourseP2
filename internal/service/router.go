package service

import (
	"net/http"
	"part2/internal/adaptor"
	"part2/internal/model"

	"github.com/gin-gonic/gin"
)

func SetupRouter(storage adaptor.ConcurrentUserStorage) *gin.Engine {
	r := gin.Default()
	r.GET("/user/:login", func(ctx *gin.Context) {
		user := storage.Get(ctx.Param("login"))
		ctx.IndentedJSON(http.StatusOK, user)
	})
	r.POST("/user/:login", func(ctx *gin.Context) {
		var newUser *model.User
		if err := ctx.BindJSON(&newUser); err != nil {
			return
		}
		storage.Update(newUser)
		ctx.IndentedJSON(http.StatusOK, newUser)
	})
	r.POST("/user/login", func(ctx *gin.Context) {

	})
	r.PUT("/user/:login", func(ctx *gin.Context) {
		var newUser *model.User
		if err := ctx.BindJSON(&newUser); err != nil {
			return
		}
		storage.Add(newUser)
		ctx.IndentedJSON(http.StatusOK, newUser)
	})
	return r
}

func AddHandlerGet(endpoint string, handlerFunc gin.HandlerFunc) {

}

// func (router *Router) Run(port string) {
// 	router.router.Run(port)
// }
