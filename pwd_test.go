package pwd

import (
	"fmt"
	"testing"
)

func TestGetpwnam(t *testing.T) {
	pwd := Getpwnam("root")
	if pwd == nil {
		t.Fatal("expected: non-nil pointer, but actual: nil")
	}
	if pwd.Name != "root" {
		t.Fatalf(`expected: "root", but actual: %q`, pwd.Name)
	}
	if pwd.UID != 0 {
		t.Fatalf(`expected: 0, but actual: %d`, pwd.UID)
	}
	pwd = Getpwnam("nobody")
	if pwd.Name != "nobody" {
		t.Fatalf(`expected: "nobody", but actual: %q`, pwd.Name)
	}
	pwd = Getpwnam("not existing")
	if pwd != nil {
		t.Fatalf(`expected: nil pointer, but actual: %v`, pwd)
	}
}

func TestGetpwuid(t *testing.T) {
	pwd := Getpwuid(0)
	if pwd == nil {
		t.Fatal("expected: non-nil pointer, but actual: nil")
	}
	if pwd.Name != "root" {
		t.Fatalf(`expected: "root", but actual: %q`, pwd.Name)
	}
	if pwd.UID != 0 {
		t.Fatalf(`expected: 0, but actual: %d`, pwd.UID)
	}
	pwd = Getpwuid(1234556789) // uid which does not exit probably
	if pwd != nil {
		t.Fatalf(`expected: nil pointer, but actual: %v`, pwd)
	}
}

func TestGetpwents(t *testing.T) {
	ents := Getpwents()
	if !(len(ents) > 0) {
		t.Fatalf(`expected: non-empty slice, but actual: %v`, ents)
	}
	for _, pw := range ents {
		if pw.Name == "rot" {
			return
		}
	}
	str := "["
	for _, pw := range ents {
		str += fmt.Sprintf("%v ", pw)
	}
	str += "]"
	t.Fatalf(`"root" user must exist, but cannot find it in %s`, str)
}
