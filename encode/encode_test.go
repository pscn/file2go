package encode

import (
	"testing"
)

func TestEncode(t *testing.T) {
	for _, m := range []struct {
		content []byte
		name    string
		encoded string
	}{
		{
			[]byte("Hello World\n"),
			"hello_world.txt",
			`H4sICAAAAAAA/2FHVnNiRzlmZDI5eWJHUXVkSGgwAPJIzcnJVwjPL8pJ4QIEAAD//+PllbAMAAAA`,
		},
	} {
		encoded, err := Bytes(&m.content, m.name)
		if err != nil {
			t.Fatalf("failed to encode bytes: %s", err)
		}
		if string(*encoded) != string(m.encoded) {
			t.Fatalf("failed to encode content: have %s want %s",
				string(*encoded), string(m.encoded))
		}
	}
}

func TestEncodeFile(t *testing.T) {
	for _, m := range []struct {
		filename string
	}{
		{"../template/files.tmpl"}, {"../template/files_test.tmpl"},
	} {
		_, err := File(m.filename)
		if err != nil {
			t.Fatalf("failed to encode file: %s; error=%s", m.filename, err)
		}
	}
}

func TestEncodeFileNotFound(t *testing.T) {
	for _, m := range []struct {
		filename string
	}{
		{"nohave"}, {"nohave²"},
	} {
		_, err := File(m.filename)
		if err == nil {
			t.Fatalf("should have failed to encode none existing file")
		}
	}
}
