package controllers

import "testing"

func Test_transTimeSecFromStr(t *testing.T) {
	type args struct {
		timeStr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{timeStr: "01:00:22"},
			3622,
		},
		{
			"case2",
			args{timeStr: "00:02:20"},
			140,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transTimeSecFromStr(tt.args.timeStr); got != tt.want {
				t.Errorf("transTimeSecFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
