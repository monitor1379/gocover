package main

/*
 * @Date: 2021-01-18 17:41:42
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2021-01-18 17:55:09
 */

import "github.com/monitor1379/gocover/cmd/gocover/internal"

func main() {
	if err := internal.New().Execute(); err != nil {
		panic(err)
	}
}
