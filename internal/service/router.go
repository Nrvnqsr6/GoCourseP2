package service

import (
	"net/http"
	"part2/internal/action"
	"part2/internal/adaptor"
	"part2/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var WRONGUSERSTRUCTUSER = []byte("Wrong structure of user type")
var TOKENERROR = []byte("Error with token header")
var UNVALIDTOKEN = []byte("Error with token header")

func SetupRouter(storage *adaptor.ConcurrentUserStorage) *gin.Engine {
	r := gin.Default()
	r.GET("/user/:login", func(ctx *gin.Context) {
		user := storage.Get(ctx.Param("login"))
		ctx.IndentedJSON(http.StatusOK, user)
	})
	r.POST("/user", func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Token")
		if tokenString == "" {
			ctx.Writer.Write(TOKENERROR)
			return
		}
		token, _ := jwt.ParseWithClaims(tokenString, &action.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return action.KEY, nil
		})
		claims, ok := token.Claims.(*action.CustomClaims)
		if !ok || !token.Valid {
			ctx.Writer.Write(UNVALIDTOKEN)
			return
		}
		user := storage.Get(claims.Login)
		userUpdates := model.User{}
		if err := ctx.BindJSON(&userUpdates); err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}
		storage.Update(user, &userUpdates)
		ctx.IndentedJSON(http.StatusOK, user)
	})
	r.POST("/user/login", func(ctx *gin.Context) {
		var user *model.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}
		err := storage.Authorize(user.Login, user.Password)
		if err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}

		token, err := action.GenerateJWT(user.Login)
		ctx.Writer.Header().Set("Token", token)
		ctx.IndentedJSON(http.StatusAccepted, token)

	})
	r.PUT("/user", func(ctx *gin.Context) {
		newUser := model.User{ID: uuid.New()}
		err := ctx.BindJSON(&newUser)
		if err != nil {
			ctx.Writer.Write(WRONGUSERSTRUCTUSER)
			return
		}
		err = storage.Add(&newUser)
		if err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}
		ctx.IndentedJSON(http.StatusOK, newUser)
	})
	return r
}

func AddHandlerGet(endpoint string, handlerFunc gin.HandlerFunc) {

}

// func (router *Router) Run(port string) {
// 	router.router.Run(port)
// }
