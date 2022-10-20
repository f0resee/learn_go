package gin

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
)

func TestGinSession(t *testing.T) {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// 理解：一个请求到达gin之后，会经过sessions.Sessions创建一个session，并保存在gin.Context中，因此
	// 在下面/hello的hanler中才能获取到一个session
	// 每个客户端有各自的session，一个客户端进行session.Set不会影响其他session中的值。但是修改session中的值似乎会导致默认的cookie失效
	i := 0
	r.GET("/hello", func(c *gin.Context) {

		session := sessions.Default(c) // 其实是获取客户端传来的cookie对应的session信息

		log.Println("session id:", session.ID(), len(session.ID()))

		val := session.Get("hello")

		log.Println(val)
		if val == nil {
			log.Println("try to set")
			session.Set("hello", i)
			i++
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
}

func TestSession2(t *testing.T) {
	r := gin.Default()

	//配置session的中间件
	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret11111"))
	// 设置 session 中间件，参数 mysession，指的是 session 的名字，也是 cookie 的名字
	// store 是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	//initMiddleware:配置路由中间件
	r.GET("/", func(c *gin.Context) {
		//设置sessions
		session := sessions.Default(c)
		session.Set("username", "成强")
		//保存sessions:给其他页面使用（必须调用）
		session.Save()

		c.String(200, "gin首页")
	})
	r.GET("/news", func(c *gin.Context) {
		//获取sessions
		session := sessions.Default(c)
		username := session.Get("username")

		c.String(200, "username=%v", username)
	})
	r.Run(":8000")
}

func Test3(t *testing.T) {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run(":8001")
}
