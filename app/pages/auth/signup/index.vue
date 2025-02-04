<script setup lang="ts">
import OauthProviders from '~/components/auth/oauth/OauthProviders.vue'
import SignupForm from '~/components/auth/signup/SignupForm.vue'
import SvgBackgroundPattern from '~/components/svgs/SvgBackgroundPattern.vue'
import SvgUseOdama from '~/components/svgs/SvgUseOdama.vue'
</script>

<template>
  <div class="relative flex h-screen w-full flex-col overflow-hidden">
    <div class="relative size-full overflow-hidden">
      <div class="h-screen w-full overflow-hidden">
        <div class="relative h-screen w-screen overflow-hidden">
          <div class="absolute inset-0 z-0 dark:hidden">
            <SvgBackgroundPattern class="size-full" />
          </div>
          <div
            class="relative z-10 flex h-screen w-screen flex-col overflow-hidden overflow-y-auto"
          >
            <div
              class="container relative flex min-w-full shrink-0 items-center justify-between px-10 pb-4 transition-all lg:px-20 xl:px-36"
            >
              <div class="flex items-center gap-x-2 py-10">
                <a
                  class="h-[30px] w-[133px]"
                  href="/"
                >
                  <SvgUseOdama />
                </a>
              </div>
              <div
                class="flex flex-col items-end text-center text-sm font-medium sm:flex-row sm:items-center sm:gap-2"
              >
                Already have an account?
                <NuxtLink
                  class="font-semibold text-brand dark:text-primary hover:underline"
                  to="/auth/signin"
                >
                  Signin
                </NuxtLink>
              </div>
            </div>
            <div
              class="container mx-auto flex h-[100vh-60px] max-w-lg grow flex-col justify-center px-10 transition-all lg:max-w-md lg:px-5"
            >
              <div class="relative flex flex-col space-y-6">
                <div class="space-y-1 text-center">
                  <h3 class="text-3xl font-bold text-brand dark:text-primary">
                    Get Started Now
                  </h3>
                  <p class="text-sm text-muted-foreground">
                    Let's get your account setup in less than 2 minutes.
                  </p>
                </div>
                <SignupForm />
                <div class="mt-4 flex items-center">
                  <hr class="w-full">
                  <p class="mx-3 shrink-0 text-center text-sm">
                    or
                  </p>
                  <hr class="w-full">
                </div>
                <OauthProviders />
                <span class="flex items-center justify-center py-6">
                  <p class="whitespace-pre-line text-center text-sm">
                    By signing in, you agree to our
                    <a
                      target="_blank"
                      rel="noopener noreferrer"
                      href="https://plane.so/legals/terms-and-conditions"
                    ><span
                      class="text-sm font-medium underline hover:cursor-pointer"
                    >Terms of Service</span></a>
                    and
                    <a
                      target="_blank"
                      rel="noopener noreferrer"
                      href="https://plane.so/legals/privacy-policy"
                    ><span
                      class="text-sm font-medium underline hover:cursor-pointer"
                    >Privacy Policy</span></a>.
                  </p>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
