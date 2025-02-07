<script setup lang="ts">
import CopyToClipboard from '~/components/homepage/CopyToClipboard.vue'
import SvgGopher from '~/components/svgs/SvgGopher.vue'
import { Button } from '~/components/ui/button'
import { githubUrl } from '~/types'
</script>

<template>
  <section
    class="mx-auto grid min-h-[calc(100vh-300px)] max-w-5xl flex-col items-center justify-center gap-4 text-center"
  >
    <div class="p-4">
      <div class="mb-2 flex items-center justify-center gap-3">
        <SvgGopher class="h-[5.5rem] w-auto" />
      </div>
      <h1
        class="text-balance bg-gradient-to-tr from-black/70 via-black to-black/60 bg-clip-text text-center text-3xl font-bold text-transparent dark:from-zinc-400/10 dark:via-white/90 dark:to-white/20 sm:text-5xl md:text-6xl lg:text-7xl"
      >
        Go Community
      </h1>
      <p
        class="mb-10 mt-4 text-balance text-center text-muted-foreground md:text-lg lg:text-xl"
      >
        Go Community is a real-time web app using Go WebSockets (Chi Router), PostgreSQL, and a Nuxt.js frontend, showcasing WebSocket integration in Go with JS frameworks (Nuxtjs 😜) 🚀.
      </p>
      <div class="mb-10">
        <div class="mx-auto max-w-[430px]">
          <CopyToClipboard :text="`git clone ${githubUrl}`" />
        </div>
      </div>
      <div class="flex justify-center gap-4">
        <Button
          size="lg"
          variant="outline"
          as-child
        >
          <NuxtLink
            target="_blank"
            :to="githubUrl"
          >
            <Icon
              name="tabler:brand-github"
              class="mr-1 size-5"
            />
            GitHub
          </NuxtLink>
        </Button>
        <Button
          size="lg"
          as-child
        >
          <NuxtLink to="/auth/signin">
            Get Started
          </NuxtLink>
        </Button>
      </div>
    </div>
  </section>
</template>
