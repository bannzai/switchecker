package main

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/bannzai/switchecker/pkg/testutil"
)

func Test_parse(t *testing.T) {
	type args struct {
		filepaths []string
	}
	tests := []struct {
		name string
		args args
		want []enum
	}{
		{
			name: "successfully parsed enum",
			args: args{
				filepaths: []string{
					testutil.CallerDirectoryPath(t) + "/testdata/parser/simple.go",
				},
			},
			want: []enum{
				{
					name:        "language",
					packageName: "testdata",
					packagePath: testutil.CallerDirectoryPath(t) + "/testdata/parser",
					patterns: []string{
						"golang",
						"swift",
						"objectivec",
						"ruby",
						"typescript",
					},
				},
			},
		},
		{
			name: "complex gofile pattern",
			args: args{
				filepaths: []string{
					testutil.CallerDirectoryPath(t) + "/testdata/parser/multiple_definition.go",
				},
			},
			want: []enum{
				{
					name:        "staticlanguage",
					packageName: "testdata",
					packagePath: testutil.CallerDirectoryPath(t) + "/testdata/parser",
					patterns: []string{
						"golang",
						"swift",
					},
				},
				{
					name:        "dynamiclanguage",
					packageName: "testdata",
					packagePath: testutil.CallerDirectoryPath(t) + "/testdata/parser",
					patterns: []string{
						"ruby",
						"python",
					},
				},
			},
		},
		{
			name: "No support define in function",
			args: args{
				filepaths: []string{
					testutil.CallerDirectoryPath(t) + "/testdata/parser/no_support_define_in_function.go",
				},
			},
			want: []enum{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.filepaths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseASTFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want *ast.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseASTFile(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseASTFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
