package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
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
	{114.141913, 22.280998},
	{114.141915, 22.280959},
	{114.141937, 22.280855},
	{114.141946, 22.280779},
	{114.141939, 22.280679},
	{114.141952, 22.280577},
	{114.14195, 22.280528},
	{114.142098, 22.28046},
	{114.142266, 22.280466},
	{114.142478, 22.280342},
	{114.142551, 22.280328},
	{114.142983, 22.280308},
	{114.143129, 22.280529},
	{114.143111, 22.280766},
	{114.143126, 22.280945},
	{114.143031, 22.281016},
	{114.142968, 22.281201},
}

//go:embed target/*
var file embed.FS

func main() {
	//go func() {
	//	r := gin.Default()
	//	r.GET("/", func(c *gin.Context) {
	//		fmt.Println("=== PING")
	//		c.File("/Users/huwenhao/Downloads/111.mp4")
	//		//c.JSON(http.StatusOK, &Body{Meta: Meta{Code: http.StatusOK}})
	//	})
	//	r.Run("0.0.0.0:8888") // listen and serve on 0.0.0.0:8080
	//}()
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("=== mock服务器启动")
	fmt.Println("=== 请输入树莓派视频访问地址(例: http://192.168.1.3:8888):")
	var target string
	fmt.Scanln(&target)
	if _, err := url.Parse(target); err != nil {
		fmt.Printf("=== 地址格式错误, err: %v\n", err)
		return
	}
	fmt.Println("=== 创建临时文件夹")
	if _, err := os.Stat("./tmp"); os.IsNotExist(err) {
		if err = os.Mkdir("./tmp", os.ModePerm); err != nil {
			fmt.Printf("%v", err)
		}
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
			Data: Data{Message: "Test Message"},
		})
	})
	r.GET("/v1/devices/:id/message-board", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Body{
			Meta: Meta{Code: http.StatusOK},
			Data: Data{Message: "Test Message"},
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
	r.POST("/v1/devices/:id/voice", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, &Body{Meta: Meta{Code: http.StatusInternalServerError}})
			fmt.Printf("%v", err)
			return
		}
		err = c.SaveUploadedFile(file, "./tmp/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusOK, &Body{Meta: Meta{Code: http.StatusInternalServerError}})
			fmt.Printf("%v", err)
			return
		}
		c.JSON(200, gin.H{
			"msg": file.Filename,
		})
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
