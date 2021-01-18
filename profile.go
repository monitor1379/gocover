/*
 * @Date: 2021-01-18 17:42:17
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2021-01-18 18:13:47
 */
package gocover

import (
	"encoding/json"
	"fmt"

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

	k, err := findPkgs(*profiles)
	if err != nil {
		return nil, err
	}

	data, _ := json.MarshalIndent(k, "", "  ")
	fmt.Println(string(data))

	return *profiles, nil
}

// Percentage Ref: https://github.com/golang/go/blob/0e85fd7561de869add933801c531bf25dee9561c/src/cmd/cover/html.go#L96-L108
func (p *Profiles) Percentage() float64 {
	return 0.0
}
