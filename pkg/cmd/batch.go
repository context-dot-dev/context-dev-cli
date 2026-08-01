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

var batchRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Check progress, and get download links once the batch finishes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			Usage:     "ID of the batch to retrieve or cancel.",
			Required:  true,
			PathParam: "batch_id",
		},
	},
	Action:          handleBatchRetrieve,
	HideHelpCommand: true,
}

var batchList = cli.Command{
	Name:    "list",
	Usage:   "List your batches from newest to oldest. Filter by status or continue with a\ncursor.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor from the previous page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Batches per page. Defaults to 25.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Free-text search term, matched against the batch id, crawl source (start URL or sitemap domain), and tags.",
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "search-type",
			Usage:     "`prefix` for as-you-type prefix matching (default), `exact` for full-token matching.",
			QueryPath: "search_type",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by status.",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "tags",
			Usage:     "Comma-separated list of tags to filter by (matches batches having any of them).",
			QueryPath: "tags",
		},
	},
	Action:          handleBatchList,
	HideHelpCommand: true,
}

var batchCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Stop a batch from starting new pages. In-progress pages finish, and unused\ncredits are refunded.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			Usage:     "ID of the batch to retrieve or cancel.",
			Required:  true,
			PathParam: "batch_id",
		},
	},
	Action:          handleBatchCancel,
	HideHelpCommand: true,
}

var batchGetResults = cli.Command{
	Name:    "get-results",
	Usage:   "Page through a finished batch's results as JSON instead of downloading the\nNDJSON files.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			Usage:     "ID of the batch to retrieve or cancel.",
			Required:  true,
			PathParam: "batch_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "next_cursor from the previous page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Records per page. Defaults to 25. A page can close early so its payload stays under ~8 MB; rely on next_cursor rather than counting records.",
			QueryPath: "limit",
		},
	},
	Action:          handleBatchGetResults,
	HideHelpCommand: true,
}

var batchSubmit = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit",
	Usage:   "Retrieve and normalize a person profile from identifiers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "identifiers",
			Usage:    "Known identifiers for the person. At least one identifier is required.",
			Required: true,
			BodyPath: "identifiers",
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
	Action:          handleBatchSubmit,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"identifiers": {
		&requestflag.InnerFlag[string]{
			Name:       "identifiers.linkedin-url",
			Usage:      "LinkedIn profile URL, e.g. https://www.linkedin.com/in/yahia-bakour/.",
			InnerField: "linkedinUrl",
		},
	},
})

func handleBatchRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("batch-id") && len(unusedArgs) > 0 {
		cmd.Set("batch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Batch.Get(ctx, cmd.Value("batch-id").(string), options...)
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
		Title:          "batch retrieve",
		Transform:      transform,
	})
}

func handleBatchList(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BatchListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Batch.List(ctx, params, options...)
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
		Title:          "batch list",
		Transform:      transform,
	})
}

func handleBatchCancel(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("batch-id") && len(unusedArgs) > 0 {
		cmd.Set("batch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Batch.Cancel(ctx, cmd.Value("batch-id").(string), options...)
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
		Title:          "batch cancel",
		Transform:      transform,
	})
}

func handleBatchGetResults(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("batch-id") && len(unusedArgs) > 0 {
		cmd.Set("batch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := contextdev.BatchGetResultsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Batch.GetResults(
		ctx,
		cmd.Value("batch-id").(string),
		params,
		options...,
	)
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
		Title:          "batch get-results",
		Transform:      transform,
	})
}

func handleBatchSubmit(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.BatchSubmitParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Batch.Submit(ctx, params, options...)
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
		Title:          "batch submit",
		Transform:      transform,
	})
}
