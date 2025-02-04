<script setup lang="ts">
import { CheckIcon, CopyIcon } from 'lucide-vue-next'
import { toast } from 'vue-sonner'
import { Button } from '../ui/button'
import { Input } from '../ui/input'
import { cn } from '~/lib/utils'

const props = defineProps<{
  text: string
}>()

const isCopied = ref(false)
const text = ref(props?.text)

const copyToClipboard = async () => {
  isCopied.value = true
  setTimeout(() => {
    isCopied.value = false
  }, 2000)

  await navigator.clipboard.writeText(props?.text)
  toast.success('Copied to clipboard', {
    position: 'top-center',
  })
}
</script>

<template>
  <div class="flex justify-center gap-3">
    <Input
      v-model="text"
      readonly
      class="bg-secondary text-muted-foreground"
    />
    <Button
      size="icon"
      @click="copyToClipboard"
    >
      <CheckIcon
        v-if="isCopied"
        :class="cn(isCopied ? 'opacity-100' : 'opacity-0', '')"
      />
      <CopyIcon
        v-else
        class="size-5"
      />
    </Button>
  </div>
</template>
