<script setup lang="ts">
import { Loader } from 'lucide-vue-next'
import { toast } from 'vue-sonner'

const { params } = useRoute()

onBeforeMount(async () => {
  try {
    const provider = params?.provider as string
    const res: {
      statusMessage: string
      statusCode: number
      data: { token: string, expires_at: number, workspace_id: string }
    } = await $api(`/auth/signin/oauth/${provider}/callback${window.location.search}`, {
      method: 'GET',
      credentials: 'include',
    })

    const nuxtApiRes = await $fetch('/api/auth/auth_token', {
      method: 'POST',
      body: {
        auth_token: res.data.token,
        expires_at: res.data.expires_at,
      },
    })

    toast.success(nuxtApiRes.message, {
      position: 'top-center',
    })

    return navigateTo(`/workspace`)
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  catch (error: any) {
    const errorMessage = error.response
      ? error.response._data.statusMessage
      : error.message

    toast.error(errorMessage, {
      position: 'top-center',
    })

    return navigateTo(`/auth/signin`)
  }
})
</script>

<template>
  <div class="grid min-h-screen place-content-center">
    <div class="flex flex-col items-center gap-y-0.5">
      <Loader class="animate-spin size-10" />
      <p class="text-sm text-muted-foreground">
        Redirecting
      </p>
    </div>
  </div>
</template>
