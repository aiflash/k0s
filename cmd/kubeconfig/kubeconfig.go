/*
Copyright 2021 k0s authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubeconfig

import (
	"github.com/k0sproject/k0s/cmd/internal"
	"github.com/k0sproject/k0s/pkg/config"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewKubeConfigCmd() *cobra.Command {
	var debugFlags internal.DebugFlags

	cmd := &cobra.Command{
		Use:              "kubeconfig [command]",
		Short:            "Create a kubeconfig file for a specified user",
		Args:             cobra.NoArgs,
		PersistentPreRun: debugFlags.Run,
		RunE:             func(*cobra.Command, []string) error { return pflag.ErrHelp }, // Enforce arg validation
	}

	pflags := cmd.PersistentFlags()
	debugFlags.AddToFlagSet(pflags)
	config.GetPersistentFlagSet().VisitAll(func(f *pflag.Flag) {
		f.Hidden = true
		f.Deprecated = "it has no effect and will be removed in a future release"
		pflags.AddFlag(f)
	})

	cmd.AddCommand(kubeconfigCreateCmd())
	cmd.AddCommand(kubeConfigAdminCmd())

	return cmd
}
