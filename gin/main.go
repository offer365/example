package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// 全局设置
func Init() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()
	// 使用默认中间件创建一个gin路由器 logger and recovery 中间件
	gin.SetMode(gin.ReleaseMode) // 生产模式 or gin.DebugMode
	r = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

// 规则匹配
func Rule() {
	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	r.GET("/test1/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println(name) // eg: /user/tom   name = tom
		c.String(http.StatusOK, name)
	})
	// eg: /user/tom    tom
	// eg: /use/tom/    tom

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	r.GET("/test2/:name/*action", func(c *gin.Context) {
		name := c.Param("name") // eg: /test2/tom/send   name = tom; action = /send
		action := c.Param("action")
		fmt.Println(name, action)             // *action 会把 / 也匹配到
		c.String(http.StatusOK, name, action) // 参数里不要带 "/"
	})
}

// 获取get参数
func GetArgs() {
	// 匹配的url格式:  /test3?name=tom&age=11
	r.GET("/test3", func(c *gin.Context) {
		name := c.DefaultQuery("name", "tom")
		age := c.Query("age")           // 是 c.Request.URL.Query().Get("lastname") 的简写
		fmt.Println(c.GetQuery("name")) // tom true
		fmt.Println(name, age)          //  tom 11
		// eg: /test3?name=周起&name=张三
		fmt.Println(c.QueryArray("name"))    // [周起 张三]
		fmt.Println(c.GetQueryArray("name")) // [周起 张三] true
		// eg: /test3?name[a]=周起&name[b]=张三
		fmt.Println(c.QueryMap("name"))    // map[a:周起 b:张三]
		fmt.Println(c.GetQueryMap("name")) // map[a:周起 b:张三] true

		c.String(http.StatusOK, "Hello %s %s", name, age)
	})
}

// 获取post参数
func PostArgs() {
	r.POST("/test4", func(c *gin.Context) {
		// curl -X POST -F msg=aa -F nick=bb 127.0.0.1/test4
		message := c.PostForm("msg")
		nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
		// curl -X POST -F msg=aa -F msg=bb 127.0.0.1/test4
		fmt.Println(c.PostFormArray("msg"))    // [aa bb]
		fmt.Println(c.GetPostFormArray("msg")) // [aa bb] true
		// curl -X POST -F msg[a]=aa -F msg[b]=bb 127.0.0.1/test4
		fmt.Println(c.PostFormMap("msg"))    // map[a:aa b:bb]
		fmt.Println(c.GetPostFormMap("msg")) // map[a:aa b:bb] true
		c.JSON(200, gin.H{
			"message": message,
			"nick":    nick,
		})
	})
}

// 单文件上传
func SingleFileUpload() {
	// 给表单限制上传大小 (默认 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/single_upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		// curl -X POST http://localhost/single_upload -F "file=@C:\Users\Administrator\Desktop\aaa.jpg" -H "Content-Type: multipart/form-data"
		// 上传文件到指定的路径
		err := c.SaveUploadedFile(file, "./"+file.Filename)
		// c.Stream()
		fmt.Println(err)

		fmt.Println(getCurrentPath())
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}

// 多文件上传
func ManyFileUpload() {
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/many_upload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		// curl -X POST http://localhost/many_upload -F "upload[]=@C:\Users\Administrator\Desktop\aaa.jpg" -F "upload[]=@C:\Users\Administrator\Desktop\aaa.jpg" -H "Content-Type: multipart/form-data"
		for _, file := range files {
			fmt.Println(file.Filename)

			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
}

// 路由分组
func routerGroup() {
	// Simple group: v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login", nil)
		v1.POST("/submit", nil)
		v1.POST("/read", nil)
	}

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", nil)
		v2.POST("/submit", nil)
		v2.POST("/read", nil)
	}
}

// 中间件
func Middleware() {
	// 全局中间件
	// 使用 Logger 中间件
	r.Use(gin.Logger())

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	// 路由添加中间件，可以添加任意多个
	r.GET("/benchmark",
		func(c *gin.Context) {
			// 中间件
			fmt.Println(c.Request.RemoteAddr)
			c.Next()
		},
		func(c *gin.Context) {
			c.String(200, "ok")
		})

	// 路由组中添加中间件
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.

	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", nil)
		authorized.POST("/submit", nil)
		authorized.POST("/read", nil)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", nil)
	}
}
func main() {
	// 全局设置
	Init()
	// 规则匹配
	Rule()
	// 获取get参数
	GetArgs()
	// 获取post参数
	PostArgs()
	// 单文件上传
	SingleFileUpload()
	// 多文件上传
	ManyFileUpload()

	r.Run(":80") // listen and serve on 0.0.0.0:8080
}

// 获取当前执行程序的路径。
func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("test")
		if cookie == "" || err != nil {
			return
		}
		c.Next()
	}
}
