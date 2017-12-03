// Copyright 2017 Prospect One https://prospectone.io/. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ProspectOne/perfops-cli/cmd/internal"
	"github.com/ProspectOne/perfops-cli/perfops"
)

var (
	creditsCmd = &cobra.Command{
		Use:     "credits",
		Short:   "Displays the remaing credits",
		Long:    `Displays the remaing credits.`,
		Example: `perfops remaining-credits`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newPerfOpsClient()
			if err != nil {
				return err
			}
			return chkRunError(runCredits(c))
		},
	}
)

func initCreditsCmd(parentCmd *cobra.Command) {
	parentCmd.AddCommand(creditsCmd)
}

func runCredits(c *perfops.Client) error {
	ctx := context.Background()

	spinner := internal.NewSpinner()
	fmt.Println("")
	spinner.Start()

	credits, err := c.DNS.RemainingCredits(ctx)
	spinner.Stop()
	if err != nil {
		return err
	}

	fmt.Printf("Remaining credits: %v\n", credits)
	return nil
}
