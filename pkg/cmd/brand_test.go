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
			"--force-language", "afrikaans",
			"--max-age-ms", "86400000",
			"--max-speed=true",
			"--timeout-ms", "1000",
		)
	})
}

func TestBrandIdentifyFromTransaction(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "identify-from-transaction",
			"--transaction-info", "transaction_info",
			"--city", "city",
			"--country-gl", "ad",
			"--force-language", "afrikaans",
			"--high-confidence-only=true",
			"--max-speed=true",
			"--mcc", "mcc",
			"--phone", "0",
			"--timeout-ms", "1000",
		)
	})
}

func TestBrandRetrieveByEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve-by-email",
			"--email", "dev@stainless.com",
			"--force-language", "afrikaans",
			"--max-age-ms", "86400000",
			"--max-speed=true",
			"--timeout-ms", "1000",
		)
	})
}

func TestBrandRetrieveByIsin(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve-by-isin",
			"--isin", "SE60513A9993",
			"--force-language", "afrikaans",
			"--max-age-ms", "86400000",
			"--max-speed=true",
			"--timeout-ms", "1000",
		)
	})
}

func TestBrandRetrieveByName(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve-by-name",
			"--name", "xxx",
			"--country-gl", "ad",
			"--force-language", "afrikaans",
			"--max-age-ms", "86400000",
			"--max-speed=true",
			"--timeout-ms", "1000",
		)
	})
}

func TestBrandRetrieveByTicker(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brand", "retrieve-by-ticker",
			"--ticker", "ticker",
			"--force-language", "afrikaans",
			"--max-age-ms", "86400000",
			"--max-speed=true",
			"--ticker-exchange", "AMEX",
			"--timeout-ms", "1000",
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
			"--timeout-ms", "1000",
		)
	})
}
