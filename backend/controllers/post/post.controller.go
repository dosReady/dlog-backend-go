package post

import (
	"net/http"

	postModel "github.com/dosReady/dlog/backend/models/post"
	"github.com/gin-gonic/gin"
)

func PostRegsiter(c *gin.Context) {
	postModel.Register(c)
	c.JSON(http.StatusOK, gin.H{"name": "성공"})
}
