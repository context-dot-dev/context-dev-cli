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

var industryRetrieveNaics = cli.Command{
	Name:    "retrieve-naics",
	Usage:   "Classify any brand into 2022 NAICS industry codes from its domain or name.",
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
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleIndustryRetrieveNaics,
	HideHelpCommand: true,
}

var industryRetrieveSic = cli.Command{
	Name:    "retrieve-sic",
	Usage:   "Classify any brand into Standard Industrial Classification (SIC) codes from its\ndomain or name. Choose between the original SIC system (`original_sic`) or the\nlatest SIC list maintained by the SEC (`latest_sec`).",
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
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Which SIC dataset to classify against. `original_sic` uses the 1987 Standard Industrial Classification system; `latest_sec` uses the current SIC list as published by the SEC. Defaults to `original_sic`.",
			Default:   "original_sic",
			QueryPath: "type",
		},
	},
	Action:          handleIndustryRetrieveSic,
	HideHelpCommand: true,
}

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
