package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Code int `json:"code"`
}

var (
	voiceSwitch = false
	videoSwitch = false
	ledSwitch   = "green"
)

type Settings struct {
	VoiceSwitch bool   `json:"voice_switch"`
	VideoSwitch bool   `json:"video_switch"`
	LEDSwitch   string `json:"led_switch"`
}

type Data struct {
	Message string `json:"message,omitempty"`
}

type Body struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Params struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Timestamp int64   `json:"timestamp"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

var Coordinates = []*Coordinate{
	{114.141812, 22.280773},
	{114.141817, 22.280751},
	{114.141833, 22.280735},
	{114.141851, 22.280716},
	{114.141883, 22.280694},
	{114.141896, 22.280664},
	{114.141892, 22.280641},
	{114.141912, 22.280606},
	{114.141928, 22.280592},
	{114.141945, 22.280557},
	{114.141954, 22.28053},
	{114.141969, 22.280519},
	{114.14201, 22.280499},
	{114.142029, 22.280488},
	{114.142054, 22.280476},
	{114.142079, 22.280474},
	{114.142092, 22.280459},
	{114.142108, 22.280458},
	{114.142142, 22.280462},
	{114.142155, 22.280458},
	{114.142175, 22.280456},
	{114.142186, 22.280458},
	{114.142205, 22.280447},
	{114.142232, 22.280444},
	{114.142246, 22.280436},
	{114.142266, 22.280433},
	{114.142287, 22.280426},
	{114.142312, 22.280408},
	{114.142313, 22.280386},
	{114.14235, 22.28037},
	{114.142348, 22.280365},
	{114.142384, 22.28035},
	{114.142423, 22.280344},
	{114.142445, 22.280337},
	{114.142484, 22.28033},
	{114.142505, 22.280333},
	{114.142555, 22.280345},
	{114.142579, 22.280354},
	{114.142627, 22.280368},
	{114.14273, 22.280356},
	{114.142745, 22.280384},
	{114.142743, 22.280411},
	{114.142746, 22.280425},
	{114.142761, 22.280439},
	{114.142746, 22.280455},
	{114.142762, 22.280469},
	{114.142773, 22.280479},
	{114.142777, 22.28049},
	{114.142781, 22.280496},
	{114.142773, 22.280511},
	{114.14279, 22.280521},
	{114.14282, 22.280517},
	{114.142843, 22.280522},
	{114.142864, 22.280523},
	{114.142887, 22.280531},
	{114.14291, 22.280528},
	{114.142932, 22.280525},
	{114.142961, 22.28053},
	{114.142963, 22.280528},
	{114.142974, 22.280526},
	{114.143021, 22.280528},
	{114.143034, 22.28054},
	{114.143029, 22.280556},
	{114.143034, 22.280576},
	{114.143037, 22.280596},
	{114.14304, 22.280606},
	{114.143037, 22.280634},
	{114.14304, 22.280646},
	{114.14304, 22.280663},
	{114.143038, 22.280678},
	{114.143042, 22.280693},
	{114.143048, 22.28072},
	{114.14305, 22.280735},
	{114.143049, 22.280762},
	{114.143044, 22.280798},
	{114.14305, 22.280823},
	{114.143065, 22.280842},
	{114.143103, 22.280863},
	{114.143138, 22.280874},
	{114.143166, 22.280888},
	{114.143195, 22.280901},
	{114.143207, 22.280914},
	{114.143227, 22.280919},
	{114.143245, 22.280926},
	{114.143272, 22.280941},
	{114.14325, 22.280953},
	{114.143235, 22.280951},
	{114.143185, 22.280968},
	{114.143178, 22.280973},
	{114.143076, 22.280984},
	{114.143062, 22.280983},
	{114.143033, 22.280988},
	{114.143002, 22.280992},
	{114.142963, 22.280983},
	{114.142938, 22.280991},
	{114.142922, 22.280992},
	{114.142879, 22.280988},
	{114.142844, 22.280997},
	{114.142844, 22.281007},
	{114.142837, 22.281041},
}

var BoardMessage = ""

//go:embed target/*
var file embed.FS

func main() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("=== mock服务器启动")
	fmt.Println("=== 请输入树莓派视频访问地址(例: http://192.168.1.3:8888)")
	var target string
	fmt.Scanln(&target)
	if _, err := url.Parse(target); err != nil {
		fmt.Printf("=== 地址格式错误, err: %v\n", err)
		return
	}
	fmt.Println("=== 尝试获取本机IP")
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println("=== 请检查网络连接")
	}
	fmt.Println("=== 本机的IP地址: ", ip)
	fmt.Printf("=== 服务器访问域名: http://%s:%s\n", ip, "9999")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 路由
	{
		web, err := fs.Sub(file, "target")
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		r.StaticFS("/web", http.FS(web))

		r.GET("/index", func(c *gin.Context) {
			// c.JSON：返回JSON格式的数据
			c.Redirect(http.StatusPermanentRedirect, "web/index.html")
		})
		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusPermanentRedirect, "/index")
		})
	}

	r.POST("/v1/devices/:id/coordinates", func(c *gin.Context) {
		id := c.Param("id")
		var p Params
		if err := c.ShouldBind(&p); err != nil {
			c.JSON(http.StatusOK, &Body{Meta: Meta{Code: http.StatusBadRequest}})
			fmt.Printf("%v", err)
			return
		}
		fmt.Printf("=== 收到设备[%s]上报信息: %+v\n", id, p)
		c.JSON(http.StatusOK, &Body{Meta: Meta{Code: http.StatusOK}})
	})
	r.GET("/v1/devices/:id/coordinates", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Body{Meta: Meta{Code: http.StatusOK}, Data: Coordinates})
	})
	r.GET("/v1/message-board", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Body{
			Meta: Meta{Code: http.StatusOK},
			Data: Data{Message: BoardMessage},
		})
	})
	r.POST("/v1/devices/:id/message-board", func(c *gin.Context) {
		BoardMessage = c.Query("msg")
		c.JSON(http.StatusOK, &Body{
			Meta: Meta{Code: http.StatusOK},
		})
	})
	r.GET("/v1/devices/:id/message-board", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Body{
			Meta: Meta{Code: http.StatusOK},
			Data: Data{Message: BoardMessage},
		})
	})
	r.GET("/v1/devices/:id/settings", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Body{
			Meta: Meta{Code: http.StatusOK},
			Data: Settings{
				VoiceSwitch: voiceSwitch,
				VideoSwitch: videoSwitch,
				LEDSwitch:   ledSwitch,
			},
		})
	})
	r.PATCH("/v1/devices/:id/settings", func(c *gin.Context) {
		switch c.Query("voice_switch") {
		case "true":
			voiceSwitch = true
		case "false":
			voiceSwitch = false
		}

		switch c.Query("video_switch") {
		case "true":
			videoSwitch = true
		case "false":
			videoSwitch = false
		}

		switch c.Query("led_switch") {
		case "red":
			ledSwitch = "red"
		case "green":
			ledSwitch = "green"
		}
		c.JSON(http.StatusOK, &Body{Meta: Meta{Code: http.StatusOK}})
	})
	r.GET("/v1/proxy", func(c *gin.Context) {
		parse, err := url.Parse(target)
		if err != nil {
			fmt.Printf("代理地址解析 %s 出错, error: %+v\n", target, err)
		}
		proxy := httputil.NewSingleHostReverseProxy(parse)
		proxy.Director = func(request *http.Request) {
			request.Host = parse.Host
			request.URL.Scheme = parse.Scheme
			request.URL.Host = parse.Host
			request.URL.Path = ""
		}
		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
			fmt.Printf("代理访问 %s 出错, error: %+v\n", target, err)
			writer.Write([]byte(err.Error()))
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})
	r.Run("0.0.0.0:9999") // listen and serve on 0.0.0.0:8080
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return strings.TrimSpace(ip), err
}
