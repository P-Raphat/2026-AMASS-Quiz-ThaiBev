<script setup lang="ts"></script>

<template>
  <router-view />
</template>

<style>
*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: 'Segoe UI', sans-serif;
  background: #f0f2f5;
  min-height: 100vh;
}

.page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.card {
  background: #fff;
  padding: 2.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 380px;
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.card h1 {
  font-size: 1.75rem;
  text-align: center;
  color: #1a1a2e;
}

.card form {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.card input {
  padding: 0.65rem 1rem;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 0.95rem;
  outline: none;
}

.card input:focus {
  border-color: #4f46e5;
}

.card button {
  padding: 0.7rem;
  background: #4f46e5;
  color: #fff;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
}

.card button:hover {
  background: #4338ca;
}

.card a {
  text-align: center;
  color: #4f46e5;
  font-size: 0.9rem;
  text-decoration: none;
}

.card a:hover {
  text-decoration: underline;
}

.error {
  color: #dc2626;
  font-size: 0.875rem;
}
</style>
