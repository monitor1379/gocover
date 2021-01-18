package internal

/*
 * @Date: 2021-01-18 17:59:03
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2021-01-18 19:22:06
 */

import (
	"fmt"

	"github.com/monitor1379/gocover"

	"github.com/spf13/cobra"
)

func newParseCmd() *cobra.Command {
	parseCmd := &cobra.Command{
		Use:   "parse",
		Short: "Parse golang cover profile file",
		Long:  "parse",
		RunE:  parseCmdRunE,
	}

	parseCmd.PersistentFlags().Bool("package")

	return parseCmd
}

func parseCmdRunE(c *cobra.Command, args []string) error {
	filePath, err := c.Parent().PersistentFlags().GetString("file")
	if err != nil {
		return err
	}
	profiles, err := gocover.ParseProfiles(filePath)
	if err != nil {
		return err
	}

	pkgs, err := profiles.Packages()
	if err != nil {
		return err
	}

	fmt.Println(pkgs)

	return nil
}
