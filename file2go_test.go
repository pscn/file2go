package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/pscn/file2go/decode"

	"github.com/pscn/file2go/encode"
)

func TestEncodeDecode(t *testing.T) {
	for _, m := range []struct {
		content []byte
		name    string
		comment string
		date    string
	}{
		{[]byte("content"), "name", "comment",
			"2019-04-11T19:13:21Z"},
		{[]byte("{{.}} some %% and more %s %d"), "!…→", "<<html>>",
			"2012-01-21T18:12:57Z"},
		{[]byte(fmt.Sprintf("%s\n%d\n\t\n", "täst", 12)), "!…→", "<<html>>",
			"2015-09-12T22:23:43Z"},
	} {
		// Mon Jan 2 15:04:05 -0700 MST 2006
		now, err := time.Parse(time.RFC3339, m.date)
		if err != nil {
			t.Fatalf("test setup failed to parse date: %s; error=%s", m.date, err)
		}
		encoded, err := encode.Bytes(&m.content, m.name, m.comment, now)
		if err != nil {
			t.Fatalf("failed to encode bytes: %s", err)
		}
		decoded, err := decode.Init(string(*encoded))
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
