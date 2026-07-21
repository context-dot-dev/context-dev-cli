// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestWebExtract(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "extract",
			"--schema", "{type: bar, properties: bar, required: bar, additionalProperties: bar}",
			"--url", "https://example.com",
			"--fact-check=true",
			"--follow-subdomains=true",
			"--include-frames=true",
			"--instructions", "instructions",
			"--max-age-ms", "0",
			"--max-depth", "0",
			"--max-pages", "1",
			"--pdf", "{end: 1, shouldParse: true, start: 1}",
			"--settle-animations=true",
			"--stop-after-ms", "10000",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
			"--wait-for-ms", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webExtract)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "extract",
			"--schema", "{type: bar, properties: bar, required: bar, additionalProperties: bar}",
			"--url", "https://example.com",
			"--fact-check=true",
			"--follow-subdomains=true",
			"--include-frames=true",
			"--instructions", "instructions",
			"--max-age-ms", "0",
			"--max-depth", "0",
			"--max-pages", "1",
			"--pdf.end", "1",
			"--pdf.should-parse=true",
			"--pdf.start", "1",
			"--settle-animations=true",
			"--stop-after-ms", "10000",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
			"--wait-for-ms", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"schema:\n" +
			"  type: bar\n" +
			"  properties: bar\n" +
			"  required: bar\n" +
			"  additionalProperties: bar\n" +
			"url: https://example.com\n" +
			"factCheck: true\n" +
			"followSubdomains: true\n" +
			"includeFrames: true\n" +
			"instructions: instructions\n" +
			"maxAgeMs: 0\n" +
			"maxDepth: 0\n" +
			"maxPages: 1\n" +
			"pdf:\n" +
			"  end: 1\n" +
			"  shouldParse: true\n" +
			"  start: 1\n" +
			"settleAnimations: true\n" +
			"stopAfterMs: 10000\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutMS: 1000\n" +
			"waitForMs: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web", "extract",
		)
	})
}

func TestWebExtractCompetitors(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "extract-competitors",
			"--domain", "xxx",
			"--num-competitors", "1",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})
}

func TestWebExtractFonts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "extract-fonts",
			"--direct-url", "https://example.com",
			"--domain", "xxx",
			"--max-age-ms", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})
}

func TestWebExtractStyleguide(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "extract-styleguide",
			"--color-scheme", "light",
			"--direct-url", "https://example.com",
			"--domain", "xxx",
			"--max-age-ms", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})
}

func TestWebScreenshot(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "screenshot",
			"--color-scheme", "light",
			"--country", "de",
			"--direct-url", "https://example.com",
			"--domain", "xxx",
			"--full-screenshot", "true",
			"--handle-cookie-popup", "'true'",
			"--max-age-ms", "0",
			"--page", "login",
			"--scroll-offset", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--viewport", "{height: 240, width: 240}",
			"--wait-for-ms", "0",
			"--zdr", "enabled",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webScreenshot)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "screenshot",
			"--color-scheme", "light",
			"--country", "de",
			"--direct-url", "https://example.com",
			"--domain", "xxx",
			"--full-screenshot", "true",
			"--handle-cookie-popup", "'true'",
			"--max-age-ms", "0",
			"--page", "login",
			"--scroll-offset", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--viewport.height", "240",
			"--viewport.width", "240",
			"--wait-for-ms", "0",
			"--zdr", "enabled",
		)
	})
}

func TestWebSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "search",
			"--query", "x",
			"--country", "af",
			"--exclude-domain", "string",
			"--freshness", "last_24_hours",
			"--include-domain", "string",
			"--markdown-options", "{enabled: true, includeFrames: true, includeImages: true, includeLinks: true, maxAgeMs: 0, pdf: {end: 1, shouldParse: true, start: 1}, shortenBase64Images: true, timeoutMS: 1000, useMainContentOnly: true, waitForMs: 0}",
			"--num-results", "10",
			"--query-fanout=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webSearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "search",
			"--query", "x",
			"--country", "af",
			"--exclude-domain", "string",
			"--freshness", "last_24_hours",
			"--include-domain", "string",
			"--markdown-options.enabled=true",
			"--markdown-options.include-frames=true",
			"--markdown-options.include-images=true",
			"--markdown-options.include-links=true",
			"--markdown-options.max-age-ms", "0",
			"--markdown-options.pdf", "{end: 1, shouldParse: true, start: 1}",
			"--markdown-options.shorten-base64-images=true",
			"--markdown-options.timeout-ms", "1000",
			"--markdown-options.use-main-content-only=true",
			"--markdown-options.wait-for-ms", "0",
			"--num-results", "10",
			"--query-fanout=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: x\n" +
			"country: af\n" +
			"excludeDomains:\n" +
			"  - string\n" +
			"freshness: last_24_hours\n" +
			"includeDomains:\n" +
			"  - string\n" +
			"markdownOptions:\n" +
			"  enabled: true\n" +
			"  includeFrames: true\n" +
			"  includeImages: true\n" +
			"  includeLinks: true\n" +
			"  maxAgeMs: 0\n" +
			"  pdf:\n" +
			"    end: 1\n" +
			"    shouldParse: true\n" +
			"    start: 1\n" +
			"  shortenBase64Images: true\n" +
			"  timeoutMS: 1000\n" +
			"  useMainContentOnly: true\n" +
			"  waitForMs: 0\n" +
			"numResults: 10\n" +
			"queryFanout: true\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutMS: 1000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web", "search",
		)
	})
}

func TestWebWebCrawlMd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-crawl-md",
			"--url", "https://example.com",
			"--country", "de",
			"--exclude-selector", "string",
			"--follow-subdomains=true",
			"--include-frames=true",
			"--include-images=true",
			"--include-links=true",
			"--include-selector", "string",
			"--max-age-ms", "0",
			"--max-depth", "0",
			"--max-pages", "1",
			"--pdf", "{end: 1, ocr: true, shouldParse: true, start: 1}",
			"--settle-animations=true",
			"--shorten-base64-images=true",
			"--stop-after-ms", "10000",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
			"--url-regex", "^https?://[^/]+/blog/",
			"--use-main-content-only=true",
			"--wait-for-ms", "0",
			"--zdr", "enabled",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webWebCrawlMd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-crawl-md",
			"--url", "https://example.com",
			"--country", "de",
			"--exclude-selector", "string",
			"--follow-subdomains=true",
			"--include-frames=true",
			"--include-images=true",
			"--include-links=true",
			"--include-selector", "string",
			"--max-age-ms", "0",
			"--max-depth", "0",
			"--max-pages", "1",
			"--pdf.end", "1",
			"--pdf.ocr=true",
			"--pdf.should-parse=true",
			"--pdf.start", "1",
			"--settle-animations=true",
			"--shorten-base64-images=true",
			"--stop-after-ms", "10000",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
			"--url-regex", "^https?://[^/]+/blog/",
			"--use-main-content-only=true",
			"--wait-for-ms", "0",
			"--zdr", "enabled",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com\n" +
			"country: de\n" +
			"excludeSelectors:\n" +
			"  - string\n" +
			"followSubdomains: true\n" +
			"includeFrames: true\n" +
			"includeImages: true\n" +
			"includeLinks: true\n" +
			"includeSelectors:\n" +
			"  - string\n" +
			"maxAgeMs: 0\n" +
			"maxDepth: 0\n" +
			"maxPages: 1\n" +
			"pdf:\n" +
			"  end: 1\n" +
			"  ocr: true\n" +
			"  shouldParse: true\n" +
			"  start: 1\n" +
			"settleAnimations: true\n" +
			"shortenBase64Images: true\n" +
			"stopAfterMs: 10000\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutMS: 1000\n" +
			"urlRegex: ^https?://[^/]+/blog/\n" +
			"useMainContentOnly: true\n" +
			"waitForMs: 0\n" +
			"zdr: enabled\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web", "web-crawl-md",
		)
	})
}

