package kata

import "testing"

func TestDisemvowel(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "codeWarsExample",
			args: args{
				comment: "This website is for losers LOL!",
			},
			want: "Ths wbst s fr lsrs LL!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Disemvowel(tt.args.comment); got != tt.want {
				t.Errorf("Disemvowel() = %v, want %v", got, tt.want)
			}
		})
	}
}
