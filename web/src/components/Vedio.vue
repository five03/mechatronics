<template>
  <div class="continer">
    <h1 class="title">Real-time Video</h1>
    <van-cell style="background-color: rgba(0, 0, 0, 0); border-radius: 10px;">
      <iframe src="http://127.0.0.1:9999/v1/proxy"
              width="100%"
              height="100%"
              frameborder="1/0"
              scrolling="auto">
      </iframe>
    </van-cell>

    <van-button type="primary"
                @click="onClick"
                round
                color="#48C7D7"
                style="margin-top: 20px"
                size="large">
      <van-icon name="phone-o"
                style="margin-right: 5px" />Real-time conversation
    </van-button>
  </div>
</template>

<script>
export default {
  methods: {
    onClick () {
      this.axios.patch(`/v1/devices/101/settings?voice_switch=true`).catch((e) => {
        console.log(e)
      })
      this.axios.get('/v1/devices/101/settings').then((response) => {
        console.log(response.data.data)
      }).catch((e) => {
        console.log(e)
      })
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
  margin-top: 180px;
}
</style>