// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hyperspell/hyperspell-go"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/stainless-sdks/hyperspell-cli/internal/apiquery"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var integrationsGoogleCalendarList = cli.Command{
	Name:            "list",
	Usage:           "List available calendars for a user. This can be used to ie. populate a dropdown\nfor the user to select a calendar.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleIntegrationsGoogleCalendarList,
	HideHelpCommand: true,
}

func handleIntegrationsGoogleCalendarList(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Integrations.GoogleCalendar.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "integrations:google-calendar list", obj, format, transform)
}
