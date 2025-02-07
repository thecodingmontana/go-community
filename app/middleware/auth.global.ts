import type { User } from '~/types'
import { $api } from '~/composables/api'
import { useUser } from '~/composables/auth'

interface APIResponse {
  statusMessage: string
  statusCode: number
  data: User
}

export default defineNuxtRouteMiddleware(async (to) => {
  const user = useUser()

  try {
    const response: APIResponse = await $api('/auth/user', {
      method: 'GET',
    })

    if (response.data) {
      user.value = response.data

      // Redirect authenticated users to chat if they try to access auth pages
      if (to.path.startsWith('/auth')) {
        return navigateTo('/chat')
      }
    }
    else {
      // Only redirect to signin if trying to access chat routes
      if (to.path.startsWith('/chat')) {
        return navigateTo('/auth/signin')
      }
    }
  }
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  catch (error) {
    // Only redirect to signin if trying to access chat routes
    if (to.path.startsWith('/chat')) {
      return navigateTo('/auth/signin')
    }
  }
})
