<template>
  <div id="app">
	  <div>{{hello}}</div>
    <hello-world :msg="hello" @rollback="rollback"></hello-world>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'app',
  components: {
    HelloWorld
  },
  data: function() {
	  return {
		  hello : '文件上传'
	  }
  },
  methods: {
	  rollback(opt) {
		  this.hello = opt
	  },
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
