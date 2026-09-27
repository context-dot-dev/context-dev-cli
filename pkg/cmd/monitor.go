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

var monitorsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Watch a page, URL inventory, or extracted website data on a schedule. A run\nstarts immediately to capture the baseline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Display name for the monitor.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "target",
			Usage:    "What to watch: a page, a sitemap, or data extracted from a site.",
			Required: true,
			BodyPath: "target",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "change-detection",
			Usage:    "How changes are judged. Defaults to `semantic` for extract targets and page targets with `instructions`, otherwise `exact`.",
			BodyPath: "change_detection",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Always `web`. Optional.",
			BodyPath: "mode",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "schedule",
			Usage:    "Run the monitor on a fixed interval defined by a frequency and a unit, e.g. every 6 hours or every 2 days. The total interval (frequency × unit) must be between 10 minutes and 1 year.",
			BodyPath: "schedule",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for filtering monitors, their changes, and their usage.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "webhook",
			Usage:    "Webhook destination and delivery settings. Null means no webhook is configured.",
			BodyPath: "webhook",
		},
	},
	Action:          handleMonitorsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"schedule": {
		&requestflag.InnerFlag[int64]{
			Name:       "schedule.frequency",
			Usage:      "Number of units between runs. The resulting interval (frequency × unit) must be at least 10 minutes and at most 1 year (e.g. minimum 10 when unit is minutes; maximum 365 when unit is days).",
			InnerField: "frequency",
		},
		&requestflag.InnerFlag[string]{
			Name:       "schedule.type",
			Usage:      "Use `interval` to run on a repeating schedule.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "schedule.unit",
			Usage:      "Time unit used with `frequency` to set the run interval.",
			InnerField: "unit",
		},
	},
	"webhook": {
		&requestflag.InnerFlag[string]{
			Name:       "webhook.url",
			Usage:      "Public HTTP(S) URL that receives events. Slack and GovSlack URLs get formatted messages.",
			InnerField: "url",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "webhook.events",
			Usage:      "Events to deliver. Defaults to `change.detected`; `run.completed` also includes unchanged runs.",
			InnerField: "events",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "webhook.retry",
			Usage:      "Webhook retry settings. Use {} for the default schedule.",
			InnerField: "retry",
		},
		&requestflag.InnerFlag[string]{
			Name:       "webhook.secret",
			Usage:      "API-generated signing secret. Visible only with full access or `monitors:write` permission.",
			InnerField: "secret",
		},
	},
})

var monitorsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a monitor’s configuration and current state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "monitor-id",
			Usage:     "ID of the monitor.",
			Required:  true,
			PathParam: "monitor_id",
		},
	},
	Action:          handleMonitorsRetrieve,
	HideHelpCommand: true,
}

var monitorsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a monitor. Changing its target or change detection replaces the baseline\nand queues a new baseline run.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "monitor-id",
			Usage:     "ID of the monitor.",
			Required:  true,
			PathParam: "monitor_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "change-detection",
			Usage:    "How changes are judged. Defaults to `semantic` for extract targets and page targets with `instructions`, otherwise `exact`.",
			BodyPath: "change_detection",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Display name for the monitor.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "schedule",
			Usage:    "Run the monitor on a fixed interval defined by a frequency and a unit, e.g. every 6 hours or every 2 days. The total interval (frequency × unit) must be between 10 minutes and 1 year.",
			BodyPath: "schedule",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Set `paused` to stop scheduled runs or `active` to resume them.",
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for filtering monitors, their changes, and their usage.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "target",
			Usage:    "What to watch: a page, a sitemap, or data extracted from a site.",
			BodyPath: "target",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "webhook",
			Usage:    "Set to null to remove the webhook. Changing `url` issues a new secret.",
			BodyPath: "webhook",
		},
	},
	Action:          handleMonitorsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"schedule": {
		&requestflag.InnerFlag[int64]{
			Name:       "schedule.frequency",
			Usage:      "Number of units between runs. The resulting interval (frequency × unit) must be at least 10 minutes and at most 1 year (e.g. minimum 10 when unit is minutes; maximum 365 when unit is days).",
			InnerField: "frequency",
		},
		&requestflag.InnerFlag[string]{
			Name:       "schedule.type",
			Usage:      "Use `interval` to run on a repeating schedule.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "schedule.unit",
			Usage:      "Time unit used with `frequency` to set the run interval.",
			InnerField: "unit",
		},
	},
	"webhook": {
		&requestflag.InnerFlag[string]{
			Name:       "webhook.url",
			Usage:      "Public HTTP(S) URL that receives events. Slack and GovSlack URLs get formatted messages.",
			InnerField: "url",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "webhook.events",
			Usage:      "Events to deliver. Defaults to `change.detected`; `run.completed` also includes unchanged runs.",
			InnerField: "events",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "webhook.retry",
			Usage:      "Webhook retry settings. Use {} for the default schedule.",
			InnerField: "retry",
		},
		&requestflag.InnerFlag[string]{
			Name:       "webhook.secret",
			Usage:      "API-generated signing secret. Visible only with full access or `monitors:write` permission.",
			InnerField: "secret",
		},
	},
})

