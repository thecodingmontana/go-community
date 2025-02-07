<script setup lang="ts">
import { Toaster } from '~/components/ui/sonner'

useHead({
  meta: [
    { property: 'og:title', content: `Go Community` },
    {
      name: 'description',
      content: 'Go Community is a real-time web app using Go WebSockets (Chi Router), PostgreSQL, and a Nuxt.js frontend, showcasing WebSocket integration in Go with JS frameworks (Nuxtjs 😜) 🚀.',
    },
    {
      property: 'og:description',
      content: 'Go Community is a real-time web app using Go WebSockets (Chi Router), PostgreSQL, and a Nuxt.js frontend, showcasing WebSocket integration in Go with JS frameworks (Nuxtjs 😜) 🚀.',
    },
    { property: 'og:url', content: process.env.NUXT_PUBLIC_SITE_URL },
    { name: 'twitter:title', content: 'Go Community' },
    {
      name: 'twitter:description',
      content: 'Go Community is a real-time web app using Go WebSockets (Chi Router), PostgreSQL, and a Nuxt.js frontend, showcasing WebSocket integration in Go with JS frameworks (Nuxtjs 😜) 🚀.',
    },
    { name: 'twitter:site', content: process.env.NUXT_PUBLIC_SITE_URL },
  ],
  title: `Go Community`,
})

defineOgImageComponent('Nuxt', {
  headline: 'Hello 👋',
  title: 'Go Community',
  description: 'Go Community is a real-time app using Go WebSockets (Chi), PostgreSQL, and Nuxt.js. 🚀',
})
</script>

<template>
  <main>
    <NuxtLoadingIndicator />
    <slot />
    <Toaster
      :rich-colors="true"
      :close-button="true"
    />
  </main>
</template>
