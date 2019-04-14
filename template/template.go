package template

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"text/template"
	"time"
)

// File keeping the name & the content
type File struct {
	Name    string
	Content *[]byte
}

// FIXME: naming → find a better name for this
func chunk(src *[]byte, chunksize int) *[]string {
	// we add "  `` +\n" to each chunk (ignoring \n)
	chunksize -= 8
	// estimate the nr of slices we need
	nr := int(math.Ceil(float64(len(*src))/float64(chunksize)) + 1)
	result := make([]string, nr)
	// hack: let the first chunk be empty to only add a "`` +" to the template
	result[0] = "``"
	for pos, i, length, end := 0, 1, len(*src), chunksize; pos < length; pos, i, end = pos+chunksize, i+1, end+chunksize {
		result[i-1] += " +\n"
		if end > length {
			end = length
		}
		result[i] = "    `" + string((*src)[pos:end]) + "`"
	}
	return &result
}

// Parse returns a go code template FIXME:
func Parse(container *[]File, arguments, pkg string, devel bool) (*[]byte, error) {
	tmplData := struct {
		Arguments string
		Pkg       string
		Container []File
		Date      string
	}{
		Arguments: arguments,
		Pkg:       pkg,
		Container: *container,
		Date:      time.Now().String(),
	}
	tmplStr := ""
	if devel {
		tmplFromFile, err := ioutil.ReadFile("template/files.tmpl")
		if err != nil {
			return nil, err
		}
		tmplStr = string(tmplFromFile)
	} else {
		dataFromGo, err := Content("template/files.tmpl")
		if err != nil {
			return nil, err
		}
		tmplStr = string(*dataFromGo)
	}
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"Chunk": func(src *[]byte, cs int) *[]string { return chunk(src, cs) },
	}).Parse(tmplStr)

	if err != nil {
		return nil, fmt.Errorf("failed to parse internal template: error=%s", err)
	}
	var b strings.Builder
	err = tmpl.Execute(&b, tmplData)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: error=%s", err)
	}
	result := []byte(b.String())
	return &result, nil
}

// ParseTest returns a go code template FIXME:
func ParseTest(container *[]File, arguments, pkg string, devel bool) (*[]byte, error) {
	tmplData := struct {
		Arguments string
		Pkg       string
		Container []File
		Date      string
	}{
		Arguments: arguments,
		Pkg:       pkg,
		Container: *container,
		Date:      time.Now().String(),
	}
	tmplStr := ""
	if devel {
		tmplFromFile, err := ioutil.ReadFile("template/files_test.tmpl")
		if err != nil {
			return nil, err
		}
		tmplStr = string(tmplFromFile)
	} else {
		dataFromGo, err := Content("template/files_test.tmpl")
		if err != nil {
			return nil, err
		}
		tmplStr = string(*dataFromGo)
	}
	tmpl, err := template.New("").Parse(tmplStr)

	if err != nil {
		return nil, fmt.Errorf("failed to parse internal template: error=%s", err)
	}
	var b strings.Builder
	err = tmpl.Execute(&b, tmplData)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: error=%s", err)
	}
	result := []byte(b.String())
	return &result, nil
}
