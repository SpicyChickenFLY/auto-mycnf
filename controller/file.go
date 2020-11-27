package controller

import (
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
)

// genFile generates cnf file in specified path
func genFile(filePath, fileText string) error {
	fw, err := os.OpenFile(filePath, syscall.O_CREAT, 0644)
	defer fw.Close()
	if err != nil {
		return err
	}
	_, err = fw.Write([]byte(fileText))
	return err
}

// GenFile is a func to process request for generating cnf file
func GenFile(c *gin.Context) {
	filePath := c.DefaultPostForm("filePath", "/tmp/my.cnf")
	fileText := c.DefaultPostForm("fileText", "")
	if err := genFile(filePath, fileText); err != nil {
		c.String(http.StatusOK, "OK, it's done")
	} else {
		c.String(http.StatusInternalServerError, "Error!")
	}
}
