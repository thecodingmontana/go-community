<script setup lang="ts">
import { Toaster } from '~/components/ui/sonner'

useHead({
  meta: [
    { property: 'og:title', content: `use-odama` },
    {
      name: 'description',
      content: 'Diversify your revenue streams. Go from creator, to business.',
    },
    {
      property: 'og:description',
      content: 'Diversify your revenue streams. Go from creator, to business.',
    },
    { property: 'og:url', content: process.env.NUXT_PUBLIC_SITE_URL },
    { name: 'twitter:title', content: 'Nuxt Nitro Chat' },
    {
      name: 'twitter:description',
      content: 'Diversify your revenue streams. Go from creator, to business.',
    },
    { name: 'twitter:site', content: process.env.NUXT_PUBLIC_SITE_URL },
  ],
  title: `use-odama`,
})

defineOgImageComponent('Nuxt', {
  headline: 'Greetings',
  title: 'Hello OG Image 👋',
  description: 'Look what at me using the Nuxt template',
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
