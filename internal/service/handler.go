package service

import (
	"net/http"
	"part2/internal/adaptor"
	"part2/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	storage *adaptor.ConcurrentUserStorage
}

func (handler *UserHandler) Get(ctx *gin.Context) {
	user := handler.storage.Get(ctx.Param("login"))
	ctx.IndentedJSON(http.StatusOK, user)
}

func (handler *UserHandler) Put(ctx *gin.Context) {
	newUser := model.User{ID: uuid.New()}
	err := ctx.BindJSON(&newUser)
	if err != nil {
		return
	}
	handler.storage.Add(&newUser)
	ctx.IndentedJSON(http.StatusOK, newUser)
}

// func (handler *UserHandler) Post(ctx *gin.Context) {
// 	var newUser *model.User
// 	if err := ctx.BindJSON(&newUser); err != nil {
// 		return
// 	}
// 	handler.storage.Update(newUser)
// 	ctx.IndentedJSON(http.StatusOK, newUser)
// }
