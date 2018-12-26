<template>
    <div>
        {{user.Email}}
        <button @click="logout">로그아웃</button>
        <button @click="test">테스트</button>
    </div>
</template>

<script>
export default {
  name: 'MainPage',
  data () {
    return {
      user: {}
    }
  },
  created () {
    let user = JSON.parse(this.$cookie.get('user'))
    this.user = user
  },
  methods: {
    async logout () {
      await this.$http('/api/user/logout')
      this.$router.push('/')
    },
    async test () {
      await this.$http('/api/test/echo')
    }
  }
}
</script>
