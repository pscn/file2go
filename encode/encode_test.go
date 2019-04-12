package encode

import (
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	for _, m := range []struct {
		content []byte
		name    string
		date    string
		encoded string
	}{
		{
			[]byte("Hello World\n"),
			"hello_world.txt",
			"2019-04-12T04:24:00Z",
			`H4sICOASsFwA/2FHVnNiRzlmZDI5eWJHUXVkSGgwAPJIzcnJVwjPL8pJ4QIEAAD//+PllbAMAAAA`,
		},
	} {
		// Mon Jan 2 15:04:05 -0700 MST 2006
		now, err := time.Parse(time.RFC3339, m.date)
		if err != nil {
			t.Fatalf("test setup failed to parse date: %s; error=%s", m.date, err)
		}
		encoded, err := Bytes(&m.content, m.name, now)
		if err != nil {
			t.Fatalf("failed to encode bytes: %s", err)
		}
		if string(*encoded) != string(m.encoded) {
			t.Fatalf("failed to encode content: have %s want %s",
				string(*encoded), string(m.encoded))
		}
	}
}
