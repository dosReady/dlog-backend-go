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

func List(c *gin.Context) []interface{} {
	if _, exists := c.Get("body"); exists {
		// body := body.(map[string]interface{})
		conn := dao.GetConnection()

		var list []interface{}
		rows, err := conn.Model(&DlogPost{}).Select("post_seq, post_title, post_content, DATE_FORMAT(create_date, '%Y-%m-%d') AS create_date, DATE_FORMAT(update_date, '%Y-%m-%d') AS update_date").Rows()
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var post struct {
				PostSeq     string `json:"post_seq,omitempty"`
				PostTitle   string `json:"post_title,omitempty"`
				PostContent string `json:"post_contnet,omitempty"`
				CreateDate  string `json:"create_date,omitempty"`
				UpdateDate  string `json:"update_date,omitempty"`
			}
			_ = conn.ScanRows(rows, &post)
			list = append(list, post)
		}
		return list
	} else {
		return nil
	}
}
