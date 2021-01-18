package internal

/*
 * @Date: 2021-01-18 17:59:03
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2021-01-18 21:45:21
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

	parseCmd.PersistentFlags().Bool("package", false, "package level")
	parseCmd.PersistentFlags().Bool("file", false, "file level")
	parseCmd.PersistentFlags().Bool("func", false, "function level")

	return parseCmd
}

func parseCmdRunE(c *cobra.Command, args []string) error {
	filePath, err := c.Parent().PersistentFlags().GetString("cover")
	if err != nil {
		return err
	}
	profiles, err := gocover.ParseProfiles(filePath)
	if err != nil {
		return err
	}

	packageLevel, err := c.PersistentFlags().GetBool("package")
	if err != nil {
		return err
	}
	fileLevel, err := c.PersistentFlags().GetBool("file")
	if err != nil {
		return err
	}
	funcLevel, err := c.PersistentFlags().GetBool("func")
	if err != nil {
		return err
	}
	if !packageLevel && !fileLevel && !funcLevel {
		fileLevel = true
	}

	// ----------------------------------------------------------------
	if packageLevel {
		packageCoverages, err := profiles.PackageLevelPercentageCovered()
		if err != nil {
			return err
		}
		for _, packageCoverage := range packageCoverages {
			fmt.Printf("%s\t%.2f%%\n", packageCoverage.PackageName, packageCoverage.Percentage())
		}
	}

	// ----------------------------------------------------------------
	if fileLevel {
		fileCoverages, err := profiles.FileLevelPercentageCovered()
		if err != nil {
			return err
		}
		for _, fileCoverage := range fileCoverages {
			fmt.Printf("%s\t%.2f%%\n", fileCoverage.Filename, fileCoverage.Percentage())
		}
	}

	// ----------------------------------------------------------------
	if funcLevel {
		funcCoverages, err := profiles.FuncLevelPercentageCovered()
		if err != nil {
			return err
		}
		for _, funcCoverage := range funcCoverages {
			fmt.Printf("%s\t%s\t%.2f%%\n", funcCoverage.Filename, funcCoverage.FuncName, funcCoverage.Percentage())
		}
	}

	return nil
}
