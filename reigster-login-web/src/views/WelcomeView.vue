<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const username = authStore.username

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <div class="page">
    <div class="card" style="align-items: center;">
      <h1>Welcome</h1>
      <p style="color: #4f46e5; font-weight: 600;">{{ username }}</p>
      <button @click="logout" style="background: #ef4444; padding: 0.7rem 2rem;">Logout</button>
    </div>
  </div>
</template>
