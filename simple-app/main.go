package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setEnv(key, def string) string {
	value, found := os.LookupEnv(key)
	if !found {
		value = def
	}
	return value
}

func main() {
	r := gin.Default()
	dir := setEnv("GIN_FILE_DIR", "./")
	static_dir := setEnv("GIN_STATIC_DIR", "./static")
	mode := os.Getenv("GIN_MODE")

	r.Static("/static", static_dir)         // serve static files from public folder
	r.GET("/healtz", func(c *gin.Context) { // healthz endpoint
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy!",
		})
	})

	r.GET("/hello", func(c *gin.Context) { // root endpoint
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World! Workgin in " + mode,
		})
	})
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dir+file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.Run() // listen and serve on port 8080
}
