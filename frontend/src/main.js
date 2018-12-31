// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import VueCookie from 'vue-cookie'
import 'es6-promise/auto'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faUser, faKey, faArrowLeft } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import VueCodemirror from 'vue-codemirror'
import VModal from 'vue-js-modal'
import 'codemirror/addon/scroll/simplescrollbars.css'
import 'codemirror/addon/scroll/simplescrollbars.js'
import 'codemirror/addon/display/placeholder.js'
import 'codemirror/theme/material.css'
import '@/assests/css/github-markdown.css'
import '@/assests/css/codemirror.css'

library.add(faUser, faKey, faArrowLeft)

Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.config.productionTip = false
Vue.prototype.$http = (url, data) => {
  return new Promise((resolve, reject) => {
    axios.post(url, data).then((result) => {
      resolve(result.data)
    }).catch((err) => {
      if (err.response.status === 401) {
        router.push('/')
      } else reject(err)
    })
  })
}
Vue.prototype.$eventBus = new Vue()
Vue.use(VueCookie)
Vue.use(VueCodemirror)
Vue.use(VModal, {
  dialog: true,
  dynamic: true
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
