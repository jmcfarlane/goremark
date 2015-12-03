package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"path/filepath"
)

// Command line arguments and constants
// Usage: https://github.com/gnab/remark/wiki/Markdown
var (
	file   = flag.String("f", "README.md", "Markdown file to consume")
	listen = flag.String("l", ":8080", "TCP Port to serve from")
	static = flag.String("s", "docs", "Directory for static files")
	footer = `
    </textarea>
    <script src="https://gnab.github.io/remark/downloads/remark-latest.min.js">
    </script>
    <script>
      var slideshow = remark.create();
    </script>
  </body>
</html>
	`

	header = `
<!DOCTYPE html>
<html>
  <head>
    <title>Title</title>
    <meta charset="utf-8">
    <style>
      @import url(https://fonts.googleapis.com/css?family=Yanone+Kaffeesatz);
      @import url(https://fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic);
      @import url(https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700,400italic);

      body { font-family: 'Droid Serif'; }
      h1, h2, h3 {
        font-family: 'Yanone Kaffeesatz';
        font-weight: normal;
      }
      .remark-code, .remark-inline-code { font-family: 'Ubuntu Mono'; }
    </style>
  </head>
  <body>
    <textarea id="source">
	`
)

// Calculate the user's home dirctory
func getHomeDirPath() string {
	usr, err := user.Current()
	if err != nil {
		panic("Unable to get the current user")
	}
	return usr.HomeDir
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Fprintf(w, "Error:\n\n", err)
	} else {
		fmt.Fprintf(w, header+string(b)+footer)
	}
}

func main() {
	flag.Parse()
	static := "/" + *static + "/"
	http.HandleFunc("/", handler)
	http.Handle(static, http.FileServer(http.Dir(filepath.Dir(*file))))
	fmt.Println("Listening on", *listen)
	http.ListenAndServe(*listen, nil)
}