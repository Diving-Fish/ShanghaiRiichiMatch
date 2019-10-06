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
    set_jwt(state, s) {
      state.jwt = s
    },
    set_school(state, s) {
      state.school = s
    },
    set_sid(state, i) {
      state.sid = i
    }
  }
})

export default store

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
