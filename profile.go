/*
 * @Date: 2021-01-18 17:42:17
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2021-01-18 21:05:33
 */
package gocover

import (
	"path"

	"golang.org/x/tools/cover"
)

type Profiles []*cover.Profile

func (ps *Profiles) Add(p *cover.Profile) {
	*ps = append(*ps, p)
}

func ParseProfiles(fileName string) (Profiles, error) {
	cProfiles, err := cover.ParseProfiles(fileName)
	if err != nil {
		return nil, err
	}
	profiles := new(Profiles)
	for _, cProfile := range cProfiles {
		profiles.Add(cProfile)
	}

	return *profiles, nil
}

func (ps *Profiles) Packages() ([]*Pkg, error) {
	namesToPkgs, err := findPkgs(*ps)
	if err != nil {
		return nil, err
	}
	pkgs := make([]*Pkg, 0)
	for _, pkg := range namesToPkgs {
		pkgs = append(pkgs, pkg)
	}
	return pkgs, nil
}

type PackageCoverage struct {
	PackageName string
	Pkg         *Pkg
	Profile     *cover.Profile
	Covered     int
	Total       int
}

func (ps *Profiles) PackageLevelPercentageCovered() ([]*PackageCoverage, error) {
	namesToPkgs, err := findPkgs(*ps)
	if err != nil {
		return nil, err
	}

	packageNameToPkgCov := make(map[string]*PackageCoverage)
	packageCoverages := make([]*PackageCoverage, 0)
	for _, profile := range *ps {
		filename := profile.FileName
		_, err := findFile(namesToPkgs, filename)
		if err != nil {
			return nil, err
		}
		var total, covered int64
		for _, blcok := range profile.Blocks {
			total += int64(blcok.NumStmt)
			if blcok.Count > 0 {
				covered += int64(blcok.NumStmt)
			}
		}
		pkgName := path.Dir(filename)
		pkgCov := packageNameToPkgCov[pkgName]
		if pkgCov == nil {
			pkgCov = &PackageCoverage{
				PackageName: pkgName,
				Pkg:         nil,
				Profile:     profile,
				Covered:     int(covered),
				Total:       int(total),
			}
			packageNameToPkgCov[pkgName] = pkgCov
			packageCoverages = append(packageCoverages, pkgCov)
		} else {
			packageNameToPkgCov[pkgName].Covered += int(covered)
			packageNameToPkgCov[pkgName].Total += int(total)
		}
	}
	return packageCoverages, nil
}

func (p *PackageCoverage) Percentage() float64 {
	total := p.Total
	if p.Total == 0 {
		total = 1 // Avoid zero denominator.
	}
	return 100.0 * float64(p.Covered) / float64(total)
}

type FileCoverage struct {
	Filename string
	Pkg      *Pkg
	Profile  *cover.Profile
	Covered  int
	Total    int
}

func (ps *Profiles) FileLevelPercentageCovered() ([]*FileCoverage, error) {
	namesToPkgs, err := findPkgs(*ps)
	if err != nil {
		return nil, err
	}

	fileCoverages := make([]*FileCoverage, 0)
	for _, profile := range *ps {
		filename := profile.FileName
		filePath, err := findFile(namesToPkgs, filename)
		if err != nil {
			return nil, err
		}
		_ = filePath
		var total, covered int64
		for _, block := range profile.Blocks {
			total += int64(block.NumStmt)
			if block.Count > 0 {
				covered += int64(block.NumStmt)
			}
		}
		fileCoverages = append(fileCoverages, &FileCoverage{
			Filename: filename,
			Pkg:      nil,
			Profile:  profile,
			Covered:  int(covered),
			Total:    int(total),
		})
	}

	return fileCoverages, nil
}

func (f *FileCoverage) Percentage() float64 {
	total := f.Total
	if f.Total == 0 {
		total = 1 // Avoid zero denominator.
	}
	return 100.0 * float64(f.Covered) / float64(total)
}

type FuncCoverage struct {
	FuncName  string
	StartLine int
	Filename  string
	Pkg       *Pkg
	Profile   *cover.Profile
	Covered   int
	Total     int
}

func (ps *Profiles) FuncLevelPercentageCovered() ([]*FuncCoverage, error) {
	namesToPkgs, err := findPkgs(*ps)
	if err != nil {
		return nil, err
	}

	var total, covered int64
	funcCoverages := make([]*FuncCoverage, 0)
	for _, profile := range *ps {
		filename := profile.FileName
		filePath, err := findFile(namesToPkgs, filename)
		if err != nil {
			return nil, err
		}
		funcs, err := findFuncs(filePath)
		if err != nil {
			return nil, err
		}
		// Now match up functions and profile blocks.
		for _, f := range funcs {
			c, t := f.coverage(profile)
			funcCoverages = append(funcCoverages, &FuncCoverage{
				FuncName:  f.name,
				StartLine: f.startLine,
				Pkg:       nil,
				Filename:  filename,
				Profile:   profile,
				Covered:   int(c),
				Total:     int(t),
			})
			total += t
			covered += c
		}
	}
	return funcCoverages, nil
}

func (f *FuncCoverage) Percentage() float64 {
	total := f.Total
	if f.Total == 0 {
		total = 1 // Avoid zero denominator.
	}
	return 100.0 * float64(f.Covered) / float64(total)
}
