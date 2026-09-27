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

var industryRetrieveNaics = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-naics",
	Usage:   "Classify a company into NAICS industry codes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "input",
			Usage:     "Brand domain or title to retrieve NAICS code for. If a valid domain is provided, it will be used for classification, otherwise, we will search for the brand using the provided title.",
			Required:  true,
			QueryPath: "input",
		},
		&requestflag.Flag[int64]{
			Name:      "max-results",
			Usage:     "Maximum number of NAICS codes to return. Must be between 1 and 10. Defaults to 5.",
			Default:   5,
			QueryPath: "maxResults",
		},
		&requestflag.Flag[int64]{
			Name:      "min-results",
			Usage:     "Minimum number of NAICS codes to return. Must be at least 1. Defaults to 1.",
			Default:   1,
			QueryPath: "minResults",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for filtering usage, e.g. `production,team-alpha`.",
			QueryPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "timeout-opts",
			Usage:     "Request deadline and what to return when it passes.",
			QueryPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:      "zdr",
			Usage:     "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			Default:   "disabled",
			QueryPath: "zdr",
		},
	},
	Action:          handleIndustryRetrieveNaics,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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

var industryRetrieveSic = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-sic",
	Usage:   "Classify a company into SIC industry codes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "input",
			Usage:     "Brand domain or title to retrieve SIC code for. If a valid domain is provided, it will be used for classification, otherwise, we will search for the brand using the provided title.",
			Required:  true,
			QueryPath: "input",
		},
		&requestflag.Flag[int64]{
			Name:      "max-results",
			Usage:     "Maximum number of SIC codes to return. Must be between 1 and 10. Defaults to 5.",
			Default:   5,
			QueryPath: "maxResults",
		},
		&requestflag.Flag[int64]{
			Name:      "min-results",
			Usage:     "Minimum number of SIC codes to return. Must be at least 1. Defaults to 1.",
			Default:   1,
			QueryPath: "minResults",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for filtering usage, e.g. `production,team-alpha`.",
			QueryPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "timeout-opts",
			Usage:     "Request deadline and what to return when it passes.",
			QueryPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "SIC dataset: `original_sic` (1987) or `latest_sec` (current SEC list).",
			Default:   "original_sic",
			QueryPath: "type",
		},
		&requestflag.Flag[string]{
			Name:      "zdr",
			Usage:     "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			Default:   "disabled",
			QueryPath: "zdr",
		},
	},
	Action:          handleIndustryRetrieveSic,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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

func handleIndustryRetrieveNaics(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := contextdev.IndustryGetNaicsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Industry.GetNaics(ctx, params, options...)
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
		Title:          "industry retrieve-naics",
		Transform:      transform,
	})
}

func handleIndustryRetrieveSic(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := contextdev.IndustryGetSicParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Industry.GetSic(ctx, params, options...)
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
		Title:          "industry retrieve-sic",
		Transform:      transform,
	})
}
