<template>
    <div class="form-gorup">
      <div class="form-field">
        <font-awesome-icon icon="user" />
        <input type="text" v-model="logindata.email" placeholder="이메일" />
      </div>
      <div class="form-field">
        <font-awesome-icon icon="key" />
        <input type="password" v-model="logindata.pwd" placeholder="비밀번호" />
      </div>
      <button @click="btnLogin">로그인</button>
      <button @click="test">테스트</button>
    </div>
</template>

<script>
import { EventBus } from '@/assets/js/eventbus'
export default {
  name: 'LoginForm',
  data () {
    return {
      logindata: {
        email: '',
        pwd: ''
      }
    }
  },
  methods: {
    async btnLogin () {
      try {
        await this.$http('/api/user/login', this.logindata)
        EventBus.$emit('setUser')
        this.$router.push('/main')
      } catch (error) {
        alert(error)
        console.log(error)
      }
    },
    async test () {
      await this.$http('/api/test/echo')
    }
  }
}
</script>

<style scoped>
</style>
