<script setup lang="ts">
import { useForm } from 'vee-validate'
import { Loader } from 'lucide-vue-next'
import { toast } from 'vue-sonner'
import ResendCodeButton from '../ResendCodeButton.vue'
import { signinFormSchema } from '~/types'
import { cn } from '~/lib/utils'
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '~/components/ui/form'

const props = defineProps<{
  email: string
  onResetForm: ({ mail, codeSent }: { mail: string, codeSent: boolean }) => void
}>()

const isSigningUp = ref(false)
const apiUrl = ref('/api/auth/signup/send-unique-code')

const form = useForm({
  validationSchema: signinFormSchema,
})

const onSubmit = form.handleSubmit(async (values) => {
  try {
    isSigningUp.value = true

    const res: {
      statusMessage: string
      statusCode: number
      data: { token: string, expires_at: number, workspace_id: string }
    } = await $api('/auth/signup', {
      method: 'POST',
      body: {
        email: props.email,
        code: values.code,
      },
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

    props?.onResetForm({
      mail: '',
      codeSent: false,
    })

    return navigateTo(`/workspace/${res.data.workspace_id}/dashboard`)
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  catch (error: any) {
    const errorMessage
      = error.response?._data?.message
        || error.response?._data?.statusMessage
        || error.message

    toast.error(errorMessage, {
      position: 'top-center',
    })
  }
  finally {
    isSigningUp.value = false
  }
})

const isSubmitting = computed(() => {
  if (!form.controlledValues.value.code || form.errors.value.code || isSigningUp.value) {
    return true
  }
  return false
})

const onClear = () => {
  props?.onResetForm({
    mail: '',
    codeSent: false,
  })
  form.resetForm()
  const emailInput = document.querySelector('input[name="email"]') as HTMLInputElement
  const codeInput = document.querySelector('input[name="code"]') as HTMLInputElement

  if (emailInput) {
    emailInput.value = ''
  }
  else if (codeInput) {
    codeInput.value = ''
  }
}
</script>

<template>
  <form
    class="mt-5 space-y-4"
    @submit.prevent="onSubmit"
  >
    <div class="space-y-2">
      <div class="space-y-1">
        <label class="text-sm text-onboarding-text-300 font-medium">Email</label>
        <div
          :class="
            cn(
              'border rounded-md relative flex items-center',
              !props.email && 'border-red-300',
            )
          "
        >
          <input
            name="email"
            type="text"
            placeholder="name@example.com"
            :value="props.email"
            disabled
            class="block bg-transparent text-sm placeholder-custom-text-400 focus:outline-none rounded-md border-custom-border-200 px-3 py-2 disable-autofill-style h-[46px] w-full placeholder:text-onboarding-text-400 border-0"
          >
          <svg
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="bg-background rounded-full absolute right-3 h-5 w-5 hover:cursor-pointer"
            @click="onClear"
          >
            <circle
              cx="12"
              cy="12"
              r="10"
            />
            <path d="m15 9-6 6" />
            <path d="m9 9 6 6" />
          </svg>
        </div>
      </div>
      <FormField
        v-slot="{ componentField }"
        name="code"
      >
        <FormItem class="space-y-1">
          <FormLabel class="text-sm text-onboarding-text-300 font-medium">
            Code
          </FormLabel>
          <FormControl>
            <div
              :class="cn('border rounded-md', form.errors.value.code && 'border-red-300')"
            >
              <input
                type="text"
                autocomplete="off"
                placeholder="gets-sets-flys"
                v-bind="componentField"
                class="block bg-transparent text-sm placeholder-custom-text-400 focus:outline-none rounded-md border-custom-border-200 px-3 py-2 h-[46px] w-full placeholder:text-onboarding-text-400 border-0 focus:bg-none active:bg-transparent"
              >
            </div>
          </FormControl>
          <div class="flex items-center gap-1 text-xs text-red-600 px-0.5">
            <svg
              v-if="form.errors.value.code"
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
          <div class="flex w-full items-center justify-between pt-1">
            <p class="flex items-center gap-1 text-xs font-medium text-green-700">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="12"
                height="12"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="lucide lucide-circle-check"
              >
                <circle
                  cx="12"
                  cy="12"
                  r="10"
                />
                <path d="m9 12 2 2 4-4" />
              </svg>
              Paste the code sent to your email
            </p>
            <ResendCodeButton
              :email="props?.email"
              :api-url="apiUrl"
            />
          </div>
        </FormItem>
      </FormField>
    </div>
    <button
      type="submit"
      :disabled="isSubmitting"
      :class="
        cn(
          'flex w-full items-center justify-center gap-1.5 whitespace-nowrap rounded px-5 py-2 text-sm font-medium text-white transition-all',
          {
            'cursor-pointer bg-brand focus:bg-brand-secondary':
              !form.errors.value.code && form.controlledValues.value.code,
            'cursor-not-allowed bg-[#9e8cce]':
              !form.controlledValues.value.code || form.errors.value.code || isSigningUp,
          },
        )
      "
    >
      <Loader
        v-if="isSigningUp"
        class="size-5 animate-spin"
      />
      Continue
    </button>
  </form>
</template>
