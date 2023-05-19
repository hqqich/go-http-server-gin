import Vue from 'vue'

import App from './App.vue'

//========= 引入 axios
import axios from 'axios'
//========= 引入 elmentUI
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 挂载一个自定义属性$http
Vue.prototype.$http = axios
// 全局配置axios请求根路径(axios.默认配置.请求根路径)
axios.defaults.baseURL = 'http://127.0.0.1:8888'

Vue.use(ElementUI);


Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
