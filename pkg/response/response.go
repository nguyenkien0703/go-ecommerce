package response

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrResponseData struct {
	Code   int         `json:"code"`   // Ma status code
	Error  string      `json:"error"`  //
	Detail interface{} `json:"detail"` // Thong bao loi
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	// message == "" set msg[code]
	if message == "" {
		message = msg[code]
	}
	c.JSON(code, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
