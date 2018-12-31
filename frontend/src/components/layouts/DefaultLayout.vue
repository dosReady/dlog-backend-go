<template>
  <div id="default-layout">
    <header class="left-container">
        <div class="login-info" v-if="!isEmptyUser">
            <label>{{user.Email}}</label>
            <button class="btn" @click="logout">로그아웃</button>
            <button class="btn" @click="test">테스트</button>
            <div class="menu-container">
              <router-link to="/post/list">포스트 목록</router-link>
              <router-link to="/post/register">포스트 작성</router-link>
              <router-link to="/">DLOG 작성</router-link>
              <router-link to="/">코드 저장소</router-link>
            </div>
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
    const cookieUser = this.$cookie.get('user')
    if (cookieUser) {
      this.user = JSON.parse(this.$cookie.get('user'))
    }
    this.$eventBus.$on('setUser', () => {
      this.user = JSON.parse(this.$cookie.get('user'))
    })
  },
  computed: {
    isEmptyUser () {
      return Object.keys(this.user).length === 0
    }
  },
  methods: {
    async logout () {
      await this.$http('/api/user/logout')
      this.user = {}
      this.$router.push('/login')
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
    .menu-container {
      margin-top: 1rem;
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
