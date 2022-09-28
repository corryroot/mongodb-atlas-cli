// Copyright 2021 MongoDB Inc
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

package availableregions

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store    store.CloudProviderRegionsLister
	provider string
	tier     string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `PROVIDER	INSTANCE SIZE	REGIONS{{range .Results}}
{{.Provider}}{{range .InstanceSizes}}	{{.Name}}{{range .AvailableRegions}}	{{.Name}}{{end}}{{end}}{{end}}
`

func (opts *ListOpts) Run() error {
	r, err := opts.store.CloudProviderRegions(opts.ConfigProjectID(), opts.tier, []*string{&opts.provider})
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas cluster(s) availableRegions list --provider provider --tier tier --projectId projectId.
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List available regions for your project.",
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVar(&opts.tier, flag.Tier, "", usage.Tier)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Tier)
	_ = cmd.MarkFlagRequired(flag.Provider)

	return cmd
}