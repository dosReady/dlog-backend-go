<template>
  <transition name="component-fade" mode="out-in">
    <component :is='switchLayout'>
      <router-view></router-view>
    </component>
  </transition>
</template>

<script>
const defaultLayout = 'default'
export default {
  name: 'LoadLayout',
  components: {
    'DefaultLayout': () => import('@/components/layouts/DefaultLayout'),
    'CommonLayout': () => import('@/components/layouts/CommonLayout')
  },
  created () {
    console.log(this.$route.meta.layout)
  },
  computed: {
    switchLayout () {
      console.log(this.$route.meta.layout)
      return (this.$route.meta.layout || defaultLayout) + '-layout'
    }
  }
}
</script>

<style scoped>
.component-fade-enter-active, .component-fade-leave-active {
  transition: opacity .3s ease;
}
.component-fade-enter, .component-fade-leave-to {
  opacity: 0;
}
</style>
