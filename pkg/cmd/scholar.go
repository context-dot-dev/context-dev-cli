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

var scholarSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Search full-text papers from arXiv and Europe PMC. Returns up to 250 relevant\npapers with matching passages and full-text links.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "A topic or question. Append one arxiv:, pmcid:, pmid:, or doi: ID to search within an indexed paper.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[*string]{
			Name:     "cursor",
			Usage:    "Use next_cursor to fetch the next page.",
			BodyPath: "cursor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter-by",
			Usage:    "Filter by publication date.",
			BodyPath: "filterBy",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Papers per page.",
			Default:  10,
			BodyPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:     "sort",
			Usage:    "Order the top 250 matches. date_desc puts newest first, with undated papers last.",
			Default:  "relevance",
			BodyPath: "sort",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for filtering usage in the dashboard.",
			BodyPath: "tags",
		},
	},
	Action:          handleScholarSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter-by": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter-by.date",
			Usage:      "Publication date range. Undated papers are excluded.",
			InnerField: "date",
		},
	},
})

func handleScholarSearch(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.ScholarSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Scholar.Search(ctx, params, options...)
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
		Title:          "scholar search",
		Transform:      transform,
	})
}
