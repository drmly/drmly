package main

// // (hack) add this func back to gin/gonic after using dep ensure (bc gin/gonic fork not started by me yet(is there better solution?))
// func (engine *Engine) LoadHTMLBinData(files []string, fileAssetFunc func(string) []byte) {
// 	if IsDebugging() {
// 		engine.HTMLRender = render.HTMLDebug{Files: files}
// 	} else {
// 		var templ *template.Template
// 		for _, filename := range files {
// 			content := fileAssetFunc(filename)

// 			name := filepath.Base(filename)

// 			var tmpl *template.Template
// 			if templ == nil {
// 				templ = template.New(name)
// 			}

// 			if name == templ.Name() {
// 				tmpl = templ
// 			} else {
// 				tmpl = templ.New(name)
// 			}

// 			_, err := tmpl.Parse(string(content))
// 			if err != nil {
// 				panic("template " + filename + " parse failed: " + err.Error())
// 			}
// 		}
// 		engine.SetHTMLTemplate(templ)
// 	}
// }
