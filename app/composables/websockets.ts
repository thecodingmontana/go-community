import { useBrowserLocation, useWebSocket } from '@vueuse/core'

export function useWs() {
  const location = useBrowserLocation()
  const config = useRuntimeConfig()
  const isSecure = location.value.protocol === 'https:'

  const host = config.public.baseURL.replace(/^https?:\/\//, '')

  const url = (isSecure ? 'wss://' : 'ws://') + host + '/ws/chat'

  const { status, data, send, open, close } = useWebSocket(url)

  return { status, data, send, open, close }
}
