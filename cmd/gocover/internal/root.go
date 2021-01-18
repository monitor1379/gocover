package internal

/*
 * @Date: 2021-01-18 17:45:38
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2021-01-18 20:07:28
 */

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gocover",
		Short: "gocover",
		Long:  "gocover",
	}

	rootCmd.PersistentFlags().StringP("cover", "c", "cover.out", "cover profile file path")

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newParseCmd())

	return rootCmd
}
