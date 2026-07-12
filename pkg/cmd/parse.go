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
	Usage:   "Converts raw text, source code, web/data, PDF, Microsoft Office, and image bytes\ninto LLM-usable Markdown. The base request costs 1 credit. When OCR runs\n(requires ocr=true), the entire call costs 5 credits; ocr=true requests where no\nOCR ends up running still cost 1 credit.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "body",
			Required:  true,
			BodyRoot:  true,
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:      "extension",
			Usage:     `Optional file extension hint. Case-insensitive; a leading dot is accepted (e.g. ".pdf").`,
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
			Usage:     "Gates all OCR. When true, PDFs get embedded-image OCR (recognized text inserted at each image's position in page reading order, preserving the text layer; pdf.start/pdf.end limit the page range), scanned PDFs with no text layer get full-document OCR, and raster images get their visible text transcribed. When false, no OCR runs: scanned PDFs may yield no content and images return only format/dimension metadata. Calls where OCR actually runs cost 5 credits instead of 1.",
			Default:   false,
			QueryPath: "ocr",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "pdf",
			Usage:     "PDF page-range controls. Use start/end to limit parsing (and OCR when ocr=true) to an inclusive 1-based page range.",
			Default:   map[string]any{},
			QueryPath: "pdf",
		},
		&requestflag.Flag[bool]{
			Name:      "shorten-base64-images",
			Usage:     "Shorten base64-encoded image data in the Markdown output",
			Default:   true,
			QueryPath: "shortenBase64Images",
		},
		&requestflag.Flag[bool]{
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
