// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/context-dot-dev/context-dev-cli/internal/mocktest"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
)

func TestPeopleEnrich(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"people", "enrich",
			"--company", "{domain: analyticalengines.example, name: Analytical Engines}",
			"--education", "{degree: x, field_of_study: x, graduation_year: 1900, institution: {domain: x, name: x}}",
			"--email", "dev@stainless.com",
			"--location", "{city: x, country: x, region: x}",
			"--name", "{first: Ada, last: Lovelace}",
			"--social-url", "https://www.linkedin.com/in/ada-lovelace/",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(peopleEnrich)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"people", "enrich",
			"--company.domain", "analyticalengines.example",
			"--company.name", "Analytical Engines",
			"--education.degree", "x",
			"--education.field-of-study", "x",
			"--education.graduation-year", "1900",
			"--education.institution", "{domain: x, name: x}",
			"--email", "dev@stainless.com",
			"--location.city", "x",
			"--location.country", "x",
			"--location.region", "x",
			"--name.first", "Ada",
			"--name.last", "Lovelace",
			"--social-url", "https://www.linkedin.com/in/ada-lovelace/",
			"--tag", "production",
			"--tag", "team-alpha",
			"--timeout-ms", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"company:\n" +
			"  domain: analyticalengines.example\n" +
			"  name: Analytical Engines\n" +
			"education:\n" +
			"  - degree: x\n" +
			"    field_of_study: x\n" +
			"    graduation_year: 1900\n" +
			"    institution:\n" +
			"      domain: x\n" +
			"      name: x\n" +
			"email: dev@stainless.com\n" +
			"location:\n" +
			"  city: x\n" +
			"  country: x\n" +
			"  region: x\n" +
			"name:\n" +
			"  first: Ada\n" +
			"  last: Lovelace\n" +
			"social_urls:\n" +
			"  - https://www.linkedin.com/in/ada-lovelace/\n" +
			"tags:\n" +
			"  - production\n" +
			"  - team-alpha\n" +
			"timeoutMS: 1000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"people", "enrich",
		)
	})
}
