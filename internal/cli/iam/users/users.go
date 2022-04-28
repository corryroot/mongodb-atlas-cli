// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package users

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	description := "Create, list and manage your Cloud Manager or Ops Manager users."
	if config.ToolName == config.AtlasCLI {
		description = "Create, list and manage your Atlas users."
	}

	const use = "users"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Users operations.",
		Long:    description,
		Aliases: cli.GenerateAliases(use),
	}

	cmd.AddCommand(
		InviteBuilder(),
		DescribeBuilder(),
	)

	if config.ToolName == config.MongoCLI {
		cmd.AddCommand(DeleteBuilder())
	}

	return cmd
}