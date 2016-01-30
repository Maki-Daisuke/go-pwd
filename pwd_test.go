package pwd

import "testing"

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