var monitorsList = cli.Command{
	Name:    "list",
	Usage:   "List your monitors with optional search and filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "change-detection-type",
			Usage:     "Filter by change detection type.",
			QueryPath: "change_detection_type",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return per page (1-100). Defaults to 25.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Free-text search term, matched against the fields named in `search_by`.",
			QueryPath: "q",
		},
		&requestflag.Flag[any]{
			Name:      "search-by",
			Usage:     "Fields to search with `q`. Defaults to all fields; page and extract targets can have instructions.",
			QueryPath: "search_by",
		},
		&requestflag.Flag[string]{
			Name:      "search-type",
			Usage:     "`prefix` for as-you-type prefix matching (default), `exact` for full-token matching.",
			QueryPath: "search_type",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter monitors by lifecycle status.",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "tag",
			Usage:     "Filter to items that have this tag.",
			QueryPath: "tag",
		},
		&requestflag.Flag[any]{
			Name:      "tag",
			Usage:     "Comma-separated list of tags to filter by (matches monitors having any of them).",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:      "target-type",
			Usage:     "Filter by target type.",
			QueryPath: "target_type",
		},
	},
	Action:          handleMonitorsList,
	HideHelpCommand: true,
}

var monitorsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a monitor and stop future runs and webhook retries.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "monitor-id",
			Usage:     "ID of the monitor.",
			Required:  true,
			PathParam: "monitor_id",
		},
	},
	Action:          handleMonitorsDelete,
	HideHelpCommand: true,
}

var monitorsGetCreditUsage = cli.Command{
	Name:    "get-credit-usage",
	Usage:   "Return usage per monitor, highest first, for up to the 10,000 most recent runs\nin the requested window.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "since",
			Usage:     "Only include items at or after this ISO 8601 timestamp.",
			QueryPath: "since",
		},
		&requestflag.Flag[any]{
			Name:      "until",
			Usage:     "Only include items before this ISO 8601 timestamp.",
			QueryPath: "until",
		},
	},
	Action:          handleMonitorsGetCreditUsage,
	HideHelpCommand: true,
}

var monitorsGetLimits = cli.Command{
	Name:            "get-limits",
	Usage:           "Retrieve your organization’s monitor allowance and usage.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMonitorsGetLimits,
	HideHelpCommand: true,
}

var monitorsListAccountChanges = cli.Command{
	Name:    "list-account-changes",
	Usage:   "List full change records across your monitors, newest first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "change-detection-type",
			Usage:     "Filter by change detection type.",
			QueryPath: "change_detection_type",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return per page (1-100). Defaults to 25.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "monitor-id",
			Usage:     "Filter changes to a single monitor.",
			QueryPath: "monitor_id",
		},
		&requestflag.Flag[any]{
			Name:      "since",
			Usage:     "Only include items at or after this ISO 8601 timestamp.",
			QueryPath: "since",
		},
		&requestflag.Flag[string]{
			Name:      "tag",
			Usage:     "Filter to items that have this tag.",
			QueryPath: "tag",
		},
		&requestflag.Flag[string]{
			Name:      "target-type",
			Usage:     "Filter by target type.",
			QueryPath: "target_type",
		},
		&requestflag.Flag[any]{
			Name:      "until",
			Usage:     "Only include items before this ISO 8601 timestamp.",
			QueryPath: "until",
		},
	},
	Action:          handleMonitorsListAccountChanges,
	HideHelpCommand: true,
}

var monitorsListAccountRuns = cli.Command{
	Name:    "list-account-runs",
	Usage:   "List runs across your monitors, newest first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return per page (1-100). Defaults to 25.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter runs by lifecycle status.",
			QueryPath: "status",
		},
	},
	Action:          handleMonitorsListAccountRuns,
	HideHelpCommand: true,
}

var monitorsListChanges = cli.Command{
	Name:    "list-changes",
	Usage:   "List full change records for a monitor, newest first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "monitor-id",
			Usage:     "ID of the monitor.",
			Required:  true,
			PathParam: "monitor_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return per page (1-100). Defaults to 25.",
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "since",
			Usage:     "Only include items at or after this ISO 8601 timestamp.",
			QueryPath: "since",
		},
		&requestflag.Flag[string]{
			Name:      "tag",
			Usage:     "Filter to items that have this tag.",
			QueryPath: "tag",
		},
		&requestflag.Flag[any]{
			Name:      "until",
			Usage:     "Only include items before this ISO 8601 timestamp.",
			QueryPath: "until",
		},
	},
	Action:          handleMonitorsListChanges,
	HideHelpCommand: true,
}

