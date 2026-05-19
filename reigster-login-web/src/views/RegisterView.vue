<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api/auth'

const router = useRouter()

const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const error = ref('')

const handleRegister = async () => {
  error.value = ''

  if (password.value !== confirmPassword.value) {
    error.value = 'Password ไม่ตรงกัน'
    return
  }

  try {
    await api.post('/register', {
      username: username.value,
      password: password.value,
    })

    alert('สมัครสมาชิกสำเร็จ')

    router.push('/login')
  } catch (err) {
    error.value = 'สมัครสมาชิกไม่สำเร็จ'
  }
}
</script>

<template>
  <div class="page">
    <div class="card">
      <h1>Register</h1>

      <form @submit.prevent="handleRegister">
        <input v-model="username" type="text" placeholder="Username" />
        <input v-model="password" type="password" placeholder="Password" />
        <input v-model="confirmPassword" type="password" placeholder="Confirm Password" />
        <p v-if="error" class="error">{{ error }}</p>
        <button type="submit">สมัครสมาชิก</button>
      </form>

      <router-link to="/login">กลับหน้า Login</router-link>
    </div>
  </div>
</template>
