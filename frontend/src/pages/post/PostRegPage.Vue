<template>
    <div class="editor-container">
      <div class="input-content">
        <codemirror v-model="data.post_content" :options="cmOption"></codemirror>
      </div>
      <div class="markdown-body" v-html="compiledMarkdown"></div>
    </div>
</template>

<script>
import marked from 'marked'
import 'codemirror/mode/markdown/markdown.js'
export default {
  name: 'PostRegPage',
  data () {
    return {
      data: {
        post_content: ''
      },
      cmOption: {
        mode: 'markdown',
        theme: 'material',
        lineNumbers: false,
        lineWrapping: true,
        // scrollbarStyle: 'overlay',
        placeholder: '내용을 입력해보세요!'
      }
    }
  },
  created () {
  },
  methods: {
  },
  computed: {
    compiledMarkdown () {
      let content = ''
      if (this.data.post_content) {
        content = this.data.post_content
      }
      return marked(content, { sanitize: true })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '~@/assests/css/github-markdown';
.editor-container{
    display: flex;
    flex-direction: column;
}
.input-content {
  background-color: #263238!important;
  color: white!important;
}
</style>
