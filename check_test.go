package main

import (
	"testing"

	"github.com/bannzai/switchecker/pkg/testutil"
)

func Test_check(t *testing.T) {
	type args struct {
		enums     []enum
		filepaths []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no_satisfy_case",
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
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/no_satisfy_case.go"},
			},
			wantErr: true,
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
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/check_wrote_all_case_function.go"},
			},
			wantErr: false,
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
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/check_all_case_when_exported_package.go"},
			},
			wantErr: false,
		},
		{
			name: "check_plural_switch_pattern",
			args: args{
				enums: []enum{
					{
						name:        "language3",
						packageName: "testdata",
						patterns: []string{
							"golang3",
							"swift3",
							"objectivec3",
							"ruby3",
							"typescript3",
						}},
				},
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/check_plural_switch_pattern.go"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.enums, tt.args.filepaths); (got != nil) != tt.wantErr {
				t.Errorf("check() = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}
