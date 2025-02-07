<script setup lang="ts">
import OauthProviders from '~/components/auth/oauth/OauthProviders.vue'
import SigninForm from '~/components/auth/signin/SigninForm.vue'
import SvgBackgroundPattern from '~/components/svgs/SvgBackgroundPattern.vue'
import SvgGopher from '~/components/svgs/SvgGopher.vue'

definePageMeta({
  layout: 'auth',
})

defineOgImageComponent('Nuxt', {
  headline: 'Hello 👋',
  title: 'Go Community - Signin',
  description: 'Go Community is a real-time app using Go WebSockets (Chi), PostgreSQL, and Nuxt.js. 🚀',
})

useHead({
  titleTemplate: '%s - Signin',
})
</script>

<template>
  <div class="relative flex h-screen w-full flex-col overflow-hidden">
    <div class="relative size-full overflow-hidden">
      <div class="h-screen w-full overflow-hidden">
        <div class="relative h-screen w-screen overflow-hidden">
          <div class="absolute inset-0 z-0 dark:hidden">
            <SvgBackgroundPattern class="size-full" />
          </div>
          <div class="relative z-10 flex h-screen w-screen flex-col overflow-hidden overflow-y-auto">
            <div
              class="container mx-auto flex h-[100vh-60px] max-w-lg grow flex-col justify-center px-10 transition-all lg:max-w-md lg:px-5"
            >
              <div class="relative flex flex-col space-y-6">
                <NuxtLink
                  class="flex flex-col items-center"
                  to="/"
                >
                  <SvgGopher class="h-16 w-auto" />
                </NuxtLink>
                <div class="space-y-1 text-center">
                  <h3 class="text-3xl font-bold text-brand dark:text-primary">
                    Hello Again!
                  </h3>
                  <p class="text-sm text-muted-foreground">
                    Sign in to access your account.
                  </p>
                </div>
                <SigninForm />
                <div class="mt-4 flex items-center">
                  <hr class="w-full">
                  <p class="mx-3 shrink-0 text-center text-sm">
                    or
                  </p>
                  <hr class="w-full">
                </div>
                <OauthProviders />
                <div class="flex flex-col items-center text-center text-sm font-medium sm:flex-row sm:items-center sm:justify-center sm:gap-2">
                  New to go-community?
                  <NuxtLink
                    class="font-semibold text-brand hover:underline dark:text-primary"
                    to="/auth/signup"
                  >
                    Signup
                  </NuxtLink>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
