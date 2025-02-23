package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusRouter struct {
}

// init router
func (p *PrometheusRouter) InitRouter(Router *gin.Engine) {
	// prometheusRouter := Router.Group("/prometheus")
	// prometheusRouter.Use(middlewares.PrometheusMiddleware())
	// {
	// 	prometheusRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// }
	/*
			Router.GET("/metrics", ...): Đây là cách định nghĩa một route trong Gin. Khi người dùng gửi yêu cầu GET tới /metrics, handler sẽ được gọi để xử lý yêu cầu đó.

			gin.WrapH(promhttp.Handler()):

			promhttp.Handler() là một handler được cung cấp bởi thư viện Prometheus Go. Nó tạo ra một HTTP handler để xuất ra các metrics của ứng dụng dưới dạng một endpoint.
			gin.WrapH(...) là một phương thức trong Gin, giúp bọc handler chuẩn (từ promhttp.Handler()) thành một handler có thể sử dụng trong Gin, vì Gin yêu cầu các handler phải có kiểu gin.HandlerFunc.
		Khi người dùng truy cập /metrics, handler này sẽ trả về các thông tin liên quan đến các chỉ số mà ứng dụng đã thu thập, chẳng hạn như số lượng yêu cầu, thời gian xử lý, lỗi, v.v. Các chỉ số này sẽ được Prometheus sử dụng để giám sát và phân tích hiệu suất của hệ thống.
	*/
	Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
