// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/hyperspell-cli/internal/apiquery"
	"github.com/stainless-sdks/hyperspell-cli/internal/requestflag"
	"github.com/stainless-sdks/hyperspell-go"
	"github.com/stainless-sdks/hyperspell-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var sessionsAdd = cli.Command{
	Name:    "add",
	Usage:   "Add an agent trace/transcript to the index.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "history",
			Usage:    "The trace history as a string. Can be a JSON array of Hyperdoc steps, a JSON array of Vercel AI SDK steps, or OpenClaw JSONL.",
			Required: true,
			BodyPath: "history",
		},
		&requestflag.Flag[any]{
			Name:     "date",
			Usage:    "Date of the trace",
			BodyPath: "date",
		},
		&requestflag.Flag[[]string]{
			Name:     "extract",
			Usage:    "What kind of memories to extract from the trace",
			Default:  []string{"procedure"},
			BodyPath: "extract",
		},
		&requestflag.Flag[any]{
			Name:     "format",
			Usage:    "Trace format: 'vercel', 'hyperdoc', or 'openclaw'. Auto-detected if not set.",
			BodyPath: "format",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Custom metadata for filtering. Keys must be alphanumeric with underscores, max 64 chars.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "session-id",
			Usage:    "Resource identifier for the trace.",
			BodyPath: "session_id",
		},
		&requestflag.Flag[any]{
			Name:     "title",
			Usage:    "Title of the trace",
			BodyPath: "title",
		},
	},
	Action:          handleSessionsAdd,
	HideHelpCommand: true,
}

func handleSessionsAdd(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.SessionAddParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Add(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions add", obj, format, transform)
}
