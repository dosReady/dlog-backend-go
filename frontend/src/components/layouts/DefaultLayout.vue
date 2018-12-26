<template>
  <div id="default-layout">
    <header class="left-container">
        <div class="login-info" v-if="user != null">
            <label>{{user.Email}}</label>
            <button class="btn" @click="logout">로그아웃</button>
            <button class="btn" @click="test">테스트</button>
        </div>
    </header>
    <div class="content-container">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DefaultLayout',
  data () {
    return {
      user: {}
    }
  },
  created () {
    let user = JSON.parse(this.$cookie.get('user'))
    this.user = user
    console.log(user)
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

<style lang="scss" scoped>
#default-layout {
  width: inherit;
  height: inherit;
  .left-container {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    width: 35%;
    height: 100%;
    background-image: url("~@/assests/images/bg.jpg");
    position: fixed;
    .login-info {
        margin: 30rem auto;
        button {
            margin-top: 1rem;
        }
    }
  }
  .content-container {
    margin-left: 35%;
    max-width: 54em;
    padding: 8em 4em 4em 4em;
    width: calc(100%-35%);
  }
}
</style>
