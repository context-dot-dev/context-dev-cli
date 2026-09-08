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

var webhooksDeliveriesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a webhook delivery, including its status and latest attempt.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "delivery-id",
			Usage:     "Delivery ID.",
			Required:  true,
			PathParam: "delivery_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
	},
	Action:          handleWebhooksDeliveriesRetrieve,
	HideHelpCommand: true,
}

var webhooksDeliveriesList = cli.Command{
	Name:    "list",
	Usage:   "List your batch or monitor webhook deliveries, newest first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Delivery source.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "batch-id",
			Usage:    "Filter by batch ID.",
			BodyPath: "batch_id",
		},
		&requestflag.Flag[any]{
			Name:     "created-after",
			Usage:    "Only include events created after this ISO 8601 timestamp.",
			BodyPath: "created_after",
		},
		&requestflag.Flag[string]{
			Name:     "cursor",
			Usage:    "The next_cursor from the previous response.",
			BodyPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Number of deliveries to return.",
			Default:  25,
			BodyPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Filter by delivery status.",
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "monitor-id",
			Usage:    "Filter by monitor ID.",
			BodyPath: "monitor_id",
		},
		&requestflag.Flag[string]{
			Name:     "run-id",
			Usage:    "Filter by monitor run ID.",
			BodyPath: "run_id",
		},
	},
	Action:          handleWebhooksDeliveriesList,
	HideHelpCommand: true,
}

var webhooksDeliveriesListAttempts = cli.Command{
	Name:    "list-attempts",
	Usage:   "List delivery attempts, newest first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "delivery-id",
			Usage:     "Delivery ID.",
			Required:  true,
			PathParam: "delivery_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "The next_cursor from the previous response.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Number of attempts to return.",
			Default:   25,
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated tags for tracking request usage. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
	},
	Action:          handleWebhooksDeliveriesListAttempts,
	HideHelpCommand: true,
}

var webhooksDeliveriesRetry = cli.Command{
	Name:    "retry",
	Usage:   "Retry a webhook delivery within seven days of creation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "delivery-id",
			Usage:     "Delivery ID.",
			Required:  true,
			PathParam: "delivery_id",
		},
		&requestflag.Flag[bool]{
			Name:     "force",
			Usage:    "Resend a delivery that already succeeded.",
			BodyPath: "force",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Usage:      "Unique key to prevent duplicate retry requests.",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleWebhooksDeliveriesRetry,
	HideHelpCommand: true,
}

func handleWebhooksDeliveriesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("delivery-id") && len(unusedArgs) > 0 {
		cmd.Set("delivery-id", unusedArgs[0])
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

	params := contextdev.WebhookDeliveryGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webhooks.Deliveries.Get(
		ctx,
		cmd.Value("delivery-id").(string),
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
		Title:          "webhooks:deliveries retrieve",
		Transform:      transform,
	})
}

func handleWebhooksDeliveriesList(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebhookDeliveryListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webhooks.Deliveries.List(ctx, params, options...)
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
		Title:          "webhooks:deliveries list",
		Transform:      transform,
	})
}

func handleWebhooksDeliveriesListAttempts(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("delivery-id") && len(unusedArgs) > 0 {
		cmd.Set("delivery-id", unusedArgs[0])
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

	params := contextdev.WebhookDeliveryListAttemptsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webhooks.Deliveries.ListAttempts(
		ctx,
		cmd.Value("delivery-id").(string),
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
		Title:          "webhooks:deliveries list-attempts",
		Transform:      transform,
	})
}

func handleWebhooksDeliveriesRetry(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("delivery-id") && len(unusedArgs) > 0 {
		cmd.Set("delivery-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := contextdev.WebhookDeliveryRetryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webhooks.Deliveries.Retry(
		ctx,
		cmd.Value("delivery-id").(string),
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
		Title:          "webhooks:deliveries retry",
		Transform:      transform,
	})
}
