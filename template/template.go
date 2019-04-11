package template

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

func Generate(filename, pkg, prefix string, content *string) ([]byte, error) {
	tmplData := struct {
		Filename string
		Pkg      string
		Prefix   string
		Content  string
		Raw      string
	}{
		Filename: strings.Join(os.Args[1:], " "),
		Pkg:      pkg,
		Prefix:   prefix,
		Content:  *content,
		Raw:      "`",
	}

	tmpl, err := template.New("").Parse(`
// Code generated by "file2go {{.Filename}}"; DO NOT EDIT.

package {{.Pkg}}

import (
  "time"

  "github.com/pscn/file2go/decode"
)

const content{{.Prefix}} = {{.Raw}}{{.Content}}{{.Raw}}

var file{{.Prefix}} *decode.File

func init() {
  var err error
  file{{.Prefix}}, err = decode.Init(content{{.Prefix}})
  if err != nil {
  	panic(err)
  }
}

func {{.Prefix}}Content() []byte    { return file{{.Prefix}}.Content() }
func {{.Prefix}}Name() string       { return file{{.Prefix}}.Name() }
func {{.Prefix}}Comment() string    { return file{{.Prefix}}.Comment() }
func {{.Prefix}}ModTime() time.Time { return file{{.Prefix}}.ModTime() }

// eof
`)
	if err != nil {
		return nil, fmt.Errorf("failed to parse internal template: error=%s", err)
	}
	var b strings.Builder
	err = tmpl.Execute(&b, tmplData)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: error=%s", err)
	}
	return []byte(b.String()), nil
}