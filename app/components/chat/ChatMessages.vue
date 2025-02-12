<script setup lang="ts">
import ChatItem from './ChatItem.vue'
import type { ChatMessage } from '~/types'

const props = defineProps<{
  messages: ChatMessage[]
}>()

const bottomRef = ref<HTMLDivElement | null>(null)

watch(props?.messages, () => {
  setTimeout(() => {
    bottomRef.value?.scrollIntoView({
      behavior: 'smooth',
    })
  }, 100)
})

onMounted(() => {
  setTimeout(() => {
    bottomRef.value?.scrollIntoView({
      behavior: 'smooth',
    })
  }, 100)
})
</script>

<template>
  <div
    class="flex-1 flex flex-col py-4 overflow-y-auto [&::-webkit-scrollbar]:w-[0.2rem] [&::-webkit-scrollbar-track]:rounded-full [&::-webkit-scrollbar-track]:bg-gray-100 [&::-webkit-scrollbar-thumb]:rounded-full [&::-webkit-scrollbar-thumb]:bg-gray-300 dark:[&::-webkit-scrollbar-track]:bg-neutral-700 dark:[&::-webkit-scrollbar-thumb]:bg-neutral-500 dark:bg-background"
  >
    <ChatItem
      v-for="message in props.messages"
      :key="message.id"
      :message="message"
    />
    <div ref="bottomRef" />
  </div>
</template>
