<template>
  <div class="register-board">
    <div class="back-arrow" @click="$router.go(-1)">
      <i class="fas fa-arrow-left"></i>
    </div>
    <h2>用户注册</h2>
    <div class="register-form">
      <div class="form-item">
        <label>手机号</label>
        <input type="text" v-model="phone" placeholder="请输入手机号">
      </div>

      <div class="form-item verification">
        <label>图形验证码</label>
        <div class="code-input">
          <input type="text" v-model="captcha" placeholder="请输入验证码">
          <div class="captcha-img" >
            <img :src="captchaUrl" @click="loadCaptcha" alt="验证码" class="captcha-image">
          </div>
        </div>
      </div>

      <div class="form-item verification">
        <label>手机验证码</label>
        <div class="code-input">
          <input type="text" v-model="phoneCode" placeholder="请输入验证码">
          <button @click="sendPhoneCode" :disabled="cooldown > 0">
            {{ cooldown > 0 ? `${cooldown}秒后重试` : '获取验证码' }}
          </button>
        </div>
      </div>

      <div class="form-item">
        <label>密码</label>
        <input type="password" v-model="password" placeholder="请输入密码">
      </div>

      <div class="form-item">
        <label>确认密码</label>
        <input type="password" v-model="confirmPassword" placeholder="请再次输入密码">
      </div>

      <button class="register-btn" @click="register">注册</button>
    </div>
  </div>
</template>

<script>

import { blogapi } from '../../blogapi/blogapi';

export default {
  data() {
    return {
        uuid: '',
      phone: '',
      phoneCode: '',
      captcha: '',
      password: '',
      confirmPassword: '',
      cooldown: 0,
      captchaUrl: '',
      backendUrl: 'http://127.0.0.1:8000/blogapi/v1'
    }
  },
  created() {
    this.loadCaptcha()
  },
  methods: {
    getUuid() {
        this.uuid = Math.random()
      return this.uuid
    },
    async loadCaptcha() {
        this.getUuid()
      this.captchaUrl = `${this.backendUrl}/home/captcha/${this.uuid}`
    },
    async sendPhoneCode() {
      // 验证手机号和图片验证码是否填写
      if (!this.phone || !this.captcha) {
        alert('请填写手机号和图片验证码')
        return
      }
      
      try {
        // 调用API发送验证码
        await blogapi.getPhoneCode({
            uuid: this.uuid,
          phone: this.phone,
          captcha: this.captcha
        })
        
        // 开始倒计时
        this.cooldown = 60
        const timer = setInterval(() => {
          this.cooldown--
          if (this.cooldown <= 0) {
            clearInterval(timer)
          }
        }, 1000)
      } catch (error) {
        alert('发送验证码失败，请重试')
        // 刷新图片验证码
        this.loadCaptcha()
      }
    },
    async register() {
      // 验证所有字段是否填写
      if (!this.phone || !this.phoneCode || !this.password || !this.confirmPassword) {
        alert('请填写所有必填项')
        return
      }

      // 验证两次密码是否一致
      if (this.password !== this.confirmPassword) {
        alert('两次输入的密码不一致')
        return
      }

      try {
        // 发送注册请求
        await blogapi.postRegister({
          phone: this.phone,
          phoneCode: this.phoneCode,
          password: this.password
        })
        
        alert('注册成功')
        // 注册成功后跳转到登录页
        this.$router.push('/home/login')
      } catch (error) {
        alert(error)
      }
    }
  }
}
</script>

<style scoped>
.register-board {
  background: #ffffff;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(76, 29, 149, 0.08);
  width: 380px;
  border: 1px solid rgba(76, 29, 149, 0.1);
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

.register-board h2 {
  text-align: center;
  color: #7c3aed;
  margin-bottom: 1.8rem;
  font-weight: 600;
  font-size: 1.5rem;
}

.register-form {
  display: flex;
  flex-direction: column;
  gap: 1.2rem;
}

.form-item {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.form-item label {
  font-size: 0.85rem;
  color: #6b21a8;
  font-weight: 500;
}

.form-item input {
  padding: 0.7rem 1rem;
  border: 1.5px solid #ddd6fe;
  border-radius: 8px;
  font-size: 0.9rem;
  transition: all 0.2s ease;
  background: #ffffff;
}

.form-item input:focus {
  border-color: #7c3aed;
  box-shadow: 0 0 0 2px rgba(124, 58, 237, 0.1);
  outline: none;
}

.verification .code-input {
  display: flex;
  gap: 0.8rem;
}

.verification .code-input input {
  flex: 1;
}

.verification .code-input button {
  padding: 0 0.6rem;
  background: #7c3aed;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  white-space: nowrap;
  font-weight: 500;
  font-size: 0.85rem;
  transition: all 0.2s ease;
}

.verification .code-input button:hover:not(:disabled) {
  background: #7c3aed;
  transform: translateY(-1px);
}

.verification .code-input button:disabled {
  background: #ddd6fe;
  cursor: not-allowed;
}

.captcha-img {
  width: 100px;
  height: 38px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f3ff;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  border: 1.5px solid #ddd6fe;
  overflow: hidden;
}

.captcha-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.captcha-img:hover {
  border-color: #7c3aed;
}

.register-btn {
  background: linear-gradient(135deg, #7c3aed 0%, #6d28d9 100%);
  color: white;
  padding: 0.8rem;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 0.95rem;
  font-weight: 600;
  transition: all 0.2s ease;
  margin-top: 0.5rem;
}

.register-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 3px 10px rgba(124, 58, 237, 0.2);
}

.register-btn:active {
  transform: translateY(0);
}
</style>