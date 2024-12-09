package controller

import (
	"blog-backend/config"
	"blog-backend/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 查
func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	var article model.Articles

	// 开启事务
	tx := config.DB.Begin()

	// 先检查文章是否存在（不带status条件）
	result := tx.Where("article_id = ?", id).First(&article)
	if result.Error != nil {
		tx.Rollback()
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error": "文章不存在",
				"code":  404,
			})
		} else {
			ctx.JSON(500, gin.H{
				"error": "查询文章失败",
				"code":  500,
			})
		}
		return
	}
	// 检查文章状态
	if article.Status != 1 {
		tx.Rollback()
		ctx.JSON(403, gin.H{
			"error": "文章已下架或待审核",
			"code":  403,
		})
		return
	}

	// 加载关联数据
	if err := tx.Preload("Tags").Preload("Category").First(&article, id).Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{
			"error": "加载文章详情失败",
			"code":  500,
		})
		return
	}

	// 更新浏览量
	if err := tx.Model(&article).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).
		Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{
			"error": "更新浏览量失败",
			"code":  500,
		})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{
			"error": "事务提交失败",
			"code":  500,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"article": article,
		"message": "获取文章成功",
		"code":    200,
	})
}
func GetCategoryById(n int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var articles []model.Articles

		// 获取分页参数
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
		offset := (page - 1) * pageSize

		// 先获取总数
		var total int64
		config.DB.Model(&model.Articles{}).
			Where("category_id = ? AND status = ?", n, 1).
			Count(&total)

		// 使用 Preload 预加载关联数据
		result := config.DB.
			Preload("Tags").
			Preload("Category").
			Where("category_id = ? AND status = ?", n, 1).
			Order("create_time DESC").
			Offset(offset).
			Limit(pageSize).
			Find(&articles)

		if result.Error != nil {
			ctx.JSON(500, gin.H{
				"error": "获取文章列表失败",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"articles":    articles,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		})
	}
}
func GetAllArticles(ctx *gin.Context) {
	var articles []model.Articles
	result := config.DB.
		Preload("Tags").
		Preload("Category").
		Where("status = ?", 1).
		Order("create_time DESC").
		Find(&articles)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "获取文章列表失败"})
		return
	}

	ctx.JSON(200, gin.H{"articles": articles})
}

// 增
func CreateArticle(ctx *gin.Context) {
	// 1. 解析请求数据
	var articleData struct {
		Title      string    `json:"title"`
		Excerpt    string    `json:"excerpt"`
		Content    string    `json:"content"`
		Category   string    `json:"category"`
		Tags       []string  `json:"tags"`
		CreateTime time.Time `json:"createTime"`
		UpdateTime time.Time `json:"updateTime"`
	}

	if err := ctx.BindJSON(&articleData); err != nil {
		ctx.JSON(400, gin.H{"error": "请求数据格式错误"})
		return
	}

	// 开启事务
	tx := config.DB.Begin()

	// 2. 获取或创建分类
	var category model.Category
	result := tx.Where("category_name = ?", articleData.Category).FirstOrCreate(&category, model.Category{
		CategoryName: articleData.Category,
	})
	if result.Error != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{"error": "分类处理失败"})
		return
	}
	// 获取最后一篇文章的ID
	var lastArticle model.Articles
	result = tx.Order("article_id desc").First(&lastArticle)
	var nextId int
	if result.Error != nil {
		// 如果没有文章,从1开始
		nextId = 1
	} else {
		nextId = lastArticle.ArticleId + 1
	}

	// 3. 创建文章
	article := model.Articles{
		Title:      articleData.Title,
		Excerpt:    articleData.Excerpt,
		Content:    articleData.Content,
		CategoryId: category.CategoryId,
		CreateTime: articleData.CreateTime,
		UpdateTime: articleData.UpdateTime,
		Status:     1,
		ViewCount:  0,
		Url:        "/articles/" + strconv.Itoa(nextId),
	}

	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{"error": "文章创建失败"})
		return
	}

	// 4. 处理标签
	var tags []model.Tags
	for _, tagName := range articleData.Tags {
		var tag model.Tags
		result := tx.Where("tag_name = ?", tagName).FirstOrCreate(&tag, model.Tags{
			TagName: tagName,
		})
		if result.Error != nil {
			tx.Rollback()
			ctx.JSON(500, gin.H{"error": "标签处理失败"})
			return
		}
		tags = append(tags, tag)
	}

	// 5. 创建文章-标签关联
	for _, tag := range tags {
		if err := tx.Create(&model.ArticleTags{
			ArticleId: article.ArticleId,
			TagId:     tag.TagId,
		}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(500, gin.H{"error": "文章标签关联失败"})
			return
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{"error": "事务提交失败"})
		return
	}

	ctx.JSON(200, gin.H{
		"ok":        true,
		"message":   "文章创建成功",
		"articleId": article.ArticleId,
	})
}

// 改
func UpdateArticle(ctx *gin.Context) {
	var articleData struct {
		ArticleId  int       `json:"articleId"`
		Title      string    `json:"title"`
		Excerpt    string    `json:"excerpt"`
		Content    string    `json:"content"`
		CategoryId int       `json:"categoryId"`
		Status     int       `json:"status"`
		Url        string    `json:"url"`
		ViewCount  int       `json:"viewCount"`
		CreateTime time.Time `json:"createTime"`
		UpdateTime time.Time `json:"updateTime"`
		Category   struct {
			CategoryId   int    `json:"categoryId"`
			CategoryName string `json:"categoryName"`
		} `json:"category"`
		Tags []struct {
			TagId   int    `json:"tagId"`
			TagName string `json:"tagName"`
		} `json:"tags"`
	}

	if err := ctx.BindJSON(&articleData); err != nil {
		ctx.JSON(400, gin.H{
			"error": "无效的请求数据",
			"code":  400,
		})
		return
	}

	tx := config.DB.Begin()

	if err := tx.Model(&model.Articles{}).Where("article_id = ?", articleData.ArticleId).Updates(map[string]interface{}{
		"title":       articleData.Title,
		"excerpt":     articleData.Excerpt,
		"content":     articleData.Content,
		"update_time": articleData.UpdateTime,
	}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{
			"error": "更新文章失败",
			"code":  500,
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{
			"error": "事务提交失败",
			"code":  500,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "文章更新成功",
		"code":    200,
	})
}

// 删
func DeleteArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	// 开启事务
	tx := config.DB.Begin()

	// 先删除文章标签关联
	if err := tx.Where("article_id = ?", articleId).Delete(&model.ArticleTags{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{"error": "删除文章标签关联失败"})
		return
	}

	// 删除文章
	if err := tx.Where("article_id = ?", articleId).Delete(&model.Articles{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{"error": "删除文章失败"})
		return
	}

	// 提交事务
	tx.Commit()

	ctx.JSON(200, gin.H{"message": "文章删除成功"})
}
