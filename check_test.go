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
			name: "check_not_enum_case_function",
			args: args{
				enums: []enum{
					{
						name:        "language1",
						packageName: "testdata",
						patterns: []string{
							"golang1",
							"swift1",
							"objectivec1",
							"ruby1",
							"typescript1",
						}},
				},
				filepath: testutil.CallerDirectoryPath(t) + "/testdata/check_not_enum_case_function.go",
			},
			want: false,
		},
		{
			name: "check_wrote_all_case_function",
			args: args{
				enums: []enum{
					{
						name:        "language2",
						packageName: "testdata",
						patterns: []string{
							"golang2",
							"swift2",
							"objectivec2",
							"ruby2",
							"typescript2",
						}},
				},
				filepath: testutil.CallerDirectoryPath(t) + "/testdata/check_wrote_all_case_function.go",
			},
			want: true,
		},
		{
			name: "check_all_case_when_exported_package",
			args: args{
				enums: []enum{
					{
						name:        "Fruit",
						packageName: "x",
						patterns: []string{
							"Apple",
							"Orange",
							"Cherry",
						}},
				},
				filepath: testutil.CallerDirectoryPath(t) + "/testdata/check_all_case_when_exported_package.go",
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
