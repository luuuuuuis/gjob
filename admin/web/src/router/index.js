import { createRouter, createWebHistory } from 'vue-router'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/home' //重定向
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('@/views/home/dashboard.vue')
    },
    {
      path: '/hello',
      name: 'hello',
      component: () => import('@/views/home/hello.vue')
    },
    {
      path: '/grouplist',
      name: 'grouplist',
      component: () => import('@/views/home/grouplist.vue')
    },
    {
      path: '/sqllist',
      name:'sqllist',
      component: () => import('@/views/home/sqllist.vue')
    },
    {
      path: '/scriptlist',
      name:'scriptlist',
      component: () => import('@/views/home/scriptlist.vue')
    },
    {
      path: '/nodelist',
      name:'nodelist',
      component: () => import('@/views/home/nodelist.vue')
    }
  ]
})

export default router
