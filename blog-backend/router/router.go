package router

import (
	"blog-backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	bv1 := r.Group("/blogapi/v1")
	{
		bv1home := bv1.Group("/home")
		{
			bv1home.GET("/latest-articles", controller.GetLatestArticles)
			bv1home.GET("/personal-summary", controller.GetPersonalSummary)
			bv1home.GET("/message-board", controller.GetMessageBoard)
			bv1home.GET("/captcha/:uuid", controller.GetCaptchaURL)
			bv1home.GET("/register/:phone/:captcha/:uuid", controller.GetPhoneCode)
			bv1home.POST("/register", controller.PostRegister)
			bv1home.POST("/login", controller.PostLogin)
			bv1home.GET("/checkLogin", controller.CheckLogin)
		}

		bv1arti := bv1.Group("/articles")
		{
			bv1arti.POST("", controller.AuthCheck(), controller.CreateArticle)       //增
			bv1arti.PUT("/:id", controller.AuthCheck(), controller.UpdateArticle)    //改
			bv1arti.DELETE("/:id", controller.AuthCheck(), controller.DeleteArticle) //删
			bv1arti.GET("/:id", controller.GetArticleById)                           //查
			bv1arti.GET("/all", controller.GetAllArticles)                           //查
			{                                                                        //分类查询
				bv1arti.GET("/network", controller.GetCategoryById(1))
				bv1arti.GET("/algorithm", controller.GetCategoryById(2))
				bv1arti.GET("/architecture", controller.GetCategoryById(3))
				bv1arti.GET("/os", controller.GetCategoryById(4))
				bv1arti.GET("/linux", controller.GetCategoryById(5))
				bv1arti.GET("/database", controller.GetCategoryById(6))
				bv1arti.GET("/golang", controller.GetCategoryById(7))
				bv1arti.GET("/leetcode", controller.GetCategoryById(8))
				bv1arti.GET("/project", controller.GetCategoryById(9))
			}

		}
	}
}
