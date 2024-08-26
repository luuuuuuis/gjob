import { createRouter, createWebHistory } from 'vue-router'
//导入进度条组件
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',  //url路径
      component: () => import('@/views/login/Login.vue'),  //视图组件
      icon: "odometer",  //图标
      meta: {title: "登录", requireAuth: false},  //meta元信息
    },
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

//递增进度条，这将获取当前状态值并添加0.2直到状态为0.994
NProgress.inc(100)
//easing 动画字符串
//speed 动画速度
//showSpinner 进度环显示隐藏
NProgress.configure({ easing: 'ease', speed: 600, showSpinner: false })

// 导入md5
import md5 from 'md5';
//路由守卫，路由拦截
router.beforeEach((to, from, next) => {
  // 启动进度条
  NProgress.start()
  //设置头部
  if (to.meta.title) {
      document.title = to.meta.title
  } else {
      document.title = "GJob"
  }
  // 放行
  if (window.location.pathname == '/login') {
      next()
  }else{
      // 获取localStorage中保存的Token和过期时间
      const storedToken = localStorage.getItem('token');
      const storedTokenExpireTime = localStorage.getItem('tokenExpireTime');
      // 如果没有保存Token或过期时间，或者Token已经过期，则跳转到登录页面

      if (!storedToken || !storedTokenExpireTime || Date.now() > parseInt(storedTokenExpireTime)) {
          // 删除localStorage中保存的Token和过期时间
          localStorage.removeItem('token');
          localStorage.removeItem('tokenExpireTime');

          // 如果当前不在登录页面，则跳转到登录页面
          if (window.location.pathname !== '/login') {
              window.location.href = '/login';
          }
      } else {
          // 验证Token是否正确
          const salt = localStorage.getItem('username')+localStorage.getItem('loginDate')
          const token = md5(salt); // 使用md5算法生成Token

          if (token === storedToken) {
              // Token正确，且在有效期内，继续进行其他操作
              // TODO: 继续访问
              next()
          } else {
              // Token错误，跳转到登录页面
              localStorage.removeItem('token');
              localStorage.removeItem('tokenExpireTime');

              // 如果当前不在登录页面，则跳转到登录页面
              if (window.location.pathname !== '/login') {
                  window.location.href = '/login';
              }
          }
      }
  }
})

router.afterEach(() => {
  // 关闭进度条
  NProgress.done()
})

// 抛出路由实例, 在 main.js 中引用
export default router
