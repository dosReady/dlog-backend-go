<template>
    <div>
        <input type="text" v-model="logindata.email" placeholder="이메일" />
        <input type="password" v-model="logindata.pwd" placeholder="비밀번호" />
        <button @click="btnLogin">로그인</button>
    </div>
</template>

<script>
const USER_NOTFOUND = 1
const USER_NOTMATCH = 2
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
      const data = await this.$http('/api/dlog/login', this.logindata)
      const token = data.accessToken
      const status = data.status
      if (!token) {
        switch (status) {
          case USER_NOTFOUND:
            alert('사용자를 찾을수 없습니다.')
            break
          case USER_NOTMATCH:
            alert('사용자 정보가 일치하지 않습니다.')
        }
      } else {
        this.$cookie.set('token', token, Infinity)
      }
    }
  }
}
</script>
