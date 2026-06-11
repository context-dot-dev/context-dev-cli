// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/context-dot-dev/context-go-sdk"
	"github.com/context-dot-dev/context-go-sdk/option"
	"github.com/stainless-sdks/context.dev-cli/internal/apiquery"
	"github.com/stainless-sdks/context.dev-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var brandRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve logos, backdrops, colors, industry, description, and more from any\ndomain",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain name to retrieve brand data for (e.g., 'example.com', 'google.com'). Cannot be used with name or ticker parameters.",
			Required:  true,
			QueryPath: "domain",
		},
		&requestflag.Flag[string]{
			Name:      "force-language",
			Usage:     "Optional parameter to force the language of the retrieved brand data.",
			QueryPath: "force_language",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   7776000000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[bool]{
			Name:      "max-speed",
			Usage:     "Optional parameter to optimize the API call for maximum speed. When set to true, the API will skip time-consuming operations for faster response at the cost of less comprehensive data. Works with all three lookup methods.",
			QueryPath: "maxSpeed",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleBrandRetrieve,
	HideHelpCommand: true,
}

var brandIdentifyFromTransaction = cli.Command{
	Name:    "identify-from-transaction",
	Usage:   "Endpoint specially designed for platforms that want to identify transaction data\nby the transaction title.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "transaction-info",
			Usage:     "Transaction information to identify the brand",
			Required:  true,
			QueryPath: "transaction_info",
		},
		&requestflag.Flag[string]{
			Name:      "city",
			Usage:     "Optional city name to prioritize when searching for the brand.",
			QueryPath: "city",
		},
		&requestflag.Flag[string]{
			Name:      "country-gl",
			Usage:     "Optional country code (GL parameter) to specify the country. This affects the geographic location used for search queries.",
			QueryPath: "country_gl",
		},
		&requestflag.Flag[string]{
			Name:      "force-language",
			Usage:     "Optional parameter to force the language of the retrieved brand data.",
			QueryPath: "force_language",
		},
		&requestflag.Flag[bool]{
			Name:      "high-confidence-only",
			Usage:     "When set to true, the API will perform an additional verification steps to ensure the identified brand matches the transaction with high confidence.",
			Default:   false,
			QueryPath: "high_confidence_only",
		},
		&requestflag.Flag[bool]{
			Name:      "max-speed",
			Usage:     "Optional parameter to optimize the API call for maximum speed. When set to true, the API will skip time-consuming operations for faster response at the cost of less comprehensive data.",
			QueryPath: "maxSpeed",
		},
		&requestflag.Flag[string]{
			Name:      "mcc",
			Usage:     "Optional Merchant Category Code (MCC) to help identify the business category/industry. ",
			QueryPath: "mcc",
		},
		&requestflag.Flag[float64]{
			Name:      "phone",
			Usage:     "Optional phone number from the transaction to help verify brand match.",
			QueryPath: "phone",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleBrandIdentifyFromTransaction,
	HideHelpCommand: true,
}

var brandRetrieveByEmail = cli.Command{
	Name:    "retrieve-by-email",
	Usage:   "Retrieve brand information using an email address while detecting disposable and\nfree email addresses. Disposable and free email addresses (like gmail.com,\nyahoo.com) will throw a 422 error.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email",
			Usage:     "Email address to retrieve brand data for (e.g., 'contact@example.com'). The domain will be extracted from the email. Free email providers (gmail.com, yahoo.com, etc.) and disposable email addresses are not allowed.",
			Required:  true,
			QueryPath: "email",
		},
		&requestflag.Flag[string]{
			Name:      "force-language",
			Usage:     "Optional parameter to force the language of the retrieved brand data.",
			QueryPath: "force_language",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   7776000000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[bool]{
			Name:      "max-speed",
			Usage:     "Optional parameter to optimize the API call for maximum speed. When set to true, the API will skip time-consuming operations for faster response at the cost of less comprehensive data.",
			QueryPath: "maxSpeed",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleBrandRetrieveByEmail,
	HideHelpCommand: true,
}

var brandRetrieveByIsin = cli.Command{
	Name:    "retrieve-by-isin",
	Usage:   "Retrieve brand information using an ISIN (International Securities\nIdentification Number).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "isin",
			Usage:     "ISIN (International Securities Identification Number) to retrieve brand data for (e.g., 'AU000000IMD5', 'US0378331005'). Must be exactly 12 characters: 2 letters followed by 9 alphanumeric characters and ending with a digit.",
			Required:  true,
			QueryPath: "isin",
		},
		&requestflag.Flag[string]{
			Name:      "force-language",
			Usage:     "Optional parameter to force the language of the retrieved brand data.",
			QueryPath: "force_language",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   7776000000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[bool]{
			Name:      "max-speed",
			Usage:     "Optional parameter to optimize the API call for maximum speed. When set to true, the API will skip time-consuming operations for faster response at the cost of less comprehensive data.",
			QueryPath: "maxSpeed",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleBrandRetrieveByIsin,
	HideHelpCommand: true,
}

var brandRetrieveByName = cli.Command{
	Name:    "retrieve-by-name",
	Usage:   "Retrieve brand information using a company name.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Company name to retrieve brand data for (e.g., 'Apple Inc', 'Microsoft Corporation'). Must be 3-30 characters.",
			Required:  true,
			QueryPath: "name",
		},
		&requestflag.Flag[string]{
			Name:      "country-gl",
			Usage:     "Optional country code hint (GL parameter) to specify the country for the company name.",
			QueryPath: "country_gl",
		},
		&requestflag.Flag[string]{
			Name:      "force-language",
			Usage:     "Optional parameter to force the language of the retrieved brand data.",
			QueryPath: "force_language",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   7776000000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[bool]{
			Name:      "max-speed",
			Usage:     "Optional parameter to optimize the API call for maximum speed. When set to true, the API will skip time-consuming operations for faster response at the cost of less comprehensive data.",
			QueryPath: "maxSpeed",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleBrandRetrieveByName,
	HideHelpCommand: true,
}

var brandRetrieveByTicker = cli.Command{
	Name:    "retrieve-by-ticker",
	Usage:   "Retrieve brand information using a stock ticker symbol.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ticker",
			Usage:     "Stock ticker symbol to retrieve brand data for (e.g., 'AAPL', 'GOOGL', 'BRK.A'). Must be 1-15 characters, letters/numbers/dots only.",
			Required:  true,
			QueryPath: "ticker",
		},
		&requestflag.Flag[string]{
			Name:      "force-language",
			Usage:     "Optional parameter to force the language of the retrieved brand data.",
			QueryPath: "force_language",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   7776000000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[bool]{
			Name:      "max-speed",
			Usage:     "Optional parameter to optimize the API call for maximum speed. When set to true, the API will skip time-consuming operations for faster response at the cost of less comprehensive data.",
			QueryPath: "maxSpeed",
		},
		&requestflag.Flag[string]{
			Name:      "ticker-exchange",
			Usage:     "Optional stock exchange for the ticker. Defaults to NASDAQ if not specified.",
			QueryPath: "ticker_exchange",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleBrandRetrieveByTicker,
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
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached brand data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   7776000000,
			QueryPath: "maxAgeMs",
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
		EmptyBody,
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

func handleBrandIdentifyFromTransaction(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BrandIdentifyFromTransactionParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Brand.IdentifyFromTransaction(ctx, params, options...)
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
		Title:          "brand identify-from-transaction",
		Transform:      transform,
	})
}

func handleBrandRetrieveByEmail(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BrandGetByEmailParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Brand.GetByEmail(ctx, params, options...)
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
		Title:          "brand retrieve-by-email",
		Transform:      transform,
	})
}

func handleBrandRetrieveByIsin(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BrandGetByIsinParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Brand.GetByIsin(ctx, params, options...)
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
		Title:          "brand retrieve-by-isin",
		Transform:      transform,
	})
}

func handleBrandRetrieveByName(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BrandGetByNameParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Brand.GetByName(ctx, params, options...)
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
		Title:          "brand retrieve-by-name",
		Transform:      transform,
	})
}

func handleBrandRetrieveByTicker(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BrandGetByTickerParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Brand.GetByTicker(ctx, params, options...)
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
		Title:          "brand retrieve-by-ticker",
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
