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

var brandRetrieve = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve logos, backdrops, colors, industry, description, and more. Provide\nexactly one lookup identifier in the request body: a domain, company name, email\naddress, stock ticker, transaction descriptor, or direct URL. Note:\n`by_direct_url` fetches brand data only from the provided URL — not from the\nentire internet.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "domain",
			Usage:    "Domain name to retrieve brand data for (e.g., 'stripe.com').",
			BodyPath: "domain",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Discriminator for domain-based brand retrieval.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[*string]{
			Name:     "force-language",
			Usage:    `Allowed values: "afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azeri", "basque", "belarusian", "bengali", "bosnian", "bulgarian", "burmese", "cantonese", "catalan", "cebuano", "chinese", "corsican", "croatian", "czech", "danish", "dutch", "english", "esperanto", "estonian", "farsi", "fijian", "finnish", "french", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian-creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "korean", "kurdish", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luxembourgish", "macedonian", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "mongolian", "nepali", "norwegian", "odia", "oromo", "pashto", "pidgin", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "scottish-gaelic", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovene", "somali", "spanish", "sundanese", "swahili", "swedish", "tagalog", "tajik", "tamil", "tatar", "telugu", "thai", "tibetan", "tigrinya", "tongan", "tswana", "turkish", "turkmen", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "wolof", "xhosa", "yiddish", "yoruba", "zulu".`,
			BodyPath: "force_language",
		},
		&requestflag.Flag[int64]{
			Name:     "max-age-ms",
			Usage:    "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Set to 0 to always perform a hard refresh. Negative values are clamped to 0; values above 1 year (31536000000 ms) are clamped to 1 year.",
			BodyPath: "maxAgeMs",
		},
		&requestflag.Flag[bool]{
			Name:     "max-speed",
			Usage:    "Optional parameter to optimize the API call for maximum speed. When set to true, the API will skip time-consuming operations for faster response at the cost of less comprehensive data.",
			BodyPath: "maxSpeed",
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
			Name:     "name",
			Usage:    "Company name to retrieve brand data for (e.g., 'Apple Inc').",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "country-gl",
			Usage:    "Optional country code hint (GL parameter) to specify the country when looking up by company name.",
			BodyPath: "country_gl",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Email address to retrieve brand data for (e.g., 'jane@stripe.com').",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "ticker",
			Usage:    "Stock ticker symbol to retrieve brand data for (e.g., 'AAPL').",
			BodyPath: "ticker",
		},
		&requestflag.Flag[string]{
			Name:     "ticker-exchange",
			Usage:    "Optional stock exchange for the ticker. Defaults to NASDAQ if not specified.",
			BodyPath: "ticker_exchange",
		},
		&requestflag.Flag[string]{
			Name:     "direct-url",
			Usage:    "Full http(s) URL to fetch brand data from (e.g., 'https://stripe.com/enterprise'). Only this URL is fetched — not the entire internet.",
			BodyPath: "direct_url",
		},
		&requestflag.Flag[string]{
			Name:     "transaction-info",
			Usage:    "Transaction information to identify the brand.",
			BodyPath: "transaction_info",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "Optional city name to prioritize when searching for the brand.",
			BodyPath: "city",
		},
		&requestflag.Flag[bool]{
			Name:     "high-confidence-only",
			Usage:    "When set to true, the API performs additional verification to ensure the identified brand matches the transaction with high confidence.",
			BodyPath: "high_confidence_only",
		},
		&requestflag.Flag[any]{
			Name:     "mcc",
			Usage:    "Optional Merchant Category Code (MCC) to help identify the business category or industry.",
			BodyPath: "mcc",
		},
		&requestflag.Flag[any]{
			Name:     "phone",
			Usage:    "Optional phone number from the transaction to help verify brand match.",
			BodyPath: "phone",
		},
	},
	Action:          handleBrandRetrieve,
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

var brandRetrieveSimplified = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-simplified",
	Usage:   "Returns a simplified version of brand data containing only essential\ninformation: domain, title, colors, logos, and backdrops. Optimized for faster\nresponses and reduced data transfer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain name to retrieve simplified brand data for",
			Required:  true,
			QueryPath: "domain",
		},
		&requestflag.Flag[*int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Set to 0 to always perform a hard refresh. Negative values are clamped to 0; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   requestflag.Ptr[int64](7776000000),
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:      "theme",
			Usage:     "Optional theme preference used when selecting brand assets.",
			QueryPath: "theme",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "timeout-opts",
			Usage:     "Optional request deadline and behavior on timeout. For GET requests, use timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded timeoutOpts object.",
			QueryPath: "timeoutOpts",
		},
	},
	Action:          handleBrandRetrieveSimplified,
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

var brandSearch = cli.Command{
	Name:    "search",
	Usage:   "Search indexed brands by name or domain",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "query",
			Usage:     "Search term, matched against the fields selected by queryBy (e.g. 'nike', 'nike.com', 'nik').",
			Required:  true,
			QueryPath: "query",
		},
		&requestflag.Flag[bool]{
			Name:      "autocomplete",
			Usage:     "Whether the search term matches by prefix, so partial words match as they are typed (e.g. 'nik' matches Nike). Set to false to match whole words only.",
			Default:   true,
			QueryPath: "autocomplete",
		},
		&requestflag.Flag[[]string]{
			Name:      "query-by",
			Usage:     "Fields to match the search term against, as a comma-separated list or repeated parameter: 'name', 'domain', or both. Defaults to both.",
			Default:   []string{"name", "domain"},
			QueryPath: "queryBy",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
		&requestflag.Flag[int64]{
			Name:      "typo-tolerance",
			Usage:     "Maximum number of typos tolerated when matching, from 0 to 2. Defaults to 0 (no typo tolerance).",
			Default:   0,
			QueryPath: "typoTolerance",
		},
	},
	Action:          handleBrandSearch,
	HideHelpCommand: true,
}

func handleBrandRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BrandGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Brand.Get(ctx, params, options...)
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
		Title:          "brand retrieve",
		Transform:      transform,
	})
}

func handleBrandRetrieveSimplified(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BrandGetSimplifiedParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Brand.GetSimplified(ctx, params, options...)
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
		Title:          "brand retrieve-simplified",
		Transform:      transform,
	})
}

func handleBrandSearch(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BrandSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Brand.Search(ctx, params, options...)
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
		Title:          "brand search",
		Transform:      transform,
	})
}
