package post

import (
	"net/http"

	postModel "github.com/dosReady/dlog/backend/models/post"
	"github.com/gin-gonic/gin"
)

func PostRegsiter(c *gin.Context) {
	if err := postModel.Register(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errmsg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": 1})
	}
}
