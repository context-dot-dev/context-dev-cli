// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
)

func TestIndustryRetrieveNaics(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"industry", "retrieve-naics",
			"--input", "input",
			"--max-results", "1",
			"--min-results", "1",
			"--timeout-ms", "1000",
		)
	})
}

func TestIndustryRetrieveSic(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"industry", "retrieve-sic",
			"--input", "input",
			"--max-results", "1",
			"--min-results", "1",
			"--timeout-ms", "1000",
			"--type", "original_sic",
		)
	})
}
