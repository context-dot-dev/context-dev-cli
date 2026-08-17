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
			"--search-by", "{entity: {name: xx, type: name}, type: entity}",
			"--cursor", "cursor",
			"--filter-by", "{articleLanguage: [ar], articleType: [editorial], date: {from: 0, to: 0}, sourceCountry: [ae], sourceDomain: [x]}",
			"--limit", "1",
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
			"--search-by.entity", "{name: xx, type: name}",
			"--search-by.type", "entity",
			"--cursor", "cursor",
			"--filter-by.article-language", "[ar]",
			"--filter-by.article-type", "[editorial]",
			"--filter-by.date", "{from: 0, to: 0}",
			"--filter-by.source-country", "[ae]",
			"--filter-by.source-domain", "[x]",
			"--limit", "1",
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
			"    name: xx\n" +
			"    type: name\n" +
			"  type: entity\n" +
			"cursor: cursor\n" +
			"filterBy:\n" +
			"  articleLanguage:\n" +
			"    - ar\n" +
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
