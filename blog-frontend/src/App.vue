<template>
  <div class="app-container" :class="{ 'dark': isDark }">
    <nav class="navbar">
      <div class="nav-brand">
        <router-link to="/" class="brand-link">zlican's Blog</router-link>
      </div>
      <div class="nav-links">
        <router-link to="/home" class="nav-link" exact>首页</router-link>
        <router-link to="/computer" class="nav-link" exact>计算机基础</router-link>
        <router-link to="/golang" class="nav-link" exact>Go语言</router-link>
        <router-link to="/leetcode" class="nav-link" exact>力扣题解</router-link>
        <router-link to="/project" class="nav-link" exact>项目相关</router-link>
      </div>
      <div class="nav-right">
        <button @click="toggleTheme" class="theme-toggle">
          <span v-if="isDark">🌞</span>
          <span v-else>🌙</span>
        </button>
      </div>
    </nav>

    <!-- 管理员入口按钮 -->
    <router-link to="/admin" class="admin-link">
      <button class="admin-btn">
        <span class="admin-icon">👨‍💻</span>
        <span class="admin-text">管理员入口</span>
      </button>
    </router-link>

    <main class="main-content" v-if="$route.path.startsWith('/')">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>

    <footer class="footer">
      <p>© 2024 zlican's Blog. All rights reserved.</p>
    </footer>

    <!-- 返回顶部按钮 -->
    <button 
      class="back-to-top"
      @click="scrollToTop"
      v-show="showBackToTop"
    >
      ⬆️
    </button>
  </div>
</template>

<script>
import { useTheme } from './composables/useTheme'
import { ref, onMounted, onUnmounted } from 'vue'

export default {
  name: 'App',
  setup() {
    const { isDark, toggleTheme } = useTheme()
    const showBackToTop = ref(false)
    
    const handleScroll = () => {
      showBackToTop.value = window.scrollY > 300
    }

    const scrollToTop = () => {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      })
    }

    onMounted(() => {
      window.addEventListener('scroll', handleScroll)
    })

    onUnmounted(() => {
      window.removeEventListener('scroll', handleScroll)
    })
    
    return {
      isDark,
      toggleTheme,
      showBackToTop,
      scrollToTop
    }
  }
}
</script>
<style>
:root {
  /* 主题色系 */
  --primary-color: #3b82f6;
  --primary-light: #60a5fa;
  --primary-dark: #2563eb;
  --accent-color: #f59e0b;
  
  /* 文字颜色 */
  --text-primary: #1f2937;
  --text-secondary: #4b5563;
  --text-tertiary: #6b7280;
  
  /* 背景颜色 */
  --bg-primary: #ffffff;
  --bg-secondary: #f3f4f6;
  --bg-tertiary: #e5e7eb;
  
  /* 边框和阴影 */
  --border-color: #e5e7eb;
  --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  --shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  
  /* 动画和过渡 */
  --transition-speed: 0.3s;
  --transition-timing: cubic-bezier(0.4, 0, 0.2, 1);
}

.dark {
  --primary-color: #60a5fa;
  --primary-light: #93c5fd;
  --primary-dark: #3b82f6;
  --accent-color: #fbbf24;
  
  --text-primary: #f3f4f6;
  --text-secondary: #d1d5db;
  --text-tertiary: #9ca3af;
  
  --bg-primary: #1f2937;
  --bg-secondary: #374151;
  --bg-tertiary: #4b5563;
  
  --border-color: #4b5563;
  --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.3);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.4);
  --shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.4);
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  color: var(--text-primary);
  background-color: var(--bg-primary);
  transition: all var(--transition-speed) var(--transition-timing);
  line-height: 1.5;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
}

/* 管理员入口按钮样式 */
.admin-link {
  position: fixed;
  top: 5rem;
  right: 1rem;
  z-index: 99;
  text-decoration: none;
}

.admin-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background-color: var(--bg-secondary);
  border: 2px solid var(--primary-color);
  border-radius: 20px;
  color: var(--text-primary);
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all var(--transition-speed) var(--transition-timing);
  box-shadow: var(--shadow-sm);
}

.admin-btn:hover {
  background-color: var(--primary-color);
  color: white;
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.admin-icon {
  font-size: 1.2rem;
}

.admin-text {
  font-size: 0.875rem;
}

.navbar {
  padding: 1rem 2rem;
  background-color: var(--bg-secondary);
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: var(--shadow-md);
  position: sticky;
  top: 0;
  z-index: 100;
  backdrop-filter: blur(8px);
}

.nav-brand .brand-link {
  font-size: 1.75rem;
  font-weight: 800;
  color: var(--primary-color);
  text-decoration: none;
  letter-spacing: -0.025em;
  transition: all var(--transition-speed) var(--transition-timing);
}

.nav-brand .brand-link:hover {
  color: var(--primary-light);
  transform: translateY(-1px);
}

.nav-links {
  display: flex;
  gap: 2rem;
  align-items: center;
}

.nav-link {
  color: var(--text-primary);
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 0.75rem;
  font-weight: 500;
  transition: all var(--transition-speed) var(--transition-timing);
  position: relative;
}

.nav-link::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  width: 0;
  height: 2px;
  background-color: var(--primary-color);
  transition: all var(--transition-speed) var(--transition-timing);
  transform: translateX(-50%);
}

.nav-link:hover {
  color: var(--primary-color);
}

.nav-link:hover::after,
.nav-link.router-link-active::after {
  width: 80%;
}

.nav-link.router-link-active {
  color: var(--primary-color);
}

.theme-toggle {
  background: none;
  border: 2px solid var(--border-color);
  cursor: pointer;
  font-size: 1.25rem;
  padding: 0.5rem;
  border-radius: 50%;
  transition: all var(--transition-speed) var(--transition-timing);
  display: flex;
  align-items: center;
  justify-content: center;
}

.theme-toggle:hover {
  background-color: var(--bg-primary);
  transform: rotate(360deg);
  border-color: var(--primary-color);
}

.main-content {
  flex: 1;
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
  animation: fadeIn 0.5s ease-out;
}

.footer {
  padding: 2rem;
  text-align: center;
  background-color: var(--bg-secondary);
  color: var(--text-secondary);
  box-shadow: var(--shadow-lg);
}

.footer p {
  font-size: 0.875rem;
  opacity: 0.8;
}

/* 返回顶部按钮样式 */
.back-to-top {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  width: 3rem;
  height: 3rem;
  border-radius: 50%;
  background-color: var(--primary-color);
  color: white;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  box-shadow: var(--shadow-lg);
  transition: all var(--transition-speed) var(--transition-timing);
  z-index: 1000;
}

.back-to-top:hover {
  background-color: var(--primary-light);
  transform: translateY(-3px);
}

/* 动画效果 */
.fade-enter-active,
.fade-leave-active {
  transition: all var(--transition-speed) var(--transition-timing);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .navbar {
    padding: 1rem;
  }
  
  .nav-links {
    gap: 1rem;
  }
  
  .nav-link {
    padding: 0.4rem 0.8rem;
    font-size: 0.9rem;
  }
  
  .main-content {
    padding: 1rem;
  }

  .back-to-top {
    bottom: 1rem;
    right: 1rem;
    width: 2.5rem;
    height: 2.5rem;
    font-size: 1.2rem;
  }

  .admin-btn {
    padding: 0.4rem 0.8rem;
  }

  .admin-text {
    display: none;
  }
}

@media (max-width: 640px) {
  .nav-brand .brand-link {
    font-size: 1.5rem;
  }
  
  .nav-links {
    display: none;
  }
}
</style>