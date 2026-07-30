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
			"--tag", "production",
			"--tag", "team-alpha",
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
			"--status", "queued",
			"--tag", "production",
			"--tag", "team-alpha",
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
			"--tag", "production",
			"--tag", "team-alpha",
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
			"--tag", "production",
			"--tag", "team-alpha",
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
			"--identifiers", "{linkedinUrl: https://www.linkedin.com/in/yahia-bakour/}",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
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
			"--identifiers.linkedin-url", "https://www.linkedin.com/in/yahia-bakour/",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"identifiers:\n" +
			"  linkedinUrl: https://www.linkedin.com/in/yahia-bakour/\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutMS: 1000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"batch", "submit",
		)
	})
}
