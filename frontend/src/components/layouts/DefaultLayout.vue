<template>
    <div class="default-layout">
        <slot></slot>
    </div>
</template>

<script>
export default {
  name: 'DefaultLayout'
}
</script>

<style scoped>
.default-layout {
    width: 100%;
    height: 100%;
    padding: 0!important;
}
</style>
