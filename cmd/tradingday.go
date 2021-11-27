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
	rootCmd.AddCommand(tradingDayCmd)
}

var tradingDayCmd = &cobra.Command{
	Use:   "tradingday [next/last]",
	Short: "Get information about the next (or last) trading day",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		var dt iex.TradeHolidayDate
		switch args[0] {
		case "next":
			dt, err = client.NextTradingDay(context.Background())
		case "last":
			dt, err = client.PreviousTradingDay(context.Background())
		}
		if err != nil {
			log.Fatalf("Error getting quote: %s", err)
		}
		b, err := json.MarshalIndent(&dt, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Next Trading day ##")
		fmt.Println(string(b))
	},
}
