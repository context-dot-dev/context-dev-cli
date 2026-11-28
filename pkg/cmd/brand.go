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
	Usage:   "Look up a brand by domain, company name, work email, stock ticker, page URL, or\ncard transaction descriptor. Returns logos, colors, backdrops, description,\nindustries, address, stock listing, and social links. A `by_direct_url` lookup\nreads only that page.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "domain",
			Usage:    "Domain of the brand, such as `stripe.com`.",
			BodyPath: "domain",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Set to `by_domain` to look up a brand by domain.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[*string]{
			Name:     "force-language",
			Usage:    "Translate the brand's `title`, `description`, and `slogan` into this language.",
			BodyPath: "force_language",
		},
		&requestflag.Flag[int64]{
			Name:     "max-age-ms",
			Usage:    "Maximum age of cached brand data, in milliseconds. `0` always fetches fresh. Defaults to 90 days (7776000000 ms). Values above 1 year (31536000000 ms) are capped at 1 year.",
			BodyPath: "maxAgeMs",
		},
		&requestflag.Flag[bool]{
			Name:     "max-speed",
			Usage:    "Skip slower enrichment steps for a faster, less detailed response.",
			BodyPath: "maxSpeed",
		},
		&requestflag.Flag[[]string]{
			Name:     "policy",
			Usage:    "Access policy IDs to apply; each URL must pass all of them. Requires access policies enabled for your organization; otherwise returns `ACCESS_POLICIES_NOT_ENABLED`.",
			BodyPath: "policies",
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
			Usage:    "`enabled` turns off request logs and traces and keeps request and response content out of usage logs. Brand data and hosted assets may still be stored. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			BodyPath: "zdr",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Company name, such as `Apple Inc`.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Lowercase ISO 3166-1 alpha-2 country code, such as `us`.",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Work email address, such as `jane@stripe.com`. The brand is looked up from its domain.",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "ticker",
			Usage:    "Stock ticker symbol, such as `AAPL`.",
			BodyPath: "ticker",
		},
		&requestflag.Flag[string]{
			Name:     "ticker-exchange",
			Usage:    "Stock exchange code, such as `NASDAQ` or `LSE`.",
			BodyPath: "ticker_exchange",
		},
		&requestflag.Flag[string]{
			Name:     "direct-url",
			Usage:    "Full `http` or `https` URL of the page to read, such as `https://stripe.com/enterprise`. Only this page is fetched.",
			BodyPath: "direct_url",
		},
		&requestflag.Flag[string]{
			Name:     "transaction-info",
			Usage:    "Transaction descriptor from a card or bank statement, such as `SQ *COFFEE SHOP`.",
			BodyPath: "transaction_info",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "City where the transaction took place. Brands in this city are preferred.",
			BodyPath: "city",
		},
		&requestflag.Flag[bool]{
			Name:     "high-confidence-only",
			Usage:    "Run extra verification so the returned brand matches the transaction with high confidence.",
			BodyPath: "high_confidence_only",
		},
		&requestflag.Flag[any]{
			Name:     "mcc",
			Usage:    "Merchant Category Code (MCC) from the transaction. Helps identify the type of business.",
			BodyPath: "mcc",
		},
		&requestflag.Flag[any]{
			Name:     "phone",
			Usage:    "Merchant phone number from the transaction. Helps confirm the match.",
			BodyPath: "phone",
		},
	},
	Action:          handleBrandRetrieve,
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
			Usage:      "`fail` returns 408 at the deadline. `return-partial` returns the results collected so far and marks the response as partial.",
			InnerField: "behavior",
		},
	},
})

var brandSearch = cli.Command{
	Name:    "search",
	Usage:   "Autocomplete brand names and domains as a user types, with up to 10 brands\nranked by popularity and a ready-to-use logo URL for each. Filter by country and\nprimary language. Pass a returned `domain` to `POST /brand/retrieve` for the\nfull brand profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "query",
			Usage:     "Search term, matched against the fields in `queryBy`, such as `nike`, `nike.com`, or `nik`.",
			Required:  true,
			QueryPath: "query",
		},
		&requestflag.Flag[bool]{
			Name:      "autocomplete",
			Usage:     "Match the search term as a prefix, so `nik` matches Nike while the user is still typing. Set to `false` to match whole words only.",
			Default:   true,
			QueryPath: "autocomplete",
		},
		&requestflag.Flag[string]{
			Name:      "country",
			Usage:     "Only return brands in this country, as an uppercase ISO 3166-1 alpha-2 code such as `US` or `GB`.",
			QueryPath: "country",
		},
		&requestflag.Flag[string]{
			Name:      "language",
			Usage:     "Only return brands with this primary language, such as `english` or `french`.",
			QueryPath: "language",
		},
		&requestflag.Flag[[]string]{
			Name:      "query-by",
			Usage:     "Fields to match the search term against: `name`, `domain`, or both. Send a comma-separated list or repeat the parameter.",
			Default:   []string{"name", "domain"},
			QueryPath: "queryBy",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for this request, such as `production,team-alpha`. Filter usage and request logs by them.",
			QueryPath: "tags",
		},
		&requestflag.Flag[int64]{
			Name:      "typo-tolerance",
			Usage:     "Maximum number of typos to tolerate when matching. `0` requires exact spelling.",
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
