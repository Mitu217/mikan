package main

import (
	"fmt"
	"strings"

	"github.com/Mitu217/mikan"
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
