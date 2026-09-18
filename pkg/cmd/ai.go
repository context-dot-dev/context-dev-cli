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

var aiExtractProduct = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract-product",
	Usage:   "Given a single URL, determines if it is a product page and extracts the product\ninformation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The product page URL to extract product data from.",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[int64]{
			Name:     "max-age-ms",
			Usage:    "Return a cached result if a prior scrape for the same parameters exists and is younger than this many milliseconds. Defaults to 7 days (604800000 ms) when omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.",
			Default:  604800000,
			BodyPath: "maxAgeMs",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Optional request deadline and behavior on timeout. For GET requests, use timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded timeoutOpts object.",
			BodyPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:     "zdr",
			Usage:    "Set to enabled to bypass shared caches and omit request and response content from retained usage logs. Asset uploads are skipped, so hosted image URLs are omitted. Requires zero data retention to be enabled for your organization (contact support@context.dev), otherwise the request fails with ZDR_NOT_ENABLED. Successful ZDR responses include X-Context-ZDR: true.",
			Default:  "disabled",
			BodyPath: "zdr",
		},
	},
	Action:          handleAIExtractProduct,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Request deadline in milliseconds. Maximum: 300000 (5 minutes).",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      `What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging credits. "return-partial" returns usable results collected so far; if none are available, the request still fails without charging credits. Partial results are not cached as complete results.`,
			InnerField: "behavior",
		},
	},
})

var aiExtractProducts = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract-products",
	Usage:   "Extract product information from a brand's website. We will analyze the website\nand return a list of products with details such as name, description, image,\npricing, features, and more.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "domain",
			Usage:    "The domain name to analyze.",
			BodyPath: "domain",
		},
		&requestflag.Flag[int64]{
			Name:     "max-age-ms",
			Usage:    "Return a cached result if a prior scrape for the same parameters exists and is younger than this many milliseconds. Defaults to 7 days (604800000 ms) when omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.",
			Default:  604800000,
			BodyPath: "maxAgeMs",
		},
		&requestflag.Flag[int64]{
			Name:     "max-products",
			Usage:    "Maximum number of products to extract.",
			BodyPath: "maxProducts",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Optional request deadline and behavior on timeout. For GET requests, use timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded timeoutOpts object.",
			BodyPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:     "direct-url",
			Usage:    "A specific URL to use directly as the starting point for extraction without domain resolution.",
			BodyPath: "directUrl",
		},
	},
	Action:          handleAIExtractProducts,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Request deadline in milliseconds. Maximum: 300000 (5 minutes).",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      `What to do at the deadline. "fail" returns 408 REQUEST_TIMEOUT without charging credits. "return-partial" returns usable results collected so far; if none are available, the request still fails without charging credits. Partial results are not cached as complete results.`,
			InnerField: "behavior",
		},
	},
})

func handleAIExtractProduct(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.AIExtractProductParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.ExtractProduct(ctx, params, options...)
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
		Title:          "ai extract-product",
		Transform:      transform,
	})
}

func handleAIExtractProducts(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.AIExtractProductsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.ExtractProducts(ctx, params, options...)
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
		Title:          "ai extract-products",
		Transform:      transform,
	})
}
