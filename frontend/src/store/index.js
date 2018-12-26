import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const stroe = new Vuex.Store({
  state: {
    user: {}
  },
  mutations: {
    setUser: (state, value) => {
      state.user = value
    }
  },
  getters: {
    getUser: (state) => state.user
  }
})
