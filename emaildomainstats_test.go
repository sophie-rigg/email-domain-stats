package emaildomainstats

import (
	"reflect"
	"testing"
)

func TestProcessCSVFile(t *testing.T) {
	type args struct {
		filePath string
		options  []Option
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		{
			name: "Without any options",
			args: args{
				filePath: "testdata/test1.csv",
				options:  nil,
			},
			want: map[string]int{
				"google.com": 2,
				"goo.gl":     1,
			},
			wantErr: false,
		},
		{
			name: "Only WithColumnHeaders, correct one",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithColumnHeaders([]string{"Email"})},
			},
			want: map[string]int{
				"google.com": 2,
				"goo.gl":     1,
			},
			wantErr: false,
		},
		{
			name: "Only WithColumnHeaders, incorrect ones",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithColumnHeaders([]string{"first_name", "last_name"})},
			},
			want:    map[string]int{},
			wantErr: false,
		},
		{
			name: "Only WithSpecifiedColumns, correct one",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithColumns([]int{2})},
			},
			want: map[string]int{
				"google.com": 2,
				"goo.gl":     1,
			},
			wantErr: false,
		},
		{
			name: "Only WithSpecifiedColumns, incorrect ones",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithColumns([]int{1, 3})},
			},
			want:    map[string]int{},
			wantErr: false,
		},
		{
			name: "Only WithColumnHeaders, headers don't exist",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithColumnHeaders([]string{"bye", "last_name"})},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Only WithSpecifiedColumns, negative column number",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithColumns([]int{-1})},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Only WithSpecifiedColumns, too big column number",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithColumns([]int{20})},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WithColumnHeaders and WithColumns, correct ones",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithColumnHeaders([]string{"Email"}), WithColumns([]int{2})},
			},
			want: map[string]int{
				"google.com": 2,
				"goo.gl":     1,
			},
			wantErr: false,
		},
		{
			name: "file doesn't exist",
			args: args{
				filePath: "testdata/badtest2.csv",
				options:  nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "With specified delimeter",
			args: args{
				filePath: "testdata/test1.csv",
				options:  []Option{WithDelimeter(',')},
			},
			want: map[string]int{
				"google.com": 2,
				"goo.gl":     1,
			},
			wantErr: false,
		},
		{
			name: "With specified ; delimeter",
			args: args{
				filePath: "testdata/test-delim1.csv",
				options:  []Option{WithDelimeter(';')},
			},
			want: map[string]int{
				"google.com": 2,
				"goo.gl":     1,
			},
			wantErr: false,
		},
		{
			name: "With specified tab delimeter",
			args: args{
				filePath: "testdata/test1.tsv",
				options:  []Option{WithDelimeter('\t')},
			},
			want: map[string]int{
				"google.com": 2,
				"goo.gl":     1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessCSVFile(tt.args.filePath, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessCSVFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessCSVFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