var monitorsListRuns = cli.Command{
	Name:    "list-runs",
	Usage:   "List a monitor’s runs, newest first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "monitor-id",
			Usage:     "ID of the monitor.",
			Required:  true,
			PathParam: "monitor_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return per page (1-100). Defaults to 25.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter runs by lifecycle status.",
			QueryPath: "status",
		},
	},
	Action:          handleMonitorsListRuns,
	HideHelpCommand: true,
}

var monitorsRetrieveChange = cli.Command{
	Name:    "retrieve-change",
	Usage:   "Retrieve a detected change, including its diff and available evidence.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "change-id",
			Usage:     "ID of the detected change.",
			Required:  true,
			PathParam: "change_id",
		},
	},
	Action:          handleMonitorsRetrieveChange,
	HideHelpCommand: true,
}

var monitorsRun = cli.Command{
	Name:    "run",
	Usage:   "Queue a run without changing the regular schedule. Paused monitors return 409.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "monitor-id",
			Usage:     "ID of the monitor.",
			Required:  true,
			PathParam: "monitor_id",
		},
	},
	Action:          handleMonitorsRun,
	HideHelpCommand: true,
}

func handleMonitorsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.MonitorNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.New(ctx, params, options...)
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
		Title:          "monitors create",
		Transform:      transform,
	})
}

func handleMonitorsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("monitor-id") && len(unusedArgs) > 0 {
		cmd.Set("monitor-id", unusedArgs[0])
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
	_, err = client.Monitors.Get(ctx, cmd.Value("monitor-id").(string), options...)
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
		Title:          "monitors retrieve",
		Transform:      transform,
	})
}

func handleMonitorsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("monitor-id") && len(unusedArgs) > 0 {
		cmd.Set("monitor-id", unusedArgs[0])
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

	params := contextdev.MonitorUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.Update(
		ctx,
		cmd.Value("monitor-id").(string),
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
		Title:          "monitors update",
		Transform:      transform,
	})
}

func handleMonitorsList(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.MonitorListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.List(ctx, params, options...)
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
		Title:          "monitors list",
		Transform:      transform,
	})
}

func handleMonitorsDelete(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("monitor-id") && len(unusedArgs) > 0 {
		cmd.Set("monitor-id", unusedArgs[0])
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
	_, err = client.Monitors.Delete(ctx, cmd.Value("monitor-id").(string), options...)
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
		Title:          "monitors delete",
		Transform:      transform,
	})
}

func handleMonitorsGetCreditUsage(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.MonitorGetCreditUsageParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.GetCreditUsage(ctx, params, options...)
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
		Title:          "monitors get-credit-usage",
		Transform:      transform,
	})
}

func handleMonitorsGetLimits(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.GetLimits(ctx, options...)
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
		Title:          "monitors get-limits",
		Transform:      transform,
	})
}

func handleMonitorsListAccountChanges(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.MonitorListAccountChangesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.ListAccountChanges(ctx, params, options...)
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
		Title:          "monitors list-account-changes",
		Transform:      transform,
	})
}

func handleMonitorsListAccountRuns(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.MonitorListAccountRunsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.ListAccountRuns(ctx, params, options...)
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
		Title:          "monitors list-account-runs",
		Transform:      transform,
	})
}

func handleMonitorsListChanges(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("monitor-id") && len(unusedArgs) > 0 {
		cmd.Set("monitor-id", unusedArgs[0])
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

	params := contextdev.MonitorListChangesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.ListChanges(
		ctx,
		cmd.Value("monitor-id").(string),
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
		Title:          "monitors list-changes",
		Transform:      transform,
	})
}

func handleMonitorsListRuns(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("monitor-id") && len(unusedArgs) > 0 {
		cmd.Set("monitor-id", unusedArgs[0])
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

	params := contextdev.MonitorListRunsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Monitors.ListRuns(
		ctx,
		cmd.Value("monitor-id").(string),
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
		Title:          "monitors list-runs",
		Transform:      transform,
	})
}

func handleMonitorsRetrieveChange(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("change-id") && len(unusedArgs) > 0 {
		cmd.Set("change-id", unusedArgs[0])
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
	_, err = client.Monitors.GetChange(ctx, cmd.Value("change-id").(string), options...)
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
		Title:          "monitors retrieve-change",
		Transform:      transform,
	})
}

func handleMonitorsRun(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("monitor-id") && len(unusedArgs) > 0 {
		cmd.Set("monitor-id", unusedArgs[0])
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
	_, err = client.Monitors.Run(ctx, cmd.Value("monitor-id").(string), options...)
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
		Title:          "monitors run",
		Transform:      transform,
	})
}
