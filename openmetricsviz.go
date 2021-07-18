package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()
	r.Static("/", "./public")
	openbrowser("http://localhost:50505")
	r.Run("0.0.0.0:50505")

}