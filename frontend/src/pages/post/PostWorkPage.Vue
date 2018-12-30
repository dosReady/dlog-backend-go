<template>
    <div id="editor-container">
      <div class="editor-wrap">
        <header>
          <input type="text" placeholder="제목을 입력하세요"/>
          <button class="btn"><font-awesome-icon icon="ellipsis-h" /></button>
        </header>
        <div class="codemirror-wrap">
          <codemirror v-model="data.post_content" :options="cmOption"></codemirror>
        </div>
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
        placeholder: '내용을 입력해보세요!',
        size: '100%'
      }
    }
  },
  created () {
  },
  methods: {
  },
  computed: {
    compiledMarkdown () {
      return marked(this.data.post_content, { sanitize: true })
    }
  }
}
</script>

<style lang="scss" scoped>
#editor-container {
  display: flex;
  height: 100%;
  .editor-wrap {
    flex: 1 1;
    background-color: #263238;
    header {
      display: flex;
      height: 3rem;
      background-color: #112530;
      input {
        flex: 0 0 80%;
        color: white;
        font-size: 1rem;
      }
      button {
        flex: 1 1;
        svg {
          color: white;
          font-size: 1rem;
        }
      }
    }
    .codemirror-wrap {
      height: 100%;
      background-color: #263238;
      padding: 1rem;
      overflow: auto;
      overflow-x: scroll;
    }
  }
  .markdown-body {
    flex: 1 1;
    box-sizing: border-box;
    margin: 0 auto;
    padding: 45px;
  }
}
</style>
