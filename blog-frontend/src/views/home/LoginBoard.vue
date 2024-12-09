<template>
  <div class="login-board">
    <div class="back-arrow" @click="$router.go(-1)">
      <i class="fas fa-arrow-left"></i>
    </div>
    <h2>用户登录</h2>
    <div class="login-form">
      <div class="form-item">
        <label>手机号</label>
        <input type="text" v-model="phone" placeholder="请输入手机号">
      </div>

      <div class="form-item">
        <label>密码</label>
        <input type="password" v-model="password" placeholder="请输入密码">
      </div>

      <button class="login-btn" @click="login">登录</button>
    </div>
  </div>
</template>

<script>
import { blogapi } from '../../blogapi/blogapi';
export default {
  data() {
    return {
      phone: '',
      password: ''
    }
  },
  methods: {
    async login() {
       const res = await blogapi.postLogin({
        phone: this.phone,
        password: this.password
      })
        this.$router.push("/home")
    }
  }
}
</script>

<style scoped>
.login-board {
  background: #ffffff;
  padding: 2.5rem;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(124, 58, 237, 0.08);
  width: 380px;
  border: 1px solid rgba(124, 58, 237, 0.1);
  position: relative;
}

.back-arrow {
  position: absolute;
  top: 1.5rem;
  left: 1.5rem;
  color: #7c3aed;
  cursor: pointer;
  font-size: 1.2rem;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: all 0.2s ease;
}

.back-arrow:hover {
  background: rgba(124, 58, 237, 0.1);
}

.login-board h2 {
  text-align: center;
  color: #7c3aed;
  margin-bottom: 2rem;
  font-weight: 600;
  font-size: 1.5rem;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 1.4rem;
}

.form-item {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-item label {
  font-size: 0.9rem;
  color: #6b21a8;
  font-weight: 500;
  margin-left: 0.2rem;
}

.form-item input {
  padding: 0.9rem 1.2rem;
  border: 1.5px solid #ddd6fe;
  border-radius: 12px;
  font-size: 0.95rem;
  transition: all 0.2s ease;
  background: #ffffff;
}

.form-item input:focus {
  border-color: #7c3aed;
  outline: none;
  box-shadow: 0 0 0 3px rgba(124, 58, 237, 0.1);
}

.form-item input::placeholder {
  color: #a1a1aa;
}

.login-btn {
  margin-top: 0.8rem;
  background: linear-gradient(135deg, #7c3aed 0%, #6d28d9 100%);
  color: white;
  padding: 1rem;
  border: none;
  border-radius: 12px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: all 0.3s ease;
  letter-spacing: 0.5px;
}

.login-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.25);
}

.login-btn:active {
  transform: translateY(0);
  background: linear-gradient(135deg, #6d28d9 0%, #5b21b6 100%);
}
</style>