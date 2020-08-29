package main

import (
	"fmt"
	"three_girls/db"

	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type msg struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Slogan string `json:"slogan"`
}

type future struct {
	Wish  string `json:"wish" form:"wh"`
	Color string `json:"color" form:"cr"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	// router.LoadHTMLFiles("./upload.html")

	router.Use(m1)

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello, world!",
		})
	})

	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index/Jack",
		})
	})

	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index/Pan",
		})
	})

	router.GET("/json", func(c *gin.Context) {
		var data map[string]interface{} = map[string]interface{}{
			"name":   "杰克爱",
			"age":    18,
			"gender": "男",
		}
		c.JSON(http.StatusOK, data)
	})

	router.GET("/another_json", func(c *gin.Context) {
		data := msg{
			"alice",
			22,
			"hello, Golang~",
		}
		c.JSON(http.StatusOK, data)
	})

	router.GET("/query", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "18")
		interest, ok := c.GetQuery("interest")
		if !ok {
			interest = "beauty"
		}
		c.JSON(http.StatusOK, gin.H{
			"name":     name,
			"age":      age,
			"interest": interest,
		})
	})

	//get parameters seperated with '/' in URL
	router.GET("/user/:account/:password", func(c *gin.Context) {
		account := c.Param("account")
		password := c.Param("password")
		c.JSON(http.StatusOK, gin.H{
			"account":  account,
			"password": password,
		})
	})

	router.GET("/binding", func(c *gin.Context) {
		wish := c.Query("wish")
		color := c.Query("color")
		future := future{
			Wish:  wish,
			Color: color,
		}
		fmt.Printf("%#v\n", future)
		c.JSON(http.StatusOK, gin.H{})
	})

	router.POST("/form", func(c *gin.Context) {
		var f future
		err := c.ShouldBind(&f)
		if err != nil {
			c.HTML(http.StatusBadRequest, "posts/index2.html", gin.H{
				"title": "bad request 404",
			})
		} else {
			fmt.Printf("%#v\n", f)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	// router.GET("/upload", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "upload.html", nil)
	// })

	// router.POST("/upload", func(c *gin.Context) {
	// 	file, err := c.FormFile("f1")
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	fmt.Println(file.Filename)
	// 	dst := fmt.Sprintf("assets/%s", file.Filename)
	// 	c.SaveUploadedFile(file, dst)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "ok",
	// 	})
	// })

	router.GET("/posts/uploadn", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/uploadn.html", nil)
	})

	router.POST("/posts/uploadn", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["f1s"]

		for _, file := range files {
			fmt.Println(file.Filename)
			dst := fmt.Sprintf("assets/%s", file.Filename)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "upload several files successfully",
		})
	})

	router.GET("/qq", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "https://www.qq.com")
	})

	router.GET("/jack", func(c *gin.Context) {
		c.Request.URL.Path = "/rose"
		router.HandleContext(c)
	})

	router.GET("/rose", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "where is the diamond?",
		})
	})

	router.GET("/login", authenticateIdentity, login)

	router.Any("/any", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"message": "any_get",
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"message": "any_default",
			})
		}
	})

	show := router.Group("/show", checkSth(true))
	{
		show.GET("/girl", func(c *gin.Context) {
			gender, ok := c.Get("gender")
			if !ok {
				gender = "undefined"
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "I love this girl",
				"gender":  gender,
			})
		})
		show.GET("/boy", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "This boy is so lovely",
			})
		})
	}

	router.GET("/sql/insert", func(c *gin.Context) {

		db.InitDatabase()
		c.JSON(http.StatusOK, gin.H{
			"message": "sql insert succeed",
		})
	})

	router.Run(":8888")
}

func m1(c *gin.Context) {
	start := time.Now()
	fmt.Println("start time", start)
	c.Next()
	end := time.Since(start)
	fmt.Println("consume time: ", end)
}

func authenticateIdentity(c *gin.Context) {
	start := time.Now()
	fmt.Println("authentication time", start)
	name := c.Query("name")
	pwd := c.Query("pwd")
	if name == "admin" && pwd == "123" {
		c.Next()
	} else {
		c.Abort()
	}
	end := time.Since(start)
	fmt.Println("consume time: ", end)
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "log in successfully",
	})
}

func checkSth(needCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if needCheck {
			fmt.Println("check something: ", time.Now())
			c.Set("gender", "male")
			c.Next()
		} else {
			c.Next()
		}
	}
}
