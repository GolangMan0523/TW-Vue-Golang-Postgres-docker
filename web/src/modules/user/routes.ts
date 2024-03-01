import { RouteRecordRaw } from 'vue-router'
import BaseLayout from '../../modules/layouts/BaseLayout.vue'
import { makeRoutesWithLayout } from '../../utils/routes'

const routes: RouteRecordRaw[] = [
  {
    path: '',
    component: () => import('./index.vue'),
  },
  {
    path: 'status/:tweetId',
    component: () => import('./status/index.vue'),
  },
]

export default makeRoutesWithLayout('/:name', BaseLayout, routes)
