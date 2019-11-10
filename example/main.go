package main

import (
	"fmt"
	"strings"

	mikan "github.com/mitu217/mikan.go"
)

func main() {
	lines := mikan.Mikan("常に最新、最高のモバイル。<Android>を開発した同じチームから。")

	fmt.Println(strings.Join(lines, "\n"))
	/*
		常に最新、最高のモバイル。
		<Android>を開発した同じ
		チームから。
	*/
}
