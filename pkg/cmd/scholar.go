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
	Usage:   "Search the full text of papers from arXiv and Europe PMC by topic or question.\nEach result includes the matching passage and a direct link to the full text.\nPage through up to 250 matches, ranked by relevance or newest first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Topic or question to search for. Append one `arxiv:`, `pmcid:`, `pmid:`, or `doi:` ID to search within that paper, such as `protein folding pmid:34265844`, or send the ID on its own.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[*string]{
			Name:     "cursor",
			Usage:    "Cursor from the previous page's `next_cursor`. Omit for the first page.",
			BodyPath: "cursor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter-by",
			Usage:    "Filter by publication date.",
			BodyPath: "filterBy",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Maximum number of papers to return per page.",
			Default:  10,
			BodyPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:     "sort",
			Usage:    "`relevance` puts the closest matches first; `date_desc` puts the newest papers first, with undated papers last. Both orders apply to the top 250 matches.",
			Default:  "relevance",
			BodyPath: "sort",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for this request. Filter usage and request logs by them in the dashboard or with `GET /org/logs`.",
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
