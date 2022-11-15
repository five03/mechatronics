<template>
  <div class="continer">
    <h1 class="title">Settings</h1>
    <div class="sub-title">
      <van-button round
                  color="#FCDFC8"
                  type="primary">
        <van-icon name="fire-o"
                  color="#FF532F"
                  size="20" />
      </van-button>
      <div class="sub-title-text">
        <b>Color of the alarm light</b>
      </div>
    </div>

    <van-field readonly
               clickable
               name="picker"
               :value="value"
               label="Color"
               placeholder="Click and select color"
               @click="showPicker = true" />
    <van-popup v-model="showPicker"
               position="bottom">
      <van-picker show-toolbar
                  :columns="columns"
                  @confirm="onConfirm"
                  @cancel="showPicker = false" />
    </van-popup>

    <div class="sub-title">
      <van-button round
                  color="#FCDFC8"
                  type="primary">
        <van-icon name="notes-o"
                  color="#FF532F"
                  size="20" />
      </van-button>
      <div class="sub-title-text">
        <b>Content of the screen</b>
      </div>
    </div>

    <van-field v-model="message"
               style="border-radius: 10px;margin-bottom: 20px"
               round
               rows="5"
               autosize
               type="textarea"
               maxlength="50"
               placeholder="Please input the content"
               show-word-limit />

    <van-button type="primary"
                round
                color="#FA6A51"
                @click="onSubmit"
                size="large">Submit</van-button>
  </div>
</template>

<script>
export default {
  data () {
    return {
      checked: false,
      value: '',
      columns: ['red', 'green'],
      showPicker: false,
      message: ""
    }
  },
  mounted () {
    this.GetSettings()
    this.GetMessageBoard()
  },
  methods: {
    onConfirm (value) {
      this.value = value;
      this.showPicker = false;
    },
    onSubmit () {
      this.axios.post(`/v1/devices/101/message-board?msg=${this.message}`).catch((e) => {
        console.log(e)
      })
      this.GetMessageBoard(`/v1/devices/101/settings?`)

      this.axios.patch(`/v1/devices/101/settings?led_switch=${this.value}`).catch((e) => {
        console.log(e)
      })
      this.GetSettings()
    },
    GetSettings () {
      this.axios.get('/v1/devices/101/settings').then((response) => {
        this.value = response.data.data.led_switch
      }).catch((e) => {
        console.log(e)
      })
    },
    GetMessageBoard () {
      this.axios.get('/v1/devices/101/message-board').then((response) => {
        this.message = response.data.data.message
      }).catch((e) => {
        console.log(e)
      })
    }
  },
}
</script>

<style scoped>
.continer {
  margin: 0px 20px;
}
.title {
  padding: 10px 0px 0px 0px;
}
.sub-title {
  padding-top: 20px;
  padding-bottom: 20px;
}
.sub-title-text {
  display: inline-block;
  height: 44px;
  vertical-align: bottom;
  line-height: 44px;
  font-size: 1.3em;
  margin-left: 10px;
}
</style>