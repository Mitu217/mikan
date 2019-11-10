# mikan.go

[![GoDoc](https://godoc.org/github.com/Mitu217/mikan?status.svg)](https://godoc.org/github.com/Mitu217/mikan)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mitu217/mikan)](https://goreportcard.com/report/github.com/Mitu217/mikan)

mikan.go is [mikan.js](https://github.com/trkbt10/mikan.js) reimplemented with golang

## example

```golang
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
```