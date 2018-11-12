package controller

import (
	"crypto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unstd/workbench/common"
	"github.com/unstd/commons/stringutils"
	"io"
	"net/http"
)

// Md5 md5(data) -> response
func Md5(c *gin.Context) {
	source := c.Request.FormValue("data")
	if stringutils.IsEmpty(source) {
		res := common.NewErrorResponse(1, "data is empty")
		c.JSON(http.StatusOK, res)
		return
	}
	h := crypto.MD5.New()
	io.WriteString(h, source)
	result := fmt.Sprintf("%x", h.Sum(nil))
	res := common.NewSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}
