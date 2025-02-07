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
  <section class="flex h-screen w-full flex-col bg-[#f8f8f8]">
    <!-- <ChatHeader :group="props?.group" :status="status" />
    <ChatMessages
      :socketMessages="messages"
      :id="group?.id as string"
      :profileId="profile?.id"
    />
    <ChatInput :send="send" /> -->
    <!-- <ChatInput :send="send" /> -->
  </section>
</template>
