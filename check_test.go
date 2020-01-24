package main

import (
	"testing"

	"github.com/bannzai/switchecker/pkg/testutil"
)

func Test_check(t *testing.T) {
	type args struct {
		enums    []enum
		filepath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "#01",
			args: args{
				enums: []enum{
					{
						name:        "language",
						packageName: "testdata",
						patterns: []string{
							"golang",
							"swift",
							"objectivec",
							"ruby",
							"typescript",
						}},
				},
				filepath: testutil.CallerDirectoryPath(t) + "/testdata/check.go",
			},
			want: false,
		},
		{
			name: "#02",
			args: args{
				enums: []enum{
					{
						name:        "language",
						packageName: "testdata",
						patterns: []string{
							"golang",
							"swift",
							"objectivec",
							"ruby",
							"typescript",
						}},
				},
				filepath: testutil.CallerDirectoryPath(t) + "/testdata/check2.go",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.enums, tt.args.filepath); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
