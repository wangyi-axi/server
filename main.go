//package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)

// 初始化程序
//func init() {
//	//读取日志文件
//	fi, err := os.Stat("./logs/stdout.log")
//	if err == nil {
//		//查看日志文件大小。。大于90000删除
//		if fi.Size() > 900000 {
//			_ = os.Remove("./logs/stdout.log")
//		}
//	}
//	//用于创建日志文件
//	LogFile, err := os.OpenFile("./logs/stdout.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err != nil {
//		fmt.Println("日志文件打开错误")
//	}
//	//设置日志输出
//	log.SetOutput(LogFile)
//	//设置日走输出格式
//	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
//	//用于数据库的连接
//	tools.DB = tools.DatabaseConn()
//}
//
//func main() {
//	//qsc.Run()
//	//SocketIo()
//	http.HandleFunc("/", testTcp)
//
//	http.ListenAndServe(":8081", nil)
//	//qsc.Connserver()
//	//tools.Systeminit()
//	//xjg, err := methods.GenerateToken("xjg", "iloveyou")
//	//fmt.Println(xjg, err)
//	//xjg1, err1 := methods.ParseToken(xjg)
//	//fmt.Println(xjg1, err1)
//	//Test()
//}
//
//func testTcp(message string) (res string) {
//	res = "eeee" + message
//	return res
//}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/methods"
)

func setUp() *gin.Engine {
	r := gin.Default()
	return r
}

func demo1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "this is a gin demo",
	})
}
func demo2(c *gin.Context) {
	e := methods.Devicelocation("1")
	c.JSON(http.StatusOK, gin.H{
		"msg": e,
	})
}

func main() {
	r := setUp()
	r.GET("/demo", demo1)
	r.GET("/demo2", demo2)
	r.Run(":8090")
}
