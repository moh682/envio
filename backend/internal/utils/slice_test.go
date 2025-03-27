package utils_test

import (
	"reflect"
	"testing"

	"github.com/moh682/envio/backend/internal/utils"
)

func TestFilterList(t *testing.T) {
	type args struct {
		list   []int
		filter utils.CallbackFunc[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "filter even numbers",
			args: args{
				list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(index int, item int) bool { return item%2 == 0 },
			},
			want: []int{2, 4, 6, 8},
		},
		{
			name: "filter odd numbers",
			args: args{
				list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(index int, item int) bool { return item%2 != 0 },
			},
			want: []int{1, 3, 5, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.Filter(tt.args.list, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterList() = %v, want %v", got, tt.want)
			}
		})
	}
}
