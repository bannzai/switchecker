package main

import (
	"go/ast"
	"reflect"
	"testing"
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
		// TODO: Add test cases.
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
