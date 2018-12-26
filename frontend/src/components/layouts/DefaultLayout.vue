<template>
  <div id="default-layout">
    <div class="left-container">
    </div>
    <div class="content-container">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DefaultLayout'
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
  }
  .content-container {
    margin-left: 35%;
    max-width: 54em;
    padding: 8em 4em 4em 4em;
    width: calc(100%-35%);
  }
}
</style>
