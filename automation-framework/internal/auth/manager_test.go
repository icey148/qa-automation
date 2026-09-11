package auth

import (
	"testing"
	"time"
)

func TestBootstrapAndValidate(t *testing.T) {
	root := t.TempDir()
	m := NewManager(root)
	if _, err := m.Bootstrap("dev", "admin", "sso", time.Hour); err != nil {
		t.Fatal(err)
	}
	got, err := m.Validate("dev", "admin")
	if err != nil {
		t.Fatal(err)
	}
	if got.Role != "admin" {
		t.Fatalf("unexpected role %q", got.Role)
	}
}
