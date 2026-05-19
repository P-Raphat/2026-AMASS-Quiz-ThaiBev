<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api/auth'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const username = ref('')
const password = ref('')
const error = ref('')

const handleLogin = async () => {
  error.value = ''

  try {
    const response = await api.post('/login', {
      username: username.value,
      password: password.value,
    })

    authStore.login(response.data.token, response.data.username)

    router.push('/welcome')
  } catch (err) {
    error.value = 'Username หรือ Password ไม่ถูกต้อง'
  }
}
</script>

<template>
  <div class="page">
    <div class="card">
      <h1>Login</h1>

      <form @submit.prevent="handleLogin">
        <input v-model="username" type="text" placeholder="Username" />
        <input v-model="password" type="password" placeholder="Password" />
        <p v-if="error" class="error">{{ error }}</p>
        <button type="submit">ลงชื่อเข้าใช้งาน</button>
      </form>

      <router-link to="/register">สมัครสมาชิก</router-link>
    </div>
  </div>
</template>
