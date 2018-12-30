<template>
    <div id="editor-container">
      <div class="editor-wrap">
        <header>
          <input type="text" placeholder="제목을 입력하세요"/>
          <button class="btn" @click="showSubmit">완료</button>
        </header>
        <div :class="submitClass">
          <div class="title">글 작성하기</div>
          <div class="btn-group field">
            <input type="text" placeholder="카테고리를 입력하세요"/>
            <button class="btn">등록</button>
          </div>
          <button class="btn field">업로드</button>
          <div class="btn-group">
            <button class="btn temp">임시저장</button>
            <button class="btn save">저장</button>
          </div>
        </div>
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
      },
      submitClass: {
        'submit-area': true,
        'hide': true
      }
    }
  },
  created () {
  },
  methods: {
    showSubmit () {
      this.submitClass.hide = !this.submitClass.hide
    }
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
    position: relative;
    .submit-area {
      &.hide {
        display: none!important;
      }
      display: flex;
      flex-direction: column;
      background-color: rgb(84, 84, 108);
      position: absolute;
      right: 2rem;
      top: 4rem;
      z-index: 15;
      padding: 1rem 2rem 1rem 2rem;
      .title {
        color: white;
        margin-bottom: 1rem;
      }
      input {
        flex: 0 0 70%;
        padding: 1rem;
        background-color: white;
        &::placeholder {
          color: gray;
        }
      }
    }
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
      height: calc(100% - 5rem);
      background-color: #263238;
      padding: 1rem 1rem 1rem 1rem;
      overflow: auto;
      overflow-y: scroll;
    }
  }
  .markdown-body {
    flex: 1 1;
    box-sizing: border-box;
    margin: 0 auto;
    padding: 45px;
    height: 100%;
  }
}
</style>
