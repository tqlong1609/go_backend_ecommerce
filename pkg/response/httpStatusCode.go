package response

const (
	SuccessCode = 200
	FailCode    = 500
)

var HttpCodeMessage = map[int]string{
	SuccessCode: "success",
	FailCode:    "fail",
}

func GetHttpCodeMessage(code int) string {
	return HttpCodeMessage[code]
}
