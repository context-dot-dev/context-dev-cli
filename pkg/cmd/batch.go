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

var batchDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently delete a finished batch and its stored results. Active batches must\nsettle first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			Usage:     "ID of the batch to retrieve or cancel.",
			Required:  true,
			PathParam: "batch_id",
		},
	},
	Action:          handleBatchDelete,
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

var batchSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Scrape 25K URLs or crawl large websites asynchronously.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "input",
			Usage:    "Choose a URL list or a site crawl.",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags stored on the batch. Filter the batch list by them later.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "URL notified when the batch finishes.",
			BodyPath: "webhookUrl",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Usage:      "Any string unique to this submission. Retries with the same key return the original batch.",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleBatchSubmit,
	HideHelpCommand: true,
}

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

func handleBatchDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Batch.Delete(ctx, cmd.Value("batch-id").(string), options...)
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
		Title:          "batch delete",
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
