// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/context-dot-dev/context-dev-cli/internal/apiquery"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
	"github.com/context-dot-dev/context-go-sdk"
	"github.com/context-dot-dev/context-go-sdk/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiAIQuery = requestflag.WithInnerFlags(cli.Command{
	Name:    "ai-query",
	Usage:   "Use AI to extract specific data points from a brand's website. The AI will crawl\nthe website and extract the requested information based on the provided data\npoints.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "data-to-extract",
			Usage:    "Array of data points to extract from the website",
			Required: true,
			BodyPath: "data_to_extract",
		},
		&requestflag.Flag[string]{
			Name:     "domain",
			Usage:    "The domain name to analyze",
			Required: true,
			BodyPath: "domain",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "specific-pages",
			Usage:    "Optional object specifying which pages to analyze",
			BodyPath: "specific_pages",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Usage:    "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			BodyPath: "timeoutMS",
		},
	},
	Action:          handleAIAIQuery,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data-to-extract": {
		&requestflag.InnerFlag[string]{
			Name:       "data-to-extract.datapoint-description",
			Usage:      "Description of what to extract",
			InnerField: "datapoint_description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-to-extract.datapoint-example",
			Usage:      "Example of the expected value",
			InnerField: "datapoint_example",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-to-extract.datapoint-name",
			Usage:      "Name of the data point to extract",
			InnerField: "datapoint_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-to-extract.datapoint-type",
			Usage:      "Type of the data point",
			InnerField: "datapoint_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-to-extract.datapoint-list-type",
			Usage:      "Type of items in the list when datapoint_type is 'list'. Defaults to 'string'. Use 'object' to extract an array of objects matching a schema.",
			InnerField: "datapoint_list_type",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data-to-extract.datapoint-object-schema",
			Usage:      "Schema definition for objects when datapoint_list_type is 'object'. Provide a map of field names to their scalar types.",
			InnerField: "datapoint_object_schema",
		},
	},
	"specific-pages": {
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.about-us",
			Usage:      "Whether to analyze the about us page",
			InnerField: "about_us",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.blog",
			Usage:      "Whether to analyze the blog",
			InnerField: "blog",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.careers",
			Usage:      "Whether to analyze the careers page",
			InnerField: "careers",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.contact-us",
			Usage:      "Whether to analyze the contact us page",
			InnerField: "contact_us",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.faq",
			Usage:      "Whether to analyze the FAQ page",
			InnerField: "faq",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.home-page",
			Usage:      "Whether to analyze the home page",
			InnerField: "home_page",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.pricing",
			Usage:      "Whether to analyze the pricing page",
			InnerField: "pricing",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.privacy-policy",
			Usage:      "Whether to analyze the privacy policy page",
			InnerField: "privacy_policy",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "specific-pages.terms-and-conditions",
			Usage:      "Whether to analyze the terms and conditions page",
			InnerField: "terms_and_conditions",
		},
	},
})

var aiExtractProduct = cli.Command{
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
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Usage:    "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			BodyPath: "timeoutMS",
		},
	},
	Action:          handleAIExtractProduct,
	HideHelpCommand: true,
}

var aiExtractProducts = cli.Command{
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
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Usage:    "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			BodyPath: "timeoutMS",
		},
		&requestflag.Flag[string]{
			Name:     "direct-url",
			Usage:    "A specific URL to use directly as the starting point for extraction without domain resolution.",
			BodyPath: "directUrl",
		},
	},
	Action:          handleAIExtractProducts,
	HideHelpCommand: true,
}

func handleAIAIQuery(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.AIAIQueryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.AIQuery(ctx, params, options...)
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
		Title:          "ai ai-query",
		Transform:      transform,
	})
}

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
