<script setup lang="ts">
definePageMeta({
  layout: 'chat',
})

defineOgImageComponent('Nuxt', {
  headline: 'Hello 👋',
  title: 'Go Community - Chat',
  description: 'Go Community is a real-time app using Go WebSockets (Chi), PostgreSQL, and Nuxt.js. 🚀',
})

useHead({
  titleTemplate: '%s - Chat',
})
</script>

<template>
  <section>
    Chat
  </section>
</template>
