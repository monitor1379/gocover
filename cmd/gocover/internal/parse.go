package internal

/*
 * @Date: 2021-01-18 17:59:03
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2021-01-18 21:05:44
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

	// pkgs, err := profiles.Packages()
	// if err != nil {
	// 	return err
	// }

	// funcCoverages, err := profiles.FuncLevelPercentageCovered()
	// if err != nil {
	// 	return err
	// }
	// for _, funcCoverage := range funcCoverages {
	// 	fmt.Println(funcCoverage.Filename, funcCoverage.FuncName, funcCoverage.Count, funcCoverage.Total, funcCoverage.Count/funcCoverage.Total)
	// }

	// fileCoverages, err := profiles.FileLevelPercentageCovered()
	// if err != nil {
	// 	return err
	// }
	// for _, fileCoverage := range fileCoverages {
	// 	fmt.Println(fileCoverage.Filename, fileCoverage.Covered, fileCoverage.Total, fileCoverage.Covered/fileCoverage.Total)
	// }

	packageCoverages, err := profiles.PackageLevelPercentageCovered()
	if err != nil {
		return err
	}
	for _, packageCoverage := range packageCoverages {
		fmt.Println(packageCoverage.PackageName, packageCoverage.Covered, packageCoverage.Total, packageCoverage.Percentage())
	}
	return nil
}
