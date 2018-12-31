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
      path: '/login',
      name: 'LoginPage',
      component: () => import('@/pages/LoginPage'),
      meta: {layout: 'default'}
    },
    {
      path: '/main',
      name: 'MainPage',
      component: () => import('@/pages/MainPage'),
      meta: {layout: 'default'}
    },
    {
      path: '/post/list',
      name: 'PostListPage',
      component: () => import('@/pages/post/PostListPage'),
      meta: {layout: 'default'}
    },
    {
      path: '/post/register',
      name: 'PostWorkPage',
      component: () => import('@/pages/post/PostWorkPage'),
      meta: {layout: 'common'}
    }
  ]
})
