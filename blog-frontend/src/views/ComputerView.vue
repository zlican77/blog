<template>
  <div class="computer-container">
    <nav class="sub-nav">
      <router-link 
        v-for="route in routes" 
        :key="route.path"
        :to="route.path"
        class="nav-item"
        active-class="active"
      >
        {{ route.name }}
      </router-link>
    </nav>

    <div class="content">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      routes: [
        { path: '/computer/network', name: '计算机网络' },
        { path: '/computer/algorithm', name: '数据结构和算法' },
        { path: '/computer/architecture', name: '组成原理' },
        { path: '/computer/os', name: '操作系统' },
        { path: '/computer/linux', name: 'Linux服务器' },
        { path: '/computer/database', name: '数据库' }
      ]
    }
  }
}
</script>

<style scoped>
.computer-container {
  padding: 2rem;
  margin-top: -2rem;
}

.page-title {
  font-size: 2rem;
  color: var(--text-primary);
  margin-bottom: 2rem;
  font-weight: 600;
  letter-spacing: -0.5px;
}

.sub-nav {
  display: flex;
  gap: 1rem;
  background: var(--bg-secondary);
  padding: 1rem;
  border-radius: 0.75rem;
  margin-bottom: 2rem;
  box-shadow: var(--shadow-md);
  overflow-x: auto;
}

.nav-item {
  color: var(--text-secondary);
  text-decoration: none;
  padding: 0.75rem 1.25rem;
  border-radius: 0.5rem;
  font-weight: 500;
  transition: all 0.3s ease;
  white-space: nowrap;
}

.nav-item:hover {
  color: var(--primary-color);
  background: var(--bg-tertiary);
}

.nav-item.active {
  color: var(--primary-color);
  background: var(--bg-tertiary);
  font-weight: 600;
}

.content {
  background: var(--bg-secondary);
  border-radius: 0.75rem;
  padding: 2rem;
  min-height: 500px;
  box-shadow: var(--shadow-md);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>