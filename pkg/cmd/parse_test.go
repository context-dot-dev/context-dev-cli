// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestParseHandle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"parse", "handle",
			"--body", mocktest.TestFile(t, "Example data"),
			"--extension", "txt",
			"--include-images=true",
			"--include-links=true",
			"--ocr=true",
			"--pdf", "{end: 1, start: 1}",
			"--shorten-base64-images=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--use-main-content-only=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(parseHandle)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"parse", "handle",
			"--body", mocktest.TestFile(t, "Example data"),
			"--extension", "txt",
			"--include-images=true",
			"--include-links=true",
			"--ocr=true",
			"--pdf.end", "1",
			"--pdf.start", "1",
			"--shorten-base64-images=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--use-main-content-only=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("Example data")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"parse", "handle",
			"--extension", "txt",
			"--include-images=true",
			"--include-links=true",
			"--ocr=true",
			"--pdf", "{end: 1, start: 1}",
			"--shorten-base64-images=true",
			"--tag", "production",
			"--tag", "team-alpha",
			"--use-main-content-only=true",
		)
	})
}
