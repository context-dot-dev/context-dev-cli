// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/context-dot-dev/context-dev-cli/internal/apiquery"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
	"github.com/context-dot-dev/context-go-sdk/v2"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var peopleEnrich = requestflag.WithInnerFlags(cli.Command{
	Name:    "enrich",
	Usage:   "Find a person from identity clues and return their profile with a match score.\nRequires a paid plan; free or disposable email addresses return 422.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "company",
			Usage:    "Company context to help identify the person. Provide a name or domain.",
			BodyPath: "company",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "education",
			Usage:    "Education history to help distinguish people with similar names.",
			BodyPath: "education",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Email address of the person to find.",
			BodyPath: "email",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "location",
			Usage:    "Location context to help identify the person. Provide a city, region, or country.",
			BodyPath: "location",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "name",
			Usage:    "Person name. Without an email or person-profile URL, provide both first and last name plus company, education, or location.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "social-url",
			Usage:    "Public profile URLs for the person. A person-profile URL can identify the person without a name.",
			BodyPath: "social_urls",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for filtering usage in the dashboard.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Request deadline and what to return when it passes.",
			BodyPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:     "zdr",
			Usage:    "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			BodyPath: "zdr",
		},
	},
	Action:          handlePeopleEnrich,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"company": {
		&requestflag.InnerFlag[string]{
			Name:       "company.domain",
			Usage:      "Website domain of a company associated with the person.",
			InnerField: "domain",
		},
		&requestflag.InnerFlag[string]{
			Name:       "company.name",
			Usage:      "Name of a company associated with the person.",
			InnerField: "name",
		},
	},
	"education": {
		&requestflag.InnerFlag[string]{
			Name:       "education.degree",
			Usage:      "Degree or qualification earned.",
			InnerField: "degree",
		},
		&requestflag.InnerFlag[string]{
			Name:       "education.field-of-study",
			Usage:      "Subject or major studied.",
			InnerField: "field_of_study",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "education.graduation-year",
			Usage:      "Four-digit graduation year.",
			InnerField: "graduation_year",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "education.institution",
			Usage:      "School or university, identified by name or domain.",
			InnerField: "institution",
		},
	},
	"location": {
		&requestflag.InnerFlag[string]{
			Name:       "location.city",
			Usage:      "City associated with the person.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "location.country",
			Usage:      "Country associated with the person.",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "location.region",
			Usage:      "State, province, or region associated with the person.",
			InnerField: "region",
		},
	},
	"name": {
		&requestflag.InnerFlag[string]{
			Name:       "name.first",
			Usage:      "First or given name.",
			InnerField: "first",
		},
		&requestflag.InnerFlag[string]{
			Name:       "name.last",
			Usage:      "Last or family name.",
			InnerField: "last",
		},
	},
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      `"fail" returns 408 at the deadline. "return-partial" returns available results; inspect the response’s partial flag.`,
			InnerField: "behavior",
		},
	},
})

func handlePeopleEnrich(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := contextdev.PersonEnrichParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.People.Enrich(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "people enrich",
		Transform:      transform,
	})
}
