package main

import (
	"reflect"
	"testing"
)

func TestModifyImports(t *testing.T) {

}

func TestCleanPaths(t *testing.T) {
	matches := [][]string{
		{"from '@already-absolute/test/path.ts"},
		{"from './test/path.ts"},
		{"from '../test/path.ts"},
		{"from './../../test/path.ts"},
	}
	expected := [][]string{
		{"from '@already-absolute/test/path.ts"},
		{"./test/path.ts"},
		{"test/path.ts"},
		{"test/path.ts"},
	}

	result := cleanPaths(matches)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result:\n%v\nexpected:\n%v", result, expected)
	}
}
