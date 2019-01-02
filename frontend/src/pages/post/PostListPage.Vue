<template>
  <div id="list-container">
    <ul class="post-list">
      <li class="card-wrap" v-for="(item, index) in list" :key="index">
        <a class="card-image" href="#">
          <img src="@/assets/images/golang.png"/>
        </a>
        <article>
          <h2 class="card-title">{{item.post_title}}</h2>
          <p>
            {{item.post_contnet}}
          </p>
        </article>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'PostListPage',
  data () {
    return {
      list: []
    }
  },
  async created () {
    const data = await this.$http('/api/post/list')
    this.list = data.list
  }
}
</script>

<style lang="scss" scoped>
#list-container {
  .card-wrap {
    padding: 1rem;
    border-radius: 5px;
    box-shadow: 8px 14px 38px rgba(39,44,49,.06), 1px 3px 8px rgba(39,44,49,.03);
    .card-title {

    }
  }
  .post-list {
    padding-top: 30px;
    padding-bottom: 30px;
  }
}

</style>
