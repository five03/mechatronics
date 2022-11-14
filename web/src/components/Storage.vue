<template>
  <div class="continer">
    <div class="title">
      <h1>Cloud storage</h1>
      <span>Device Id: 101</span>
      <van-divider>
        <van-icon name="tv-o" />
      </van-divider>
    </div>

    <div>
      <van-cell-group inset
                      @click="showPopup">
        <van-cell clickable
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong" />
        <van-cell clickable
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong" />
        <van-cell clickable
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong" />
        <van-cell clickable
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong" />
        <van-cell clickable
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong" />
        <van-cell clickable
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong" />
      </van-cell-group>
    </div>

    <van-popup v-model="popup"
               closeable
               close-icon="close"
               :style="{width: '100%', 'background-color': 'rgba(0, 0, 0, 0)'}">

        <video width="100%"
               height="100%"
               controls
               autoplay>
          <source src="movie.ogg"
                  type="video/ogg">
          <source src="movie.mp4"
                  type="video/mp4">
          <source src="movie.webm"
                  type="video/webm">
          <object data="movie.mp4"
                  width="320"
                  height="240">
            <embed width="320"
                   height="240"
                   src="movie.swf">
          </object>
        </video>

    </van-popup>
  </div>
</template>

<script>
export default {
  data () {
    return {
      popup: false,
    }
  },
  methods: {
    showPopup () {
      this.popup = true;
    },
  },
}
</script>

<style scoped>
.continer {
  margin: 0px 20px;
}
.title {
  padding: 10px 0px 20px 0px;
}
h1 {
  margin-bottom: 0;
}
</style>