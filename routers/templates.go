/*
 * Copyright 2018 mritd <mritd1234@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package routers

import (
	"html/template"
	"io/ioutil"

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
	htmlBox := packr.NewBox("../templates")
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
