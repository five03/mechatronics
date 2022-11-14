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
      //此处不声明 map 对象，可以直接使用 this.map赋值或者采用非响应式的普通对象来存储。
      //map:null,
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
        this.map = new AMap.Map('container', {
          //设置地图容器id
          viewMode: '3D', //是否为3D地图模式
          zoom: 18, //初始化地图级别
          pitch: 45,
          center: [114.142718, 22.280121], //初始化地图中心点位置
          layers: [//使用多个图层
            new AMap.TileLayer.Satellite(),
            // new AMap.TileLayer.RoadNet()
          ],
        })

        // this.map.setLang('zh-en');

        var marker1 = new AMap.Marker({
          position: new AMap.LngLat(114.142718, 22.280121),//位置
          title: '香港大学1'
        })
        var marker2 = new AMap.Marker({
          position: new AMap.LngLat(114.142858, 22.280321),//位置
          title: '香港大学2'
        })
        var marker3 = new AMap.Marker({
          position: new AMap.LngLat(114.142988, 22.280521),//位置
          title: '香港大学3'
        })
        var markerList = [marker1, marker2, marker3];
        this.map.add(markerList);//添加到地图

        var path = [
          new AMap.LngLat(114.142718, 22.280121),
          new AMap.LngLat(114.142858, 22.280321),
          new AMap.LngLat(114.142988, 22.280521),
        ];
        var polyline = new AMap.Polyline({
          path: path,
          strokeWeight: 8,
          strokeColor: 'red', // 线条颜色
          lineJoin: 'round' // 折线拐点连接处样式
        });

        // 将折线添加至地图实例
        this.map.add(polyline);
        // this.map.addControl(new AMap.Scale());
        // this.map.addControl(new AMap.MapType());
        // this.map.addControl(new AMap.HawkEye({ isOpen: true }));
      })
        .catch((e) => {
          console.log(e)
        })
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
