<template>
  <div>
    <div class="home-container">
      <div class="top-section">
        <PersonalSummary class="section" />
        <div>
          <router-view v-slot="{ Component }" v-if="$route.path === '/home/login' || $route.path === '/home/register'">
            <transition name="fade" mode="out-in">
              <component :is="Component" />
            </transition>
          </router-view>
          <MessageBoard v-else />
        </div>
      </div>
      <LatestArticles class="section" />
    </div>
  </div>
</template>

<script>
import PersonalSummary from './home/PersonalSummary.vue'
import MessageBoard from './home/MessageBoard.vue'
import LatestArticles from './home/LatestArticles.vue'

export default {
  components: {
    PersonalSummary,
    MessageBoard,
    LatestArticles
  }
}
</script>

<style>
.home-container {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  padding: 2rem;
}

.top-section {
  display: flex;
  gap: 2rem;
  margin-top: -2rem;
}

.section {
  background: #fff;
  border-radius: 8px;
  padding: 1.5rem;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  flex: 1;
}
</style>