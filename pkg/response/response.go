package response

import "github.com/gin-gonic/gin"

type GetResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, GetResponse{
		Code:    code,
		Message: GetHttpCodeMessage(code),
		Data:    data,
	})
}

func FailResponse(c *gin.Context, code int, msg ...string) {
	message := GetHttpCodeMessage(code)
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(code, GetResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
