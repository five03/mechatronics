<!-- @format -->

<template>
  <div class="index">
    <div id="container"></div>
  </div>
</template>

<script>
import AMapLoader from '@amap/amap-jsapi-loader'
export default {
  name: 'HelloWorld',
  props: {
    msg: String,
  },
  data () {
    return {
      markerList: [],
      path: []
    }
  },
  mounted () {
    //DOM初始化完成进行地图初始化
    this.initMap()
  },
  methods: {
    initMap () {
      AMapLoader.load({
        key: '80c0c6dad135569de66ac5f241004b23', // 申请好的Web端开发者Key，首次调用 load 时必填
      }).then((AMap) => {
        var map = new AMap.Map('container', {
          //设置地图容器id
          viewMode: '3D', // 地图模式
          terrain: true, // 开启地形图
          version: "2.0",
          zoom: 18, //初始化地图级别
          pitch: 25,
          lang: 'en',
          center: [114.142718, 22.280121], //初始化地图中心点位置
          "AMapUI": {             // 是否加载 AMapUI，缺省不加载
            "version": '1.1',   // AMapUI 版本
            "plugins": ['overlay/SimpleMarker'],       // 需要加载的 AMapUI ui插件
          },
        })
        this.map = map;

        this.axios.get('/v1/devices/101/coordinates').then((response) => {
          var markerList = []
          var path = []
          response.data.data.forEach((item, index) => {
            if (index % 4 == 0) {
              markerList.push(new AMap.Marker({
                position: [item.longitude, item.latitude],
              }))
            }
            path.push([item.longitude, item.latitude]);
          });
          var polyline = new AMap.Polyline({
            path: path,
            strokeWeight: 4,
            strokeColor: 'red', // 线条颜色
            lineJoin: 'round' // 折线拐点连接处样式
          });
          //添加到地图
          this.map.add(polyline);
          this.map.add(markerList);
        }).catch((response) => {
          console.log(response);
        })
      }).catch((e) => {
        console.log(e)
      })
    },
  },
}
</script>

<style scoped>
.index {
  height: 100%;
}

#container {
  margin: 0px;
  width: 100%;
  height: 100%;
}
</style>
