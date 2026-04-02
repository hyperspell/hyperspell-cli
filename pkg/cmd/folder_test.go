// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/hyperspell-cli/internal/mocktest"
)

func TestFoldersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"folders", "list",
			"--connection-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--parent-id", "parent_id",
		)
	})
}

func TestFoldersDeletePolicy(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"folders", "delete-policy",
			"--connection-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--policy-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFoldersListPolicies(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"folders", "list-policies",
			"--connection-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFoldersSetPolicies(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"folders", "set-policies",
			"--connection-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--provider-folder-id", "provider_folder_id",
			"--sync-mode", "sync",
			"--folder-name", "folder_name",
			"--folder-path", "folder_path",
			"--parent-folder-id", "parent_folder_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"provider_folder_id: provider_folder_id\n" +
			"sync_mode: sync\n" +
			"folder_name: folder_name\n" +
			"folder_path: folder_path\n" +
			"parent_folder_id: parent_folder_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"folders", "set-policies",
			"--connection-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
