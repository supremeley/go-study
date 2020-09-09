package main

import (
	"studygo/array"
	"studygo/cancel"
	"studygo/channel"
	"studygo/client"
	"studygo/csp"
	"studygo/def"
	emptyinterface "studygo/empty_interface"
	"studygo/error"
	"studygo/extension"
	firstres "studygo/first_res"
	"studygo/fun"
	"studygo/groutine"
	"studygo/iface"
	"studygo/loop"
	"studygo/mapgo"
	"studygo/once"
	"studygo/panic"
	"studygo/polymorphism"
	"studygo/pool"
	"studygo/str"
	"studygo/structs"
	syncpool "studygo/sync_pool"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/configor"
)

func main() {
	r := gin.Default()
	// r.GET("/user", func(c *gin.Context) {
	// 	// name := c.Param("name")
	// 	a := c.Query("a")
	// 	b := c.DefaultQuery("b", "123")
	// 	// cc := c.PostForm("c")
	// 	// dd := c.DefaultPostForm("d", "dd")
	// 	// c.String(http.StatusOK, "hello %s %s", a, b)
	// 	// c.String(http.StatusOK, "hello %s %s", c, d)
	// 	c.JSON(200, gin.H{
	// 		"a": a,
	// 		"b": b,
	// 		// "c": cc,
	// 		// "d": dd,
	// 	})

	// })
	v1 := r.Group("/v1")
	{
		v1.POST("/user", func(c *gin.Context) {
			a := c.Query("a")
			b := c.DefaultQuery("b", "123")
			cc := c.PostForm("c")
			dd := c.DefaultPostForm("d", "dd")
			// c.String(http.StatusOK, "hello %s %s", a, b)
			// c.String(http.StatusOK, "hello %s %s", c, d)
			c.JSON(200, gin.H{
				"a": a,
				"b": b,
				"c": cc,
				"d": dd,
			})

		})
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/user", func(c *gin.Context) {
			a := c.Query("a")
			b := c.DefaultQuery("b", "123")
			cc := c.PostForm("c")
			dd := c.DefaultPostForm("d", "dd")
			// c.String(http.StatusOK, "hello %s %s", a, b)
			// c.String(http.StatusOK, "hello %s %s", c, d)
			c.JSON(200, gin.H{
				"a": a,
				"b": b,
				"c": cc,
				"d": dd,
			})

		})
	}

	r.Run(":8081")
	// fmt.Println("out", configor.Config{})
	// fmt.Println("in", array.Main(123))
	array.Main()
	loop.Main()
	mapgo.Main()
	str.Main()
	fun.Main()
	def.Main()
	structs.Main()
	iface.Main()
	extension.Main()
	emptyinterface.Main()
	polymorphism.Main()
	error.Main()
	panic.Main()
	client.Main()
	groutine.Main()
	csp.Main()
	channel.Main()
	cancel.Main()
	once.Main()
	firstres.Main()
	pool.Main()
	syncpool.Main()
}
