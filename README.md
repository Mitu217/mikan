# mikan.go

mikan.go is [mikan.js](https://github.com/trkbt10/mikan.js) reimplemented with golang

## example

```golang
func main() {
	lines := mikan.Mikan("常に最新、最高のモバイル。<Android>を開発した同じチームから。")

	fmt.Println(strings.Join(lines, "\n"))
	/*
		常に最新、最高のモバイル。
		<Android>を開発した同じ
		チームから。
	*/
}
```