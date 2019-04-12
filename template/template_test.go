package template

import (
	"testing"
)

func TestChunk(t *testing.T) {
	for _, m := range []struct {
		src       []byte
		chunksize int
		result    []string
	}{
		{
			[]byte("Hello World\n"), 11, []string{
				"`` +\n",
				"    `Hel` +\n",
				"    `lo ` +\n",
				"    `Wor` +\n",
				"    `ld\n`"},
		},
	} {
		result := chunk(&m.src, m.chunksize)
		if len(*result) != len(m.result) {
			t.Fatalf("chunk unexpected number of chunks: want=%d; have=%d",
				len(m.result), len(*result))
		}
		for i, c := range *result {
			if m.result[i] != c {
				t.Fatalf("chunk wrong chunk at index %d: want=%s; have=%s",
					i, m.result[i], c)
			}
		}
	}
}
