package routers

import (
	"html/template"
	"io/ioutil"

	packr2 "github.com/gobuffalo/packr/v2"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/mritd/ginmvc/utils"
)

// init html template router
func init() {
	registerWithWeight("templates", 100, func(router *gin.Engine) {
		t, err := loadTemplate()
		utils.CheckAndExit(err)
		router.SetHTMLTemplate(t)
	})
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	htmlBox := packr2.New("templates", "../templates")
	err := htmlBox.Walk(func(s string, file packr.File) error {
		fileInfo, err := file.FileInfo()
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() {
			h, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			t, err = t.New(s).Parse(string(h))
			if err != nil {
				return err
			}
		}
		return nil
	})
	return t, err

}
