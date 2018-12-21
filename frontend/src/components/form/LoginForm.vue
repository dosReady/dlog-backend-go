<template>
    <div>
        <input type="text" v-model="logindata.email" placeholder="이메일" />
        <input type="password" v-model="logindata.pwd" placeholder="비밀번호" />
        <button @click="btnLogin">로그인</button>
        <button @click="test">테스트</button>
    </div>
</template>

<script>
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
      await this.$http('/api/dlog/login', this.logindata)
      console.log(this.$cookie.get('token'))
    },
    async test () {
      const data = await this.$http('/api/test/echo')
      console.log(this.$cookie.get('token'))
      console.log(data)
    }
  }
}
</script>
