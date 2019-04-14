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
