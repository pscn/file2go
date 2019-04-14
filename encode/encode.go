package encode

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path"
)

// File GZIPs the file given by filename as string, embeing the base of
// filename as name
func File(filename string) (*[]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: filename=%s; error=%s",
			filename, err)
	}
	return Bytes(&data, path.Base(filename))
}

// Bytes GZIPs the []bytes and returns a BASE64 encoded string, embeding
// and name
func Bytes(data *[]byte, name string) (*[]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	// only Latin in GZIP headers
	zw.Name = base64.StdEncoding.EncodeToString([]byte(name))
	_, err := zw.Write(*data)
	if err != nil {
		return nil, fmt.Errorf("failed to compress data: name=%s; error=%s",
			name, err)
	}
	// do not defer! otherwise the base64 encode will not see anything
	err = zw.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close GZIP writer: error=%s", err)
	}
	result := []byte(base64.StdEncoding.EncodeToString(buf.Bytes()))
	return &result, nil
}
