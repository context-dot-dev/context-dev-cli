// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestUtilityPrefetch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"utility", "prefetch",
			"--identifier", "{domain: stripe.com}",
			"--type", "brand",
			"--policy", "pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts", "{milliseconds: 1000, behavior: fail}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(utilityPrefetch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"utility", "prefetch",
			"--identifier", "{domain: stripe.com}",
			"--type", "brand",
			"--policy", "pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts.milliseconds", "1000",
			"--timeout-opts.behavior", "fail",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"identifier:\n" +
			"  domain: stripe.com\n" +
			"type: brand\n" +
			"policies:\n" +
			"  - pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutOpts:\n" +
			"  milliseconds: 1000\n" +
			"  behavior: fail\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"utility", "prefetch",
		)
	})
}
