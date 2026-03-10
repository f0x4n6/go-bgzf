package bgzf

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	f, err := os.Open("./testdata/moby.txt.gz")

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = f.Close()
	}()

	r, err := NewReader(f, 0)

	if err != nil {
		t.Fatal(err)
	}

	b, err := io.ReadAll(r)

	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(string(b), "The Project Gutenberg eBook of Moby Dick") {
		t.Fatal("invalid data")
	}
}
