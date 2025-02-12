<script setup lang="ts">
import ChatUserMessage from './ChatUserMessage.vue'
import type { ChatMessage } from '~/types'

const props = defineProps<{
  message: ChatMessage
}>()

const user = useAuthenticatedUser()
</script>

<template>
  <div>
    <div
      class="w-full p-4 space-y-2"
    >
      <div
        v-if="user.id !== props?.message.user.id"
        class="grid"
      >
        <!-- Sender Message -->
        <ChatSenderMessage
          :message="props?.message"
        />
        <!-- End Of Sender Message -->
      </div>
      <!-- User Message -->
      <div
        v-else
        class="flex gap-2.5 justify-end"
      >
        <ChatUserMessage :message="message" />
      </div>
      <!-- End of User Message -->
    </div>
  </div>
</template>
