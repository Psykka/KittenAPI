package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Path      string
	Route     string
	DirSize   int
	Extension string
}

var assetsDir = []Route{
	{Path: "./assets/gif", Route: "/gif", DirSize: 0, Extension: "gif"},
	{Path: "./assets/img", Route: "/img", DirSize: 0, Extension: "jpg"},
	{Path: "./assets/video", Route: "/video", DirSize: 0, Extension: "mp4"},
}

func getDirSize(path string) int {
	fileCount := 0

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileCount++
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return fileCount
}

func getRandomFile(max int) string {
	n := rand.Int()%max + 1

	return fmt.Sprintf("%05d", n)
}

func init() {
	for i := range assetsDir {
		assetsDir[i].DirSize = getDirSize(assetsDir[i].Path)
	}
}

func main() {
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	})

	// Routes
	for dir := range assetsDir {
		r.GET("/"+assetsDir[dir].Route, func(c *gin.Context) {
			randFile := getRandomFile(assetsDir[dir].DirSize)
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
			c.File(assetsDir[dir].Path + "/" + randFile + "." + assetsDir[dir].Extension)
		})
	}

	r.GET("/", func(c *gin.Context) {
		message := ""
		for i := range assetsDir {
			message += assetsDir[i].Route + " - " + fmt.Sprint(assetsDir[i].DirSize) + " files\n"
		}

		c.String(200, message)
	})

	r.Run()
}
