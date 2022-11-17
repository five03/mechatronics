import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import router from './router'
import { Button } from 'vant';
import { Icon } from 'vant';
import { Form } from 'vant';
import { Field } from 'vant';
import { Switch } from 'vant';
import { Cell, CellGroup } from 'vant';
import { Picker } from 'vant';
import { Popup } from 'vant';
import { Tabbar, TabbarItem } from 'vant';
import { Divider } from 'vant';
import { NavBar } from 'vant';
import axios from 'axios'
import VueAxios from 'vue-axios'

// 配置请求的根路径
axios.defaults.baseURL = '/'

Vue.use(VueAxios,axios);

Vue.use(VueRouter)
Vue.use(Button);
Vue.use(Icon);
Vue.use(Form);
Vue.use(Field);
Vue.use(Switch);
Vue.use(Picker);
Vue.use(Cell);
Vue.use(CellGroup);
Vue.use(NavBar);
Vue.use(Divider);
Vue.use(Tabbar);
Vue.use(TabbarItem);
Vue.use(Popup);

new Vue({
  el: "#app",
  render: h => h(App),
  router,
}).$mount('#app')
