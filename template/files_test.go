// Code generated by "file2go -v -d -t -o template/files.go template/*.tmpl"; DO NOT EDIT.

// Testing for included files:
// → template/files.tmpl
// → template/files_test.tmpl

package template

import (
	"testing"
)

func TestContentDoesNotExist(t *testing.T) {
	_, err := Content("")
	if err == nil {
		t.Fatalf("Content: returned no error")
	}
}

func TestContentExists(t *testing.T) {
	var err error

	_, err = Content("template/files.tmpl")
	if err != nil {
		t.Fatalf("Content \"template/files.tmpl\" not found: %s", err)
	}
	_, err = Content("template/files_test.tmpl")
	if err != nil {
		t.Fatalf("Content \"template/files_test.tmpl\" not found: %s", err)
	}
}

func TestContentMustDoesNotExist(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("ContentMust: should have panic'd")
		}
	}()
	_ = ContentMust("")
	t.Fatalf("ContentMust: should have panic'd")
}

func TestContentMustExist(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("ContentMust: should not have panic'd")
		}
	}()
	_ = ContentMust("template/files.tmpl")
	_ = ContentMust("template/files_test.tmpl")
}

func TestFilenames(t *testing.T) {
	filenames := Filenames()
	i := 0
	if filenames[i] != "template/files.tmpl" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "template/files.tmpl")
	}
	i++
	if filenames[i] != "template/files_test.tmpl" {
		t.Fatalf("Filenames: mismatch got '%s' want '%s'",
			filenames[i], "template/files_test.tmpl")
	}
	i++
}

// eof
