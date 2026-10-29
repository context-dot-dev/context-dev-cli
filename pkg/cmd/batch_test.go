// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestBatchRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batch", "retrieve",
			"--batch-id", "batch_9f2c8a",
		)
	})
}

func TestBatchList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batch", "list",
			"--cursor", "cursor",
			"--limit", "1",
			"--q", "batch_1a2b",
			"--search-type", "exact",
			"--should-include-unified=true",
			"--status", "queued",
			"--tags", "docs,competitor",
		)
	})
}

func TestBatchDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batch", "delete",
			"--batch-id", "batch_9f2c8a",
		)
	})
}

func TestBatchCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batch", "cancel",
			"--batch-id", "batch_9f2c8a",
		)
	})
}

func TestBatchGetResults(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batch", "get-results",
			"--batch-id", "batch_9f2c8a",
			"--cursor", "321669910225:155771193",
			"--limit", "1",
		)
	})
}

func TestBatchSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batch", "submit",
			"--input", "{data: {formats: {bytes: true, highlights: true, html: true, images: true, json: true, markdown: true, parse: true, product: true, screenshot: true}, urls: [{url: https://example.com/products/anvil, itemId: sku-1, meta: {category: bar}}, {url: https://example.com/products/hammer, itemId: sku-2, meta: {foo: bar}}], highlightsParams: {query: x, maxCharacters: 100}, imageParams: {dedupe: none, enrich: [dimensions]}, jsonParams: {schema: {type: bar, properties: bar, required: bar, additionalProperties: bar}, instructions: instructions}, markdownParams: {includeImages: true, includeLinks: true, inlineImages: placeholder, sourcePreference: [html-conversion]}, maxAgeMs: 0, parseParams: {rules: {title: h1, links: {selector: a, output: '@href', type: list}}, type: custom}, productParams: {dedupeImages: true, useAIFallback: true}, screenshotParams: {area: viewport, format: png}, sharedParams: {actions: [{action: Click the product details tab, type: perform}], country: US, dismissCookies: true, dismissPopups: true, excludeSelectors: [P], headers: {Accept-Language: en-US}, includeFrames: true, includeSelectors: [P], mainContentOnly: true, parsers: {pdf: {endPage: 1, includeNumberedPages: true, ocr: 'off', startPage: 1}}, profile: {profileId: prof_210b9798eb53baa4e69d31c1071cf03d, type: existing}, settleAnimations: true, theme: light, viewport: {height: 240, width: 240}, waitFor: 500}, tags: [production, team-alpha], timeoutOpts: {milliseconds: 1, behavior: fail}, zdr: disabled}, mode: scrape}",
			"--policy", "pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e",
			"--tag", "docs",
			"--tag", "competitor",
			"--webhook", "{url: https://example.com, retry: {delays_seconds: [10, 60, 300, 1800, 7200, 21600, 57600]}}",
			"--webhook-url", "https://example.com",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(batchSubmit)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batch", "submit",
			"--input", "{data: {formats: {bytes: true, highlights: true, html: true, images: true, json: true, markdown: true, parse: true, product: true, screenshot: true}, urls: [{url: https://example.com/products/anvil, itemId: sku-1, meta: {category: bar}}, {url: https://example.com/products/hammer, itemId: sku-2, meta: {foo: bar}}], highlightsParams: {query: x, maxCharacters: 100}, imageParams: {dedupe: none, enrich: [dimensions]}, jsonParams: {schema: {type: bar, properties: bar, required: bar, additionalProperties: bar}, instructions: instructions}, markdownParams: {includeImages: true, includeLinks: true, inlineImages: placeholder, sourcePreference: [html-conversion]}, maxAgeMs: 0, parseParams: {rules: {title: h1, links: {selector: a, output: '@href', type: list}}, type: custom}, productParams: {dedupeImages: true, useAIFallback: true}, screenshotParams: {area: viewport, format: png}, sharedParams: {actions: [{action: Click the product details tab, type: perform}], country: US, dismissCookies: true, dismissPopups: true, excludeSelectors: [P], headers: {Accept-Language: en-US}, includeFrames: true, includeSelectors: [P], mainContentOnly: true, parsers: {pdf: {endPage: 1, includeNumberedPages: true, ocr: 'off', startPage: 1}}, profile: {profileId: prof_210b9798eb53baa4e69d31c1071cf03d, type: existing}, settleAnimations: true, theme: light, viewport: {height: 240, width: 240}, waitFor: 500}, tags: [production, team-alpha], timeoutOpts: {milliseconds: 1, behavior: fail}, zdr: disabled}, mode: scrape}",
			"--policy", "pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e",
			"--tag", "docs",
			"--tag", "competitor",
			"--webhook.url", "https://example.com",
			"--webhook.retry", "{delays_seconds: [10, 60, 300, 1800, 7200, 21600, 57600]}",
			"--webhook-url", "https://example.com",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input:\n" +
			"  data:\n" +
			"    formats:\n" +
			"      bytes: true\n" +
			"      highlights: true\n" +
			"      html: true\n" +
			"      images: true\n" +
			"      json: true\n" +
			"      markdown: true\n" +
			"      parse: true\n" +
			"      product: true\n" +
			"      screenshot: true\n" +
			"    urls:\n" +
			"      - url: https://example.com/products/anvil\n" +
			"        itemId: sku-1\n" +
			"        meta:\n" +
			"          category: bar\n" +
			"      - url: https://example.com/products/hammer\n" +
			"        itemId: sku-2\n" +
			"        meta:\n" +
			"          foo: bar\n" +
			"    highlightsParams:\n" +
			"      query: x\n" +
			"      maxCharacters: 100\n" +
			"    imageParams:\n" +
			"      dedupe: none\n" +
			"      enrich:\n" +
			"        - dimensions\n" +
			"    jsonParams:\n" +
			"      schema:\n" +
			"        type: bar\n" +
			"        properties: bar\n" +
			"        required: bar\n" +
			"        additionalProperties: bar\n" +
			"      instructions: instructions\n" +
			"    markdownParams:\n" +
			"      includeImages: true\n" +
			"      includeLinks: true\n" +
			"      inlineImages: placeholder\n" +
			"      sourcePreference:\n" +
			"        - html-conversion\n" +
			"    maxAgeMs: 0\n" +
			"    parseParams:\n" +
			"      rules:\n" +
			"        title: h1\n" +
			"        links:\n" +
			"          selector: a\n" +
			"          output: '@href'\n" +
			"          type: list\n" +
			"      type: custom\n" +
			"    productParams:\n" +
			"      dedupeImages: true\n" +
			"      useAIFallback: true\n" +
			"    screenshotParams:\n" +
			"      area: viewport\n" +
			"      format: png\n" +
			"    sharedParams:\n" +
			"      actions:\n" +
			"        - action: Click the product details tab\n" +
			"          type: perform\n" +
			"      country: US\n" +
			"      dismissCookies: true\n" +
			"      dismissPopups: true\n" +
			"      excludeSelectors:\n" +
			"        - P\n" +
			"      headers:\n" +
			"        Accept-Language: en-US\n" +
			"      includeFrames: true\n" +
			"      includeSelectors:\n" +
			"        - P\n" +
			"      mainContentOnly: true\n" +
			"      parsers:\n" +
			"        pdf:\n" +
			"          endPage: 1\n" +
			"          includeNumberedPages: true\n" +
			"          ocr: 'off'\n" +
			"          startPage: 1\n" +
			"      profile:\n" +
			"        profileId: prof_210b9798eb53baa4e69d31c1071cf03d\n" +
			"        type: existing\n" +
			"      settleAnimations: true\n" +
			"      theme: light\n" +
			"      viewport:\n" +
			"        height: 240\n" +
			"        width: 240\n" +
			"      waitFor: 500\n" +
			"    tags:\n" +
			"      - production\n" +
			"      - team-alpha\n" +
			"    timeoutOpts:\n" +
			"      milliseconds: 1\n" +
			"      behavior: fail\n" +
			"    zdr: disabled\n" +
			"  mode: scrape\n" +
			"policies:\n" +
			"  - pol_3f9c2a7e5b8d4c1f9a6e0b2d7c4f8a1e\n" +
			"tags:\n" +
			"  - docs\n" +
			"  - competitor\n" +
			"webhook:\n" +
			"  url: https://example.com\n" +
			"  retry:\n" +
			"    delays_seconds:\n" +
			"      - 10\n" +
			"      - 60\n" +
			"      - 300\n" +
			"      - 1800\n" +
			"      - 7200\n" +
			"      - 21600\n" +
			"      - 57600\n" +
			"webhookUrl: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"batch", "submit",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}
