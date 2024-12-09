<template>
    <article-list 
      category-name="《项目实战》" 
      :articles="articles"
    />
  </template>
  
  <script>
  import ArticleList from '@/components/ArticleList.vue'
  import { blogapi } from '../blogapi/blogapi';
  
  export default {
    components: {
      ArticleList
    },
    data() {
      return {
        articles: []
      }
    },
    async created() {
      try {
        const response = await blogapi.getProjectArticles()
        //更新时间格式
        response.data.articles.forEach(article => {
          article.CreateTime = new Date(article.CreateTime).toLocaleString()
          article.UpdateTime = new Date(article.UpdateTime).toLocaleString()
        })
        this.articles = response.data.articles
      } catch (error) {
        console.error('获取文章列表失败:', error)
      }
    }
  }
  </script>