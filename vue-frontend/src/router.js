import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import Login from './components/Login.vue'
import Group from './components/Group.vue'
import Manage from './components/Manage.vue'
import Admin from './components/Admin.vue'
import Stat from './components/Stat.vue'
import Admin2 from './components/Admin2.vue'
import axios from 'axios'
import store from './main'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: "/",
      name: "index",
      component: HelloWorld,
      title: '主页'
    },
    {
      path: "/login",
      name: "login",
      component: Login,
      title: '登录'
    },
    {
      path: "/manage",
      name: "manage",
      component: Manage,
      title: "队伍管理"
    },
    {
      path: "/admin",
      name: "admin",
      component: Admin,
      title: "管理员"
    },
    {
      path: "/admin2",
      name: "admin2",
      component: Admin2,
      title: "管理员"
    },
    {
      path: "/group",
      name: "group",
      component: Group
    },
    {
      path: "/stat",
      name: "stat",
      component: Stat
    }
  ]
})

router.beforeEach((to, _from, next) => {
  if (to.path == '/manage') {
    axios.get('http://47.100.50.175:8088/api/player/status', {
      headers: { Authorization: store.state.jwt }
    }).then(response => {
      store.commit('set_school', response.data.school)
      store.commit('set_sid', response.data.sid)
      next()
    }).catch(() => {
      next('/login')
    })
  } else if (to.path == '/login') {
    axios.get('http://47.100.50.175:8088/api/player/status', {
      headers: { Authorization: store.state.jwt }
    }).then(response => {
      store.commit('set_school', response.data.school)
      store.commit('set_sid', response.data.sid)
      next('/manage')
    }).catch(() => {
      next()
    })
  }
  next()
})

export default router