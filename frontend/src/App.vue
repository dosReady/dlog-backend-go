<template>
  <div id="app">
    <load-layout></load-layout>
  </div>
</template>

<script>
import LoadLayout from '@/components/layouts/LoadLayout'
export default {
  name: 'App',
  components: {
    LoadLayout
  }
}
</script>

<style lang="scss">
@import "~@/assets/scss/main.scss";
</style>
