<template>
  <div class="manage-articles">
    <div class="articles-table">
      <table>
        <thead>
          <tr>
            <th class="article-id-col">文章ID</th>
            <th class="title-col">标题</th>
            <th class="category-col">分类</th>
            <th class="date-col">发布时间</th>
            <th class="date-col">更新时间</th>
            <th class="view-col">浏览量</th>
            <th class="action-col">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="article in articles" :key="article.ArticleId">
            <td class="article-id-col">{{ article.ArticleId }}</td>
            <td class="title-col">{{ article.Title }}</td>
            <td class="category-col">{{ article.Category.CategoryName }}</td>
            <td class="date-col">{{ formatDate(article.CreateTime) }}</td>
            <td class="date-col">{{ formatDate(article.UpdateTime) }}</td>
            <td class="view-col">{{ article.ViewCount }}</td>
            <td class="action-col">
              <div class="actions-stack">
                <button @click="editArticle(article)" class="edit-btn">
                  <i class="fas fa-edit"></i> 编辑
                </button>
                <button @click="confirmDelete(article)" class="delete-btn">
                  <i class="fas fa-trash"></i> 删除
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 编辑文章对话框 -->
    <div v-if="showEditDialog" class="edit-dialog">
      <div class="dialog-content">
        <h3>编辑文章</h3>
        <input v-model="editingArticle.Title" placeholder="文章标题" />
        <textarea v-model="editingArticle.Content" placeholder="文章内容"></textarea>
        <div class="dialog-buttons">
          <button @click="saveEdit" class="save-btn">保存</button>
          <button @click="cancelEdit" class="cancel-btn">取消</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { blogapi } from '../../blogapi/blogapi'

export default {
  setup() {
    const articles = ref([])
    const showEditDialog = ref(false)
    const editingArticle = ref(null)

    const fetchArticles = async () => {
      try {
        const response = await blogapi.getAllArticles()
        articles.value = response.data.articles
      } catch (error) {
        console.error('获取文章列表失败:', error)
      }
    }

    const editArticle = (article) => {
      editingArticle.value = { ...article }
      editingArticle.value.UpdateTime = new Date()
      showEditDialog.value = true
    }

    const saveEdit = async () => {
      try {
        await blogapi.updateArticle(editingArticle.value)
        showEditDialog.value = false
        fetchArticles()
      } catch (error) {
        if (error.response.data.code === 401) {
          alert("收手吧！你还不具备管理员资格！")
        }
        else {
          alert("更新文章失败")
        }
      }
    }

    const cancelEdit = () => {
      showEditDialog.value = false
      editingArticle.value = null
    }

    const confirmDelete = async (article) => {
      if (confirm('确定要删除这篇文章吗？')) {
        try {
          await blogapi.deleteArticle(article.ArticleId)
          fetchArticles()
        } catch (error) {
          if (error.response.data.code === 401) {
            alert("收手吧！你还不具备管理员资格！")
          }
          else {
            alert("删除文章失败")
          }
        }
      }
    }

    const formatDate = (date) => {
      return new Date(date).toLocaleString()
    }

    // 初始加载文章列表
    fetchArticles()

    return {
      articles,
      showEditDialog,
      editingArticle,
      editArticle,
      saveEdit,
      cancelEdit,
      confirmDelete,
      formatDate
    }
  }
}
</script>

<style scoped>
.manage-articles {
  padding-top: 2.5rem;
  padding-left: 0.5rem;
  padding-right: 0.5rem;
  background-color: #f8fafc;
  min-height: 100vh;
  margin-top: -80px;
}



.articles-table {
  background: #ffffff;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
  overflow: hidden;
  transition: all 0.3s ease;
}

table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0;
}

th, td {
  padding: 0.75rem 1rem;
  text-align: left;
  font-size: 0.9rem;
  line-height: 1.4;
}

/* 设置列宽 */

.article-id-col {
  width: 10%;
  text-align: center;
}
.title-col {
  width: 22%;
  text-align: center;
}

.category-col {
  width: 10%;
  text-align: center;
}

.date-col {
  width: 12%;
  text-align: center;
}

.view-col {
  width: 12%;
  text-align: center;
}

.action-col {
  width: 15%;
  text-align: center;
}

th {
  background: #f1f5f9;
  color: #475569;
  font-weight: 500;
  letter-spacing: 0.025em;
  border-bottom: 1px solid #e2e8f0;
}

td {
  color: #334155;
  border-bottom: 1px solid #f1f5f9;
}

tr:hover {
  background-color: #f8fafc;
}

.actions-stack {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.edit-btn, .delete-btn {
  padding: 0.35rem 0.75rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.8rem;
  font-weight: 500;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.3rem;
  transition: all 0.2s ease;
  width: 100%;
}

.edit-btn {
  background: #6366f1;
  color: white;
}

.edit-btn:hover {
  background: #4f46e5;
  transform: translateY(-1px);
}

.delete-btn {
  background: #f87171;
  color: white;
}

.delete-btn:hover {
  background: #ef4444;
  transform: translateY(-1px);
}

.edit-dialog {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.4);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.dialog-content {
  background: white;
  padding: 2.5rem;
  border-radius: 16px;
  width: 90%;
  max-width: 700px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
}

.dialog-content h3 {
  color: #1e293b;
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
}

.dialog-content input,
.dialog-content textarea {
  width: 100%;
  margin: 1rem 0;
  padding: 0.75rem 1rem;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 0.95rem;
  transition: all 0.2s ease;
}

.dialog-content input:focus,
.dialog-content textarea:focus {
  outline: none;
  border-color: #6366f1;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.1);
}

.dialog-content textarea {
  height: 300px;
  resize: vertical;
  line-height: 1.6;
}

.dialog-buttons {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 2rem;
}

.save-btn, .cancel-btn {
  padding: 0.75rem 2rem;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease;
}

.save-btn {
  background: #6366f1;
  color: white;
}

.save-btn:hover {
  background: #4f46e5;
  transform: translateY(-1px);
}

.cancel-btn {
  background: #e2e8f0;
  color: #475569;
}

.cancel-btn:hover {
  background: #cbd5e1;
  transform: translateY(-1px);
}
</style>