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
			name: "satisfy_case",
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
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/satisfy_case.go"},
			},
			wantErr: false,
		},
		{
			name: "satisfy_case_when_multipy_caller",
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
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/satisfy_case_when_multipy_caller.go"},
			},
			wantErr: true,
		},
		{
			name: "satisfy_case_for_exported_package",
			args: args{
				enums: []enum{
					{
						name:        "Language4",
						packageName: "x",
						patterns: []string{
							"Golang4",
							"Swift4",
							"Objectivec4",
							"Ruby4",
							"Typescript4",
						}},
				},
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/satisfy_case_for_exported_package.go"},
			},
			wantErr: false,
		},
		{
			name: "no_satisfy_case_for_exported_package",
			args: args{
				enums: []enum{
					{
						name:        "Language4",
						packageName: "x",
						patterns: []string{
							"Golang4",
							"Swift4",
							"Objectivec4",
							"Ruby4",
							"Typescript4",
						}},
				},
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/no_satisfy_case_for_exported_package.go"},
			},
			wantErr: true,
		},
		{
			name: "satisfy_case_of_single_line_definition",
			args: args{
				enums: []enum{
					{
						name:        "language6",
						packageName: "testdata",
						patterns: []string{
							"golang6",
							"swift6",
							"objectivec6",
							"ruby6",
							"typescript6",
						}},
				},
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/satisfy_case_of_single_line_definition.go"},
			},
			wantErr: false,
		},
		{
			name: "satisfy_case_of_complex_definition",
			args: args{
				enums: []enum{
					{
						name:        "language7",
						packageName: "testdata",
						patterns: []string{
							"golang7",
							"swift7",
							"objectivec7",
							"ruby7",
							"typescript7",
						}},
				},
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/satisfy_case_of_complex_definition.go"},
			},
			wantErr: false,
		},
		{
			name: "no_satisfy_case_of_single_line_definition",
			args: args{
				enums: []enum{
					{
						name:        "language8",
						packageName: "testdata",
						patterns: []string{
							"golang8",
							"swift8",
							"objectivec8",
							"ruby8",
							"typescript8",
						}},
				},
				filepaths: []string{testutil.CallerDirectoryPath(t) + "/testdata/no_satisfy_case_of_single_line_definition.go"},
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
