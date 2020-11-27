package controller

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// getCnfPara is a func to find all specified k-v pair in configure file
// FIXME: for now it only support single instance conf
func getCnfPara(srcCnfFile string) (string, error) {
	// fmt.Println(srcCnfFile)
	result := ""

	fr, err := os.Open(srcCnfFile)
	if err != nil {
		return "", err
	}
	defer fr.Close()

	br := bufio.NewReader(fr)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			// fmt.Println(result)
			return result, nil
		} else if err != nil {
			fmt.Println(string(line))
			return "", err
		}

		// ignore blank
		if len(line) == 0 {
			continue
		}
		// [xxx] is a chapter mark
		if line[0] == '[' {
			result += fmt.Sprintf(
				`<hr class="mb-4">
				<h4 data-type="chapter" class="mb-3 param">%s</h4>`,
				string(line))
			continue
		}
		// ### xxxx is a module mark
		if line[0] == '#' && line[1] == '#' && line[2] == '#' {
			result += fmt.Sprintf(
				`<hr class="mb-4">
				<h5 class="mb-3">%s</h5>`,
				string(line[4:]))
			continue
		}
		// ## xxxx is a usual comment
		if line[0] == '#' && line[1] == '#' {
			result += fmt.Sprintf(
				"<small class=\"text-muted\">%s</small><br/>",
				string(line[3:]))
			continue
		}

		// fmt.Println(string(line))
		lineStr := string(line)
		activated := `checked="checked"`
		// # xxxx is an unactivated parameter
		if line[0] == '#' {
			activated = ""
			lineStr = string(line[2:])
		}
		strArr := strings.Split(lineStr, "#")
		parameter := strArr[0]
		comment := ""
		if len(strArr) > 1 {
			comment = strArr[1]
		}

		signIndex := strings.Index(parameter, "=") // find "="
		if signIndex < 0 {
			// single parameter
			result += fmt.Sprintf(
				`<div class="row">
					<div data-type="single" class="col-md-12 mb-3 param">
						<input type="checkbox" %s />
						<label>%s</label>
						<small class="text-muted">%s</small>
					</div>
				</div>`, activated, parameter, comment)
		} else {
			// kv pair
			key := string(parameter[:signIndex])
			value := string(parameter[signIndex+1:])
			result += fmt.Sprintf(
				`<div class="row">
					<div data-type="kv" class="col-md-12 mb-3 param">
						<input type="checkbox" %s />
						<label>%s</label><small class="text-muted">%s</small><br/>
						<input class="form-control" type="text" value='%s' />
					</div>
				</div>`, activated, key, comment, value)
		}
	}
}

// GetCnf is a func of controller process conf html request
func GetCnf(c *gin.Context) {
	if result, err := getCnfPara("static/mycnf.template"); err != nil {
		c.HTML(http.StatusOK, "error.tmpl",
			gin.H{"error": err})
	} else {
		c.HTML(http.StatusOK, "index.tmpl",
			gin.H{"html": template.HTML(result)})
	}
}
