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

var parseHandle = cli.Command{
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
			Name:      "base-url",
			Usage:     "Optional HTTP(S) source document URL used to resolve relative links and image references. Relative references remain relative when omitted.",
			QueryPath: "baseUrl",
		},
		&requestflag.Flag[string]{
			Name:      "extension",
			Usage:     "Optional file extension hint, such as pdf, docx, xlsx, pptx, html, json, csv, md, py, rtf, jpg, png, or txt.",
			QueryPath: "extension",
		},
		&requestflag.Flag[string]{
			Name:      "filename",
			Usage:     "Optional filename hint used to infer the extension when extension is omitted.",
			QueryPath: "filename",
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
			Usage:     "When true for PDF inputs, detect and OCR images embedded in the selected pages, inserting recognized text at each image's position in page reading order while preserving the PDF text layer. pdfStart/pdfEnd limit the inclusive page range. This is separate from automatic scanned-PDF OCR fallback.",
			Default:   false,
			QueryPath: "ocr",
		},
		&requestflag.Flag[int64]{
			Name:      "pdf-end",
			Usage:     "Last 1-based PDF page to parse. When omitted, parsing ends at the last page. Must be greater than or equal to pdfStart when both are provided.",
			QueryPath: "pdfEnd",
		},
		&requestflag.Flag[int64]{
			Name:      "pdf-start",
			Usage:     "First 1-based PDF page to parse. When omitted, parsing starts at the first page.",
			QueryPath: "pdfStart",
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
}

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
