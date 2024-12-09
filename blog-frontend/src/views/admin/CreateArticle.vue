<template>
    <div class="article-editor">
      <div class="editor-header">
        <input 
          v-model="article.title" 
          class="title-input" 
          placeholder="请输入文章标题"
        />
        
        <select v-model="article.category" class="category-select">
          <option value="">请选择文章分类</option>
          <option value="network">计算机网络</option>
          <option value="algorithm">数据结构和算法</option>
          <option value="architecture">计算机组成原理</option>
          <option value="os">操作系统</option>
          <option value="linux">Linux服务器</option>
          <option value="database">数据库</option>
          <option value="golang">Golang</option>
          <option value="leetcode">力扣题解</option>
          <option value="project">项目</option>
        </select>

        <textarea
          v-model="article.excerpt"
          class="excerpt-input"
          placeholder="请输入文章简介"
          rows="3"
        ></textarea>
  
        <input 
          v-model="article.tags" 
          class="tags-input" 
          placeholder="输入标签，用空格或逗号分隔"
        />
      </div>
  
      <MarkdownEditor
        v-model="article.content"
        @change="handleContentChange"
        class="md-editor"
      />
  
      <div class="editor-footer">
        <button @click="saveArticle" class="save-btn">发布文章</button>
      </div>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  import { blogapi } from '@/blogapi/blogapi'
  import MarkdownEditor from '@/components/MarkdownEditor.vue'
  
  export default {
    components: {
      MarkdownEditor
    },
    setup() {
      const router = useRouter()
      const article = ref({
        title: '',
        excerpt: '',
        content: '',
        category: '',
        tags: '',
        createTime: '',
        updateTime: ''
      })
  
      const handleContentChange = (value) => {
        article.value.content = value
      }
  
      const saveArticle = async () => {
        try {
          const articleData = {
            ...article.value,
            tags: article.value.tags.split(/[,，\s]+/).map(tag => tag.trim()),
            createTime: new Date().toISOString(),
            updateTime: new Date().toISOString()
          }

          const response = await blogapi.createArticle(articleData) 
  
          if (response.data.ok) {
            // 根据分类决定路由路径
            const baseRoute = ['golang', 'leetcode', 'project'].includes(article.value.category) 
              ? `/${article.value.category}`
              : `/computer/${article.value.category}`
            
            router.push(baseRoute)
          }
        } catch (error) {
          if (error.response.data.code === 401) {
            alert("收手吧！你还不具备管理员资格！")
          }
          else {
            alert("创建文章失败")
          }
        }
      }
  
      return {
        article,
        handleContentChange,
        saveArticle
      }
    }
  }
  </script>
  
  <style scoped>
  .article-editor {
    padding: 2.5rem;
    max-width: 1200px;
    margin: 0 auto;
    margin-top: -50px;
    background-color: #ffffff;
    border-radius: 16px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.06);
  }
  
  .editor-header {
    margin-bottom: 2.5rem;
    display: grid;
    gap: 1.25rem;
  }
  
  .title-input,
  .category-select,
  .tags-input,
  .excerpt-input {
    width: 100%;
    padding: 0.875rem 1.25rem;
    border: 1px solid #e2e8f0;
    border-radius: 12px;
    outline: none;
    transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
    background-color: #fafafa;
  }
  
  .title-input:focus,
  .category-select:focus,
  .tags-input:focus,
  .excerpt-input:focus {
    border-color: #8b5cf6;
    box-shadow: 0 0 0 4px rgba(139, 92, 246, 0.08);
    background-color: #ffffff;
  }
  
  .title-input {
    font-size: 1.625rem;
    font-weight: 600;
    color: #2d3748;
    letter-spacing: -0.025em;
  }
  
  .category-select {
    background-color: #fafafa;
    color: #4a5568;
    cursor: pointer;
    font-size: 1rem;
  }
  
  .tags-input {
    color: #4a5568;
    font-size: 1rem;
  }

  .excerpt-input {
    color: #4a5568;
    font-size: 1rem;
    resize: vertical;
    min-height: 100px;
  }
  
  .md-editor {
    min-height: 550px;
    margin-bottom: 2.5rem;
    border: 1px solid #e2e8f0;
    border-radius: 12px;
    background-color: #fafafa;
  }
  
  .editor-footer {
    display: flex;
    justify-content: flex-end;
    padding-top: 1.5rem;
    border-top: 1px solid #edf2f7;
  }
  
  .save-btn {
    background-color: #8b5cf6;
    color: white;
    padding: 0.875rem 2.5rem;
    border-radius: 12px;
    border: none;
    cursor: pointer;
    font-weight: 600;
    font-size: 1.0625rem;
    transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
    letter-spacing: 0.025em;
  }
  
  .save-btn:hover {
    background-color: #7c3aed;
    transform: translateY(-2px);
    box-shadow: 0 8px 16px rgba(139, 92, 246, 0.15);
  }
  
  .save-btn:active {
    transform: translateY(0);
    box-shadow: 0 4px 8px rgba(139, 92, 246, 0.1);
  }
  </style>