package decode

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"time"
)

// File stores name, content, comment and modTime of the decoded string
type File struct {
	content *[]byte
	comment *string
	name    *string
	modTime *time.Time
}

// Content of the *File
func (f *File) Content() *[]byte {
	return f.content
}

// Comment for the *File
func (f *File) Comment() *string {
	return f.comment
}

// Name of the *File
func (f *File) Name() *string {
	return f.name
}

// ModTime of the *File
func (f *File) ModTime() *time.Time {
	return f.modTime
}

func gzipReader(base64Encoded *string) (*gzip.Reader, error) {
	gzipEncoded, err := base64.StdEncoding.DecodeString(*base64Encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data(BASE64): %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.Write(gzipEncoded)
	if err != nil {
		return nil, fmt.Errorf("failed buffer decode data: %s", err)
	}
	zr, err := gzip.NewReader(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader from buffer: %s", err)
	}
	return zr, nil
}

// Content decodes the BASE64 encoded GZIP encoded data from string to string
func content(zr *gzip.Reader) ([]byte, error) {
	decoded, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data(GZIP): %s", err)
	}
	return decoded, nil
}

// Init populates *File with data decoded from base64Encoded string
func Init(base64Encoded string) (*File, error) {
	zr, err := gzipReader(&base64Encoded)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	c, err := content(zr)
	if err != nil {
		return nil, err
	}
	bName, err := base64.StdEncoding.DecodeString(zr.Name)
	if err != nil {
		return nil, err
	}
	name := string(bName)
	bComment, err := base64.StdEncoding.DecodeString(zr.Comment)
	if err != nil {
		return nil, err
	}
	comment := string(bComment)
	return &File{
		content: &c,
		name:    &name,
		comment: &comment,
		modTime: &zr.ModTime,
	}, nil
}
