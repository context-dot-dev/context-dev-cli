// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestNewsSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"news", "search",
			"--search-by", "{entity: {domain: stripe.com, type: domain}, type: entity}",
			"--cursor", "cursor",
			"--filter-by", "{articleLanguage: [am], articleType: [editorial], date: {from: 0, to: 0}, sourceCountry: [ae], sourceDomain: [x]}",
			"--limit", "1",
			"--policy", "pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e",
			"--sort-by", "{type: relevance}",
			"--tag", "production",
			"--tag", "team-alpha",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(newsSearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"news", "search",
			"--search-by", "{entity: {domain: stripe.com, type: domain}, type: entity}",
			"--cursor", "cursor",
			"--filter-by.article-language", "[am]",
			"--filter-by.article-type", "[editorial]",
			"--filter-by.date", "{from: 0, to: 0}",
			"--filter-by.source-country", "[ae]",
			"--filter-by.source-domain", "[x]",
			"--limit", "1",
			"--policy", "pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e",
			"--sort-by.type", "relevance",
			"--tag", "production",
			"--tag", "team-alpha",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"searchBy:\n" +
			"  entity:\n" +
			"    domain: stripe.com\n" +
			"    type: domain\n" +
			"  type: entity\n" +
			"cursor: cursor\n" +
			"filterBy:\n" +
			"  articleLanguage:\n" +
			"    - am\n" +
			"  articleType:\n" +
			"    - editorial\n" +
			"  date:\n" +
			"    from: 0\n" +
			"    to: 0\n" +
			"  sourceCountry:\n" +
			"    - ae\n" +
			"  sourceDomain:\n" +
			"    - x\n" +
			"limit: 1\n" +
			"policies:\n" +
			"  - pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e\n" +
			"sortBy:\n" +
			"  type: relevance\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"news", "search",
		)
	})
}
