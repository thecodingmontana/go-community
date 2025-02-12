<script setup lang="ts">
import { useBrowserLocation, useWebSocket } from '@vueuse/core'
import ChatHeader from '~/components/chat/ChatHeader.vue'
import ChatMessages from '~/components/chat/ChatMessages.vue'
import ChatInput from '~/components/chat/ChatInput.vue'
import type { ChatMessage, SocketData, Stats } from '~/types'

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

const location = useBrowserLocation()
const config = useRuntimeConfig()
const isSecure = location.value.protocol === 'https:'
const messages = ref<ChatMessage[]>([])
const stats = ref<Stats>({ onlineUsers: 0, totalUsers: 0 })

const host = config.public.baseURL.replace(/^https?:\/\//, '')

const url = (isSecure ? 'wss://' : 'ws://') + host + '/ws/chat'

const { status, send } = useWebSocket(url, {
  onMessage: (ws, event) => {
    const data: SocketData = JSON.parse(event.data)
    handleWebSocketMessage(data)
  },
})

const handleWebSocketMessage = (data: SocketData) => {
  switch (data.type) {
    case 'message':
      messages.value.push(data.payload as ChatMessage)
      break
    case 'history':
      messages.value = data.payload as ChatMessage[]
      console.log(messages.value)
      break
    case 'stats':
      console.log(data.payload)
      break
  }
}
</script>

<template>
  <section class="flex h-screen w-full flex-col bg-[#f8f8f8]">
    <ChatHeader
      :status="status"
      :stats="stats"
    />
    <ChatMessages :messages="messages" />
    <ChatInput :send="send" />
  </section>
</template>
