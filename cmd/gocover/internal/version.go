package internal

/*
 * @Date: 2021-01-18 17:47:01
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2021-01-18 17:55:43
 */

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v0.1.0"
)

func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version of gocover",
		RunE:  versionCmdRunE,
	}
	return versionCmd
}

func versionCmdRunE(c *cobra.Command, args []string) error {
	fmt.Printf("gocover version %s\n", version)
	return nil
}
