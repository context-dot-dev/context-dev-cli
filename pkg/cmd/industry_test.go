// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestIndustryRetrieveNaics(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"industry", "retrieve-naics",
			"--input", "xxxx",
			"--max-results", "1",
			"--min-results", "1",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts", "{milliseconds: 1000, behavior: fail}",
			"--zdr", "enabled",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(industryRetrieveNaics)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"industry", "retrieve-naics",
			"--input", "xxxx",
			"--max-results", "1",
			"--min-results", "1",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts.milliseconds", "1000",
			"--timeout-opts.behavior", "fail",
			"--zdr", "enabled",
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
			"--input", "xxxx",
			"--max-results", "1",
			"--min-results", "1",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts", "{milliseconds: 1000, behavior: fail}",
			"--type", "original_sic",
			"--zdr", "enabled",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(industryRetrieveSic)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"industry", "retrieve-sic",
			"--input", "xxxx",
			"--max-results", "1",
			"--min-results", "1",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts.milliseconds", "1000",
			"--timeout-opts.behavior", "fail",
			"--type", "original_sic",
			"--zdr", "enabled",
		)
	})
}
