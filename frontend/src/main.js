// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import VueCookie from 'vue-cookie'
import 'es6-promise/auto'

Vue.config.productionTip = false
Vue.prototype.$http = (url, data) => {
  return new Promise((resolve, reject) => {
    axios.post(url, data).then((result) => {
      resolve(result.data)
    }).catch((err) => {
      reject(err)
    })
  })
}
Vue.use(VueCookie)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
