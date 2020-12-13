package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/configor"
)

// Login is a struct
type Login struct {
	User string `form:"user" json:"user" xml:"user"  binding:"required"`
}

// User is a struct
type User struct {
	id   int64          `db:"id"`
	name sql.NullString `db:"name"`
	sex  int64          `db:"sex"`
}

func main() {
	DB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Printf("Open mysql failed, err:%v\n", err)
	}

	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	queryOne(DB)
	queryMulti(DB)
	insertData(DB)
	updateData(DB)
	deleteData(DB)
}

func queryOne(DB *sql.DB) {
	user := new(User)
	row := DB.QueryRow("select * from test where id = ?", 2)
	err := row.Scan(&user.id, &user.name, &user.sex)
	if err != nil {
		fmt.Printf("queryOne scan failde, err:%v\n", err)
		return
	}
	fmt.Println(*user)
}

func queryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select * from test where id > ?", 0)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("queryMulti query failde, err:%v\n", err)
		return
	}
	for rows.Next() {
		err := rows.Scan(&user.id, &user.name, &user.sex)
		if err != nil {
			fmt.Printf("queryMulti scan failde, err:%v\n", err)
			return
		}
		fmt.Println(*user)
	}
}

func insertData(Db *sql.DB) {
	res, err := Db.Exec("insert INTO test(id, name, sex) values(?,?,?)", 1, "aa", 0)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Insertdata get lastinsertid failed, err:%v\n", err)
		return
	}
	fmt.Println("Insertdata lastInsertID:", lastInsertID)
	rowsAddected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Insertdata get rowsAddected failed, err:%v\n", err)
		return
	}
	fmt.Println("Insertdata rowsAddected:", rowsAddected)
}

func updateData(DB *sql.DB) {
	res, err := DB.Exec("update test set name=? where id=?", "fr", 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	rowsAddected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("update get rowsAddected failed, err:%v\n", err)
		return
	}
	fmt.Println("update rowsAddected:", rowsAddected)
}

func deleteData(DB *sql.DB) {
	res, err := DB.Exec("delete from test where id=?", 4)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	rowsAddected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("delete get rowsAddected failed, err:%v\n", err)
		return
	}
	fmt.Println("delete rowsAddected:", rowsAddected)
}

// func main() {
// 	gin.DisableConsoleColor()

// 	f, _ := os.Create("gin.log")
// 	gin.DefaultWriter = io.MultiWriter(f)

// 	r := gin.Default()
// 	// r.GET("/user", func(c *gin.Context) {
// 	// 	// name := c.Param("name")
// 	// 	a := c.Query("a")
// 	// 	b := c.DefaultQuery("b", "123")
// 	// 	// cc := c.PostForm("c")
// 	// 	// dd := c.DefaultPostForm("d", "dd")
// 	// 	// c.String(http.StatusOK, "hello %s %s", a, b)
// 	// 	// c.String(http.StatusOK, "hello %s %s", c, d)
// 	// 	c.JSON(200, gin.H{
// 	// 		"a": a,
// 	// 		"b": b,
// 	// 		// "c": cc,
// 	// 		// "d": dd,
// 	// 	})

// 	// })
// 	v1 := r.Group("v1")
// 	// v1.Use(gin.Logger())
// 	{
// 		v1.GET("/user", func(c *gin.Context) {
// 			var json Login
// 			if err := c.ShouldBindJSON(&json); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			if json.User != "manu" {
// 				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
// 				return
// 			}

// 			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
// 			// if err := c.ShouldBindJSON(&json); err != nil {
// 			// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			// 	return
// 			// }
// 			// c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
// 			// a := c.Query("a")
// 			// c.String(http.StatusOK, "user s%", a)
// 		})
// 	}
// 	// v1 := r.Group("/v1")
// 	// {
// 	// 	v1.POST("/user", func(c *gin.Context) {
// 	// 		a := c.Query("a")
// 	// 		b := c.DefaultQuery("b", "123")
// 	// 		cc := c.PostForm("c")
// 	// 		dd := c.DefaultPostForm("d", "dd")
// 	// 		// c.String(http.StatusOK, "hello %s %s", a, b)
// 	// 		// c.String(http.StatusOK, "hello %s %s", c, d)
// 	// 		c.JSON(200, gin.H{
// 	// 			"a": a,
// 	// 			"b": b,
// 	// 			"c": cc,
// 	// 			"d": dd,
// 	// 		})

// 	// 	})
// 	// }

// 	v2 := r.Group("/v2")
// 	{
// 		v2.POST("/user", func(c *gin.Context) {
// 			a := c.Query("a")
// 			b := c.DefaultQuery("b", "123")
// 			cc := c.PostForm("c")
// 			dd := c.DefaultPostForm("d", "dd")
// 			// c.String(http.StatusOK, "hello %s %s", a, b)
// 			// c.String(http.StatusOK, "hello %s %s", c, d)
// 			c.JSON(200, gin.H{
// 				"a": a,
// 				"b": b,
// 				"c": cc,
// 				"d": dd,
// 			})

// 		})
// 	}

// 	r.Run(":8081")
// 	// fmt.Println("out", configor.Config{})
// 	// fmt.Println("in", array.Main(123))
// 	array.Main()
// 	loop.Main()
// 	mapgo.Main()
// 	str.Main()
// 	fun.Main()
// 	def.Main()
// 	structs.Main()
// 	iface.Main()
// 	extension.Main()
// 	emptyinterface.Main()
// 	polymorphism.Main()
// 	error.Main()
// 	panic.Main()
// 	client.Main()
// 	groutine.Main()
// 	csp.Main()
// 	channel.Main()
// 	cancel.Main()
// 	once.Main()
// 	firstres.Main()
// 	pool.Main()
// 	syncpool.Main()
// }
