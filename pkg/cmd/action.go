// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/hyperspell/hyperspell-cli/internal/apiquery"
	"github.com/hyperspell/hyperspell-cli/internal/requestflag"
	"github.com/hyperspell/hyperspell-go"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var actionsAddReaction = cli.Command{
	Name:    "add-reaction",
	Usage:   "Add an emoji reaction to a message on a connected integration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "channel",
			Usage:    "Channel ID containing the message",
			Required: true,
			BodyPath: "channel",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Emoji name without colons (e.g., thumbsup)",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "Integration provider (e.g., slack)",
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "timestamp",
			Usage:    "Message timestamp to react to",
			Required: true,
			BodyPath: "timestamp",
		},
		&requestflag.Flag[*string]{
			Name:     "connection",
			Usage:    "Connection ID. If omitted, auto-resolved from provider + user.",
			BodyPath: "connection",
		},
	},
	Action:          handleActionsAddReaction,
	HideHelpCommand: true,
}

var actionsSendMessage = cli.Command{
	Name:    "send-message",
	Usage:   "Send a message to a channel or conversation on a connected integration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "Integration provider (e.g., slack)",
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Message text",
			Required: true,
			BodyPath: "text",
		},
		&requestflag.Flag[*string]{
			Name:     "channel",
			Usage:    "Channel ID (required for Slack)",
			BodyPath: "channel",
		},
		&requestflag.Flag[*string]{
			Name:     "connection",
			Usage:    "Connection ID. If omitted, auto-resolved from provider + user.",
			BodyPath: "connection",
		},
		&requestflag.Flag[*string]{
			Name:     "parent",
			Usage:    "Parent message ID for threading (thread_ts for Slack)",
			BodyPath: "parent",
		},
	},
	Action:          handleActionsSendMessage,
	HideHelpCommand: true,
}

func handleActionsAddReaction(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := hyperspell.ActionAddReactionParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Actions.AddReaction(ctx, params, options...)
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
		Title:          "actions add-reaction",
		Transform:      transform,
	})
}

func handleActionsSendMessage(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := hyperspell.ActionSendMessageParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Actions.SendMessage(ctx, params, options...)
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
		Title:          "actions send-message",
		Transform:      transform,
	})
}
