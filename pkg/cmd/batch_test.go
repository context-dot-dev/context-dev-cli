// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
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
			"--cursor", "cursor",
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
			"--input", "{data: {format: markdown, urls: [{url: https://example.com/products/anvil, itemId: sku-1, meta: {category: bar}}, {url: https://example.com/products/hammer, itemId: sku-2, meta: {foo: bar}}], options: {country: de, excludeSelectors: [x], includeHTML: true, includeImages: true, includeLinks: true, includeSelectors: [x], maxAgeMs: 0, pdf: {end: 1, ocr: true, shouldParse: true, start: 1}, settleAnimations: true, shortenBase64Images: true, useMainContentOnly: true, waitForMs: 0}}, mode: scrape}",
			"--tag", "docs",
			"--tag", "competitor",
			"--webhook-url", "webhookUrl",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input:\n" +
			"  data:\n" +
			"    format: markdown\n" +
			"    urls:\n" +
			"      - url: https://example.com/products/anvil\n" +
			"        itemId: sku-1\n" +
			"        meta:\n" +
			"          category: bar\n" +
			"      - url: https://example.com/products/hammer\n" +
			"        itemId: sku-2\n" +
			"        meta:\n" +
			"          foo: bar\n" +
			"    options:\n" +
			"      country: de\n" +
			"      excludeSelectors:\n" +
			"        - x\n" +
			"      includeHTML: true\n" +
			"      includeImages: true\n" +
			"      includeLinks: true\n" +
			"      includeSelectors:\n" +
			"        - x\n" +
			"      maxAgeMs: 0\n" +
			"      pdf:\n" +
			"        end: 1\n" +
			"        ocr: true\n" +
			"        shouldParse: true\n" +
			"        start: 1\n" +
			"      settleAnimations: true\n" +
			"      shortenBase64Images: true\n" +
			"      useMainContentOnly: true\n" +
			"      waitForMs: 0\n" +
			"  mode: scrape\n" +
			"tags:\n" +
			"  - docs\n" +
			"  - competitor\n" +
			"webhookUrl: webhookUrl\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"batch", "submit",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}
