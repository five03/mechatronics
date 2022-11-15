<template>
  <div id="app">

    <router-view></router-view>

    <van-tabbar v-model="active"
                route>
      <van-tabbar-item eplace
                       to="/settings"
                       icon="setting-o">Settings</van-tabbar-item>
      <van-tabbar-item eplace
                       to="/index"
                       icon="shield-o">Monitoring</van-tabbar-item>
      <van-tabbar-item eplace
                       to="/vedio"
                       icon="video-o">Vedio</van-tabbar-item>
      <van-tabbar-item eplace
                       to="/storage"
                       icon="send-gift-o">Storage</van-tabbar-item>
    </van-tabbar>

    <div style="position: absolute;bottom: 90px; right:50px;">
      <van-button style="border-radius: 50px; height: 80px; box-shadow: 1px 1px 5px #888888;"
                  color="#FF1C24"
                  type="primary"
                  @click="onClick">
        <b>ALARM</b>
      </van-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      active: 1,
    }
  },
  methods: {
    onClick () {
      this.axios.patch(`/v1/devices/101/settings?video_switch=true`).catch((e) => {
        console.log(e)
      })
      this.axios.get('/v1/devices/101/settings').then((response) => {
        console.log(response.data.data)
      }).catch((e) => {
        console.log(e)
      })
    }
  },
}
</script>

<style>
html,
body {
  height: 100%;
  background-color: rgb(251, 249, 234);
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100%;
}
.van-tabbar {
  height: 70px;
}
.amap-logo {
  display: none !important;
  visibility: hidden !important;
}

.amap-copyright {
  display: none !important;
  visibility: hidden !important;
}
</style>
