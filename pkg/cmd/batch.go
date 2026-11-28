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
	Usage:   "Retrieve a batch's status, progress, and credit usage, plus links to download\nits results once it finishes. Result files are deleted 180 days after the batch\nfinishes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			Usage:     "ID of the batch.",
			Required:  true,
			PathParam: "batch_id",
		},
	},
	Action:          handleBatchRetrieve,
	HideHelpCommand: true,
}

var batchList = cli.Command{
	Name:    "list",
	Usage:   "List your batches, newest first. Filter by status or tags, or search by batch\nID, crawl source, or tag.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor from the previous page's `next_cursor`. Omit for the first page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of batches to return per page.",
			Default:   25,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Search term matched against batch IDs, crawl sources (start URL or sitemap domain), and tags.",
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "search-type",
			Usage:     "How `q` matches: `prefix` matches the start of a term, for search-as-you-type, and `exact` matches whole terms. Defaults to `prefix`.",
			QueryPath: "search_type",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by batch status.",
			QueryPath: "status",
		},
		&requestflag.Flag[any]{
			Name:      "tags",
			Usage:     "Comma-separated tags. Returns batches that have any of them. You can also repeat the `tags` parameter.",
			QueryPath: "tags",
		},
	},
	Action:          handleBatchList,
	HideHelpCommand: true,
}

var batchDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently delete a finished batch and its results. Its webhook deliveries can\nno longer be retried.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			Usage:     "ID of the batch.",
			Required:  true,
			PathParam: "batch_id",
		},
	},
	Action:          handleBatchDelete,
	HideHelpCommand: true,
}

var batchCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Stop a batch from starting new pages. Pages already in progress finish, then the\nbatch becomes `cancelled` and unused credits are refunded.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			Usage:     "ID of the batch.",
			Required:  true,
			PathParam: "batch_id",
		},
	},
	Action:          handleBatchCancel,
	HideHelpCommand: true,
}

var batchGetResults = cli.Command{
	Name:    "get-results",
	Usage:   "Page through a finished batch's results as JSON. Each record's `result` has the\nsame outputs and per-format errors as `POST /web/scrape`. Results are kept for\n180 days after the batch finishes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			Usage:     "ID of the batch.",
			Required:  true,
			PathParam: "batch_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor from the previous page's `next_cursor`. Omit for the first page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return per page. A page can hold fewer to stay under about 8 MB, so page with `next_cursor` instead of counting records.",
			Default:   25,
			QueryPath: "limit",
		},
	},
	Action:          handleBatchGetResults,
	HideHelpCommand: true,
}

var batchSubmit = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit",
	Usage:   "Scrape up to 100,000 URLs, or crawl a site, asynchronously with every Scrape\nAnything format and option from `POST /web/scrape`. Poll `GET /batch/{batch_id}`\nor pass `webhook` to be notified when the batch finishes. URLs and credits are\nchecked after the batch is accepted, so a batch that fails those checks ends as\n`failed` with a `failure` code instead of returning an error.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "input",
			Usage:    "What to scrape: a list of URLs (`scrape`) or a site crawl (`crawl`).",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[[]string]{
			Name:     "policy",
			Usage:    "Access policy IDs to apply to every page; each URL must pass all of them. Requires access policies enabled for your organization; otherwise returns `ACCESS_POLICIES_NOT_ENABLED`.",
			BodyPath: "policies",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to store on the batch. Filter `GET /batch/list` by them later.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "webhook",
			Usage:    "Webhook to notify when the batch finishes. Omit `retry` for a single attempt, or pass `retry: {}` for the default retry schedule.",
			BodyPath: "webhook",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Usage:      "Unique key for this submission. Retrying with the same key and body returns the original batch instead of creating another; a different body returns `409`.",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleBatchSubmit,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"webhook": {
		&requestflag.InnerFlag[string]{
			Name:       "webhook.url",
			Usage:      "Public HTTP or HTTPS URL that receives the `batch.completed`, `batch.failed`, or `batch.cancelled` event.",
			InnerField: "url",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "webhook.retry",
			Usage:      "Automatic retry settings. Pass `{}` for the default schedule.",
			InnerField: "retry",
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
