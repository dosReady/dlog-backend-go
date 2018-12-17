<template>
    <div>
        <input type="text" v-model="logindata.email" placeholder="이메일" />
        <input type="password" v-model="logindata.pwd" placeholder="비밀번호" />
        <button @click="btnLogin">로그인</button>
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
      const data = await this.$http.post('/api/dlog/login', {'login': this.logindata})
      alert(data)
    }
  }
}
</script>
