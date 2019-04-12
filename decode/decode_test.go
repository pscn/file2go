package decode

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	for _, m := range []struct {
		encoded string
		content []byte
		name    string
		comment string
		date    string
	}{
		{
			`H4sIGOASsFwA/2FHVnNiRzlmZDI5eWJHUXVkSGgwAFpXNWpiMlJsWkNCaWVTQm1hV3hsTW1kdgDySM3JyVcIzy/KSeECBAAA///j5ZWwDAAAAA==`,
			[]byte("Hello World\n"),
			"hello_world.txt",
			"encoded by file2go",
			"2019-04-12T04:24:00Z"},
	} {
		// Mon Jan 2 15:04:05 -0700 MST 2006
		now, err := time.Parse(time.RFC3339, m.date)
		if err != nil {
			t.Fatalf("test setup failed to parse date: %s; error=%s", m.date, err)
		}
		decoded, err := Init(m.encoded)
		if err != nil {
			t.Fatalf("failed to decode string: %s", err)
		}
		if string(*decoded.Content()) != string(m.content) {
			t.Fatalf("failed to decode content: have %s want %s",
				*decoded.Content(), m.content)
		}
		if *decoded.Name() != m.name {
			t.Fatalf("failed to decode name: have %s want %s",
				*decoded.Name(), m.name)
		}
		if *decoded.Comment() != m.comment {
			t.Fatalf("failed to decode comment: have %s want %s",
				*decoded.Comment(), m.comment)
		}
		if decoded.ModTime().UTC() != now {
			t.Fatalf("failed to decode modTime: have %s want %s",
				decoded.ModTime().UTC(), now)
		}
	}
}
