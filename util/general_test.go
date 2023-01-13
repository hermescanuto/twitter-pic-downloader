package util

import "testing"

func TestGetFolder(t *testing.T) {

	SetFolder()

	newfolder := GetFolder()

	if *Folder != newfolder {
		t.Fatalf("got %s, wanted %s", *Folder, newfolder)
	}
}
