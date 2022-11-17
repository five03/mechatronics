<template>
  <div class="continer">
    <div class="title">
      <h1>Cloud storage</h1>
      <span>Device Id: 101</span>
      <van-divider>
        <van-icon name="tv-o"/>
      </van-divider>
    </div>

    <div>
      <van-cell-group inset
                      @click="showPopup">
        <van-cell clickable
                  @click="src='/web/video/1.mp4'"
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong"/>
        <van-cell clickable
                  @click="src='/web/video/2.mp4'"
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong"/>
        <van-cell clickable
                  @click="src='/web/video/3.mp4'"
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong"/>
        <van-cell clickable
                  @click="src='/web/video/4.mp4'"
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong"/>
        <van-cell clickable
                  @click="src='/web/video/5.mp4'"
                  title="2022-11-12 12:12:11"
                  value="Vedio"
                  label="location: Hong Kong"/>
      </van-cell-group>
    </div>

    <van-popup v-model="popup"
               closeable
               close-icon="close"
               :style="{width: '100%', 'background-color': 'rgba(0, 0, 0, 0)'}">

      <video id="my-video"
             width="100%"
             height="100%"
             controls
             autoplay>
        <source :src="src"
                type="video/mp4">
      </video>

    </van-popup>
  </div>
</template>

<script>
export default {
  data() {
    return {
      popup: false,
      src: ""
    }
  },
  methods: {
    showPopup() {
      this.popup = true;
    },
    closePopup() {
      const video = document.getElementById("my-video")
      video.pause()
      this.popup = false;
    }
  },
  watch: {
    popup(newName, oldName) {
      if (!newName) {
        this.closePopup()
      }
    }
  }
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