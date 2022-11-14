<template>
  <div class="continer">
    <h1 class="title">Real-time Video</h1>
    <van-cell style="background-color: rgba(0, 0, 0, 0); border-radius: 10px;">
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
    </van-cell>

    <van-button type="primary"
                round
                color="#48C7D7"
                style="margin-top: 20px"
                size="large"><van-icon name="phone-o" style="margin-right: 5px" />Real-time conversation</van-button>
  </div>
</template>

<script>
export default {

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