package mapping

import (
	"reflect"
	"strconv"
	"testing"
)

func TestParallelMap(t *testing.T) {
	type args struct {
		source    interface{}
		transform interface{}
		t         reflect.Type
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "argument is not an array",
			args:    args{source: 1, transform: nil, t: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "transform function is nil",
			args:    args{source: []int{1, 2, 3}, transform: nil, t: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "transform is not a function",
			args:    args{source: []int{1, 2, 3}, transform: 1, t: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "T is not supplied",
			args: args{source: []int{1, 2, 3}, transform: func(num int) int {
				return num + 1
			}, t: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid transform",
			args: args{source: []int{1, 2, 3}, transform: func(num int) int {
				return num + 1
			}, t: reflect.TypeOf(1)},
			want:    []int{2, 3, 4},
			wantErr: false,
		},
		{
			name: "valid transform",
			args: args{source: []int{1, 2, 3}, transform: func(num int) string {
				return strconv.Itoa(num)
			}, t: reflect.TypeOf("")},
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParallelMap(tt.args.source, tt.args.transform, tt.args.t)

			if (err != nil) != tt.wantErr {
				t.Errorf("map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("map() = %v, want %v", got, tt.want)
			}
		})
	}
}
