import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'RootPage',
      component: () => import('@/pages/RootPage'),
      meta: {layout: 'default'}
    },
    {
      path: '/main',
      name: 'MainPage',
      component: () => import('@/pages/MainPage'),
      meta: {layout: 'default'}
    },
    {
      path: '/post/register',
      name: 'PostRegPage',
      component: () => import('@/pages/post/PostRegPage'),
      meta: {layout: 'common'}
    }
  ]
})
