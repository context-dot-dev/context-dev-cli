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

var brandRetrieve = cli.Command{
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
			Usage:    "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
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
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Usage:    "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			BodyPath: "timeoutMS",
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
}

var brandRetrieveSimplified = cli.Command{
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
			Usage:     "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   requestflag.Ptr[int64](7776000000),
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Optional comma-separated caller-defined tags for tracking this request. Tags are recorded on the request's usage log and can be used to filter usage on the dashboard usage page. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:      "theme",
			Usage:     "Optional theme preference used when selecting brand assets.",
			QueryPath: "theme",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleBrandRetrieveSimplified,
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
