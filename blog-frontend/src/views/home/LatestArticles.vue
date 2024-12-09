<template>
  <div class="latest-articles">
    <h1 class="title">最新文章</h1>
    <div class="articles-grid">
      <div v-for="article in articles" :key="article.id" class="article-card">
        <div class="article-category" :class="article.category">
          {{ article.categoryName }}
        </div>
        <h2 class="article-title">{{ article.title }}</h2>
        <p class="article-excerpt">{{ article.excerpt }}</p>
        <div class="article-meta">
          <span class="publish-date">
            <i class="far fa-calendar"></i>
            {{ article.publishDate }}
          </span>
          <span class="read-count">
            <i class="far fa-eye"></i>
            {{ article.readCount }} 次阅读
          </span>
        </div>
        <router-link :to="article.url" class="read-more">
          阅读全文
          <i class="fas fa-arrow-right"></i>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      articles: [
        {
          id: 1,
          title: '深入理解TCP协议',
          excerpt: 'TCP是互联网最重要的协议之一,本文将深入剖析TCP协议的工作原理、三次握手、四次挥手等核心概念...',
          categoryName: '计算机基础',
          category: 'computer',
          publishDate: '2024-01-15',
          readCount: 1250,
          url: '/computer/tcp-deep-dive'
        },
        {
          id: 2, 
          title: 'Go语言并发编程实践',
          excerpt: '本文介绍Go语言goroutine和channel的使用方法,以及常见的并发模式和注意事项...',
          categoryName: 'Go语言相关',
          category: 'golang',
          publishDate: '2024-01-14',
          readCount: 986,
          url: '/golang/concurrency'
        },
        {
          id: 3,
          title: '动态规划解题技巧',
          excerpt: '动态规划是算法中的重要思想,本文总结常见的动态规划题型和解题方法...',
          categoryName: '算法相关',
          category: 'algorithm',
          publishDate: '2024-01-13',
          readCount: 756,
          url: '/algorithm/dynamic-programming'
        },
        {
          id: 4,
          title: '从零搭建微服务框架',
          excerpt: '本文将介绍如何使用Go语言从零开始搭建一个微服务框架,包括服务发现、负载均衡等功能...',
          categoryName: '项目相关',
          category: 'projects',
          publishDate: '2024-01-12',
          readCount: 1102,
          url: '/projects/microservice-framework'
        }
      ]
    }
  },
  methods: {
    async fetchLatestArticles() {
      try {
        const response = await fetch('/blogapi/v1/home/latest-articles')
        const data = await response.json()
        this.articles = data.articles
      } catch (error) {
        console.error('获取最新文章失败:', error)
      }
    }
  },
  created() {
    this.fetchLatestArticles()
  }
}
</script>

<style scoped>
.latest-articles {
  padding: 2rem;
}

.title {
  font-size: 1.5rem;
  color: #6d28d9;
  margin-bottom: 2rem;
  font-weight: 600;
  text-align: left;
}

.articles-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 2rem;
}

.article-card {
  background: white;
  border-radius: 1rem;
  padding: 1.5rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease;
  position: relative;
}

.article-card:hover {
  transform: translateY(-5px);
}

.article-category {
  display: inline-block;
  padding: 0.5rem 1rem;
  border-radius: 2rem;
  font-size: 0.875rem;
  font-weight: 500;
  margin-bottom: 1rem;
}

.article-category.computer {
  background: #e0f2fe;
  color: #0369a1;
}

.article-category.golang {
  background: #dcfce7;
  color: #15803d;
}

.article-category.algorithm {
  background: #fef3c7;
  color: #b45309;
}

.article-category.projects {
  background: #f3e8ff;
  color: #7e22ce;
}

.article-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 1rem;
  line-height: 1.4;
}

.article-excerpt {
  color: #4b5563;
  font-size: 0.875rem;
  line-height: 1.6;
  margin-bottom: 1rem;
}

.article-meta {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1rem;
  font-size: 0.875rem;
  color: #6b7280;
}

.read-more {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  color: #6d28d9;
  font-weight: 500;
  text-decoration: none;
  font-size: 0.875rem;
  transition: color 0.2s ease;
}

.read-more:hover {
  color: #5b21b6;
}

.read-more i {
  transition: transform 0.2s ease;
}

.read-more:hover i {
  transform: translateX(4px);
}
</style>