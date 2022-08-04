package main

import "testing"

type testData struct {
	args    string
	want    string
	wantErr bool
}

func TestUnpackString(t *testing.T) {

	tests := []testData{
		{
			args:    "a4bc2d5e",
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			args:    "a4b0c2d5e",
			want:    "aaaaccddddde",
			wantErr: false,
		},
		{
			args:    "abcd",
			want:    "abcd",
			wantErr: false,
		},
		{
			args:    "45",
			want:    "",
			wantErr: true,
		},
		{
			args:    "",
			want:    "",
			wantErr: false,
		},
		{
			args:    `qwe\4\5`,
			want:    "qwe45",
			wantErr: false,
		},
		{
			args:    `qwe\45`,
			want:    "qwe44444",
			wantErr: false,
		},
		{
			args:    `qwe\\5`,
			want:    `qwe\\\\\`,
			wantErr: false,
		},
		{
			args:    `qwe\r5`,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {

		got, err := UnpackString(tt.args)
		if err != nil != tt.wantErr {
			t.Errorf("UnpackString() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if got != tt.want {
			t.Errorf("UnpackString() = %v, want %v", got, tt.want)
		}
	}
}
