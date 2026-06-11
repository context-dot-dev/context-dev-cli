// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/context.dev-cli/internal/mocktest"
	"github.com/stainless-sdks/context.dev-cli/internal/requestflag"
)

func TestAIAIQuery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "ai-query",
			"--data-to-extract", "{datapoint_description: datapoint_description, datapoint_example: datapoint_example, datapoint_name: datapoint_name, datapoint_type: text, datapoint_list_type: string, datapoint_object_schema: {testimonial_text: string, testimonial_author: string}}",
			"--domain", "domain",
			"--specific-pages", "{about_us: true, blog: true, careers: true, contact_us: true, faq: true, home_page: true, pricing: true, privacy_policy: true, terms_and_conditions: true}",
			"--timeout-ms", "1000",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiAIQuery)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "ai-query",
			"--data-to-extract.datapoint-description", "datapoint_description",
			"--data-to-extract.datapoint-example", "datapoint_example",
			"--data-to-extract.datapoint-name", "datapoint_name",
			"--data-to-extract.datapoint-type", "text",
			"--data-to-extract.datapoint-list-type", "string",
			"--data-to-extract.datapoint-object-schema", "{testimonial_text: string, testimonial_author: string}",
			"--domain", "domain",
			"--specific-pages.about-us=true",
			"--specific-pages.blog=true",
			"--specific-pages.careers=true",
			"--specific-pages.contact-us=true",
			"--specific-pages.faq=true",
			"--specific-pages.home-page=true",
			"--specific-pages.pricing=true",
			"--specific-pages.privacy-policy=true",
			"--specific-pages.terms-and-conditions=true",
			"--timeout-ms", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data_to_extract:\n" +
			"  - datapoint_description: datapoint_description\n" +
			"    datapoint_example: datapoint_example\n" +
			"    datapoint_name: datapoint_name\n" +
			"    datapoint_type: text\n" +
			"    datapoint_list_type: string\n" +
			"    datapoint_object_schema:\n" +
			"      testimonial_text: string\n" +
			"      testimonial_author: string\n" +
			"domain: domain\n" +
			"specific_pages:\n" +
			"  about_us: true\n" +
			"  blog: true\n" +
			"  careers: true\n" +
			"  contact_us: true\n" +
			"  faq: true\n" +
			"  home_page: true\n" +
			"  pricing: true\n" +
			"  privacy_policy: true\n" +
			"  terms_and_conditions: true\n" +
			"timeoutMS: 1000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai", "ai-query",
		)
	})
}

func TestAIExtractProduct(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "extract-product",
			"--url", "https://example.com",
			"--max-age-ms", "0",
			"--timeout-ms", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com\n" +
			"maxAgeMs: 0\n" +
			"timeoutMS: 1000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai", "extract-product",
		)
	})
}

func TestAIExtractProducts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "extract-products",
			"--domain", "domain",
			"--max-age-ms", "0",
			"--max-products", "1",
			"--timeout-ms", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domain: domain\n" +
			"maxAgeMs: 0\n" +
			"maxProducts: 1\n" +
			"timeoutMS: 1000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai", "extract-products",
		)
	})
}
