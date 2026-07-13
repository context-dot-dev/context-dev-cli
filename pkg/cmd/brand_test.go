// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
)

func TestBrandRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve",
			"--domain", "domain",
			"--type", "by_domain",
			"--force-language", "afrikaans",
			"--max-age-ms", "0",
			"--max-speed=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domain: domain\n" +
			"type: by_domain\n" +
			"force_language: afrikaans\n" +
			"maxAgeMs: 0\n" +
			"maxSpeed: true\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutMS: 1000\n")
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
			"--domain", "domain",
			"--max-age-ms", "86400000",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})
}
