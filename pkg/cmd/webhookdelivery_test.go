// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
)

func TestWebhooksDeliveriesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:deliveries", "retrieve",
			"--delivery-id", "whd_210b9798eb53baa4e69d31c1071cf03d",
			"--tag", "production",
			"--tag", "team-alpha",
		)
	})
}

func TestWebhooksDeliveriesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:deliveries", "list",
			"--batch-id", "batch_id",
			"--cursor", "whd_210b9798eb53baa4e69d31c1071cf03d",
			"--limit", "1",
			"--monitor-id", "monitor_id",
			"--run-id", "run_id",
			"--status", "pending",
			"--tag", "production",
			"--tag", "team-alpha",
		)
	})
}

func TestWebhooksDeliveriesListAttempts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:deliveries", "list-attempts",
			"--delivery-id", "whd_210b9798eb53baa4e69d31c1071cf03d",
			"--cursor", "wha_210b9798eb53baa4e69d31c1071cf03d",
			"--limit", "1",
			"--tag", "production",
			"--tag", "team-alpha",
		)
	})
}

func TestWebhooksDeliveriesRetry(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks:deliveries", "retry",
			"--delivery-id", "whd_210b9798eb53baa4e69d31c1071cf03d",
			"--force=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"force: true\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks:deliveries", "retry",
			"--delivery-id", "whd_210b9798eb53baa4e69d31c1071cf03d",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}
