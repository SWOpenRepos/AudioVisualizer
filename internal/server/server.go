package server

import (
	"github.com/gin-gonic/gin"

	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

func Start() int {
	appPath, err := os.Executable()
	if err != nil {
		return 1
	}
	dir := filepath.Dir(appPath)

	_, err1 := os.Stat(filepath.Join(dir, "control.js"))
	_, err2 := os.Stat(filepath.Join(dir, "control.wasm"))
	if err1 != nil || err2 != nil {
		return 2
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		data, err := fs.ReadFile(publicFS, "index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "index.html not found")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	r.GET("/local/control.js", func(c *gin.Context) {
		data, err := os.ReadFile(filepath.Join(dir, "control.js"))
		if err != nil {
			c.String(http.StatusNotFound, "control.js not found")
			return
		}
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", data)
	})

	r.GET("/local/control.wasm", func(c *gin.Context) {
		data, err := os.ReadFile(filepath.Join(dir, "control.wasm"))
		if err != nil {
			c.String(http.StatusNotFound, "control.wasm not found")
			return
		}
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", data)
	})

	r.StaticFS("/assets", http.FS(assetsFS))

	go r.Run(":1516")

	return 0
}
