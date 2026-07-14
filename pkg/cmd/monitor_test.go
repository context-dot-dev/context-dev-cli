// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestMonitorsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "create",
			"--change-detection", "{type: exact}",
			"--name", "Acme pricing page",
			"--schedule", "{frequency: 6, type: interval, unit: hours}",
			"--target", "{type: page, url: https://acme.com/pricing, normalize_whitespace: true}",
			"--mode", "web",
			"--tag", "pricing",
			"--tag", "competitor",
			"--webhook", "{url: https://example.com/webhook, events: [change.detected, run.completed]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(monitorsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "create",
			"--change-detection", "{type: exact}",
			"--name", "Acme pricing page",
			"--schedule.frequency", "6",
			"--schedule.type", "interval",
			"--schedule.unit", "hours",
			"--target", "{type: page, url: https://acme.com/pricing, normalize_whitespace: true}",
			"--mode", "web",
			"--tag", "pricing",
			"--tag", "competitor",
			"--webhook.url", "https://example.com/webhook",
			"--webhook.events", "[change.detected, run.completed]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"change_detection:\n" +
			"  type: exact\n" +
			"name: Acme pricing page\n" +
			"schedule:\n" +
			"  frequency: 6\n" +
			"  type: interval\n" +
			"  unit: hours\n" +
			"target:\n" +
			"  type: page\n" +
			"  url: https://acme.com/pricing\n" +
			"  normalize_whitespace: true\n" +
			"mode: web\n" +
			"tags:\n" +
			"  - pricing\n" +
			"  - competitor\n" +
			"webhook:\n" +
			"  url: https://example.com/webhook\n" +
			"  events:\n" +
			"    - change.detected\n" +
			"    - run.completed\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"monitors", "create",
		)
	})
}

func TestMonitorsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "retrieve",
			"--monitor-id", "mon_123",
		)
	})
}

func TestMonitorsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "update",
			"--monitor-id", "mon_123",
			"--change-detection", "{type: exact}",
			"--name", "Acme pricing monitor",
			"--schedule", "{frequency: 1, type: interval, unit: hours}",
			"--status", "active",
			"--tag", "pricing",
			"--tag", "competitor",
			"--target", "{type: page, url: https://acme.com/pricing, normalize_whitespace: true}",
			"--webhook", "{url: https://example.com/webhook, events: [change.detected, run.completed]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(monitorsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "update",
			"--monitor-id", "mon_123",
			"--change-detection", "{type: exact}",
			"--name", "Acme pricing monitor",
			"--schedule.frequency", "1",
			"--schedule.type", "interval",
			"--schedule.unit", "hours",
			"--status", "active",
			"--tag", "pricing",
			"--tag", "competitor",
			"--target", "{type: page, url: https://acme.com/pricing, normalize_whitespace: true}",
			"--webhook.url", "https://example.com/webhook",
			"--webhook.events", "[change.detected, run.completed]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"change_detection:\n" +
			"  type: exact\n" +
			"name: Acme pricing monitor\n" +
			"schedule:\n" +
			"  frequency: 1\n" +
			"  type: interval\n" +
			"  unit: hours\n" +
			"status: active\n" +
			"tags:\n" +
			"  - pricing\n" +
			"  - competitor\n" +
			"target:\n" +
			"  type: page\n" +
			"  url: https://acme.com/pricing\n" +
			"  normalize_whitespace: true\n" +
			"webhook:\n" +
			"  url: https://example.com/webhook\n" +
			"  events:\n" +
			"    - change.detected\n" +
			"    - run.completed\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"monitors", "update",
			"--monitor-id", "mon_123",
		)
	})
}

func TestMonitorsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "list",
			"--change-detection-type", "exact",
			"--cursor", "cursor",
			"--limit", "1",
			"--q", "pricing",
			"--search-by", "[name]",
			"--search-type", "exact",
			"--status", "active",
			"--tag", "pricing",
			"--tag", "[x]",
			"--target-type", "page",
		)
	})
}

func TestMonitorsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "delete",
			"--monitor-id", "mon_123",
		)
	})
}

func TestMonitorsListAccountChanges(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "list-account-changes",
			"--change-detection-type", "exact",
			"--cursor", "cursor",
			"--limit", "1",
			"--monitor-id", "mon_123",
			"--since", "'2026-06-01T00:00:00Z'",
			"--tag", "pricing",
			"--target-type", "page",
			"--until", "'2026-06-28T00:00:00Z'",
		)
	})
}

func TestMonitorsListAccountRuns(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "list-account-runs",
			"--cursor", "cursor",
			"--limit", "1",
			"--status", "queued",
		)
	})
}

func TestMonitorsListChanges(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "list-changes",
			"--monitor-id", "mon_123",
			"--cursor", "cursor",
			"--limit", "1",
			"--since", "'2026-06-01T00:00:00Z'",
			"--tag", "pricing",
			"--until", "'2026-06-28T00:00:00Z'",
		)
	})
}

func TestMonitorsListRuns(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "list-runs",
			"--monitor-id", "mon_123",
			"--cursor", "cursor",
			"--limit", "1",
			"--status", "queued",
		)
	})
}

func TestMonitorsRetrieveChange(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "retrieve-change",
			"--change-id", "chg_123",
		)
	})
}

func TestMonitorsRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"monitors", "run",
			"--monitor-id", "mon_123",
		)
	})
}