func TestWebWebScrapeHTML(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-scrape-html",
			"--url", "https://example.com",
			"--country", "de",
			"--exclude-selector", "[x]",
			"--headers", "{foo: J!}",
			"--include-frames", "'true'",
			"--include-selector", "[x]",
			"--max-age-ms", "0",
			"--pdf", "{end: 1, ocr: 'true', shouldParse: 'true', start: 1}",
			"--settle-animations", "'true'",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--use-main-content-only", "'true'",
			"--wait-for-ms", "0",
			"--zdr", "enabled",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webWebScrapeHTML)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-scrape-html",
			"--url", "https://example.com",
			"--country", "de",
			"--exclude-selector", "[x]",
			"--headers", "{foo: J!}",
			"--include-frames", "'true'",
			"--include-selector", "[x]",
			"--max-age-ms", "0",
			"--pdf.end", "1",
			"--pdf.ocr", "true",
			"--pdf.should-parse", "true",
			"--pdf.start", "1",
			"--settle-animations", "'true'",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--use-main-content-only", "'true'",
			"--wait-for-ms", "0",
			"--zdr", "enabled",
		)
	})
}

func TestWebWebScrapeImages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-scrape-images",
			"--url", "https://example.com",
			"--dedupe", "'true'",
			"--enrichment", "{classification: 'true', hostedUrl: 'true', maxTimePerMs: 1, resolution: 'true'}",
			"--headers", "{foo: J!}",
			"--max-age-ms", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--wait-for-ms", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webWebScrapeImages)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-scrape-images",
			"--url", "https://example.com",
			"--dedupe", "'true'",
			"--enrichment.classification", "true",
			"--enrichment.hosted-url", "true",
			"--enrichment.max-time-per-ms", "1",
			"--enrichment.resolution", "true",
			"--headers", "{foo: J!}",
			"--max-age-ms", "0",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--wait-for-ms", "0",
		)
	})
}

func TestWebWebScrapeMd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-scrape-md",
			"--url", "https://example.com",
			"--country", "de",
			"--exclude-selector", "[x]",
			"--headers", "{foo: J!}",
			"--include-frames", "'true'",
			"--include-images", "'true'",
			"--include-links", "'true'",
			"--include-selector", "[x]",
			"--max-age-ms", "0",
			"--pdf", "{end: 1, ocr: 'true', shouldParse: 'true', start: 1}",
			"--settle-animations", "'true'",
			"--shorten-base64-images", "'true'",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--use-main-content-only", "'true'",
			"--wait-for-ms", "0",
			"--zdr", "enabled",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webWebScrapeMd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-scrape-md",
			"--url", "https://example.com",
			"--country", "de",
			"--exclude-selector", "[x]",
			"--headers", "{foo: J!}",
			"--include-frames", "'true'",
			"--include-images", "'true'",
			"--include-links", "'true'",
			"--include-selector", "[x]",
			"--max-age-ms", "0",
			"--pdf.end", "1",
			"--pdf.ocr", "true",
			"--pdf.should-parse", "true",
			"--pdf.start", "1",
			"--settle-animations", "'true'",
			"--shorten-base64-images", "'true'",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--use-main-content-only", "'true'",
			"--wait-for-ms", "0",
			"--zdr", "enabled",
		)
	})
}

func TestWebWebScrapeSitemap(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web", "web-scrape-sitemap",
			"--domain", "xxx",
			"--headers", "{foo: J!}",
			"--max-links", "1",
			"--sitemap-url", "https://example.com",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1",
			"--url-regex", "^https?://[^/]+/blog/",
			"--zdr", "enabled",
		)
	})
}
