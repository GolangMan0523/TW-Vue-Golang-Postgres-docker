import { createRouter, createWebHistory } from 'vue-router'
import { store } from './store'
import routes from './modules/routes'
import { Action } from './modules/storeActionTypes'

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to, from, next) => {
  if (store.getters['accessToken']) {
    await store.dispatch(Action.AuthActionTypes.REFRESH_AUTH_TOKEN)
    await store.dispatch(Action.AuthActionTypes.GET_USER_DATA)
  }

  if (to.matched.some((route) => route.meta.requiresAuth)) {
    return store.getters['isLoggedIn'] ? next() : next({ path: '/login', query: { redirect: to.fullPath } })
  }
  return next()
})

