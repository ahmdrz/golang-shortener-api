package googleshortener

import (
	"testing"
)

var key = "-------" //google developer key

func TestConnection(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	if key == "" {
		t.Fatal("Could't find key, aborting.")
	}

	_, err := New(key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestShort(t *testing.T) {
	if key == "" {
		t.Fatal("Could't find key, aborting.")
	}

	google, err := New(key)
	if err != nil {
		t.Fatal(err)
	}

	res, err := google.MakeShort("http://yahoo.com")
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Id) == 0 {
		t.Fatal("Empty response id")
	}
}
