<script setup lang="ts">
import SignupUniqueCodeForm from './SignupUniqueCodeForm.vue'
import SignupCompleteForm from './SignupCompleteForm.vue'

const email = ref('')
const isCodeSent = ref(false)

const onResetForm = ({ mail, codeSent }: { mail: string, codeSent: boolean }) => {
  email.value = mail
  isCodeSent.value = codeSent
}
</script>

<template>
  <SignupCompleteForm
    v-if="isCodeSent"
    :email="email"
    :on-reset-form="onResetForm"
  />
  <SignupUniqueCodeForm
    v-else
    :on-reset-form="onResetForm"
  />
</template>
