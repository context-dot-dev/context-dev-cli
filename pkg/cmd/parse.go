// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/context-dot-dev/context-dev-cli/internal/apiquery"
	"github.com/context-dot-dev/context-dev-cli/internal/binaryparam"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
	"github.com/context-dot-dev/context-go-sdk/v2"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var parseHandle = requestflag.WithInnerFlags(cli.Command{
	Name:    "handle",
	Usage:   "Convert uploaded file bytes into Markdown and optional HTML.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "body",
			Required:  true,
			BodyRoot:  true,
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:      "client",
			Usage:     "Optional client identifier used for usage attribution.",
			QueryPath: "client",
		},
		&requestflag.Flag[string]{
			Name:      "extension",
			Usage:     "Optional file extension hint, such as pdf, docx, xlsx, pptx, html, json, csv, md, py, rtf, jpg, png, or txt.",
			QueryPath: "extension",
		},
		&requestflag.Flag[bool]{
			Name:      "include-images",
			Usage:     "Include image references in Markdown output",
			Default:   false,
			QueryPath: "includeImages",
		},
		&requestflag.Flag[bool]{
			Name:      "include-links",
			Usage:     "Preserve hyperlinks in Markdown output",
			Default:   true,
			QueryPath: "includeLinks",
		},
		&requestflag.Flag[bool]{
			Name:      "ocr",
			Usage:     "Read text from images and scanned PDF pages. PDF page ranges still apply.",
			Default:   false,
			QueryPath: "ocr",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "pdf",
			Usage:     `PDF page-range options as a JSON object, e.g. {"start": 2, "end": 5}.`,
			QueryPath: "pdf",
		},
		&requestflag.Flag[bool]{
			Name:      "shorten-base64-images",
			Usage:     "Shorten base64-encoded image data in the Markdown output",
			Default:   true,
			QueryPath: "shortenBase64Images",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for filtering usage, e.g. `production,team-alpha`.",
			QueryPath: "tags",
		},
		&requestflag.Flag[bool]{
			Name:      "use-main-content-only",
			Usage:     "Extract only the main content from HTML-like inputs",
			Default:   false,
			QueryPath: "useMainContentOnly",
		},
		&requestflag.Flag[string]{
			Name:      "zdr",
			Usage:     "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			Default:   "disabled",
			QueryPath: "zdr",
		},
	},
	Action:          handleParseHandle,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pdf": {
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.end",
			Usage:      "Last PDF page to parse (1-based, inclusive). Defaults to the final page. Must be >= start.",
			InnerField: "end",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.start",
			Usage:      "First 1-based PDF page to parse.",
			InnerField: "start",
		},
	},
})

func handleParseHandle(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("body") && len(unusedArgs) > 0 {
		cmd.Set("body", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	bodyReader, stdinInUse, err := binaryparam.FileOrStdin(os.Stdin, cmd.Value("body").(string))
	if err != nil {
		return fmt.Errorf("Failed on param '%s': %w", "body", err)
	}
	defer bodyReader.Close()

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationOctetStream,
		stdinInUse,
	)
	if err != nil {
		return err
	}

	params := contextdev.ParseHandleParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Parse.Handle(
		ctx,
		bodyReader,
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
		Title:          "parse handle",
		Transform:      transform,
	})
}
