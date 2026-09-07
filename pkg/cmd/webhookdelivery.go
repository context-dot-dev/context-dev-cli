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
	Usage:   "Get the live status, retry policy, latest attempt, and replay expiration for a\nretained delivery. Use the attempts endpoint for its complete paginated history.\nThis endpoint costs no credits.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "delivery-id",
			Required:  true,
			PathParam: "delivery_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Optional comma-separated caller-defined tags for tracking this request. Tags are recorded on the request's usage log and can be used to filter usage on the dashboard usage page. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
	},
	Action:          handleWebhooksDeliveriesRetrieve,
	HideHelpCommand: true,
}

var webhooksDeliveriesList = cli.Command{
	Name:    "list",
	Usage:   "List retained batch and monitor webhook deliveries for your organization, newest\nfirst. Filter by at most one of batch_id, monitor_id, or run_id, optionally\ncombined with status. Historical events without retained payloads are not\nlisted. This endpoint costs no credits.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "batch-id",
			QueryPath: "batch_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   25,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "monitor-id",
			QueryPath: "monitor_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			QueryPath: "run_id",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "pending", "delivering", "retrying", "delivered", "failed", "cancelled".`,
			QueryPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Optional comma-separated caller-defined tags for tracking this request. Tags are recorded on the request's usage log and can be used to filter usage on the dashboard usage page. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
	},
	Action:          handleWebhooksDeliveriesList,
	HideHelpCommand: true,
}

var webhooksDeliveriesListAttempts = cli.Command{
	Name:    "list-attempts",
	Usage:   "List individual HTTP attempts for a delivery, newest first, including their\ndestination, timestamps, HTTP status, and error. An interrupted attempt may have\nreached the endpoint even when its outcome is unknown. This endpoint costs no\ncredits.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "delivery-id",
			Required:  true,
			PathParam: "delivery_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   25,
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Optional comma-separated caller-defined tags for tracking this request. Tags are recorded on the request's usage log and can be used to filter usage on the dashboard usage page. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
	},
	Action:          handleWebhooksDeliveriesListAttempts,
	HideHelpCommand: true,
}

var webhooksDeliveriesRetry = cli.Command{
	Name:    "retry",
	Usage:   "Queue an immediate attempt without rerunning or billing the underlying batch or\nmonitor. A waiting retry is brought forward. A failed delivery gets one\nadditional attempt without restarting its automatic retry budget. Set force:\ntrue to resend an acknowledged delivery. An in-progress attempt cannot be\nduplicated. The stored event body, event ID, and creation time remain unchanged;\neach attempt receives a fresh signature. Monitor retries use the current URL and\nsecret; removing the webhook cancels pending deliveries. Batch result URLs in\nold payloads may have expired: retrieve the batch to get fresh URLs. Replay is\navailable for seven days. A successful attempt cancels remaining automatic\nretries. Idempotency-Key is scoped to your organization and retained with the\ndelivery metadata; repeating the same key and input returns the original\naccepted response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "delivery-id",
			Required:  true,
			PathParam: "delivery_id",
		},
		&requestflag.Flag[bool]{
			Name:     "force",
			BodyPath: "force",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
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
		EmptyBody,
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
