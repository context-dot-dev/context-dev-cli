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

var logsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get one logged API call, including its request input and response body.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "request-id",
			Usage:     "The request ID of the logged API call.",
			Required:  true,
			PathParam: "request_id",
		},
	},
	Action:          handleLogsRetrieve,
	HideHelpCommand: true,
}

var logsList = cli.Command{
	Name:    "list",
	Usage:   "List your organization's API requests, newest first. Defaults to the last 24\nhours.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "error-code",
			Usage:     "Filter by the `error_code` returned in the response.",
			QueryPath: "error_code",
		},
		&requestflag.Flag[bool]{
			Name:      "errors-only",
			Usage:     "Only include requests that returned a 4xx or 5xx status.",
			Default:   false,
			QueryPath: "errors_only",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "Only include requests at or after this ISO 8601 timestamp. Defaults to 24 hours before `to`.",
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "key-id",
			Usage:     "Filter by the API key that made the request.",
			QueryPath: "key_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Number of log entries per page.",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number, starting at 1.",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "path",
			Usage:     "Filter by endpoint path, with or without the /v1 prefix.",
			QueryPath: "path",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Case-insensitive substring match against the request query and body, e.g. a domain.",
			QueryPath: "search",
		},
		&requestflag.Flag[int64]{
			Name:      "status-code",
			Usage:     "Filter by exact HTTP status code.",
			QueryPath: "status_code",
		},
		&requestflag.Flag[string]{
			Name:      "tags",
			Usage:     "Comma-separated request tags. Matches requests carrying any of them. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "Only include requests at or before this ISO 8601 timestamp. Defaults to now.",
			QueryPath: "to",
		},
	},
	Action:          handleLogsList,
	HideHelpCommand: true,
}

func handleLogsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("request-id") && len(unusedArgs) > 0 {
		cmd.Set("request-id", unusedArgs[0])
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
	_, err = client.Logs.Get(ctx, cmd.Value("request-id").(string), options...)
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
		Title:          "logs retrieve",
		Transform:      transform,
	})
}

func handleLogsList(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.LogListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Logs.List(ctx, params, options...)
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
		Title:          "logs list",
		Transform:      transform,
	})
}
