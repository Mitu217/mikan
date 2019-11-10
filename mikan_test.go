package mikan

import (
	"reflect"
	"testing"
)

func TestAnalyze(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "japanese",
			args: args{
				str: "常に最新、最高のモバイル。Androidを開発した同じチームから。",
			},
			want: []string{
				"常に", "最新、", "最高の", "モバイル。", "Androidを", "開発した", "同じ", "チームから。",
			},
		},
		{
			name: "hankana",
			args: args{
				str: "ﾊﾛｰﾜｰﾙﾄﾞ",
			},
			want: []string{
				"ﾊﾛｰﾜｰﾙﾄﾞ",
			},
		},
		{
			name: "inner '・'",
			args: args{
				str: "ﾊﾛｰ・ﾜｰﾙﾄﾞ",
			},
			want: []string{
				"ﾊﾛｰ・ﾜｰﾙﾄﾞ",
			},
		},
		{
			name: "english",
			args: args{
				str: "Always the latest and best mobile. From the same team that developed Android.",
			},
			want: []string{
				"Always", " ", "the", " ", "latest", " ", "and", " ", "best", " ", "mobile.", " ", "From", " ", "the", " ", "same", " ", "team", " ", "that", " ", "developed", " ", "Android.",
			},
		},
		{
			name: "french",
			args: args{
				str: "Toujours le dernier et le meilleur mobile. De la même équipe qui a développé Android.",
			},
			want: []string{
				"Toujours", " ", "le", " ", "dernier", " ", "et", " ", "le", " ", "meilleur", " ", "mobile.", " ", "De", " ", "la", " ", "même", " ", "équipe", " ", "qui", " ", "a", " ", "développé", " ", "Android.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Analyze(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
