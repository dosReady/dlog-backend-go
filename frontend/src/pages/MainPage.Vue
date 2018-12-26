<template>
    <div>
        메인
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
  }
}
</script>
