// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    '@nuxtjs/color-mode',
    '@nuxtjs/tailwindcss',
    'shadcn-nuxt',
    '@nuxt/icon',
    '@nuxt/eslint',
    '@nuxtjs/google-fonts',
    'nuxt-og-image',
  ],
  devtools: { enabled: true },
  site: {
    url: process.env.NUXT_PUBLIC_SITE_URL,
    name: process.env.NUXT_PUBLIC_SITE_NAME,
  },
  colorMode: {
    preference: 'system', // default value of $colorMode.preference
    fallback: 'light', // fallback value if not system preference found
    hid: 'nuxt-color-mode-script',
    globalName: '__NUXT_COLOR_MODE__',
    componentName: 'ColorScheme',
    classPrefix: '',
    classSuffix: '',
    storage: 'localStorage', // or 'sessionStorage' or 'cookie'
    storageKey: 'go-community-color-mode',
  },
  runtimeConfig: {
    public: {
      baseURL: process.env.NUXT_PUBLIC_API_URL,
    },
  },
  future: {
    compatibilityVersion: 4,
  },
  compatibilityDate: '2024-11-01',
  nitro: {
    devProxy: {
      '/api/v1/**': {
        target: process.env.NUXT_PUBLIC_API_TARGET,
        changeOrigin: true,
      },
    },
  },
  eslint: {
    config: {
      stylistic: true,
    },
  },
  googleFonts: {
    families: {
      Geist: true,
    },
  },
  shadcn: {
    /**
     * Prefix for all the imported component
     */
    prefix: '',
    /**
     * Directory that the component lives in.
     * @default "./components/ui"
     */
    componentDir: './components/ui',
  },
})
