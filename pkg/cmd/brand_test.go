// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestBrandRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve",
			"--domain", "xxx",
			"--type", "by_domain",
			"--force-language", "afrikaans",
			"--max-age-ms", "0",
			"--max-speed=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts", "{milliseconds: 1000, behavior: fail}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(brandRetrieve)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve",
			"--domain", "xxx",
			"--type", "by_domain",
			"--force-language", "afrikaans",
			"--max-age-ms", "0",
			"--max-speed=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-opts.milliseconds", "1000",
			"--timeout-opts.behavior", "fail",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domain: xxx\n" +
			"type: by_domain\n" +
			"force_language: afrikaans\n" +
			"maxAgeMs: 0\n" +
			"maxSpeed: true\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutOpts:\n" +
			"  milliseconds: 1000\n" +
			"  behavior: fail\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"brand", "retrieve",
		)
	})
}

func TestBrandRetrieveSimplified(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve-simplified",
			"--domain", "xxx",
			"--max-age-ms", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--theme", "light",
			"--timeout-opts", "{milliseconds: 1000, behavior: fail}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(brandRetrieveSimplified)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve-simplified",
			"--domain", "xxx",
			"--max-age-ms", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--theme", "light",
			"--timeout-opts.milliseconds", "1000",
			"--timeout-opts.behavior", "fail",
		)
	})
}

func TestBrandSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "search",
			"--query", "x",
			"--autocomplete=true",
			"--query-by", "name",
			"--tag", "production",
			"--tag", "team-alpha",
			"--typo-tolerance", "0",
		)
	})
}
