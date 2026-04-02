// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hyperspell/hyperspell-cli/internal/apiquery"
	"github.com/hyperspell/hyperspell-cli/internal/requestflag"
	"github.com/hyperspell/hyperspell-go"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var authDeleteUser = cli.Command{
	Name:            "delete-user",
	Usage:           "Endpoint to delete user.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAuthDeleteUser,
	HideHelpCommand: true,
}

var authMe = cli.Command{
	Name:            "me",
	Usage:           "Endpoint to get basic user data.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAuthMe,
	HideHelpCommand: true,
}

var authUserToken = cli.Command{
	Name:    "user-token",
	Usage:   "Use this endpoint to create a user token for a specific user. This token can be\nsafely passed to your user-facing front-end.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
			BodyPath: "user_id",
		},
		&requestflag.Flag[any]{
			Name:     "expires-in",
			Usage:    "Token lifetime, e.g., '30m', '2h', '1d'. Defaults to 24 hours if not provided.",
			BodyPath: "expires_in",
		},
		&requestflag.Flag[any]{
			Name:     "origin",
			Usage:    "Origin of the request, used for CSRF protection. If set, the token will only be valid for requests originating from this origin.",
			BodyPath: "origin",
		},
	},
	Action:          handleAuthUserToken,
	HideHelpCommand: true,
}

func handleAuthDeleteUser(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Auth.DeleteUser(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "auth delete-user", obj, format, transform)
}

func handleAuthMe(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Auth.Me(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "auth me", obj, format, transform)
}

func handleAuthUserToken(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.AuthUserTokenParams{}

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
	_, err = client.Auth.UserToken(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "auth user-token", obj, format, transform)
}
