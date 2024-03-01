<script lang="ts">
import { computed, defineComponent, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import IconSearch from '../../icons/IconSearch.vue'

export default defineComponent({
  components: { IconSearch },
  name: 'TrendingSidebar',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const searchFocused = ref(false)
    const searchQuery = ref('')

    const isInSearchPage = computed(() => route.path === '/search')

    function redirectWithSearchQuery() {
      router.push({
        path: '/search',
        query: { q: searchQuery.value },
      })
      return
    }

    return {
      searchFocused,
      searchQuery,
      redirectWithSearchQuery,
      isInSearchPage,
    }
  },
})
</script>

<template>
  <div
    class="
      md:block
      hidden
      w-1/2
      h-full
      border-l border-lighter
      dark:border-dark
      py-2
      px-6
      overflow-y-scroll
      relative
    "
  >
    <form v-show="!isInSearchPage" @submit.prevent="redirectWithSearchQuery">
      <input
        class="
          pl-12
          rounded-full
          w-full
          p-3
          bg-lightest
          dark:bg-darkest
          dark:text-light
          text-sm
          mb-4
          focus:bg-white
          dark:focus:bg-black
          focus:outline-none
          border-2 border-lighter
          dark:border-darkest
          focus:border-blue
          dark:focus:border-blue
          dark:focus:text-lightest
          transition
          duration-150
        "
        @focus="searchFocused = true"
        @blur="searchFocused = false"
        v-model="searchQuery"
        type="search"
        placeholder="Search T2"
      />
      <IconSearch
        :size="24"
        class="absolute left-0 top-0 mt-5 ml-10"
        :class="searchFocused ? 'text-blue' : 'text-light'"
      />
      <input type="submit" class="hidden" />
    </form>
    <div
      class="
        w-full
        rounded-2xl
        bg-lightest
        dark:bg-darkest
        my-4
        hidden
        lg:block
      "
    >
      <div class="p-4">
        <br />
      <p class="text-sm leading-tight dark:text-lightest">
            Everything you write is publicly visible.
      </p> <br />
      <p class="text-sm leading-tight dark:text-lightest">
            Be kind to each other.
      </p> <br />
      <p class="text-sm leading-tight dark:text-lightest">
            This is an early prototype.
      </p> <br />
      </div>

    </div>
  </div>
</template>
