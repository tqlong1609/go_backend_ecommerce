package response

import "github.com/gin-gonic/gin"

type GetResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(200, GetResponse{
		Code:    code,
		Message: GetHttpCodeMessage(code),
		Data:    data,
	})
}

func FailResponse(c *gin.Context, code int) {
	c.JSON(401, GetResponse{
		Code:    code,
		Message: GetHttpCodeMessage(code),
		Data:    nil,
	})
}
