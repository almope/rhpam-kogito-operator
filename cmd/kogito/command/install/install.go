// Copyright 2019 Red Hat, Inc. and/or its affiliates
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

package install

import (
	"github.com/kiegroup/kogito-cloud-operator/cmd/kogito/command/context"
	"github.com/spf13/cobra"
)

type installCommand struct {
	context.CommandContext
	command *cobra.Command
	Parent  *cobra.Command
}

func newInstallCommand(ctx *context.CommandContext, parent *cobra.Command) context.KogitoCommand {
	cmd := installCommand{
		CommandContext: *ctx,
		Parent:         parent,
	}
	cmd.RegisterHook()
	cmd.InitHook()
	return &cmd
}

func (i *installCommand) Command() *cobra.Command {
	return i.command
}

func (i *installCommand) RegisterHook() {
	i.command = &cobra.Command{
		Use:    "install",
		Short:  "Install all sorts of infrastructure components to your Kogito project",
		PreRun: i.CommonPreRun,
	}
}

func (i *installCommand) InitHook() {
	i.Parent.AddCommand(i.command)
}
