package main

import (
	"fmt"
	"strings"

	"github.com/Mitu217/mikan"
)

func main() {
	mikan := mikan.NewMikan(
		mikan.RuneWidth(30),
	)
	lines := mikan.Split("常に最新、最高のモバイル。Androidを開発した同じチームから。")
	fmt.Println(strings.Join(lines, "\n"))
	/*
		常に最新、最高のモバイル。
		<Android>を開発した同じ
		チームから。
	*/
}
