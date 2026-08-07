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
	Usage:   "Finds and normalizes the best available person candidate from additive identity\nclues, then assigns an identity match score from 0 to 100. Available on Pro and\nScale plans. Successful requests cost 20 credits. Disposable and free email\naddresses (like gmail.com, yahoo.com) will throw a 422 error.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "company",
			BodyPath: "company",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "education",
			BodyPath: "education",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "location",
			BodyPath: "location",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "social-url",
			BodyPath: "social_urls",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.",
			BodyPath: "tags",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Usage:    "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			BodyPath: "timeoutMS",
		},
	},
	Action:          handlePeopleEnrich,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"company": {
		&requestflag.InnerFlag[string]{
			Name:       "company.domain",
			InnerField: "domain",
		},
		&requestflag.InnerFlag[string]{
			Name:       "company.name",
			InnerField: "name",
		},
	},
	"education": {
		&requestflag.InnerFlag[string]{
			Name:       "education.degree",
			InnerField: "degree",
		},
		&requestflag.InnerFlag[string]{
			Name:       "education.field-of-study",
			InnerField: "field_of_study",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "education.graduation-year",
			InnerField: "graduation_year",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "education.institution",
			InnerField: "institution",
		},
	},
	"location": {
		&requestflag.InnerFlag[string]{
			Name:       "location.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "location.country",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "location.region",
			InnerField: "region",
		},
	},
	"name": {
		&requestflag.InnerFlag[string]{
			Name:       "name.first",
			InnerField: "first",
		},
		&requestflag.InnerFlag[string]{
			Name:       "name.last",
			InnerField: "last",
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
