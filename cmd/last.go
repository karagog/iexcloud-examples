// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/goinvest/iexcloud-examples/domain"
	iex "github.com/goinvest/iexcloud/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lastCmd)
}

var lastCmd = &cobra.Command{
	Use:   "last [stocks]",
	Short: "Retrieve the last available price from TOPS",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		prices, err := client.Last(context.Background(), args)
		if err != nil {
			log.Fatalf("Error getting account metadata: %s", err)
		}
		b, err := json.MarshalIndent(prices, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling AccountMetadata into JSON: %s", err)
		}
		fmt.Println("## Last Prices ##")
		fmt.Println(string(b))
	},
}
