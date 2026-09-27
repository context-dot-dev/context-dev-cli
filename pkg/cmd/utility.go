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
	Usage:   "Queue brand or styleguide data so a later lookup can return sooner.",
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
			Usage:    "Data to prefetch.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for filtering usage in the dashboard.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Request deadline and what to return when it passes.",
			BodyPath: "timeoutOpts",
		},
	},
	Action:          handleUtilityPrefetch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      `Only "fail" is supported: return 408 at the deadline.`,
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
