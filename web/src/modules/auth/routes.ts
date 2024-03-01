import { RouteRecordRaw } from 'vue-router'
import Login from './login.vue'
import Forgot from './forgot.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    component: Login,
  },
  {
    path: '/reset',
    component: Forgot,
  },
]

export default routes
