package testcontext

import (
	"os"
	"testing"

	"github.com/icey148/qa-automation/automation-framework/pkg/api"
)

type Context struct {
	API *api.Client
}

func New(t testing.TB) *Context {
	t.Helper()
	baseURL := os.Getenv("QA_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return &Context{API: api.New(baseURL)}
}
