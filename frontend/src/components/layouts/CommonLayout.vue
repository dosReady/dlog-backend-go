<template>
  <div id="common-layout">
      <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'CommonLayout'
}
</script>

<style lang="scss" scoped>
#common-layout {
  width: inherit;
  height: inherit;
}
</style>
