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
	Usage:   "Converts raw text, source code, web/data, PDF, Microsoft Office, and image bytes\ninto LLM-usable Markdown.",
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
		&requestflag.Flag[any]{
			Name:      "include-images",
			Usage:     "Include image references in Markdown output",
			Default:   false,
			QueryPath: "includeImages",
		},
		&requestflag.Flag[any]{
			Name:      "include-links",
			Usage:     "Preserve hyperlinks in Markdown output",
			Default:   true,
			QueryPath: "includeLinks",
		},
		&requestflag.Flag[any]{
			Name:      "ocr",
			Usage:     "When true for PDF inputs, detect and OCR images embedded in the selected pages, inserting recognized text at each image's position in page reading order while preserving the PDF text layer. pdf.start/pdf.end limit the inclusive page range. When false, all OCR is disabled, including the automatic scanned-PDF fallback.",
			Default:   false,
			QueryPath: "ocr",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "pdf",
			Usage:     `PDF page-range options as a JSON object, e.g. {"start": 2, "end": 5}.`,
			Default:   map[string]any{},
			QueryPath: "pdf",
		},
		&requestflag.Flag[any]{
			Name:      "shorten-base64-images",
			Usage:     "Shorten base64-encoded image data in the Markdown output",
			Default:   true,
			QueryPath: "shortenBase64Images",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Optional comma-separated caller-defined tags for tracking this request. Tags are recorded on the request's usage log and can be used to filter usage on the dashboard usage page. Up to 20 tags, each 1-50 characters.",
			QueryPath: "tags",
		},
		&requestflag.Flag[any]{
			Name:      "use-main-content-only",
			Usage:     "Extract only the main content from HTML-like inputs",
			Default:   false,
			QueryPath: "useMainContentOnly",
		},
	},
	Action:          handleParseHandle,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pdf": {
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.end",
			Usage:      "Last 1-based PDF page to parse. When omitted, parsing ends at the last page. Must be greater than or equal to start when both are provided.",
			InnerField: "end",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.start",
			Usage:      "First 1-based PDF page to parse. When omitted, parsing starts at the first page.",
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
