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

var utilityPrefetch = requestflag.WithInnerFlags(cli.Command{
	Name:    "prefetch",
	Usage:   "Signal that you may fetch data soon to improve latency. The type field selects\nwhat to prefetch ('brand' queues a brand data fetch, 'styleguide' queues a\nstyleguide extraction) and identifier carries exactly one lookup key: a domain,\nor an email whose domain is extracted and validated (free email providers and\ndisposable email addresses are not allowed).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "identifier",
			Usage:    "Identifier of the target to prefetch. Provide exactly one of domain or email.",
			Required: true,
			BodyPath: "identifier",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "What to prefetch: 'brand' warms the brand data cache, 'styleguide' warms the styleguide cache.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Optional request deadline and behavior on timeout. For GET requests, use timeoutOpts[milliseconds]=30000&timeoutOpts[behavior]=fail or a JSON-encoded timeoutOpts object.",
			BodyPath: "timeoutOpts",
		},
	},
	Action:          handleUtilityPrefetch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Request deadline in milliseconds. Maximum: 300000 (5 minutes).",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      `What to do at the deadline. This endpoint supports "fail": return 408 REQUEST_TIMEOUT without charging credits.`,
			InnerField: "behavior",
		},
	},
})

func handleUtilityPrefetch(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.UtilityPrefetchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Utility.Prefetch(ctx, params, options...)
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
		Title:          "utility prefetch",
		Transform:      transform,
	})
}
