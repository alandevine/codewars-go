package kata

import "testing"

func TestSeriesSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "n = 1",
			args: args{
				n: 1,
			},
			want: "1.00",
		},
		{
			name: "n = 2",
			args: args{
				n: 2,
			},
			want: "1.25",
		},
		{
			name: "n = 3",
			args: args{
				n: 3,
			},
			want: "1.39",
		},
		{
			name: "n = 4",
			args: args{
				n: 4,
			},
			want: "1.49",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SeriesSum(tt.args.n); got != tt.want {
				t.Errorf("SeriesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
