package post

import (
	"fmt"
	"time"

	"github.com/dosReady/dlog/backend/models/common"
	"github.com/dosReady/dlog/backend/modules/dao"
	"github.com/gin-gonic/gin"
)

type DlogPost struct {
	PostSeq     string `json:"post_seq"`
	PostTitle   string `json:"post_title,omitempty"`
	PostContent string `json:"post_contnet,omitempty"`
	common.Base
}

func Register(c *gin.Context) error {
	if body, exists := c.Get("body"); exists {
		body := body.(map[string]interface{})

		conn := dao.GetConnection()
		if err := conn.Create(DlogPost{
			PostTitle:   body["post_title"].(string),
			PostContent: body["post_content"].(string),
			Base: common.Base{
				CreateDate: time.Now(),
				UpdateDate: time.Now(),
			},
		}).Error; err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
