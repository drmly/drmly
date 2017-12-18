package main

//go:generate go-bindata -prefix "frontend/" -pkg main -o bindata.go frontend/...

// func static_handler(rw http.ResponseWriter, req *http.Request) {
// 	var path string = req.URL.Path
// 	if path == "" {
// 		path = "index.html"
// 	} else if path == "downloads" {
// 		path = "dl.html"
// 	}
// 	if bs, err := Asset(path); err != nil {
// 		rw.WriteHeader(http.StatusNotFound)
// 	} else {
// 		var reader = bytes.NewBuffer(bs)
// 		io.Copy(rw, reader)
// 	}
// }

// func web() {
// 	http.Handle("/d", http.StripPrefix("/", http.HandlerFunc(static_handler)))
// 	http.ListenAndServe(":8080", nil)
// }
