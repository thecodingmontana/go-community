import type { NitroFetchOptions, NitroFetchRequest } from 'nitropack'

export async function $api<
  DefaultT = unknown,
  DefaultR extends NitroFetchRequest = NitroFetchRequest,
  T = DefaultT,
  R extends NitroFetchRequest = DefaultR,
  O extends NitroFetchOptions<R> = NitroFetchOptions<R>,
>(url: R, options?: O) {
  const config = useRuntimeConfig()
  const auth_token = useCookie('auth_token')

  return await $fetch<T>(url, {
    ...options,
    baseURL: config.public.baseURL,
    headers: useRequestHeaders(['authorization', 'x-tenant-id']),
    onRequest({ options }) {
      // Set the request headers
      options.headers = {
        ...options.headers,
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        authorization: `Bearer ${auth_token.value}`,
      }
    },
    onResponseError({ request, response, options }) {
      const auth_token = useCookie('auth_token')

      if (response._data.errorCode === 401) {
        auth_token.value = ''
        navigateTo('/session-expired')
      }
      console.log(request, response, options)
    },
  })
}
