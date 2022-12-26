package adaptor

import (
	"net/http"
	"part2/internal/model"

	"github.com/gin-gonic/gin"
)

type DataBaseAdaptor struct {
	//storage service.UserStorage
}

func (db *DataBaseAdaptor) AddUserJSONToResponse(c *gin.Context, user *model.User) {
	c.IndentedJSON(http.StatusOK, user)
}
