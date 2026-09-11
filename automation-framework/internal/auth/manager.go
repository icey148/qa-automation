package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var ErrSessionMissing = errors.New("authentication session is missing")

type Session struct {
	Role      string    `json:"role"`
	Provider  string    `json:"provider"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type Manager struct {
	root string
}

func NewManager(root string) *Manager { return &Manager{root: root} }

func (m *Manager) SessionPath(env, role string) string {
	return filepath.Join(m.root, env, role+".json")
}

func (m *Manager) Validate(env, role string) (Session, error) {
	path := m.SessionPath(env, role)
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Session{}, fmt.Errorf("%w: %s", ErrSessionMissing, path)
		}
		return Session{}, err
	}
	var s Session
	if err := json.Unmarshal(b, &s); err != nil {
		return Session{}, fmt.Errorf("parse session metadata: %w", err)
	}
	if !s.ExpiresAt.IsZero() && time.Now().After(s.ExpiresAt) {
		return Session{}, fmt.Errorf("authentication session expired at %s", s.ExpiresAt.Format(time.RFC3339))
	}
	return s, nil
}

func (m *Manager) Bootstrap(env, role, provider string, ttl time.Duration) (Session, error) {
	// This stores metadata only. Real browser SSO/MFA integration belongs in a provider adapter.
	s := Session{Role: role, Provider: provider, ExpiresAt: time.Now().Add(ttl)}
	path := m.SessionPath(env, role)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return Session{}, err
	}
	b, _ := json.MarshalIndent(s, "", "  ")
	if err := os.WriteFile(path, b, 0o600); err != nil {
		return Session{}, err
	}
	return s, nil
}
