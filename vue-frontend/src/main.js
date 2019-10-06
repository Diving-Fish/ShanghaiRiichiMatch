import Vue from 'vue'
import Vuex from 'vuex'
import App from './App.vue'
import './plugins/element.js'
import router from './router'

Vue.config.productionTip = false

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    jwt: 'default',
    team_id: 0,
    admin_key: ''
  },
  mutations: {
    set_jwt(state, j) {
      state.jwt = j
    },
    set_team_id(state, i) {
      state.team_id = i
    },
    set_admin_key(state, admin_key) {
      state.admin_key = admin_key
    }
  }
})

export default store

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
