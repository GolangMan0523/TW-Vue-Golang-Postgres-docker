<script lang="ts">
import { computed, defineComponent, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import IconTwitter from '../../icons/IconTwitter.vue'
import LoadingSpinner from '../../shared/LoadingSpinner.vue'
import { useStore } from '../../store'
import { Action } from '../storeActionTypes'
import { resetPassword } from './service'

export default defineComponent({
  name: 'Forgot',
  components: {
    IconTwitter,
    LoadingSpinner,
  },
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()

    const email = ref('')
    const loading = ref(false)
    const message = ref("")

    const input = reactive({ email })

    const inputEmpty = computed(
      () => input.email === ''
    )

    async function handleResetPassword() {
      loading.value = true
      const response = await resetPassword(email.value)
      console.log(response)
    }
    return { input, loading, handleResetPassword, inputEmpty }
  },
})
</script>

<template>
  <div
    class="
      flex
      justify-center
      container
      mx-auto
      h-screen
      w-1/4
      px-4
      lg:px-0
      py-8
    "
  >
    <div>
      <IconTwitter :size="60" class="h-12 w-12 text-blue" />
      <h1 class="pt-12 text-4xl dark:text-lightest font-bold">
        Forgot Password
      </h1>
      <form @submit.prevent="handleResetPassword" class="w-full text-center">
        <input
          v-model="input.email"
          type="text"
          placeholder="Email"
          class="
            w-full
            px-2
            py-4
            my-6
            border-2 border-lighter
            text-xl
            rounded
            dark:border-dark
            focus:outline-none
            dark:bg-black
            dark:text-light
            focus:border-blue
            dark:focus:border-blue
            transition-colors
            duration-75
          "
        />
        <label>If your email matches our records, we will send you a password reset link.</label>
        <button
          type="submit"
          class="bg-blue rounded-full focus:outline-none w-full h-auto p-4 mt-4"
          :class="
            inputEmpty
              ? 'cursor-not-allowed'
              : 'cursor-pointer hover:bg-darkblue'
          "
          :disabled="inputEmpty"
        >
          <LoadingSpinner v-if="loading" color="white" size="36px" />
          <span v-else class="text-lightest text-lg font-semibold">Send password reset email</span>
        </button>
        <div class="mt-4">
          <router-link to="/login">
            <span class="text-blue">Log into T2</span>
          </router-link>
        </div>
        <div class="mt-2">
          <router-link to="/">
            <span class="text-blue">Sign up for T2</span>
          </router-link>
        </div>
      </form>
    </div>
  </div>
</template>
