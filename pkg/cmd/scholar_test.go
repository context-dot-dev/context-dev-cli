// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestScholarSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"scholar", "search",
			"--query", "CRISPR off-target effects in human cells",
			"--cursor", "cursor",
			"--filter-by", "{date: {from: 0, to: 0}}",
			"--limit", "1",
			"--sort", "relevance",
			"--tag", "production",
			"--tag", "team-alpha",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(scholarSearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"scholar", "search",
			"--query", "CRISPR off-target effects in human cells",
			"--cursor", "cursor",
			"--filter-by.date", "{from: 0, to: 0}",
			"--limit", "1",
			"--sort", "relevance",
			"--tag", "production",
			"--tag", "team-alpha",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: CRISPR off-target effects in human cells\n" +
			"cursor: cursor\n" +
			"filterBy:\n" +
			"  date:\n" +
			"    from: 0\n" +
			"    to: 0\n" +
			"limit: 1\n" +
			"sort: relevance\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"scholar", "search",
		)
	})
}
