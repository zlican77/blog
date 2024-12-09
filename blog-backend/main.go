package main

import (
	"blog-backend/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 配置Redis作为session存储
	//配置session:初始化容器， 设置加密密钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//设置cookie key，加密成session
	r.Use(sessions.Sessions("mysession", store))

	// 配置CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"}
	config.AllowHeaders = []string{
		"Content-Type",
		"Content-Length",
		"Authorization",
		"Accept",
		"Origin",
		"Referer",
		"User-Agent",
		"X-Requested-With",
	}
	config.MaxAge = 86400

	r.Use(cors.New(config))

	router.SetupRouter(r)

	r.Run(":8000")
}
