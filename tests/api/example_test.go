package api_test

import (
    "testing"

    "github.com/icey148/qa-automation/automation-framework/pkg/assertion"
)

// TestCase: TC-EXAMPLE-001
func TestFrameworkWiring(t *testing.T) {
    assertion.Equal(t, 2, 1+1)
}
