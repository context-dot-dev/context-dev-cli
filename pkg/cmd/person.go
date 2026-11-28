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
	Usage:   "Match a person from a work email, a profile URL, or a name with company,\neducation, or location context, and get their profile with an identity-match\nscore from 0 to 100. The best candidate is returned even when the match is weak,\nso set a score threshold for your workflow. `not_found` results are free.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "company",
			Usage:    "Company associated with the person. Send a `name`, a `domain`, or both.",
			BodyPath: "company",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "education",
			Usage:    "Education history to help distinguish people with similar names.",
			BodyPath: "education",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Work email address of the person. Free and disposable email providers return 422.",
			BodyPath: "email",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "location",
			Usage:    "Location associated with the person. Send at least one of `city`, `region`, or `country`.",
			BodyPath: "location",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "name",
			Usage:    "Name of the person. Without an email or profile URL, send both `first` and `last` plus `company`, `education`, or `location`.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "social-url",
			Usage:    "Public profile URLs for the person. A person-profile URL is enough to identify them without a name.",
			BodyPath: "social_urls",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for this request. Filter usage and request logs by them in the dashboard or with `GET /org/logs`.",
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
			Usage:      "`fail` returns 408 at the deadline. `return-partial` returns the results collected so far and marks the response as partial.",
			InnerField: "behavior",
		},
	},
})

var peopleSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Find people from a plain-language description, such as\n`VPs of sales at solar companies in California`. Each result includes the\ncurrent role and employer, location, years of experience, skills, schools, and\npublic profile URLs. Pages with no matches are free.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Plain-language description of the people to find. Up to 500 characters in `fast` mode and 5000 in `ultra` mode.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "cursor",
			Usage:    "Cursor from the previous page's `next_cursor`. Omit for the first page. Send the same `query`, `mode`, and `filters` with it.",
			BodyPath: "cursor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filters",
			Usage:    "Constraints every result must match, in addition to `query`.",
			BodyPath: "filters",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "`fast` matches people on basic criteria such as role, location, or background; `ultra` runs a deeper search, checks every candidate against each criterion in your query, and adds a `match` score. Complex `ultra` searches can take up to 30 seconds, so set your client timeout to match.",
			Default:  "fast",
			BodyPath: "mode",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for this request. Filter usage and request logs by them in the dashboard or with `GET /org/logs`.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "zdr",
			Usage:    "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			BodyPath: "zdr",
		},
	},
	Action:          handlePeopleSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filters": {
		&requestflag.InnerFlag[[]string]{
			Name:       "filters.countries",
			Usage:      "Only return people who live in these countries, as ISO 3166-1 alpha-2 codes such as `US`.",
			InnerField: "countries",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filters.years-of-experience-max",
			Usage:      "Maximum years of professional experience.",
			InnerField: "years_of_experience_max",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filters.years-of-experience-min",
			Usage:      "Minimum years of professional experience.",
			InnerField: "years_of_experience_min",
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

func handlePeopleSearch(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.PersonSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.People.Search(ctx, params, options...)
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
		Title:          "people search",
		Transform:      transform,
	})
}
