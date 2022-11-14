import VueRouter from 'vue-router'

import HelloWorld from '../components/HelloWorld.vue'
import Settings from '../components/Settings.vue'
import Vedio from '../components/Vedio.vue'
import Storage from '../components/Storage.vue'

export default new VueRouter({
    routes: [
        {
            path: "/index",
            component: HelloWorld
        },
        {
            path: "/settings",
            component: Settings
        },
        {
            path: "/vedio",
            component: Vedio
        },
        {
            path: "/storage",
            component: Storage
        }
    ]
})