<template>
  <div class="admin-container">
    <nav class="admin-nav">
      <div class="nav-header">
        <h1>管理员控制台</h1>
      </div>
      <div class="nav-links">
        <router-link to="/admin/create" class="nav-item">
          <span class="icon">📝</span>
          <span>上传文章</span>
        </router-link>
        <router-link to="/admin/articles" class="nav-item">
          <span class="icon">📚</span>
          <span>文章管理</span>
        </router-link>
        <router-link to="/admin/tags" class="nav-item">
          <span class="icon">🏷️</span>
          <span>标签管理</span>
        </router-link>
        <router-link to="/admin/messages" class="nav-item">
          <span class="icon">💬</span>
          <span>留言管理</span>
        </router-link>
      </div>
    </nav>

    <main class="admin-content">
      <router-view></router-view>
    </main>
  </div>
</template>

<script>
export default {
  name: 'AdminView'
}
</script>

<style scoped>
.admin-container {
  display: flex;
  flex-direction: row;
  min-height: 100vh;
  background-color: var(--bg-secondary);
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
  position: relative; /* 添加相对定位 */
  padding-top: 20px;
}

.admin-nav {
  width: 240px;
  height: 480px;
  background-color: var(--bg-primary);
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.05);
  position: relative;
  left: 50%; /* 修改为相对于容器居中 */
  transform: translateX(calc(-50% - 420px)); /* 向左偏移以对齐内容区域 */
  top: 0;
  display: flex;
  flex-direction: column;
  border-right: 1px solid rgba(0, 0, 0, 0.05);
  backdrop-filter: blur(10px);
}

.nav-header {
  padding: 2rem;
  border-bottom: 1px solid rgba(0, 0, 0, 0.05);
  text-align: center;
}

.nav-header h1 {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--text-primary);
  letter-spacing: -0.5px;
  background: linear-gradient(120deg, var(--primary-color), var(--accent-color));
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.nav-links {
  display: flex;
  flex-direction: column;
  padding: 1.5rem 0;
  gap: 0.5rem;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 1rem 1.75rem;
  color: var(--text-secondary);
  text-decoration: none;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  border-radius: 12px;
  margin: 0 0.75rem;
  font-weight: 500;
}

.nav-item:hover {
  background-color: rgba(var(--primary-color-rgb), 0.08);
  color: var(--primary-color);
  transform: translateX(4px);
}

.nav-item.router-link-active {
  background: linear-gradient(120deg, var(--primary-color), var(--primary-light));
  color: white;
  box-shadow: 0 4px 12px rgba(var(--primary-color-rgb), 0.25);
}

.icon {
  margin-right: 1rem;
  font-size: 1.35rem;
  opacity: 0.9;
  transition: transform 0.3s ease;
}

.nav-item:hover .icon {
  transform: scale(1.1);
}

.admin-content {
  flex: 1;
  margin-left: 10px;
  padding: 2.5rem;
  overflow-y: auto;
  background: linear-gradient(to bottom right, 
    rgba(var(--primary-color-rgb), 0.03),
    rgba(var(--accent-color-rgb), 0.03));
  min-height: 100vh;
}


</style>