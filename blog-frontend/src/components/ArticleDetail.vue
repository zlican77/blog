<template>
  <div class="article-detail" v-if="article">
    <div class="article-header">
      <h1 class="article-title">{{ article.Title }}</h1>
      
      <div class="article-meta">
        <span class="meta-item">
          <i class="far fa-calendar-plus"></i>
          发布于 {{ article.CreateTime }}
        </span>
        <span class="meta-item">
          <i class="far fa-calendar-check"></i>
          更新于 {{ article.UpdateTime }}
        </span>
        <span class="meta-item">
          <i class="far fa-eye"></i>
          {{ article.ViewCount }} 次浏览
        </span>
      </div>

      <div class="article-tags">
        <span v-for="tag in article.Tags" :key="tag.TagId" class="tag">
          {{ tag.TagName }}
        </span>
      </div>
    </div>

    <div class="article-content markdown-body" v-html="htmlContent" ref="contentRef"></div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { marked } from 'marked'
import Prism from 'prismjs'
import 'prismjs/themes/prism-okaidia.css'
import 'prismjs/components/prism-javascript'
import 'prismjs/components/prism-typescript'
import 'prismjs/components/prism-python'
import 'prismjs/components/prism-go'
import 'prismjs/components/prism-bash'
import 'prismjs/components/prism-sql'
import ClipboardJS from 'clipboard'
import { blogapi } from '../blogapi/blogapi'

export default {
  setup() {
    const route = useRoute()
    const article = ref(null)
    const htmlContent = ref('')
    const contentRef = ref(null)

    // 配置marked
    marked.setOptions({
      highlight: function(code, lang) {
        if (Prism.languages[lang]) {
          return Prism.highlight(code, Prism.languages[lang], lang)
        }
        return code
      }
    })

    // 添加复制按钮到代码块
    const addCopyButtons = () => {
      const codeBlocks = contentRef.value.querySelectorAll('pre code')
      codeBlocks.forEach((codeBlock, index) => {
        // 创建复制按钮容器
        const buttonContainer = document.createElement('div')
        buttonContainer.className = 'code-header'
        
        // 创建语言标签
        const langTag = document.createElement('span')
        langTag.className = 'code-lang'
        langTag.textContent = codeBlock.className.split('-')[1] || 'plaintext'
        
        // 创建复制按钮
        const copyButton = document.createElement('button')
        copyButton.className = 'copy-button'
        copyButton.setAttribute('data-clipboard-target', `#code-${index}`)
        copyButton.innerHTML = '<i class="fas fa-copy"></i>'
        
        // 添加到容器
        buttonContainer.appendChild(langTag)
        buttonContainer.appendChild(copyButton)
        
        // 给代码块添加id
        codeBlock.id = `code-${index}`
        
        // 将按钮容器添加到pre标签前
        codeBlock.parentElement.insertBefore(buttonContainer, codeBlock)
      })

      // 初始化clipboard
      const clipboard = new ClipboardJS('.copy-button')
      clipboard.on('success', (e) => {
        const button = e.trigger
        button.innerHTML = '<i class="fas fa-check"></i>'
        setTimeout(() => {
          button.innerHTML = '<i class="fas fa-copy"></i>'
        }, 2000)
      })
    }

    const fetchArticle = async () => {
      try {
        const response = await blogapi.getArticleById(route.params.id)
        article.value = response.data.article
        
        // 格式化时间
        article.value.CreateTime = new Date(article.value.CreateTime).toLocaleString()
        article.value.UpdateTime = new Date(article.value.UpdateTime).toLocaleString()
        
        // 转换markdown为HTML
        htmlContent.value = marked(article.value.Content)
        
        // 等待DOM更新后添加复制按钮
        setTimeout(() => {
          addCopyButtons()
          Prism.highlightAll()
        }, 0)
      } catch (error) {
        console.error('获取文章详情失败:', error)
      }
    }

    onMounted(() => {
      fetchArticle()
    })

    return {
      article,
      htmlContent,
      contentRef
    }
  }
}
</script>

<style scoped>
.article-detail {
  max-width: 900px;
  margin: 0 auto;
  padding: 2rem;
  background: white;
  border-radius: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.article-header {
  margin-bottom: 2rem;
  padding-bottom: 1.5rem;
  border-bottom: 1px solid #e5e7eb;
}

.article-title {
  font-size: 2rem;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 1rem;
  line-height: 1.3;
}

.article-meta {
  display: flex;
  gap: 1.5rem;
  color: #6b7280;
  font-size: 0.9rem;
  margin-bottom: 1rem;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.meta-item i {
  color: #6d28d9;
}

.article-tags {
  display: flex;
  gap: 0.75rem;
  flex-wrap: wrap;
}

.tag {
  background: #f3e8ff;
  color: #6d28d9;
  padding: 0.25rem 0.75rem;
  border-radius: 1rem;
  font-size: 0.85rem;
  font-weight: 500;
}

/* Markdown样式 */
:deep(.markdown-body) {
  color: #374151;
  line-height: 1.8;
  font-size: 1.1rem;
}

:deep(.markdown-body h1),
:deep(.markdown-body h2),
:deep(.markdown-body h3) {
  margin-top: 2rem;
  margin-bottom: 1rem;
  font-weight: 600;
  color: #1f2937;
}

:deep(.markdown-body p) {
  margin-bottom: 1.5rem;
}

:deep(.markdown-body pre) {
  position: relative;
  background: #1e1e1e;
  padding: 1rem;
  border-radius: 0.5rem;
  overflow-x: auto;
  margin: 1.5rem 0;
}

:deep(.markdown-body code) {
  font-family: 'Fira Code', monospace;
  font-size: 0.9em;
  color: #e5e7eb;
}

:deep(.markdown-body img) {
  max-width: 100%;
  border-radius: 0.5rem;
  margin: 1.5rem 0;
}

:deep(.markdown-body blockquote) {
  border-left: 4px solid #6d28d9;
  padding-left: 1rem;
  color: #6b7280;
  margin: 1.5rem 0;
}

:deep(.markdown-body ul),
:deep(.markdown-body ol) {
  padding-left: 1.5rem;
  margin: 1.5rem 0;
}

:deep(.markdown-body li) {
  margin-bottom: 0.5rem;
}

/* 代码块样式 */
:deep(.code-header) {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5em 1em;
  background: #2d2d2d;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
  border-bottom: 1px solid #404040;
}

:deep(.code-lang) {
  color: #b3b3b3;
  font-size: 0.85em;
  font-family: 'Fira Code', monospace;
}

:deep(.copy-button) {
  background: transparent;
  border: none;
  color: #b3b3b3;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s ease;
}

:deep(.copy-button:hover) {
  background: #404040;
  color: #fff;
}

:deep(.copy-button i) {
  font-size: 0.9em;
}

/* 代码块滚动条样式 */
:deep(pre::-webkit-scrollbar) {
  width: 8px;
  height: 8px;
}

:deep(pre::-webkit-scrollbar-thumb) {
  background: #404040;
  border-radius: 4px;
}

:deep(pre::-webkit-scrollbar-track) {
  background: #2d2d2d;
  border-radius: 4px;
}
</style>