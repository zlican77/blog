<template>
    <div class="article-list">
      <div class="category-header">
        <h1 class="category-title">{{ categoryName }}</h1>
        <div class="category-stats">
          <span class="total-articles">
            <i class="fas fa-book"></i>
            共 {{ articles.length }} 篇文章
          </span>
          <span class="total-views">
            <i class="fas fa-eye"></i>
            总浏览 {{ totalViews }} 次
          </span>
        </div>
      </div>
  
      <div class="articles">
        <div v-for="article in articles" :key="article.ArticleId" class="article-card">
          <div class="article-meta">
            <span class="create-time">
              <i class="far fa-calendar-plus"></i>
              发布于 {{ article.CreateTime }}
            </span>
            <span class="update-time">
              <i class="far fa-calendar-check"></i>
              更新于 {{ article.UpdateTime }}
            </span>
            <span class="view-count">
              <i class="far fa-eye"></i>
              {{ article.ViewCount }} 次浏览
            </span>
          </div>
  
          <h2 class="article-title">{{ article.Title }}</h2>
          <p class="article-excerpt">{{ article.Excerpt }}</p>
          
          <div class="article-footer">
            <div class="tags">
              <span v-for="tag in article.Tags" :key="tag" class="tag">
                {{ tag.TagName }}
              </span>
            </div>
            <router-link :to="article.Url" class="read-more">
              阅读全文
              <i class="fas fa-arrow-right"></i>
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      categoryName: {
        type: String,
        required: true
      },
      articles: {
        type: Array,
        default: () => []
      }
    },
    computed: {
      totalViews() {
        return this.articles.reduce((sum, article) => sum + article.ViewCount, 0) //计算总浏览量
      }
    }
  }
  </script>
  
  <style scoped>
  .article-list {
    color: var(--text-primary);
  }
  
  .category-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
    padding-bottom: 1rem;
    border-bottom: 2px solid #e5e7eb;
  }
  
  .category-title {
    font-size: 1.8rem;
    font-weight: 600;
    color: #6d28d9;
    letter-spacing: -0.5px;
  }
  
  .category-stats {
    display: flex;
    gap: 1.5rem;
    color: #6b7280;
    font-size: 0.95rem;
  }
  
  .category-stats i {
    margin-right: 0.5rem;
    color: #6d28d9;
  }
  
  .articles {
    display: grid;
    gap: 1.5rem;
  }
  
  .article-card {
    background: white;
    border-radius: 1rem;
    padding: 1.5rem;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
    transition: all 0.3s ease;
  }
  
  .article-card:hover {
    transform: translateY(-3px);
    box-shadow: 0 8px 12px rgba(0, 0, 0, 0.1);
  }
  
  .article-meta {
    display: flex;
    gap: 1.5rem;
    font-size: 0.85rem;
    color: #6b7280;
    margin-bottom: 1rem;
  }
  
  .article-meta i {
    margin-right: 0.35rem;
    color: #6d28d9;
  }
  
  .article-title {
    font-size: 1.35rem;
    font-weight: 600;
    color: #1f2937;
    margin-bottom: 0.75rem;
    line-height: 1.4;
  }
  
  .article-excerpt {
    color: #4b5563;
    line-height: 1.6;
    margin-bottom: 1.5rem;
    font-size: 0.95rem;
  }
  
  .article-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .tags {
    display: flex;
    gap: 0.5rem;
  }
  
  .tag {
    background: #f3e8ff;
    color: #6d28d9;
    padding: 0.25rem 0.75rem;
    border-radius: 1rem;
    font-size: 0.85rem;
    font-weight: 500;
  }
  
  .read-more {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    color: #6d28d9;
    text-decoration: none;
    font-size: 0.95rem;
    font-weight: 500;
    transition: all 0.2s ease;
  }
  
  .read-more:hover {
    color: #5b21b6;
  }
  
  .read-more:hover i {
    transform: translateX(4px);
  }
  
  .read-more i {
    transition: transform 0.2s ease;
  }
  </style>