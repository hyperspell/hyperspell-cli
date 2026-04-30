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

var foldersList = cli.Command{
	Name:    "list",
	Usage:   "List one level of folders from the user's connected source.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
		&requestflag.Flag[*string]{
			Name:      "parent-id",
			Usage:     "Parent folder ID. Omit for root-level folders.",
			QueryPath: "parent_id",
		},
	},
	Action:          handleFoldersList,
	HideHelpCommand: true,
}

var foldersDeletePolicy = cli.Command{
	Name:    "delete-policy",
	Usage:   "Delete a folder policy for a specific connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "policy-id",
			Required: true,
		},
	},
	Action:          handleFoldersDeletePolicy,
	HideHelpCommand: true,
}

var foldersListPolicies = cli.Command{
	Name:    "list-policies",
	Usage:   "List all folder policies for a specific connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
	},
	Action:          handleFoldersListPolicies,
	HideHelpCommand: true,
}

var foldersSetPolicies = cli.Command{
	Name:    "set-policies",
	Usage:   "Create or update a folder policy for a specific connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "provider-folder-id",
			Usage:    "Folder ID from the source provider",
			Required: true,
			BodyPath: "provider_folder_id",
		},
		&requestflag.Flag[string]{
			Name:     "sync-mode",
			Usage:    "Sync mode for this folder",
			Required: true,
			BodyPath: "sync_mode",
		},
		&requestflag.Flag[*string]{
			Name:     "folder-name",
			Usage:    "Display name of the folder",
			BodyPath: "folder_name",
		},
		&requestflag.Flag[*string]{
			Name:     "folder-path",
			Usage:    "Display path of the folder",
			BodyPath: "folder_path",
		},
		&requestflag.Flag[*string]{
			Name:     "parent-folder-id",
			Usage:    "Parent folder's provider ID for inheritance resolution",
			BodyPath: "parent_folder_id",
		},
	},
	Action:          handleFoldersSetPolicies,
	HideHelpCommand: true,
}

func handleFoldersList(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.FolderListParams{}

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
	_, err = client.Folders.List(
		ctx,
		cmd.Value("connection-id").(string),
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
		Title:          "folders list",
		Transform:      transform,
	})
}

func handleFoldersDeletePolicy(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("policy-id") && len(unusedArgs) > 0 {
		cmd.Set("policy-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.FolderDeletePolicyParams{
		ConnectionID: cmd.Value("connection-id").(string),
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
	_, err = client.Folders.DeletePolicy(
		ctx,
		cmd.Value("policy-id").(string),
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
		Title:          "folders delete-policy",
		Transform:      transform,
	})
}

func handleFoldersListPolicies(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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
	_, err = client.Folders.ListPolicies(ctx, cmd.Value("connection-id").(string), options...)
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
		Title:          "folders list-policies",
		Transform:      transform,
	})
}

func handleFoldersSetPolicies(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.FolderSetPoliciesParams{}

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
	_, err = client.Folders.SetPolicies(
		ctx,
		cmd.Value("connection-id").(string),
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
		Title:          "folders set-policies",
		Transform:      transform,
	})
}
