<template>
  <div class="message-board">
    <div class="message-board-header">
      <h1 class="title">留言板</h1>
      <div v-if="!username" class="auth-buttons">
        <router-link to="/home/login">  
          <button class="auth-btn login-btn">登录</button>
        </router-link>
        <router-link to="/home/register">
          <button class="auth-btn register-btn">注册</button>
        </router-link>
      </div>
      <div v-else class="user-info">
        <div class="user-info-content">
          <i class="fas fa-user-circle user-icon"></i>
          <span class="welcome">欢迎, {{ username }}</span>
          <div class="user-status-dot"></div>
        </div>
      </div>
    </div>
    
    <!-- 发表留言区域 -->
    <div class="post-message">
      <textarea 
        v-model="newMessage" 
        placeholder="写下你的留言..."
        class="message-input"
      ></textarea>
      <button @click="postMessage" class="post-btn">发表留言</button>
    </div>

    <!-- 留言列表 -->
    <div class="message-list-container">
      <div class="message-list">
        <div v-for="message in messages" :key="message.id" class="message-item">
          <div class="message-header">
            <img :src="message.avatar" :alt="message.username" class="user-avatar">
            <span class="username">{{ message.username }}</span>
            <span class="time">{{ message.time }}</span>
          </div>
          
          <div class="message-content">{{ message.content }}</div>
          
          <div class="message-actions">
            <button @click="likeMessage(message.id)" class="action-btn">
              <i class="fas fa-heart" :class="{ 'liked': message.isLiked }"></i>
              {{ message.likes }}
            </button>
            <button @click="showReplyInput(message.id)" class="action-btn">
              <i class="fas fa-reply"></i> 回复
            </button>
          </div>

          <!-- 回复输入框 -->
          <div v-if="replyingTo === message.id" class="reply-input">
            <textarea 
              v-model="replyContent" 
              placeholder="写下你的回复..."
              class="reply-textarea"
            ></textarea>
            <button @click="submitReply(message.id)" class="reply-btn">回复</button>
          </div>

          <!-- 回复列表 -->
          <div class="replies">
            <div v-for="reply in message.replies" :key="reply.id" class="reply-item">
              <div class="reply-header">
                <img :src="reply.avatar" :alt="reply.username" class="user-avatar">
                <span class="username">{{ reply.username }}</span>
                <span class="time">{{ reply.time }}</span>
              </div>
              <div class="reply-content">{{ reply.content }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { blogapi } from '../../blogapi/blogapi';

export default {
  data() {
    return {
      username: '',
      messages: [
        {
          id: 1,
          username: '张三',
          avatar: 'https://i.pravatar.cc/150?img=1',
          content: '博主写的文章很有深度,学到了很多!',
          time: '2024-01-14 10:23',
          likes: 12,
          isLiked: false,
          replies: [
            {
              id: 1,
              username: '李四',
              avatar: 'https://i.pravatar.cc/150?img=2', 
              content: '同感,特别是Go并发那篇文章写得很棒',
              time: '2024-01-14 10:45'
            }
          ]
        },
        {
          id: 2,
          username: '王五',
          avatar: 'https://i.pravatar.cc/150?img=3',
          content: '期待更多算法相关的文章分享!',
          time: '2024-01-14 09:15',
          likes: 8,
          isLiked: true,
          replies: []
        }
      ],
      newMessage: '',
      replyContent: '',
      replyingTo: null
    }
  },
  methods: {
    async checkLoginStatus() {
      try {
        const response = await blogapi.checkLogin()
        if(response.data.username) {
          this.username = response.data.username
        }
      } catch (error) {
        console.error('检查登录状态失败:', error)
      }
    },
    async fetchMessages() {
      try {
        const response = await blogapi.checkLogin()
        const data = await response.json()
        this.messages = data.messages
      } catch (error) {
        console.error('获取留言失败:', error)
      }
    },
    async postMessage() {
      if (!this.newMessage.trim()) return
      // 发送留言到服务器的逻辑
      this.newMessage = ''
    },
    async likeMessage(messageId) {
      // 点赞逻辑
    },
    showReplyInput(messageId) {
      this.replyingTo = this.replyingTo === messageId ? null : messageId
      this.replyContent = ''
    },
    async submitReply(messageId) {
      if (!this.replyContent.trim()) return
      // 提交回复到服务器的逻辑
      this.replyContent = ''
      this.replyingTo = null
    }
  },
  created() {
    this.checkLoginStatus()
    this.fetchMessages()
  }
}
</script>

<style scoped>
.message-board {
  background: #ffffff;
  padding: 2rem;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(124, 58, 237, 0.08);
  border: 1px solid rgba(124, 58, 237, 0.1);
}

.message-board-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.auth-buttons {
  display: flex;
  gap: 0.75rem;
}

.user-info {
  background: linear-gradient(135deg, #6d28d9 0%, #8b5cf6 100%);
  padding: 0.6rem 1.2rem;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(109, 40, 217, 0.2);
  transition: all 0.3s ease;
}

.user-info:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 15px rgba(109, 40, 217, 0.25);
}

.user-info-content {
  display: flex;
  align-items: center;
  gap: 0.8rem;
}

.user-icon {
  color: white;
  font-size: 1.2rem;
}

.welcome {
  color: white;
  font-size: 0.95rem;
  font-weight: 500;
  letter-spacing: 0.3px;
}

.user-status-dot {
  width: 8px;
  height: 8px;
  background-color: #4ade80;
  border-radius: 50%;
  border: 2px solid white;
}

.auth-btn {
  padding: 0.4rem 1.25rem;
  border-radius: 6px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.login-btn {
  background-color: #6d28d9;
  color: white;
  border: none;
}

.login-btn:hover {
  background-color: #5b21b6;
}

.register-btn {
  background-color: white;
  color: #6d28d9;
  border: 2px solid #6d28d9;
}

.register-btn:hover {
  background-color: #f3f4f6;
}

.title {
  font-size: 1.5rem;
  color: #6d28d9;
  font-weight: 600;
  letter-spacing: -0.5px;
  margin: 0;
}

.post-message {
  margin-bottom: 1.5rem;
  position: relative;
}

.message-list-container {
  max-height: 500px;
  overflow-y: auto;
  padding-right: 0.75rem;
}

.message-list-container::-webkit-scrollbar {
  width: 6px;
  opacity: 0;
  transition: opacity 0.2s;
}

.message-list-container:hover::-webkit-scrollbar,
.message-list-container::-webkit-scrollbar:hover {
  opacity: 1;
}

.message-list-container::-webkit-scrollbar-track {
  background: #e5e7eb;
  border-radius: 3px;
}

.message-list-container::-webkit-scrollbar-thumb {
  background: #60a5fa;
  border-radius: 3px;
}

.message-list-container::-webkit-scrollbar-thumb:hover {
  background: #3b82f6;
}

.message-input {
  width: 100%;
  height: 76px;
  padding: 0.75rem;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  resize: none;
  margin-bottom: 0.75rem;
  font-size: 0.95rem;
  font-weight: 400;
  transition: border-color 0.2s ease;
  padding-bottom: 2.5rem;
}

.message-input::-webkit-scrollbar {
  width: 6px;
  opacity: 0;
  transition: opacity 0.2s;
}

.message-input:hover::-webkit-scrollbar,
.message-input::-webkit-scrollbar:hover {
  opacity: 1;
}

.message-input::-webkit-scrollbar-track {
  background: #e5e7eb;
  border-radius: 3px;
}

.message-input::-webkit-scrollbar-thumb {
  background: #60a5fa;
  border-radius: 3px;
}

.message-input::-webkit-scrollbar-thumb:hover {
  background: #3b82f6;
}

.message-input:focus {
  border-color: #6d28d9;
  outline: none;
}

.post-btn {
  background-color: #6d28d9;
  color: white;
  border: none;
  padding: 0.4rem 0.7rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.85rem;
  font-weight: 500;
  transition: background-color 0.2s ease;
  position: absolute;
  bottom: 1.25rem;
  right: 0.5rem;
}

.post-btn:hover {
  background-color: #5b21b6;
}

.message-list {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.message-item {
  background-color: #f8fafc;
  border-radius: 12px;
  padding: 1.5rem;
  border: 1px solid rgba(124, 58, 237, 0.08);
  box-shadow: 0 2px 8px rgba(124, 58, 237, 0.05);
}

.message-header {
  display: flex;
  align-items: center;
  margin-bottom: 0.75rem;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 1rem;
}

.username {
  font-weight: 600;
  margin-right: 1rem;
  font-size: 0.95rem;
}

.time {
  color: #666;
  font-size: 0.85rem;
  font-weight: 400;
}

.message-content {
  margin-bottom: 0.75rem;
  line-height: 1.5;
  font-size: 0.95rem;
}

.message-actions {
  display: flex;
  gap: 1rem;
  margin-bottom: 0.75rem;
}

.action-btn {
  background: none;
  border: none;
  color: #666;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.85rem;
  font-weight: 500;
  transition: color 0.2s ease;
}

.action-btn:hover {
  color: #6d28d9;
}

.action-btn i.liked {
  color: #6d28d9;
}

.reply-input {
  margin: 0.75rem 0;
}

.reply-textarea {
  width: 100%;
  height: 80px;
  padding: 0.75rem;
  border: 1px solid #e0e0e0;
  border-radius: 6px;
  resize: none;
  margin-bottom: 0.5rem;
  font-size: 0.9rem;
  transition: border-color 0.2s ease;
}

.reply-textarea::-webkit-scrollbar {
  width: 6px;
  opacity: 0;
  transition: opacity 0.2s;
}

.reply-textarea:hover::-webkit-scrollbar,
.reply-textarea::-webkit-scrollbar:hover {
  opacity: 1;
}

.reply-textarea::-webkit-scrollbar-track {
  background: #e5e7eb;
  border-radius: 3px;
}

.reply-textarea::-webkit-scrollbar-thumb {
  background: #60a5fa;
  border-radius: 3px;
}

.reply-textarea::-webkit-scrollbar-thumb:hover {
  background: #3b82f6;
}

.reply-textarea:focus {
  border-color: #6d28d9;
  outline: none;
}

.reply-btn {
  background-color: #6d28d9;
  color: white;
  border: none;
  padding: 0.5rem 1.25rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  font-weight: 500;
  transition: background-color 0.2s ease;
}

.reply-btn:hover {
  background-color: #5b21b6;
}

.replies {
  margin-left: 2rem;
  margin-top: 0.75rem;
}

.reply-item {
  background-color: #ffffff;
  border-radius: 10px;
  padding: 1.25rem;
  border: 1px solid rgba(124, 58, 237, 0.08);
  box-shadow: 0 1px 4px rgba(124, 58, 237, 0.03);
}

.reply-header {
  display: flex;
  align-items: center;
  margin-bottom: 0.5rem;
}

.reply-content {
  color: #2c3e50;
  line-height: 1.4;
  font-size: 0.9rem;
  font-weight: 400;
}
</style>