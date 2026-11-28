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
	Usage:   "Retrieve a request's metadata, retained input, and response. A request that just\nfinished can return 404 `NOT_FOUND` for a few seconds, so retry shortly.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "request-id",
			Usage:     "ID of the request, from its response's `request_id` or `X-Request-Id` header.",
			Required:  true,
			PathParam: "request_id",
		},
	},
	Action:          handleLogsRetrieve,
	HideHelpCommand: true,
}

var logsList = cli.Command{
	Name:    "list",
	Usage:   "List your organization's API requests, newest first. Filter by time, endpoint,\nAPI key, status, error code, tags, or text. Includes batch settlements and\nmonitor runs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "error-code",
			Usage:     "Filter by the `error_code` returned in the response.",
			QueryPath: "error_code",
		},
		&requestflag.Flag[bool]{
			Name:      "errors-only",
			Usage:     "Return only requests that failed with a 4xx or 5xx status.",
			Default:   false,
			QueryPath: "errors_only",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "Return requests at or after this time. Defaults to 24 hours before `to`.",
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "key-id",
			Usage:     "Filter by the ID of the API key that made the request.",
			QueryPath: "key_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of log entries to return per page.",
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
			Usage:     "Filter by endpoint path, with or without the `/v1` prefix.",
			QueryPath: "path",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Case-insensitive text to find in the request query or body, such as a domain.",
			QueryPath: "search",
		},
		&requestflag.Flag[int64]{
			Name:      "status-code",
			Usage:     "Filter by exact HTTP status code.",
			QueryPath: "status_code",
		},
		&requestflag.Flag[string]{
			Name:      "tags",
			Usage:     "Comma-separated tags. Returns log entries that have any of them. Up to 20 tags of 1 to 50 characters each.",
			QueryPath: "tags",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "Return requests at or before this time. Defaults to now.",
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
