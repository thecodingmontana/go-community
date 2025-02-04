<script setup lang="ts">
const config = useRuntimeConfig()

const oauthSignin = (provider: string) => {
  window.location.href = `${config.public.baseURL}/auth/signin/oauth/${provider}`
}
</script>

<template>
  <div class="mt-7 grid gap-4 overflow-hidden">
    <div class="flex h-[42px] items-center !overflow-hidden">
      <button
        class="flex h-[42px] w-full items-center justify-center gap-2 rounded border px-2 text-sm font-medium duration-300 hover:border-purple-200"
        @click="oauthSignin('google')"
      >
        <Icon
          name="devicon:google"
          class="size-4"
        />
        Continue with Google
      </button>
    </div>
    <button
      class="flex h-[42px] w-full items-center justify-center gap-2 rounded border px-2 text-sm font-medium duration-300 hover:border-purple-200"
      @click="oauthSignin('github')"
    >
      <Icon
        name="akar-icons:github-fill"
        class="size-4"
      />
      Continue with Github
    </button>
  </div>
</template>
