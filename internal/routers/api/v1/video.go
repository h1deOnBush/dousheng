package v1

import "github.com/gin-gonic/gin"

type Video struct {

}

func NewVideo() Video {
	return Video{}
}

func (v Video) Publish(c *gin.Context) {

}

func (v Video) PublishList(c *gin.Context) {

}

func (v Video) Feed(c *gin.Context) {

}
