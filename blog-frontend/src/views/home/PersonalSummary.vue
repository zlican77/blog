<template>
  <div class="personal-summary">
    <h1 class="title">个人简介</h1>
    
    <div class="info-section">
      <div class="avatar-container">
        <img :src="personalInfo.avatar" alt="头像" class="avatar">
        <h2>{{ personalInfo.name }}</h2>
        <div class="contact-info">
          <a :href="personalInfo.github" target="_blank" class="contact-item">
            <i class="fab fa-github"></i>
            <span>{{personalInfo.github}}</span>
          </a>
          <div class="contact-item">
            <i class="fas fa-envelope"></i>
            <span>{{ personalInfo.email }}</span>
          </div>
          <div class="contact-item">
            <i class="fas fa-map-marker-alt"></i>
            <span>{{ personalInfo.location }}</span>
          </div>
        </div>
      </div>

      <div class="bio-container">
        <p class="bio">{{ personalInfo.bio }}</p>
      </div>
    </div>

    <div class="skills-section">
      <h3>技术栈</h3>
      <div class="skills-container">
        <div v-for="(skill, index) in skills" :key="index" class="skill-item">
          <span class="skill-name">{{ skill.name }}</span>
          <div class="skill-level">
            <div class="skill-progress" :style="{ width: skill.level + '%' }"></div>
          </div>
        </div>
      </div>
    </div>

    <div class="github-section">
      <h3>GitHub 项目</h3>
      <div class="repos-container">
        <div v-for="(repo, index) in githubRepos" :key="index" class="repo-item">
          <h4>{{ repo.name }}</h4>
          <p>{{ repo.description }}</p>
          <div class="repo-stats">
            <span><i class="fas fa-star"></i> {{ repo.stars }}</span>
            <span><i class="fas fa-code-branch"></i> {{ repo.forks }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      personalInfo: {
        name: '陈子陵',
        avatar: 'https://avatars.githubusercontent.com/u/116709865?v=4',
        bio: '全栈开发工程师，热爱技术，喜欢分享。专注于Go、Vue.js和微服务架构。',
        github: 'https://github.com/zlican77',
        email: 'zlican77@gmail.com',
        location: '安徽-合肥'
      },
      skills: [
        { name: 'Go', level: 90 },
        { name: 'JS', level: 80 },
        { name: 'Vue.js', level: 85 },
        { name: 'Docker', level: 80 },
        { name: 'Kubernetes', level: 75 },
        { name: 'MySQL', level: 85 }
      ],
      githubRepos: [
        {
          name: 'go-microservices',
          description: '基于Go语言的微服务框架示例',
          stars: 128,
          forks: 45
        },
        {
          name: 'vue-blog',
          description: '使用Vue.js开发的个人博客系统',
          stars: 89,
          forks: 32
        },
      ]
    }
  },
  async created() {
    try {
      const response = await fetch('/blogapi/v1/home/personal-summary')
      const data = await response.json()
      this.personalInfo = data.personalInfo
      this.skills = data.skills
      this.githubRepos = data.githubRepos
    } catch (error) {
      console.error('获取个人信息失败:', error)
    }
  }
}
</script>

<style scoped>
.personal-summary {
  background: #ffffff;
  padding: 2rem;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(124, 58, 237, 0.08);
  border: 1px solid rgba(124, 58, 237, 0.1);
}

.title {
  font-size: 1.5rem;
  margin-bottom: 1.25rem;
  color: #6d28d9;
  font-weight: 600;
  letter-spacing: -0.5px;
}

.info-section {
  display: grid;
  grid-template-columns: 150px 1fr;
  gap: 1.5rem;
  margin-bottom: 1.25rem;
}

.avatar-container {
  text-align: center;
}

.avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 0.75rem;
  border: 3px solid #6d28d9;
}

.avatar-container h2 {
  font-size: 1.1rem;
  margin: 0 0 0.75rem 0;
}

.bio-container {
  display: flex;
  align-items: center;
}

.bio {
  color: #666;
  margin: 0;
  font-size: 0.95rem;
  line-height: 1.5;
  font-weight: 400;
}

.contact-info {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 0.35rem;
  color: #666;
  text-decoration: none;
  font-size: 0.85rem;
  font-weight: 400;
}

.contact-item i {
  color: #6d28d9;
  width: 16px;
}

.skills-section, .github-section {
  margin-top: 1.25rem;
}

.skills-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 0.75rem;
}

.skill-item {
  background: #f8fafc;
  padding: 1rem;
  border-radius: 12px;
  border: 1px solid rgba(124, 58, 237, 0.08);
  box-shadow: 0 2px 8px rgba(124, 58, 237, 0.05);
}

.skill-name {
  display: block;
  margin-bottom: 0.35rem;
  font-size: 0.9rem;
  font-weight: 500;
}

.skill-level {
  height: 6px;
  background: #ddd;
  border-radius: 3px;
  overflow: hidden;
}

.skill-progress {
  height: 100%;
  background: #6d28d9;
  border-radius: 3px;
  transition: width 0.3s ease;
}

.repos-container {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 0.75rem;
}

.repo-item {
  background: #f8fafc;
  padding: 1rem;
  border-radius: 12px;
  border: 1px solid rgba(124, 58, 237, 0.08);
  box-shadow: 0 2px 8px rgba(124, 58, 237, 0.05);
}

.repo-item h4 {
  color: #6d28d9;
  margin: 0 0 0.35rem 0;
  font-size: 1rem;
  font-weight: 600;
}

.repo-stats {
  display: flex;
  gap: 0.75rem;
  margin-top: 0.35rem;
  color: #666;
  font-size: 0.9rem;
}

.repo-stats i {
  color: #6d28d9;
}
</style>