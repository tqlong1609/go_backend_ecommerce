package response

const (
	SuccessCode      = 200
	FailCode         = 500
	UnauthorizedCode = 401
	BadRequestCode   = 400
)

var HttpCodeMessage = map[int]string{
	SuccessCode:      "success",
	FailCode:         "fail",
	UnauthorizedCode: "unauthorized",
	BadRequestCode:   "bad request",
}

func GetHttpCodeMessage(code int) string {
	return HttpCodeMessage[code]
}
