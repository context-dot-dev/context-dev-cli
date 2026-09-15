// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestAIExtractProduct(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "extract-product",
			"--url", "https://example.com",
			"--max-age-ms", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts", "{milliseconds: 1000, behavior: fail}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiExtractProduct)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "extract-product",
			"--url", "https://example.com",
			"--max-age-ms", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts.milliseconds", "1000",
			"--timeout-opts.behavior", "fail",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com\n" +
			"maxAgeMs: 0\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutOpts:\n" +
			"  milliseconds: 1000\n" +
			"  behavior: fail\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai", "extract-product",
		)
	})
}

func TestAIExtractProducts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "extract-products",
			"--domain", "domain",
			"--max-age-ms", "0",
			"--max-products", "1",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts", "{milliseconds: 1000, behavior: fail}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiExtractProducts)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "extract-products",
			"--domain", "domain",
			"--max-age-ms", "0",
			"--max-products", "1",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts.milliseconds", "1000",
			"--timeout-opts.behavior", "fail",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domain: domain\n" +
			"maxAgeMs: 0\n" +
			"maxProducts: 1\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutOpts:\n" +
			"  milliseconds: 1000\n" +
			"  behavior: fail\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai", "extract-products",
		)
	})
}
