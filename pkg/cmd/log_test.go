// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
)

func TestLogsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"logs", "retrieve",
			"--request-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLogsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"logs", "list",
			"--error-code", "WEBSITE_ACCESS_ERROR",
			"--errors-only=true",
			"--from", "'2026-09-10T00:00:00Z'",
			"--key-id", "x",
			"--limit", "1",
			"--page", "1",
			"--path", "/brand/retrieve",
			"--search", "acme.com",
			"--status-code", "500",
			"--tags", "nightly-import,team-alpha",
			"--to", "'2026-09-11T00:00:00Z'",
		)
	})
}
