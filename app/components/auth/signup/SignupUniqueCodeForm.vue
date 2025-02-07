<script setup lang="ts">
import { Loader } from 'lucide-vue-next'
import { useForm } from 'vee-validate'
import { toast } from 'vue-sonner'
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '~/components/ui/form'
import { cn } from '~/lib/utils'
import { sendUniqueCodeSchema } from '~/types'

const props = defineProps<{
  onResetForm: ({ mail, codeSent }: { mail: string, codeSent: boolean }) => void
}>()

const isSendingCode = ref(false)

const sendUniqueCodeForm = useForm({
  validationSchema: sendUniqueCodeSchema,
})

const isSubmitting = computed(() => {
  if (
    !sendUniqueCodeForm.controlledValues.value.email
    || sendUniqueCodeForm.errors.value.email
    || isSendingCode.value
  ) {
    return true
  }
  return false
})

const onSendUniqueCode = sendUniqueCodeForm.handleSubmit(async (values) => {
  try {
    isSendingCode.value = true

    const res: { statusMessage: string, statusCode: number } = await $api(
      '/auth/signup/send-unique-code',
      {
        method: 'POST',
        body: {
          email: values.email,
        },
      },
    )

    props?.onResetForm({
      mail: values.email,
      codeSent: true,
    })

    const message = res.statusMessage
      ? res.statusMessage
      : 'Check your email for the verification code!'

    return toast.success(message, {
      position: 'top-center',
    })
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  catch (error: any) {
    const errorMessage = error.response
      ? error.response._data.statusMessage
      : error.message

    toast.error(errorMessage, {
      position: 'top-center',
    })
  }
  finally {
    isSendingCode.value = false
  }
})
</script>

<template>
  <form
    class="mt-5 space-y-4"
    @submit.prevent="onSendUniqueCode"
  >
    <FormField
      v-slot="{ componentField }"
      name="email"
    >
      <FormItem class="space-y-1">
        <FormLabel class="text-onboarding-text-300 text-sm font-medium">
          Email
        </FormLabel>
        <FormControl>
          <div
            :class="
              cn(
                'border rounded-md',
                sendUniqueCodeForm.errors.value.email && 'border-red-300',
              )
            "
          >
            <input
              type="text"
              placeholder="name@example.com"
              v-bind="componentField"
              class="placeholder:text-custom-text-400 border-custom-border-200 placeholder:text-onboarding-text-400 block h-[46px] w-full rounded-md border-0 bg-transparent px-3 py-2 text-sm focus:bg-none focus:outline-none active:bg-transparent"
            >
          </div>
        </FormControl>
        <div class="flex items-center gap-1 px-0.5 text-xs text-red-600">
          <svg
            v-if="sendUniqueCodeForm.errors.value.email"
            xmlns="http://www.w3.org/2000/svg"
            width="12"
            height="12"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="lucide lucide-circle-alert"
          >
            <circle
              cx="12"
              cy="12"
              r="10"
            />
            <line
              x1="12"
              x2="12"
              y1="8"
              y2="12"
            />
            <line
              x1="12"
              x2="12.01"
              y1="16"
              y2="16"
            />
          </svg>
          <FormMessage />
        </div>
      </FormItem>
    </FormField>
    <button
      type="submit"
      :disabled="isSubmitting"
      :class="
        cn(
          'flex w-full items-center justify-center gap-1.5 whitespace-nowrap rounded px-5 py-2 text-sm font-medium text-white dark:text-black transition-all',
          {
            'cursor-pointer bg-brand dark:bg-primary focus:bg-brand-secondary':
              sendUniqueCodeForm.controlledValues.value.email
              && !sendUniqueCodeForm.errors.value.email,
            'cursor-not-allowed bg-[#9e8cce] dark:bg-muted-foreground':
              !sendUniqueCodeForm.controlledValues.value.email
              || sendUniqueCodeForm.errors.value.email
              || isSendingCode,
          },
        )
      "
    >
      <Loader
        v-if="isSendingCode"
        class="size-5 animate-spin"
      />
      Continue
    </button>
  </form>
</template>
