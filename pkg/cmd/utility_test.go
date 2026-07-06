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
			"--identifier", "{domain: domain, email: dev@stainless.com}",
			"--type", "brand",
			"--timeout-ms", "1000",
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
			"--identifier.domain", "domain",
			"--identifier.email", "dev@stainless.com",
			"--type", "brand",
			"--timeout-ms", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"identifier:\n" +
			"  domain: domain\n" +
			"  email: dev@stainless.com\n" +
			"type: brand\n" +
			"timeoutMS: 1000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"utility", "prefetch",
		)
	})
}
